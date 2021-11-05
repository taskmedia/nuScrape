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

// ParseGesamtspielplan will parse a HTML table from nuLiga to Matches
func ParseGesamtspielplan(html colly.HTMLElement) (sport.Matches, error) {
	var matches sport.Matches
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

		// loop through the columns of the table
		tr.Find("td").Each(func(td_i int, td *goquery.Selection) {
			t := standardizeSpaces(td.Text())

			// switch through the column elements
			// each column must be considered and parsed separately
			switch td_i {
			// date
			case 1:
				if t != "" {
					// date has to be cached because not every column will have a date field
					// the date will be parsed together with the time
					cachedDate = t
				}

			// time
			case 2:
				match := regexp.MustCompile(`\d{2}:\d{2}`).FindStringSubmatch(t)
				m.Date = parseGermanTime(cachedDate, match[0])

				// check if date annotation is available
				// this happens when e.g. a game has been postponed
				attr, isAttr := td.Attr("title")
				if isAttr {
					m.Annotation.Date = attr
				}

			// location ID
			case 3:
				location, err := strconv.Atoi(t)
				if err != nil {
					log.WithFields(
						log.Fields{
							"locationId": t,
							"error":      err,
						}).Warning("can not parse location ID")
				} else {
					m.LocationId = location
				}

			// game ID
			case 4:
				game, err := strconv.Atoi(t)
				if err != nil {
					log.WithFields(
						log.Fields{
							"gameId": t,
							"error":  err,
						}).Warning("can not parse game ID")
				} else {
					m.LocationId = game
				}

			// hometeam
			case 5:
				m.Team.Home = t

			// guestteam
			case 6:
				m.Team.Guest = t

			// result / result annotation / referee
			case 7:
				goalsHome, goalsGuest, annotation, referee, err := parseResult(t, td)
				if err != nil {
					log.WithFields(
						log.Fields{
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
