package scrape

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// test func scrapeTableResultset
func TestScrape(t *testing.T) {
	u, _ := url.Parse("https://bhv-handball.liga.nu/cgi-bin/WebObjects/nuLigaHBDE.woa/wa/groupPage?championship=AV+2021%2F22&displayDetail=meetings&displayTyp=gesamt&group=281103")

	html_scrape, err := scrape(*u, "table.result-set")
	assert.Equal(t, html_scrape.Response.StatusCode, http.StatusOK, "expecting HTTP status 200 when scraping nuLiga table.result-set")
	assert.Equal(t, err, nil, "expecting no error when scraping nuLiga website")

	u, _ = url.Parse("https://task.media")
	html_scrape, err = scrape(*u, "html")
	if err == nil {
		t.Error("scraping on different URL worked but should not")
	}
}
