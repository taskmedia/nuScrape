package championship

import (
	"errors"
	"reflect"
	"strings"
)

type Championship string

const (
	BHV = Championship("Bayerischer Handball-Verband")
	UF  = Championship("Unterfranken")
	OF  = Championship("Oberfranken")
	MF  = Championship("Mittelfranken")
	OS  = Championship("Ostbayern")
	SW  = Championship("Schwaben")
	AB  = Championship("Altbayern")
	AV  = Championship("Alpenvorland")
	OB  = Championship("Oberbayern")
)

// func GetAbbreviation returns short name of a championship
func (c Championship) GetAbbreviation() string {
	switch c {
	case BHV:
		return "BHV"
	case UF:
		return "UF"
	case OF:
		return "OF"
	case MF:
		return "MF"
	case OS:
		return "OS"
	case SW:
		return "SW"
	case AB:
		return "AB"
	case AV:
		return "AV"
	case OB:
		return "OB"
	default:
		return "invalid championship (abbreviation)"
	}
}

// func GetName returns the full name of a championship
func (c Championship) GetName() string {
	return reflect.ValueOf(c).String()
}

// func ParseAbbreviation converts a given abbreviation to a championship
func ParseAbbreviation(s string) (Championship, error) {
	switch strings.ToUpper(s) {
	case
		BHV.GetAbbreviation():
		return BHV, nil

	case
		UF.GetAbbreviation():
		return UF, nil

	case
		OF.GetAbbreviation():
		return OF, nil

	case
		MF.GetAbbreviation():
		return MF, nil

	case
		OS.GetAbbreviation():
		return OS, nil

	case
		SW.GetAbbreviation():
		return SW, nil

	case
		AB.GetAbbreviation():
		return AB, nil

	case
		AV.GetAbbreviation():
		return AV, nil

	case
		OB.GetAbbreviation():
		return OB, nil

	default:
		return "error", errors.New("could not parse class type (unknown)")
	}
}
