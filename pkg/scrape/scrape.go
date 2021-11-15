package scrape

import (
	"net/url"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// scrapeTableResultset will scrape the requested website and searches for given objects
// the return will be a map of collyHTMLElement where the key is the search string
// this will enable to seachr multiple elements in one scrape
func scrape(u url.URL, htmlElements ...string) (map[string]colly.HTMLElement, error) {
	content := make(map[string]colly.HTMLElement)
	var return_error error

	c := colly.NewCollector(
		colly.AllowedDomains("bhv-handball.liga.nu"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.WithField("url", u).Debug("scraping url")
	})

	for _, htmlEl := range htmlElements {
		c.OnHTML(htmlEl, func(e *colly.HTMLElement) {
			content[htmlEl] = *e
		})
	}

	c.OnError(func(_ *colly.Response, err error) {
		return_error = err
	})

	c.Visit(u.String())

	return content, return_error
}
