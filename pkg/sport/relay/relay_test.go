package relay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func Parse if it returns the correct relay
func TestParse(t *testing.T) {
	testRelays := map[string]Relay{
		"Nord":     N,
		"Nordwest": NW,
		"Süd-Ost":  SO,
		"Mitte":    M,
		"NW":       NW,
		"B":        B,
	}

	for check, expected := range testRelays {
		c, err := Parse(check)
		assert.Equal(t, expected, c, "Expected other constant from parsing relay from value %s", check)
		assert.Equal(t, nil, err, "Expected no error from parsing relay from value %s", check)
	}
}

// todo other tests
