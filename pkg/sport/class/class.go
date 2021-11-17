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
func GetAbbreviation(name string) string {
	if len(name) <= 4 {
		return name
	}

	for n, a := range classMap {
		if name == n {
			return a
		}
	}

	log.Warning("could not find abbreviation for class %s", name)

	return name
}

// GetFullname returns the complete name of a class
// input will be the abbreviation of that class
func GetFullname(abbreviation string) string {
	for n, a := range classMap {
		if a == abbreviation {
			return n
		}
	}

	log.Warning("could not find fullname for class %s", abbreviation)

	return abbreviation
}
