package parser

import (
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

// test func getMeetingReport
func TestGetMeetingReport(t *testing.T) {
	htmlCode := "<html><td><a target=\"_blank\" class=\"picto-pdf\" href=\"/cgi-bin/WebObjects/nuLigaDokumentHBDE.woa/wa/nuDokument?dokument=meetingReportHB&meeting=7013920\"></a></td></html>"
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader((htmlCode)))

	id, ok := getMeetingReport(doc.Find("html"))

	assert.Equal(t, true, ok, "expecting that id was found")
	assert.Equal(t, 7013920, id, "expecting different id from meeting report")

	htmlCode = "<html><td><span class=\"picto-pdf-disabled\"></span></td></html>"
	doc, _ = goquery.NewDocumentFromReader(strings.NewReader((htmlCode)))

	id, ok = getMeetingReport(doc.Find("html"))

	assert.Equal(t, false, ok, "expecting that id was not found")
	assert.Equal(t, -1, id, "expecting different id from meeting report")
}

// test func parseGermanTime
func TestParseGermanTime(t *testing.T) {
	time_actual := parseGermanTime("09.10.2021", "19:15")
	assert.Equal(t, 2021, time_actual.Year(), "expecting other value from year")
	assert.Equal(t, time.October, time_actual.Month(), "expecting other value from mont")
	assert.Equal(t, 9, time_actual.Day(), "expecting other value from day")
	assert.Equal(t, 19, time_actual.Hour(), "expecting other value from hour")
	assert.Equal(t, 15, time_actual.Minute(), "expecting other value from minute")
}

// test func standardizeSpaces
func TestStandardizeSpaces(t *testing.T) {
	testMap := make(map[string]string)

	testMap["Hello World"] = " Hello    World"
	testMap["Hello nuLiga"] = " Hello 		   nuLiga  	"
	testMap["hello"] = "	hello	"
	testMap["09.10.2021"] = `
	            09.10.2021
	          `

	for expected, actual := range testMap {
		assert.Equal(t, expected, standardizeSpaces(actual), "standardize space not removing spaces as expected")
	}
}
