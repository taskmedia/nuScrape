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

	html_tag := "table.result-set"
	html_scrape, err := scrape(*u, html_tag)
	assert.Equal(t, html_scrape[html_tag].Response.StatusCode, http.StatusOK, "expecting HTTP status 200 when scraping nuLiga table.result-set")
	assert.Equal(t, err, nil, "expecting no error when scraping nuLiga website")

	u, _ = url.Parse("https://task.media")
	html_tag = "html"
	_, err = scrape(*u, html_tag)
	if err == nil {
		t.Error("scraping on different URL worked but should not")
	}
}
