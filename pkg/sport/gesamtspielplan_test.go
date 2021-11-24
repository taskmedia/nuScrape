package sport

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskmedia/roundrobintournament"
)

// Test func GetDistinctTeams
// This test will check if the returned teams are correct
func TestGetDistinctTeams(t *testing.T) {

	teamCount := 4

	var gsp Gesamtspielplan

	// generate matches fot test Gesamtspielplan
	generatedMatches := roundrobintournament.GenerateRoundRobinTournamentMatchesByNumber(teamCount)

	// place test matches in gsp struct
	for _, gm := range generatedMatches {
		gsp.Matches = append(gsp.Matches, Match{Team: matchTeam{Home: gm[0], Guest: gm[1]}})
	}

	dt := gsp.GetDistinctTeams()

	// check length of distinct teams
	assert.Equal(t, teamCount, len(dt), "expected length of of distinct team differs")

	// check values of distinct teams
	for i := 0; i < teamCount; i++ {
		team := fmt.Sprintf("Team %d", i+1)
		fmt.Println(team)
		if !stringSliceContains(dt, team) {
			t.Fatalf("expected team '%s' was not in slice of distinct teams", team)
		}
	}
}

// func stringSliceContains will return true if a string is inside a slice of strings
func stringSliceContains(sl []string, str string) bool {
	for _, s := range sl {
		if s == str {
			return true
		}
	}
	return false
}
