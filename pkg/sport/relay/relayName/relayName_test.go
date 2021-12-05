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
		"Nord-Ost":  NO,
		"Nord-West": NW,
		"Ost":       O,
		"Süd":       S,
		"Süd-Ost":   SO,
		"Süd-West":  SW,
		"West":      W,
		"Mitte":     M,
		"A":         A,
		"B":         B,
		"C":         C,
		"D":         D,
		"E":         E,
		"F":         F,
	}

	for check, expected := range testRelayName {
		fmt.Println(check)
		rn, err := Parse(check)

		assert.Equal(t, expected, rn, "Expected other constant from parsing relayName from value %s", check)
		assert.Equal(t, nil, err, "Expected no error from parsing relayName from value %s", check)
	}

	// test for error
	// test for error
	_, err := Parse("unknown-type")
	if err == nil {
		t.Error("expected string `unknown-type` not be able to be parsed")
	} else {
		assert.Equal(t, "could not parse relayName type (unknown)", err.Error(), "expected string `unknown-type` not output different error")
	}
}

// Test func GetAbbreviation if it returns the correct abbreviation of a relayName
func TestGetAbbreviation(t *testing.T) {
	assert.Equal(t, "N", N.GetAbbreviation(), "expected other abbreviation from relayName N")
	assert.Equal(t, "NO", NO.GetAbbreviation(), "expected other abbreviation from relayName NO")
	assert.Equal(t, "NW", NW.GetAbbreviation(), "expected other abbreviation from relayName NW")
	assert.Equal(t, "O", O.GetAbbreviation(), "expected other abbreviation from relayName O")
	assert.Equal(t, "S", S.GetAbbreviation(), "expected other abbreviation from relayName S")
	assert.Equal(t, "SO", SO.GetAbbreviation(), "expected other abbreviation from relayName SO")
	assert.Equal(t, "SW", SW.GetAbbreviation(), "expected other abbreviation from relayName SW")
	assert.Equal(t, "W", W.GetAbbreviation(), "expected other abbreviation from relayName W")
	assert.Equal(t, "M", M.GetAbbreviation(), "expected other abbreviation from relayName M")
	assert.Equal(t, "A", A.GetAbbreviation(), "expected other abbreviation from relayName A")
	assert.Equal(t, "B", B.GetAbbreviation(), "expected other abbreviation from relayName B")
	assert.Equal(t, "C", C.GetAbbreviation(), "expected other abbreviation from relayName C")
	assert.Equal(t, "D", D.GetAbbreviation(), "expected other abbreviation from relayName D")
	assert.Equal(t, "E", E.GetAbbreviation(), "expected other abbreviation from relayName E")
	assert.Equal(t, "F", F.GetAbbreviation(), "expected other abbreviation from relayName F")
}

// Test func GetName if it returns the correct full name of a relayName
func TestGetName(t *testing.T) {
	assert.Equal(t, "Nord", N.GetName(), "expected other full name form relayName N")
	assert.Equal(t, "Nord-Ost", NO.GetName(), "expected other full name form relayName NO")
	assert.Equal(t, "Nord-West", NW.GetName(), "expected other full name form relayName NW")
	assert.Equal(t, "Ost", O.GetName(), "expected other full name form relayName O")
	assert.Equal(t, "Süd", S.GetName(), "expected other full name form relayName S")
	assert.Equal(t, "Süd-Ost", SO.GetName(), "expected other full name form relayName SO")
	assert.Equal(t, "Süd-West", SW.GetName(), "expected other full name form relayName SW")
	assert.Equal(t, "West", W.GetName(), "expected other full name form relayName W")
	assert.Equal(t, "Mitte", M.GetName(), "expected other full name form relayName M")
	assert.Equal(t, "A", A.GetName(), "expected other full name form relayName A")
	assert.Equal(t, "B", B.GetName(), "expected other full name form relayName B")
	assert.Equal(t, "C", C.GetName(), "expected other full name form relayName C")
	assert.Equal(t, "D", D.GetName(), "expected other full name form relayName D")
	assert.Equal(t, "E", E.GetName(), "expected other full name form relayName E")
	assert.Equal(t, "F", F.GetName(), "expected other full name form relayName F")
}
