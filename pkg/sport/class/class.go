package class

import (
	"errors"
	"reflect"
	"strings"
)

type Class string

const (
	BL    = Class("Bayernliga")
	LL    = Class("Landesliga")
	BOL   = Class("Bezirksoberliga")
	BZL   = Class("Bezirksliga")
	BZK   = Class("Bezirksklasse")
	UeBOL = Class("Übergreifende Bezirksoberliga")
	UeBL  = Class("Übergreifende Bezirksliga")
	UeBK  = Class("Übergreifende Bezirksklasse")
)

// replaces dash and space
var replacer = strings.NewReplacer("-", "", " ", "")

// func GetAbbreviation returns short name of a class
func (c Class) GetAbbreviation() string {
	switch c {
	case BL:
		return "BL"
	case LL:
		return "LL"
	case BOL:
		return "BOL"
	case BZL:
		return "BZL"
	case BZK:
		return "BZK"
	case UeBOL:
		return "ÜBOL"
	case UeBL:
		return "ÜBL"
	case UeBK:
		return "ÜBK"
	default:
		return "invalid class (abbreviation)"
	}
}

// func GetName returns the full name of a class
func (c Class) GetName() string {
	return reflect.ValueOf(c).String()
}

// func Parse converts a given string to a Class
// it tries to convert different styles of classes to a Class type
func Parse(s string) (Class, error) {
	switch unifyString(s) {
	case
		unifyString(BL.GetName()),
		unifyString(BL.GetAbbreviation()):
		return BL, nil

	case
		unifyString(LL.GetName()),
		unifyString(LL.GetAbbreviation()):
		return LL, nil

	case
		unifyString(BOL.GetName()),
		unifyString(BOL.GetAbbreviation()):
		return BOL, nil

	case
		unifyString(BZL.GetName()),
		unifyString(BZL.GetAbbreviation()):
		return BZL, nil

	case
		unifyString(BZK.GetName()),
		unifyString(BZK.GetAbbreviation()):
		return BZK, nil

	case
		unifyString(UeBOL.GetName()),
		unifyString(UeBOL.GetAbbreviation()):
		return UeBOL, nil

	case
		unifyString(UeBL.GetName()),
		unifyString(UeBL.GetAbbreviation()):
		return UeBL, nil

	case
		unifyString(UeBK.GetName()),
		unifyString(UeBK.GetAbbreviation()):
		return UeBK, nil
	default:
		return "error", errors.New("todo")
	}
}

// func unifyString returns the value removed from dash or spaces in lowercase
// this will be used to compare strings with each other
func unifyString(s string) string {
	return strings.ToLower(replacer.Replace(s))
}
