package scrape

import (
	"net/url"

	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// scrapeTableResultset will scrape the requested website and searches for given objects
// the return will be a map of collyHTMLElement where the key is the search string
// this will enable to search multiple elements in one scrape
func scrape(u url.URL, htmlElements ...string) (map[string]*colly.HTMLElement, error) {
	content := make(map[string]*colly.HTMLElement)
	var return_error error

	c := colly.NewCollector(
		colly.AllowedDomains("bhv-handball.liga.nu"),
	)

	c.OnRequest(func(r *colly.Request) {
		log.WithField("url", u).Debug("scraping url")
	})

	for _, htmlEl := range htmlElements {
		c.OnHTML(htmlEl, func(e *colly.HTMLElement) {
			suffix := ""
			suffix_id := e.Attr("id")
			if suffix_id != "" {
				suffix += "#" + suffix_id
			}
			suffix_class := e.Attr("class")
			if suffix_class != "" {
				suffix += "." + suffix_class
			}

			// adding HTMLElement to htmlEl (generated)
			// issue: htmlEl can not be used here because it will not (correctly) available in the function
			// therefore the htmlEl has to be generated manually
			//
			// please keep in mind that a goselector has to be one tag:
			// e.g. 'div#content-col1 h1' will result in h1 and would differ to htmlEl
			content[e.Name+suffix] = e
			return
		})
	}

	c.OnError(func(_ *colly.Response, err error) {
		return_error = err
	})

	c.Visit(u.String())

	return content, return_error
}
