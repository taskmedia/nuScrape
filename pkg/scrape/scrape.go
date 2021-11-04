package scrape

import (
	"errors"
	"net/url"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// scrapeTableResultset will scrape the requested website and searches for table.result-set object
func scrape(u url.URL, htmlElement string) (colly.HTMLElement, error) {
	var content colly.HTMLElement
	var return_error error

	c := colly.NewCollector(
		colly.AllowedDomains("bhv-handball.liga.nu"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.WithField("url", u).Debug("scraping url")
	})

	c.OnHTML(htmlElement, func(e *colly.HTMLElement) {
		content = *e
	})

	c.OnError(func(_ *colly.Response, err error) {
		return_error = err
	})

	c.Visit(u.String())

	// check content object is empty
	if content.Response == nil {
		return_error = errors.New("scraping website was not successful")
	}

	return content, return_error
}
