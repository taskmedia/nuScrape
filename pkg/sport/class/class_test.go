package class

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func GetAbbreviation if it returns the correct class
func TestParse(t *testing.T) {
	testClasses := map[string]Class{
		"Bayernliga":      BL,
		"Bezirksoberliga": BOL,
		"ÜBOL":            UeBOL,
		"LL":              LL,
	}

	for check, expected := range testClasses {
		c, err := Parse(check)
		assert.Equal(t, expected, c, "expected other constant from parsing class from value %s", check)
		assert.Equal(t, nil, err, "expected no error from parsing class from value %s", check)
	}
}

// Test func GetAbbreviation if it returns the correct abbreviation of a class
func TestGetAbbreviation(t *testing.T) {
	assert.Equal(t, "BL", BL.GetAbbreviation(), "expected other abbreviation from class BL")
	assert.Equal(t, "LL", LL.GetAbbreviation(), "expected other abbreviation from class LL")
	assert.Equal(t, "ÜBOL", UeBOL.GetAbbreviation(), "expected other abbreviation from class UeBOL")
	assert.Equal(t, "BZK", BZK.GetAbbreviation(), "expected other abbreviation from class BZK")
}

// Test func GetName if it returns the correct full name of a class
func TestGetName(t *testing.T) {
	assert.Equal(t, "Bayernliga", BL.GetName(), "expected other full name from class BL")
	assert.Equal(t, "Landesliga", LL.GetName(), "expected other full name from class LL")
	assert.Equal(t, "Bezirksoberliga", BOL.GetName(), "expected other full name from class BOL")
	assert.Equal(t, "Übergreifende Bezirksliga", UeBL.GetName(), "expected other full name from class UeBL")
}
