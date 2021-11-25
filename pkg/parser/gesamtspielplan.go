package parser

import (
	"errors"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/sport"
	"github.com/taskmedia/nuScrape/pkg/sport/ageCategory"
	"github.com/taskmedia/nuScrape/pkg/sport/class"
	"github.com/taskmedia/nuScrape/pkg/sport/relay"
)

type Parse colly.HTMLElement

var re_ageCategory = regexp.MustCompile(`(Männer|Frauen|(?:[mw][ABCDEF].Jgd.)|(?:männliche|weibliche) [ABCDEF](?:\s|\-)Jugend)`)
var re_class = regexp.MustCompile(`((?:Bayern|Landes)liga|ÜBOL|ÜBL|Bezirks(?:ober)?(?:liga|klasse))`)
var re_relay1 = regexp.MustCompile(`((?:(?:[Nn]ord|[Oo]st|[Ss]üd|[Ww]est|Mitte)(?:-|\s\d)?){1,2})`)
var re_relay2 = regexp.MustCompile(`(?:Staffel )([ABCDEF])`)

// ParseGesamtspielplanInfo will parse a HTML (h1) group description from nuLiga to ageCategory, class, relay and error
func ParseGesamtspielplanInfo(html *colly.HTMLElement) (ageCategory.AgeCategory, class.Class, relay.Relay, error) {
	searchString := html.DOM.Find("h1").First().Text()

	ageCategoryString := ""
	f := re_ageCategory.FindStringSubmatch(searchString)
	if len(f) > 1 {
		ageCategoryString = f[1]
	}

	ac, err := ageCategory.Parse(ageCategoryString)
	if err != nil {
		return ageCategory.AgeCategory{}, "", relay.Relay{}, err
	}

	classString := ""
	f = re_class.FindStringSubmatch(searchString)
	if len(f) > 1 {
		classString = f[1]
	}

	class, err := class.Parse(classString)
	if err != nil {
		return ac, "", relay.Relay{}, err
	}

	// relay has two regex pattern because the structure is not really standardized
	relayString := ""
	f = re_relay1.FindStringSubmatch(searchString)
	if len(f) > 1 {
		relayString = f[1]
	} else {
		f = re_relay2.FindStringSubmatch(searchString)
		if len(f) > 1 {
			relayString = f[1]
		}
	}

	var r relay.Relay
	if relayString != "" {
		r, err = relay.Parse(relayString)
		if err != nil {
			return ac, "", r, err
		}
	} else {
		r = relay.Relay{}
	}

	// check if ageCategory and class are present otherwise return error
	// relay is not always present and optional
	if class == "" {
		err := errors.New("class not found in Gesamtspielplan info")
		return ac, class, r, err
	}

	return ac, class, r, nil
}

// ParseGesamtspielplanTable will parse a HTML table from nuLiga to Matches
func ParseGesamtspielplanTable(html *colly.HTMLElement) ([]sport.Match, error) {
	var matches []sport.Match
	cachedDate := ""
	skippedTableHeader := false

	// loop through the rows of the table
	html.DOM.Find("tr").Each(func(_ int, tr *goquery.Selection) {
		// check if the header of the table has to be skipped
		if !skippedTableHeader {
			skippedTableHeader = true
			return
		}

		m := sport.Match{}
		dateNotAvailable := 0

		// loop through the columns of the table
		tr.Find("td").Each(func(td_i int, td *goquery.Selection) {
			t := standardizeSpaces(td.Text())

			// switch through the column elements
			// each column must be considered and parsed separately
			switch td_i {
			case 0:
				// if no date is set in column 0 Termin offen will be set
				// the column span will be two - because of this the columns would not match anymore
				if t == "Termin offen" {
					dateNotAvailable = -1
					return
				}

			// date
			case 1 + dateNotAvailable:
				if t != "" {
					cachedDate = t
				}

			// time
			case 2 + dateNotAvailable:
				match := regexp.MustCompile(`\d{2}:\d{2}`).FindStringSubmatch(t)
				if len(match) >= 1 {
					matchDate, err := parseGermanTime(cachedDate, match[0])
					if err != nil {
						log.WithFields(log.Fields{
							"date":  cachedDate,
							"time":  match[0],
							"error": err,
						}).Warning("can not parse date/time")
					}
					m.Date = matchDate
				}

				// check if date annotation is available
				// this happens when e.g. a game has been postponed
				attr, isAttr := td.Attr("title")
				if isAttr {
					m.Annotation.Date = attr
				}

			// location ID
			case 3 + dateNotAvailable:
				location, err := strconv.Atoi(t)
				if err != nil {
					log.WithFields(log.Fields{
						"locationId": t,
						"error":      err,
					}).Warning("can not parse location ID")
				} else {
					m.LocationId = location
				}

			// game ID
			case 4 + dateNotAvailable:
				game, err := strconv.Atoi(t)
				if err != nil {
					log.WithFields(log.Fields{
						"gameId": t,
						"error":  err,
					}).Warning("can not parse game ID")
				} else {
					m.Id = game
				}

			// hometeam
			case 5 + dateNotAvailable:
				m.Team.Home = t

			// guestteam
			case 6 + dateNotAvailable:
				m.Team.Guest = t

			// result / result annotation / referee
			case 7 + dateNotAvailable:
				goalsHome, goalsGuest, annotation, referee, err := parseResult(t, td)
				if err != nil {
					log.WithFields(log.Fields{
						"goalsHome":  goalsHome,
						"goalsGuest": goalsGuest,
						"annotation": annotation,
						"referee":    referee,
						"err":        err,
					}).Warning("can not parse result")
				} else {
					m.Goal.Home = goalsHome
					m.Goal.Guest = goalsGuest
					m.Annotation.Result = annotation
					m.Referee = referee
				}
			}
		})

		// adding meeting report ID (link to PDF) if available
		mr, ok := getMeetingReport(tr)
		if ok {
			m.ReportId = mr
		}

		if m.Id != 0 {
			matches = append(matches, m)
		}
	})

	return matches, nil
}

// func getMeetingReport checks if in a html element the hyperlink to the meeting report (Spielbericht) is present
// if it is present the func will return the ID of this report
func getMeetingReport(html_element *goquery.Selection) (int, bool) {
	ln, ok := html_element.Find("a.picto-pdf").First().Attr("href")

	if ok {
		q, err := url.ParseQuery(ln)
		if err != nil {
			log.WithField("html_element", html_element).Warning("could not parse meeting report")
			return 0, false
		}

		id, err := strconv.Atoi(q["meeting"][0])
		if err != nil {
			log.WithField("query", q).Warning("could not convert meeting id to integer")
			return 0, false
		}

		return id, true
	}
	return 0, false
}

// func parseGermanTime will use the given time format from nuLiga and parse it into Time
// warning: the time will have not timezone information
func parseGermanTime(d, t string) (time.Time, error) {
	split_date := strings.Split(d, ".")

	year, err := strconv.Atoi(split_date[2])
	if err != nil {
		return time.Time{}, errors.New("could not parse year from date " + d)
	}

	mont_int, err := strconv.Atoi(split_date[1])
	if err != nil {
		return time.Time{}, errors.New("could not parse month from date " + d)
	}
	month := time.Month(mont_int)

	day, err := strconv.Atoi(split_date[0])
	if err != nil {
		return time.Time{}, errors.New("could not parse day from date " + d)
	}

	split_time := strings.Split(t, ":")

	hour, err := strconv.Atoi(split_time[0])
	if err != nil {
		return time.Time{}, errors.New("could not parse hour from time " + t)
	}

	minute, err := strconv.Atoi(split_time[1])
	if err != nil {
		return time.Time{}, errors.New("could not parse minute from time " + t)
	}

	timeLocation, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		log.Fatal(err)
	}
	timeInUtc := time.Date(year, month, day, hour, minute, 0, 0, timeLocation)

	return timeInUtc, nil
}

// func parseResult will parse the input from  result column and parse it to different information
// this field has not only the goals inside. also an annotation and the referees can be available.
func parseResult(resultString string, html_element *goquery.Selection) (int, int, string, []string, error) {
	home := 0
	guest := 0
	annotation := ""
	var referee []string

	if regexp.MustCompile(`(WH|WG|NH|NG|ZH|ZG)`).MatchString(resultString) {
		// check if result annotation is available
		// this may be the case if e.g. the game has been postponed
		attr, isAttr := html_element.Attr("alt")
		if isAttr {
			annotation = attr
		}
	} else if regexp.MustCompile(`\d{1,2}:\d{1,2}`).MatchString(resultString) {
		// check if result is available
		goals := strings.Split(resultString, ":")
		var err error
		home, err = strconv.Atoi(goals[0])
		if err != nil {
			return home, guest, annotation, referee, errors.New("could not parse home team goals")
		}
		guest, err = strconv.Atoi(goals[1])
		if err != nil {
			return home, guest, annotation, referee, errors.New("could not parse guest team goals")
		}
	} else if regexp.MustCompile(`[a-zA-z]+`).MatchString(resultString) {
		// check if referee is available
		// this may be the case if the match has not yet taken place
		attr, isAttr := html_element.Find("span").Attr("title")
		if isAttr {
			refSlice := strings.Split(attr, "/")
			// remove whitespace from referees
			for i := range refSlice {
				refSlice[i] = strings.TrimSpace(refSlice[i])
			}
			referee = refSlice
		}
	}

	return home, guest, annotation, referee, nil
}

// func standardizeSpaces will remove any whitespace from the given string
func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
