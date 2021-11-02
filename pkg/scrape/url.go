package scrape

import (
	"net/url"

	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// generateUrlGesamtspielplan will generate a URL for scrapping gesamtspielplan
func generateUrlGesamtspielplan(s season.Season, c string, g group.Group) url.URL {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = "bhv-handball.liga.nu"
	u.Path = "/cgi-bin/WebObjects/nuLigaHBDE.woa/wa/groupPage"

	query := url.Values{}
	query.Add("displayTyp", "gesamt")
	query.Add("displayDetail", "meetings")
	query.Add("championship", c+" "+s.FormatSeasonUrlNuLiga()) // todo: test this
	query.Add("group", g.String())
	u.RawQuery = query.Encode()

	return u
}
