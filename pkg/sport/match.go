package sport

import (
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/taskmedia/nuScrape/pkg/sport/annotationResult"
)

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
	Result annotationResult.AnnotationResult `json:"result"`
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

// func GetDescription will return a formatted description (multi line) of the match
func (m Match) GetDescription() string {
	// `Liga: %s
	// Klasse: %s
	// Altersklasse:

	desc := ""
	desc += "Heim: " + m.Team.Home + "\n"
	desc += "Gast: " + m.Team.Guest + "\n"
	if m.Goal.Home != 0 {
		desc += "Ergebnis: " + strconv.Itoa(m.Goal.Home) + ":" + strconv.Itoa(m.Goal.Guest) + "\n"
	}
	desc += "Spielnummer: " + strconv.Itoa(m.Id) + "\n"
	if m.LocationId != 0 {
		desc += "Hallennummer: " + strconv.Itoa(m.LocationId) + "\n"
	}
	if m.Referee != nil {
		desc += "Schiedsrichter:\n"
		for _, ref := range m.Referee {
			desc += "  - " + ref + "\n"
		}
	}
	if m.Annotation.Date != "" || m.Annotation.Result != "" {
		desc += "Anmerkungen:\n"
		if m.Annotation.Date != "" {
			desc += "  - " + m.Annotation.Date + "\n"
		}
		if m.Annotation.Result != "" {
			desc += "  - " + m.Annotation.Result.GetName() + "\n"
		}
	}

	return desc
}

// func GetReportUrl will return a URL if a Report is available
func (m Match) GetReportUrl() (url.URL, error) {
	u := url.URL{}

	if m.ReportId == 0 {
		return u, errors.New("no report available")
	}

	u.Scheme = "https"
	u.Host = "bhv-handball.liga.nu"
	u.Path = "/cgi-bin/WebObjects/nuLigaDokumentHBDE.woa/wa/nuDokument"

	query := url.Values{}
	query.Add("dokument", "meetingReportHB")
	query.Add("meeting", strconv.Itoa(m.ReportId))
	u.RawQuery = query.Encode()

	return u, nil
}
