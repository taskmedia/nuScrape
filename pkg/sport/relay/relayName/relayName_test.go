package relayName

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func Parse if it returns the correct relayName
func TestParse(t *testing.T) {
	testRelayName := map[string]RelayName{
		"Nord":      N,
		"Nord-West": NW,
		"Süd":       S,
		"Süd-Ost":   SO,
		"Südost":    SO,
		"Mitte":     M,
		"A":         A,
	}

	for check, expected := range testRelayName {
		fmt.Println(check)
		rn, err := Parse(check)

		assert.Equal(t, expected, rn, "Expected other constant from parsing relayName from value %s", check)
		assert.Equal(t, nil, err, "Expected no error from parsing relayName from value %s", check)
	}
}

// Test func GetAbbreviation if it returns the correct abbreviation of a relayName
func TestGetAbbreviation(t *testing.T) {
	assert.Equal(t, "N", N.GetAbbreviation(), "expected other abbreviation from relayName N")
	assert.Equal(t, "NW", NW.GetAbbreviation(), "expected other abbreviation from relayName NW")
	assert.Equal(t, "S", S.GetAbbreviation(), "expected other abbreviation from relayName S")
	assert.Equal(t, "M", M.GetAbbreviation(), "expected other abbreviation from relayName M")
	assert.Equal(t, "C", C.GetAbbreviation(), "expected other abbreviation from relayName C")
}

// Test func GetName if it returns the correct full name of a relayName
func TestGetName(t *testing.T) {
	assert.Equal(t, "Nord", N.GetName(), "expected other full name from relayName N")
	assert.Equal(t, "Süd-West", SW.GetName(), "expected other full name from relayName SW")
	assert.Equal(t, "Mitte", M.GetName(), "expected other full name from relayName M")
	assert.Equal(t, "B", B.GetName(), "expected other full name from relayName B")
}
