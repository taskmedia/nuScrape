package scrape

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test func scrapeTableResultset
func TestScrapeTableResultset(t *testing.T) {
	u, _ := url.Parse("https://bhv-handball.liga.nu/cgi-bin/WebObjects/nuLigaHBDE.woa/wa/groupPage?championship=AV+2021%2F22&displayDetail=meetings&displayTyp=gesamt&group=281103")

	html, err := scrapeTableResultset(*u)
	assert.Equal(t, nil, err, "scraping nuLiga url failed")
	assert.Equal(t, http.StatusOK, html.Response.StatusCode, "scraping of nuLiga url failed (status code)")

	u, _ = url.Parse("https://task.media")
	html, err = scrapeTableResultset(*u)
	assert.Equal(t, "scraping website was not successful", err.Error(), "scraping nuLiga url failed")
}
