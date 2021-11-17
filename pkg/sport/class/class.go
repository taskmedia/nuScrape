package class

import log "github.com/sirupsen/logrus"

var classMap = map[string]string{
	"Bayernliga":      "BL",
	"Landesliga":      "LL",
	"Bezirksoberliga": "BOL",
	"Bezirksliga":     "BZL",
	"Bezirksklasse":   "BZK",

	// the following classes are mostly not referenced with fullname
	"BZL":  "BZL",  // Bezirksliga
	"BZK":  "BZK",  // Bezirksklasse
	"ÜBOL": "ÜBOL", // Übergreifende Bezirksoberliga
	"ÜBL":  "ÜBL",  // Übergreifende Bezirksliga
	"ÜBK":  "ÜBK",  // Übergreifende Bezirksklasse

	// never used - but due list completion added
	"Übergreifende Bezirksoberliga": "ÜBOL",
	"Übergreifende Bezirksliga":     "ÜBL",
	"Übergreifende Bezirksklasse":   "ÜBK",
}

// GetAbbreviation returns the abbreviation of a class
// input will be the complete name of the class
func GetAbbreviation(fullname string) string {
	if len(fullname) <= 4 {
		return fullname
	}

	for n, a := range classMap {
		if fullname == n {
			return a
		}
	}

	log.WithField("fullname", fullname).Warning("could not find abbreviation for class")

	return fullname
}

// GetFullname returns the complete name of a class
// input will be the abbreviation of that class
func GetFullname(abbreviation string) string {
	for n, a := range classMap {
		if a == abbreviation {
			return n
		}
	}

	log.WithField("abbreviation", abbreviation).Warning("could not find fullname for class")

	return abbreviation
}
