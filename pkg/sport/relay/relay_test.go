package relay

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskmedia/nuScrape/pkg/sport/relay/relayName"
)

// Test func Parse if it returns the correct relay
func TestParse(t *testing.T) {
	testRelays := map[string]Relay{
		"Nord":        Relay{Name: relayName.N, Id: -1},
		"Nordwest":    Relay{Name: relayName.NW, Id: -1},
		"Süd-Ost":     Relay{Name: relayName.SO, Id: -1},
		"Mitte":       Relay{Name: relayName.M, Id: -1},
		"NW":          Relay{Name: relayName.NW, Id: -1},
		"B":           Relay{Name: relayName.B, Id: -1},
		"Nord-West 2": Relay{Name: relayName.NW, Id: 2},
	}

	for check, expected := range testRelays {
		c, err := Parse(check)
		assert.Equal(t, expected, c, "Expected other constant from parsing relay from value %s", check)
		assert.Equal(t, nil, err, "Expected no error from parsing relay from value %s", check)
	}
}
