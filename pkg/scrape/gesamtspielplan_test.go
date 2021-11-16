package scrape

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// test func scrapeTableResultset
func TestScrapeGesamtspielplan(t *testing.T) {
	c := "AV"
	s, _ := season.New("2021_22")
	g, _ := group.New("281103")

	html_info, _, err := ScrapeGesamtspielplan(s, c, g)
	assert.Equal(t, nil, err, "scraping nuLiga url failed")
	assert.Equal(t, http.StatusOK, html_info.Response.StatusCode, "scraping of nuLiga url failed (status code)")
}
