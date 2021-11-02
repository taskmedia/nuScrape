package season

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Season is a type that represents a nuLiga season (year)
// The format should be YYYY_YY
type Season string

// New creates a new Season and checks if it is valid otherwise returns an error.
// Checks for the format YYYY_YY: e.g. 2021_22 for 2021/22
func New(s string) (Season, error) {
	if !regexp.MustCompile(`\d{4}_\d{2}`).MatchString(s) {
		return "", errors.New("season pattern not valid (expected 6 digits with underscore separated)")
	}

	// check if years match to each other
	// valid: e.g. 2021_22, 1999_00
	// not valid: e.g. 2021_23
	hr, _ := strconv.Atoi(s[2:4])
	rr, _ := strconv.Atoi(s[5:])
	if !(((hr + 1) == rr) || (hr == 99 && rr == 0)) {
		return "", errors.New("year patter not valid (years must follow each other in season)")
	}

	return Season(s), nil
}

// FormatSeasonUrlNuLiga formats a Season to a string which is used in the URL of nuLiga.
func (y Season) FormatSeasonUrlNuLiga() string {
	return fmt.Sprintf("%s/%s", y[:4], y[5:])
}
