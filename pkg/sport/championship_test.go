package sport

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func ValidateChampionshipAbb
// This test will test different abbreviations and check the return value
func TestValidateChampionshipAbb(t *testing.T) {
	expectedSuccess := []string{"BHV", "AV", "MF"}
	expectedFailure := []string{"XX", "YZ", "Bhv", "av"}

	for _, test := range expectedSuccess {
		assert.Equal(t, true, ValidateChampionshipAbb(test), "Expected abbreviation (%s) should be valid for championship", test)
	}

	for _, test := range expectedFailure {
		assert.Equal(t, false, ValidateChampionshipAbb(test), "Expected abbreviation (%s) should be NOT valid for championship", test)
	}
}
