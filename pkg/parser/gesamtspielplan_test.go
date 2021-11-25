package parser

import (
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/stretchr/testify/assert"
	"github.com/taskmedia/nuScrape/pkg/sport/ageCategory"
	"github.com/taskmedia/nuScrape/pkg/sport/class"
	"github.com/taskmedia/nuScrape/pkg/sport/relay"
	"github.com/taskmedia/nuScrape/pkg/sport/relay/relayName"
)

// struct for testing func ParseGesamtspielplanInfo
// this struct will represent different values used for one test
type testStructParseGesamtspielplanInfo struct {
	html        string
	ageCategory ageCategory.AgeCategory
	class       class.Class
	relay       relay.Relay
	err         error
}

// test func parseGesamtspielplanInfo
func TestParseGesamtspielplanInfo(t *testing.T) {
	testGspInfo := []testStructParseGesamtspielplanInfo{
		testStructParseGesamtspielplanInfo{
			html: `<div id="content-col1">
  <h1>
    Alpenvorland 2021/22
    <br>
    Bezirksklasse Männer Staffel Nord-West
    <br>
    Tabelle und Spielplan (Aktuell)
  </h1>
      <h2>Schiedsrichterkosten</h2>
      <p>
        <label>Durchschnitt Staffel:</label> 36,68 €
      </p>
</div>`,
			ageCategory: ageCategory.AgeCategory{Sex: "m", Age: ""},
			class:       class.BZK,
			relay:       relay.Relay{Name: relayName.NW, Id: 0},
			err:         nil,
		},
		testStructParseGesamtspielplanInfo{
			html:        `<div id="content-col1"><h1>BHV 2021/22<br>ÜBOL wC-Jgd. Südwest 1<br>Tabelle und Spielplan (Aktuell)</h1></div>`,
			ageCategory: ageCategory.AgeCategory{Sex: "w", Age: "C"},
			class:       class.UeBOL,
			relay:       relay.Relay{Name: relayName.SW, Id: 1},
			err:         nil,
		},
		testStructParseGesamtspielplanInfo{
			html:        `<div id="content-col1"><h1>BHV 2021/22<br>Bayernliga männliche A-Jugend Nordwest<br>Tabelle und Spielplan (Aktuell)</h1></div>`,
			ageCategory: ageCategory.AgeCategory{Sex: "m", Age: "A"},
			class:       class.BL,
			relay:       relay.Relay{Name: relayName.NW, Id: 0},
			err:         nil,
		},
		testStructParseGesamtspielplanInfo{
			html:        `<div id="content-col1"><h1>Schwaben 2021/22<br>Bezirksliga Frauen<br>Spielplan (gesamt)</h1></div>`,
			ageCategory: ageCategory.AgeCategory{Sex: "w", Age: ""},
			class:       class.BZL,
			relay:       relay.Relay{Name: "", Id: 0},
			err:         nil,
		},
		testStructParseGesamtspielplanInfo{
			html:        `<div id="content-col1"><h1>Schwaben 2021/22<br>Bezirksoberliga Männer - Staffel A<br>Tabelle und Spielplan (Aktuell)</h1></div>`,
			ageCategory: ageCategory.AgeCategory{Sex: "m", Age: ""},
			class:       class.BOL,
			relay:       relay.Relay{Name: "A", Id: 0},
			err:         nil,
		},
	}

	for _, result := range testGspInfo {
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader((result.html)))
		sel := doc.Find("div#content-col1").First()
		htmlElement := colly.HTMLElement{
			DOM: sel,
		}

		ac, c, r, err := ParseGesamtspielplanInfo(&htmlElement)
		assert.Equal(t, result.ageCategory, ac, "expected other ageCategory from html '%s'", result.html)
		assert.Equal(t, result.class, c, "expected other class from html '%s'", result.html)
		assert.Equal(t, result.relay, r, "expected other relay from html '%s'", result.html)
		assert.Equal(t, result.err, err, "expected other error from html '%s'", result.html)
	}

}

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
	assert.Equal(t, 0, id, "expecting different id from meeting report")
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

// struct for testing func ParseResult
// this struct will represent different values used for one test
type testStructParseResult struct {
	html       string
	expected   string
	goalsHome  int
	goalsGuest int
	annotation string
	referee    []string
	err        error
}

// test func ParseResult
func TestParseResult(t *testing.T) {
	testResults := []testStructParseResult{
		testStructParseResult{
			html: `<div id="wrapper"><span alt="10:9 zur Halbzeit" title="11:12 zur Halbzeit">  22:24  	 </span></div>`,
			goalsHome:  22,
			goalsGuest: 24,
		},
		testStructParseResult{
			html:    `<div id="wrapper"><span title="Mustermann Max">Must.</span></div>`,
			referee: []string{"Mustermann Max"},
		},
		testStructParseResult{
			html:    `<div id="wrapper"><span title="Doe John / Nordmann Ola">Doe/Nord.</span></div>`,
			referee: []string{"Doe John", "Nordmann Ola"},
		},
		testStructParseResult{
			html:       `<div id="wrapper"><span><span>0:0</span></span></div>`,
			goalsHome:  0,
			goalsGuest: 0,
		},
		testStructParseResult{
			html: `<div id="wrapper">
          
           
	          
          	&nbsp;
          
        </div>`,
		},
		testStructParseResult{
			// workaround for test - using div directly instead of td
			html:       `<div id="wrapper" alt="Wertung gegen Gastmannschaft" class="center" title="Wertung gegen Gastmannschaft">WG</div>`,
			annotation: "Wertung gegen Gastmannschaft",
		},
	}

	for _, result := range testResults {
		doc, _ := goquery.NewDocumentFromReader(strings.NewReader((result.html)))
		sel := doc.Find("div#wrapper").First()
		text := standardizeSpaces(sel.Text())

		goalsHome, goalsGuest, annotation, referee, err := parseResult(text, sel)
		assert.Equal(t, result.goalsHome, goalsHome, "expected other goalsHome from html '%s'", result.html)
		assert.Equal(t, result.goalsGuest, goalsGuest, "expected other goalsGuest from html '%s'", result.html)
		assert.Equal(t, result.annotation, annotation, "expected other annotation from html '%s'", result.html)
		assert.Equal(t, result.referee, referee, "expected other referee from html '%s'", result.html)
		assert.Equal(t, result.err, err, "expected other err from html '%s'", result.html)
	}

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
