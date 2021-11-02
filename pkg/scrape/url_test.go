package scrape

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

func TestGenerateUrlGesamtspielplan(t *testing.T) {
	expectedUrlString := "https://bhv-handball.liga.nu/cgi-bin/WebObjects/nuLigaHBDE.woa/wa/groupPage?championship=AV+2021%2F22&displayDetail=meetings&displayTyp=gesamt&group=281103"

	// input
	s, _ := season.New("2021_22")
	c := "AV"
	g, _ := group.New("281103")
	// , c string, g Group)

	u := generateUrlGesamtspielplan(s, c, g)

	require.Equal(t, expectedUrlString, u.String(), "The two URLs should match.")
}
