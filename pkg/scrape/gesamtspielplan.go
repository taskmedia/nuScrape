package scrape

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// GenerateGesamtspielplan will scrape Gesamtspielplan and return Gesamtspielplan Info and Table HTMLElement
func ScrapeGesamtspielplan(s season.Season, c string, g group.Group) (*colly.HTMLElement, *colly.HTMLElement, error) {
	log.WithFields(log.Fields{
		"season":       s,
		"championship": c,
		"group":        g,
	},
	).Debug("generating new gesamtspielplan")

	url := generateUrlGesamtspielplan(s, c, g)

	// scrape and return table and gesamtspielplan info (class, relay, age, sex)
	htmlInfo := "div#content-col1"
	htmlTable := "table.result-set"
	scrapeMap, err := scrape(url, htmlInfo, htmlTable)
	if err != nil {
		return nil, nil, err
	}

	return scrapeMap[htmlInfo], scrapeMap[htmlTable], err
}
