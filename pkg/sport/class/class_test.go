package class

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func TestGetAbbreviation if it returns the correct abbreviation
func TestParse(t *testing.T) {
	testClasses := map[string]Class{
		"Bayernliga":      BL,
		"Bezirksoberliga": BOL,
		"ÜBOL":            UeBOL,
		"LL":              LL,
	}

	for check, expected := range testClasses {
		c, err := Parse(check)
		assert.Equal(t, expected, c, "Expected other constant from parsing class from value %s", check)
		assert.Equal(t, nil, err, "Expected no error from parsing class from value %s", check)
	}
}

// todo other tests
