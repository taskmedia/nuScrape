package scrape

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// GenerateGesamtspielplan will scrape and generate Matches for a given group
func ScrapeGesamtspielplan(s season.Season, c string, g group.Group) (colly.HTMLElement, error) {
	log.WithFields(
		log.Fields{
			"season":       s,
			"championship": c,
			"group":        g,
		},
	).Debug("generating new gesamtspielplan")

	url := generateUrlGesamtspielplan(s, c, g)

	// scrape and return table result-set
	return scrape(url, "table.result-set")
}
