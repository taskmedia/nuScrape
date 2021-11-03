package scrape

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// Test func GenerateUrlGesamtspielplan
// This test will check if a generated URL will be correctly generated
func TestGenerateUrlGesamtspielplan(t *testing.T) {
	expectedUrlString := "https://bhv-handball.liga.nu/cgi-bin/WebObjects/nuLigaHBDE.woa/wa/groupPage?championship=AV+2021%2F22&displayDetail=meetings&displayTyp=gesamt&group=281103"

	// input
	s, _ := season.New("2021_22")
	c := "AV"
	g, _ := group.New("281103")

	u := generateUrlGesamtspielplan(s, c, g)

	assert.Equal(t, expectedUrlString, u.String(), "The two URLs should match")
}
