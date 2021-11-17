package relay

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

var relayMap = map[string]string{
	"Nord":      "N",
	"Nord-Ost":  "NO",
	"Nord-West": "NW",

	"Ost": "O",

	"Süd":      "S",
	"Süd-Ost":  "SO",
	"Süd-West": "SW",

	"West": "W",

	"Mitte": "M",

	"A": "A",
	"B": "B",
	"C": "C",
	"D": "D",
	"E": "E",
	"F": "F",
}

// GetAbbreviation returns the abbreviation of a relay
// input will be the complete name of the relay
func GetAbbreviation(fullname string) string {
	for n, a := range relayMap {
		// dash will be removed because sometimes it is not provided
		if strings.ReplaceAll(strings.ToLower(fullname), "-", "") == strings.ReplaceAll(strings.ToLower(n), "-", "") {
			return a
		}
	}

	log.WithField("fullname", fullname).Warning("could not find abbreviation for relay")

	return fullname
}

// GetFullname returns the complete name of a relay
// input will be the abbreviation of that relay
func GetFullname(abbreviation string) string {
	for n, a := range relayMap {
		if strings.ToLower(a) == strings.ToLower(abbreviation) {
			return n
		}
	}

	log.WithField("abbreviation", abbreviation).Warning("could not find fullname for relay")

	return abbreviation
}
