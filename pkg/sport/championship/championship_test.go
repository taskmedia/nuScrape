package championship

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func GetAbbreviation if it returns the correct championship
func TestParse(t *testing.T) {
	testChampionship := map[string]Championship{
		"bhv": BHV,
		"UF":  UF,
		"Of":  OF,
		"mF":  MF,
		"OS":  OS,
		"SW":  SW,
		"AB":  AB,
		"AV":  AV,
		"OB":  OB,
	}

	for check, expected := range testChampionship {
		c, err := ParseAbbreviation(check)
		assert.Equal(t, expected, c, "expected other constant from parsing championship from value %s", check)
		assert.Equal(t, nil, err, "expected no error from parsing championship from value %s", check)
	}
}

// Test func GetAbbreviation if it returns the correct abbreviation of a championship
func TestGetAbbreviation(t *testing.T) {
	assert.Equal(t, "BHV", BHV.GetAbbreviation(), "expected other abbreviation from championship BHV")
	assert.Equal(t, "UF", UF.GetAbbreviation(), "expected other abbreviation from championship UF")
	assert.Equal(t, "OF", OF.GetAbbreviation(), "expected other abbreviation from championship OF")
	assert.Equal(t, "MF", MF.GetAbbreviation(), "expected other abbreviation from championship MF")
	assert.Equal(t, "OS", OS.GetAbbreviation(), "expected other abbreviation from championship OS")
	assert.Equal(t, "SW", SW.GetAbbreviation(), "expected other abbreviation from championship SW")
	assert.Equal(t, "AB", AB.GetAbbreviation(), "expected other abbreviation from championship AB")
	assert.Equal(t, "AV", AV.GetAbbreviation(), "expected other abbreviation from championship AV")
	assert.Equal(t, "OB", OB.GetAbbreviation(), "expected other abbreviation from championship OB")
}

// Test func GetName if it returns the correct full name of a championship
func TestGetName(t *testing.T) {
	assert.Equal(t, "Bayerischer Handball-Verband", BHV.GetName(), "expected other full name from championship BHV")
	assert.Equal(t, "Unterfranken", UF.GetName(), "expected other full name from championship UF")
	assert.Equal(t, "Oberfranken", OF.GetName(), "expected other full name from championship OF")
	assert.Equal(t, "Mittelfranken", MF.GetName(), "expected other full name from championship MF")
	assert.Equal(t, "Ostbayern", OS.GetName(), "expected other full name from championship OS")
	assert.Equal(t, "Schwaben", SW.GetName(), "expected other full name from championship SW")
	assert.Equal(t, "Altbayern", AB.GetName(), "expected other full name from championship AB")
	assert.Equal(t, "Alpenvorland", AV.GetName(), "expected other full name from championship AV")
	assert.Equal(t, "Oberbayern", OB.GetName(), "expected other full name from championship OB")
}
