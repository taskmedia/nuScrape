package sport

import (
	"github.com/taskmedia/nuScrape/pkg/sport/ageCategory"
	"github.com/taskmedia/nuScrape/pkg/sport/class"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/relay"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// Matches represents a slice of multiple Match structs.
type Gesamtspielplan struct {
	Matches      []Match                 `json:"matches" binding:"required"`
	Season       season.Season           `json:"season" binding:"required"`
	Championship string                  `json:"championship" binding:"required"`
	Group        group.Group             `json:"group" binding:"required"`
	AgeCategory  ageCategory.AgeCategory `json:"agecategory" binding:"required"`
	Class        class.Class             `json:"class" binding:"required"`
	Relay        relay.Relay             `json:"relay"`
}
