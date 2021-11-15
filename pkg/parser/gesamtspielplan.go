package parser

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/sport"
)

type Parse colly.HTMLElement

var re_ageCategory = regexp.MustCompile(`(Männer|Frauen|(?:[mw][ABCDEF].Jgd.)|(?:männliche|weibliche) [ABCDEF](?:\s|\-)Jugend)`)
var re_class = regexp.MustCompile(`((?:Bayern|Landes)liga|ÜBOL|ÜBL|Bezirks(?:ober)?(?:liga|klasse))`)
var re_relay1 = regexp.MustCompile(`((?:(?:[Nn]ord|[Oo]st|[Ss]üd|[Ww]est|Mitte)(?:-|\s\d)?){1,2})`)
var re_relay2 = regexp.MustCompile(`(?:Staffel )([ABCDEF])`)

// ParseGesamtspielplanInfo will parse a HTML (h1) group description from nuLiga to ageCategory, class, relay
func ParseGesamtspielplanInfo(html *colly.HTMLElement) (string, string, string, error) {
	searchString := html.DOM.Find("h1").First().Text()

	ageCategory := ""
	f := re_ageCategory.FindStringSubmatch(searchString)
	if len(f) > 1 {
		ageCategory = f[1]
	}

	class := ""
	f = re_class.FindStringSubmatch(searchString)
	if len(f) > 1 {
		class = f[1]
	}

	// relay has two regex pattern because the structure is not really standardized
	relay := ""
	f = re_relay1.FindStringSubmatch(searchString)
	if len(f) > 1 {
		relay = f[1]
	} else {
		f = re_relay2.FindStringSubmatch(searchString)
		if len(f) > 1 {
			relay = f[1]
		}
	}

	return ageCategory, class, relay, nil
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
					m.Date = parseGermanTime(cachedDate, match[0])
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

		matches = append(matches, m)
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
			return -1, false
		}

		id, err := strconv.Atoi(q["meeting"][0])
		if err != nil {
			log.WithField("query", q).Warning("could not convert meeting id to integer")
			return -1, false
		}

		return id, true
	}
	return -1, false
}

// func parseGermanTime will use the given time format from nuLiga and parse it into Time
// warning: the time will have not timezone information
func parseGermanTime(d, t string) time.Time {
	split_date := strings.Split(d, ".")

	datetimeFormatted := fmt.Sprintf("%s-%s-%sT%s:00.000Z", split_date[2], split_date[1], split_date[0], t)

	dt, err := time.Parse(time.RFC3339, datetimeFormatted)
	if err != nil {
		fmt.Println(err)
	}

	return dt
}

// func parseResult will parse the input from  result column and parse it to different information
// this field has not only the goals inside. also an annotation and the referees can be available.
func parseResult(resultString string, html_element *goquery.Selection) (int, int, string, []string, error) {
	home := -1
	guest := -1
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
			referee = strings.Split(attr, "/")
		}
	}

	return home, guest, annotation, referee, nil
}

// func standardizeSpaces will remove any whitespace from the given string
func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
