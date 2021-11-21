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
