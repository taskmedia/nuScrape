package class

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func TestGetAbbreviation if it returns the correct abbreviation
func TestGetAbbreviation(t *testing.T) {
	assert.Equal(t, "BL", GetAbbreviation("Bayernliga"), "Expected class abbreviation was not correct for Bayernliga")
	assert.Equal(t, "BOL", GetAbbreviation("Bezirksoberliga"), "Expected class abbreviation was not correct for Bezirksoberliga")
	assert.Equal(t, "ÜBOL", GetAbbreviation("ÜBOL"), "Expected class abbreviation was not correct for ÜBOL")
}

// Test func TestGetFullname if it returns the correct complete name of a class
func TestGetFullname(t *testing.T) {
	assert.Equal(t, "Landesliga", GetFullname("LL"), "Expected class fullname was not correct for LL")
	// assert.Equal(t, "Bezirksklasse", GetFullname("BZK"), "Expected fullname was not correct for BZK") // not possible because duplicate entries
}
