package relay

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/taskmedia/nuScrape/pkg/sport/relay/relayName"
)

type Relay struct {
	Name relayName.RelayName
	Id   int
}

// replaces dash and space
var replacer = strings.NewReplacer("-", "", " ", "")

// searches for a number in unified relay
var re_relayNumber = regexp.MustCompile(`(.*)(\d)$`)

// func GetAbbreviation returns short name of a relay
// if no relay given a empty string will be returned
func (r Relay) GetAbbreviation() string {
	if r.Name == "" {
		return ""
	}

	// check if no ID is provided
	if r.Id == 0 {
		return r.Name.GetAbbreviation()
	}

	return fmt.Sprintf("%s %d", r.Name.GetAbbreviation(), r.Id)
}

// func GetName returns the full name of a relay
// if no relay given a empty string will be returned
func (r Relay) GetName() string {
	if r.Name == "" {
		return ""
	}

	// check if no ID is provided
	if r.Id == 0 {
		return r.Name.GetName()
	}

	return fmt.Sprintf("%s %d", r.Name.GetName(), r.Id)
}

// func Parse converts a given string to a Relay
// it tries to convert different styles of relays to a Relay type
func Parse(s string) (Relay, error) {
	searchString := unifyString(s)
	classNumber := 0

	// check if number
	re := re_relayNumber.FindStringSubmatch(searchString)
	if len(re) > 2 {
		searchString = re[1]
		cn, err := strconv.Atoi(re[2])
		if err != nil {
			return Relay{}, err
		}
		classNumber = cn
	}

	rn, err := relayName.Parse(searchString)
	if err != nil {
		return Relay{}, err
	}

	return Relay{Name: rn, Id: classNumber}, nil
}

// func unifyString returns the value removed from dash or spaces in lowercase
// this will be used to compare strings with each other
func unifyString(s string) string {
	return strings.ToLower(replacer.Replace(s))
}
