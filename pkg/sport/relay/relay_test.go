package relay

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskmedia/nuScrape/pkg/sport/relay/relayName"
)

// Test func GetAbbreviation if it returns the correct abbreviation of a relay
func TestGetAbbreviation(t *testing.T) {
	testRelay := map[string]Relay{
		"N":   Relay{Name: relayName.N, Id: 0},
		"SW":  Relay{Name: relayName.SW, Id: 0},
		"W 2": Relay{Name: relayName.W, Id: 2},
		"M":   Relay{Name: relayName.M, Id: 0},
		"B":   Relay{Name: relayName.B, Id: 0},
		"":    Relay{},
	}

	for check, expected := range testRelay {
		assert.Equal(t, check, expected.GetAbbreviation(), "expected other abbreviation from relay %s", check)
	}
}

// Test func GetName if it returns the correct full name of a relay
func TestGetName(t *testing.T) {
	testRelay := map[string]Relay{
		"Nord":     Relay{Name: relayName.N, Id: 0},
		"Süd-West": Relay{Name: relayName.SW, Id: 0},
		"West 2":   Relay{Name: relayName.W, Id: 2},
		"Mitte":    Relay{Name: relayName.M, Id: 0},
		"B":        Relay{Name: relayName.B, Id: 0},
		"":         Relay{},
	}

	for check, expected := range testRelay {
		assert.Equal(t, check, expected.GetName(), "expected other full name from relay %s", check)
	}
}

// Test func Parse if it returns the correct relay
func TestParse(t *testing.T) {
	testRelays := map[string]Relay{
		"Nord":        Relay{Name: relayName.N, Id: 0},
		"Nordwest":    Relay{Name: relayName.NW, Id: 0},
		"Süd-Ost":     Relay{Name: relayName.SO, Id: 0},
		"Mitte":       Relay{Name: relayName.M, Id: 0},
		"NW":          Relay{Name: relayName.NW, Id: 0},
		"B":           Relay{Name: relayName.B, Id: 0},
		"Nord-West 2": Relay{Name: relayName.NW, Id: 2},
	}

	for check, expected := range testRelays {
		c, err := Parse(check)
		assert.Equal(t, expected, c, "Expected other constant from parsing relay from value %s", check)
		assert.Equal(t, nil, err, "Expected no error from parsing relay from value %s", check)
	}
}
