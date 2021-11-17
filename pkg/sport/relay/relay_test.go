package relay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func TestGetAbbreviation if it returns the correct abbreviation
func TestGetAbbreviation(t *testing.T) {
	assert.Equal(t, "N", GetAbbreviation("Nord"), "Expected relay abbreviation was not correct for Nord")
	assert.Equal(t, "NW", GetAbbreviation("Nord-West"), "Expected relay abbreviation was not correct for Nord-West")
	assert.Equal(t, "SO", GetAbbreviation("Südost"), "Expected relay abbreviation was not correct for Südost")
	assert.Equal(t, "B", GetAbbreviation("B"), "Expected relay abbreviation was not correct for B")
}

// Test func TestGetFullname if it returns the correct complete name of a class
func TestGetFullname(t *testing.T) {
	assert.Equal(t, "Nord-West", GetFullname("NW"), "Expected relay fullname was not correct for NW")
	assert.Equal(t, "Süd-Ost", GetFullname("SO"), "Expected relay fullname was not correct for SO")
	assert.Equal(t, "Mitte", GetFullname("M"), "Expected relay fullname was not correct for M")
	assert.Equal(t, "A", GetFullname("A"), "Expected relay fullname was not correct for A")

}
