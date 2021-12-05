package sport

import (
	"github.com/taskmedia/nuScrape/pkg/sport/ageCategory"
	"github.com/taskmedia/nuScrape/pkg/sport/championship"
	"github.com/taskmedia/nuScrape/pkg/sport/class"
	"github.com/taskmedia/nuScrape/pkg/sport/group"
	"github.com/taskmedia/nuScrape/pkg/sport/relay"
	"github.com/taskmedia/nuScrape/pkg/sport/season"
)

// Matches represents a slice of multiple Match structs.
type Gesamtspielplan struct {
	Matches      []Match                   `json:"matches" binding:"required"`
	Season       season.Season             `json:"season" binding:"required"`
	Championship championship.Championship `json:"championship" binding:"required"`
	Group        group.Group               `json:"group" binding:"required"`
	AgeCategory  ageCategory.AgeCategory   `json:"agecategory" binding:"required"`
	Class        class.Class               `json:"class" binding:"required"`
	Relay        relay.Relay               `json:"relay"`
}

// func GetDistinctTeams will return a list of all teams in a Gesamtspielplan
// the returned string slice will not contain duplicates.
func (gsp Gesamtspielplan) GetDistinctTeams() []string {
	var dt []string

	// add first team
	dt = append(dt, gsp.Matches[0].Team.Home)

	// loop through all Guest teams from first team and add them to the list
	// this will probably have a better performance because no slice comparison will be required
	// possible because matches are of type double round robin tournament
	for _, m := range gsp.Matches {
		if m.Team.Home == dt[0] {
			dt = append(dt, m.Team.Guest)
		}
	}
	return dt
}

// func GetDescription will return a formatted description (multi line) of the GSP (without matches)
func (gsp Gesamtspielplan) GetDescription() string {
	desc := ""
	desc += "Liga: " + gsp.Championship.GetName() + "\n"
	desc += "Klasse: " + gsp.Class.GetName() + " " + gsp.Relay.GetName() + "\n"
	desc += "Gruppennummer: " + gsp.Group.String() + "\n"
	desc += "Altersklasse: " + gsp.AgeCategory.GetName() + "\n"
	desc += "Saison: " + string(gsp.Season) + "\n"

	return desc
}
