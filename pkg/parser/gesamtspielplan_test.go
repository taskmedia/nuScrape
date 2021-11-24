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
	testTime, err := parseGermanTime("09.10.2021", "19:15")
	assert.Equal(t, nil, err, "expecting no error from parseGermanTime")
	assert.Equal(t, 2021, testTime.Year(), "expecting other value from year")
	assert.Equal(t, time.October, testTime.Month(), "expecting other value from mont")
	assert.Equal(t, 9, testTime.Day(), "expecting other value from day")
	assert.Equal(t, 19, testTime.Hour(), "expecting other value from hour")
	assert.Equal(t, 15, testTime.Minute(), "expecting other value from minute")

	testTime, err = parseGermanTime("11.11.2021", "19:30")
	if err != nil {
		assert.Equal(t, nil, err, "expecting no error from parseGermanTime CET")
	}
	assert.Equal(t, "2021-11-11 19:30:00 +0100 CET", testTime.String(), "expected CET time not matching")

	testTime, err = parseGermanTime("12.08.2021", "19:30")
	if err != nil {
		assert.Equal(t, nil, err, "expecting no error from parseGermanTime CEST")
	}
	assert.Equal(t, "2021-08-12 19:30:00 +0200 CEST", testTime.String(), "expected CEST time not matching")
}

// test func standardizeSpaces
func TestStandardizeSpaces(t *testing.T) {
	testStandardize := map[string]string{
		"Hello World": " Hello    World",
		"Hello nuLiga": " Hello 		   nuLiga  	",
		"hello": "	hello	",
		"09.10.2021": `
	            09.10.2021
	          `,
	}

	for expected, actual := range testStandardize {
		assert.Equal(t, expected, standardizeSpaces(actual), "standardize space not removing spaces as expected")
	}
}
