package annotationResult

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func GetAbbreviation if it returns the correct annotationResult
func TestParse(t *testing.T) {
	testAnnotationResults := map[string]AnnotationResult{
		"u":  Annotation_u,
		"U":  Annotation_U,
		"NH": Annotation_NH,
		"WH": Annotation_WH,
		"ZH": Annotation_ZH,
	}

	for check, expected := range testAnnotationResults {
		c, err := Parse(check)
		assert.Equal(t, expected, c, "expected other constant from parsing annotationResult from value %s", check)
		assert.Equal(t, nil, err, "expected no error from parsing annotationResult from value %s", check)
	}
}

// Test func GetAbbreviation if it returns the correct abbreviation of a annotationResult
func TestGetAbbreviation(t *testing.T) {
	assert.Equal(t, "u", Annotation_u.GetAbbreviation(), "expected other abbreviation from annotationResult u")
	assert.Equal(t, "U", Annotation_U.GetAbbreviation(), "expected other abbreviation from annotationResult U")
	assert.Equal(t, "NH", Annotation_NH.GetAbbreviation(), "expected other abbreviation from annotationResult NH")
	assert.Equal(t, "WH", Annotation_WH.GetAbbreviation(), "expected other abbreviation from annotationResult WH")
	assert.Equal(t, "ZH", Annotation_ZH.GetAbbreviation(), "expected other abbreviation from annotationResult ZH")
}

// Test func GetName if it returns the correct full name of a annotationResult
func TestGetName(t *testing.T) {
	assert.Equal(t, "Spiel wird in/auf noch unbekannte/n Halle/Termin verlegt", Annotation_u.GetName(), "expected other full name from annotationResult u")
	assert.Equal(t, "Begegnung wurde umgewertet", Annotation_U.GetName(), "expected other full name from annotationResult U")
	assert.Equal(t, "Heimmannschaft nicht angegreten", Annotation_NH.GetName(), "expected other full name from annotationResult NH")
	assert.Equal(t, "Wertung gegen Heimmannschaft", Annotation_WH.GetName(), "expected other full name from annotationResult WH")
	assert.Equal(t, "Heimmannschaft zurückgezogen", Annotation_ZH.GetName(), "expected other full name from annotationResult ZH")
}
