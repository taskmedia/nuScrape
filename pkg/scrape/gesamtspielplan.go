package scrape

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	. "github.com/taskmedia/nuScrape/pkg/sport/group"
	. "github.com/taskmedia/nuScrape/pkg/sport/season"
)

// GenerateGesamtspielplan will scrape and generate Matches for a given group
func GenerateGesamtspielplan(s Season, c string, g Group) {
	log.WithFields(
		log.Fields{
			"season":       s,
			"championship": c,
			"group":        g,
		},
	).Debug("generating new gesamtspielplan")

	url := generateUrlGesamtspielplan(s, c, g)

	// testing url
	fmt.Println(url.String())
}
