package ageCategory

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func Parse if it returns the correct ageCategory
func TestParse(t *testing.T) {
	testAgeCategory := map[string]AgeCategory{
		"Männer":             AgeCategory{Sex: "m", Age: ""},
		"Frauen":             AgeCategory{Sex: "w", Age: ""},
		"wA":                 AgeCategory{Sex: "w", Age: "A"},
		"männliche A-Jugend": AgeCategory{Sex: "m", Age: "A"},
		"wB-Jgd.":            AgeCategory{Sex: "w", Age: "B"},
		"weibliche C-Jgd.":   AgeCategory{Sex: "w", Age: "C"},
	}

	for check, expected := range testAgeCategory {
		ag, err := Parse(check)

		assert.Equal(t, expected, ag, "Expected other constant from parsing ageCategory from value %s", check)
		assert.Equal(t, nil, err, "Expected no error from parsing ageCategory from value %s", check)
	}
}

// Test func GetAbbreviation if it returns the correct abbreviation of a ageCategory
func TestGetAbbreviation(t *testing.T) {
	testAgeCategory := map[string]AgeCategory{
		"M":  AgeCategory{Sex: "m", Age: ""},
		"F":  AgeCategory{Sex: "w", Age: ""},
		"wA": AgeCategory{Sex: "w", Age: "A"},
		"mA": AgeCategory{Sex: "m", Age: "A"},
		"wB": AgeCategory{Sex: "w", Age: "B"},
		"wC": AgeCategory{Sex: "w", Age: "C"},
	}

	for check, expected := range testAgeCategory {
		assert.Equal(t, check, expected.GetAbbreviation(), "expected other abbreviation from ageCategory %s", check)
	}
}

// Test func GetName if it returns the correct full name of a ageCategory
func TestGetName(t *testing.T) {
	testAgeCategory := map[string]AgeCategory{
		"Männer":             AgeCategory{Sex: "m", Age: ""},
		"Frauen":             AgeCategory{Sex: "w", Age: ""},
		"weibliche A-Jugend": AgeCategory{Sex: "w", Age: "A"},
		"männliche A-Jugend": AgeCategory{Sex: "m", Age: "A"},
		"weibliche B-Jugend": AgeCategory{Sex: "w", Age: "B"},
		"weibliche C-Jugend": AgeCategory{Sex: "w", Age: "C"},
	}

	for check, expected := range testAgeCategory {
		assert.Equal(t, check, expected.GetName(), "expected other full name from ageCategory %s", check)
	}
}
