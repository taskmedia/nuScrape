package parser

import (
	"github.com/gocolly/colly"
	"github.com/taskmedia/nuScrape/pkg/sport"
)

type Parse colly.HTMLElement

// ParseGesamtspielplan will parse a HTML table from nuLiga to Matches
func ParseGesamtspielplan(html colly.HTMLElement) (sport.Matches, error) {
	// todo
	return nil, nil
}
