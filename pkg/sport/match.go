package sport

import (
	"time"

	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// Matches represents a slice of multiple Match structs.
type Gesamtspielplan struct {
	Matches      []Match       `json:"matches" binding:"required"`
	Season       season.Season `json:"season" binding:"required"`
	Championship string        `json:"championship" binding:"required"`
	Group        group.Group   `json:"group" binding:"required"`
}

// Match represents a nuLiga match (game)
type Match struct {
	// Date represents the date the match is scheduled
	Date time.Time `json:"date"`

	// Team represents the teams participating on the match
	Team matchTeam `json:"team"`

	// Goal represents the achieved goals of the match
	Goal matchGoal `json:"goal"`

	// LocationId represents the ID number of the location where the match takes place
	LocationId int `json:"location"`

	// Id represents the unique ID of the match
	Id int `json:"id" binding:"required"`

	// Annotation represents the annotations deposited for a match
	Annotation matchAnnotation `json:"annotation"`

	// MatchReport represents the URL to the PDF file containing the report of a match
	ReportId int `json:"report"`

	// Referee represents a slice of strings containing each referee related to the match
	Referee []string `json:"referee"`
}

// matchAnnotation represents annotation to a Match
type matchAnnotation struct {
	// Date represents matchAnntations dedicated to the date or time of a Match
	// e.g. if a match was postponed
	Date string `json:"date"`

	// Result represents matchAnntations dedicated to the result of a Match
	// e.g. a judging by the referees
	Result string `json:"result"`
}

// matchTeam represents the teams participating in the match
type matchTeam struct {
	// Home represents the home team of a match
	Home string `json:"home"`

	// Guest represents the guest team of a match
	Guest string `json:"guest"`
}

// matchGoal represents the goals of a match
type matchGoal struct {
	// Home represents the achieved goals of the home team
	Home int `json:"home"`

	// Guest represents the achieved goals of the guest team
	Guest int `json:"guest"`
}
