package scrape

import (
	log "github.com/sirupsen/logrus"
	"github.com/taskmedia/nuScrape/pkg/parser"
	"github.com/taskmedia/nuScrape/pkg/sport"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// GenerateGesamtspielplan will scrape and generate Matches for a given group
func GenerateGesamtspielplan(s season.Season, c string, g group.Group) (sport.Matches, error) {
	log.WithFields(
		log.Fields{
			"season":       s,
			"championship": c,
			"group":        g,
		},
	).Debug("generating new gesamtspielplan")

	url := generateUrlGesamtspielplan(s, c, g)

	// scrape website
	html_scrape, err := scrapeTableResultset(url)
	if err != nil {
		return nil, err
	}

	// parse website content to Matches
	parser.ParseGesamtspielplan(html_scrape)

	return nil, nil
}
