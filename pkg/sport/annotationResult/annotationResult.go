package annotationResult

import (
	"errors"
	"reflect"
	"strings"
)

type AnnotationResult string

const (
	// using prefix `Annotation_` to be able to differ between u and U
	Annotation_v  = AnnotationResult("Spiel wurde in/auf genannte/n Halle/Termin verlegt")
	Annotation_u  = AnnotationResult("Spiel wird in/auf noch unbekannte/n Halle/Termin verlegt")
	Annotation_h  = AnnotationResult("Halle wurde geändert")
	Annotation_t  = AnnotationResult("Heimrecht wurde getauscht")
	Annotation_NH = AnnotationResult("Heimmannschaft nicht angegreten")
	Annotation_NG = AnnotationResult("Gastmannschaft nicht angetreten")
	Annotation_N2 = AnnotationResult("Heim- und Gastmannschaft nicht angetreten")
	Annotation_WH = AnnotationResult("Wertung gegen Heimmannschaft")
	Annotation_WG = AnnotationResult("Wertung gegen Gastmannschaft")
	Annotation_W2 = AnnotationResult("Wertung gegen Heim- und Gastmannschaft")
	Annotation_U  = AnnotationResult("Begegnung wurde umgewertet")
	Annotation_ZH = AnnotationResult("Heimmannschaft zurückgezogen")
	Annotation_ZG = AnnotationResult("Gastmannschaft zurückgezogen")

	// regex of allowed annotation abbreviations
	RegexAnnotations = `^(v|u|h|t|NH|NG|N2|WH|WG|W2|U|ZH|ZG)$`
)

// func GetAbbreviation returns short name of a annotationResult
func (ar AnnotationResult) GetAbbreviation() string {
	switch ar {
	case Annotation_v:
		return "v"
	case Annotation_u:
		return "u"
	case Annotation_h:
		return "h"
	case Annotation_t:
		return "t"
	case Annotation_NH:
		return "NH"
	case Annotation_NG:
		return "NG"
	case Annotation_N2:
		return "N2"
	case Annotation_WH:
		return "WH"
	case Annotation_WG:
		return "WG"
	case Annotation_W2:
		return "W2"
	case Annotation_U:
		return "U"
	case Annotation_ZH:
		return "ZH"
	case Annotation_ZG:
		return "ZG"

	default:
		return "invalid annotationResult (abbreviation)"
	}
}

// func GetName returns the full name of a annotationResult
func (c AnnotationResult) GetName() string {
	return reflect.ValueOf(c).String()
}

// func Parse converts a given string to a Class
// it tries to convert different styles of classes to a Class type
func Parse(s string) (AnnotationResult, error) {
	switch strings.TrimSpace(s) {
	case
		strings.TrimSpace(Annotation_v.GetName()),
		strings.TrimSpace(Annotation_v.GetAbbreviation()):
		return Annotation_v, nil
	case
		strings.TrimSpace(Annotation_u.GetName()),
		strings.TrimSpace(Annotation_u.GetAbbreviation()):
		return Annotation_u, nil
	case
		strings.TrimSpace(Annotation_h.GetName()),
		strings.TrimSpace(Annotation_h.GetAbbreviation()):
		return Annotation_h, nil
	case
		strings.TrimSpace(Annotation_t.GetName()),
		strings.TrimSpace(Annotation_t.GetAbbreviation()):
		return Annotation_t, nil
	case
		strings.TrimSpace(Annotation_NH.GetName()),
		strings.TrimSpace(Annotation_NH.GetAbbreviation()):
		return Annotation_NH, nil
	case
		strings.TrimSpace(Annotation_NG.GetName()),
		strings.TrimSpace(Annotation_NG.GetAbbreviation()):
		return Annotation_NG, nil
	case
		strings.TrimSpace(Annotation_N2.GetName()),
		strings.TrimSpace(Annotation_N2.GetAbbreviation()):
		return Annotation_N2, nil
	case
		strings.TrimSpace(Annotation_WH.GetName()),
		strings.TrimSpace(Annotation_WH.GetAbbreviation()):
		return Annotation_WH, nil
	case
		strings.TrimSpace(Annotation_WG.GetName()),
		strings.TrimSpace(Annotation_WG.GetAbbreviation()):
		return Annotation_WG, nil
	case
		strings.TrimSpace(Annotation_W2.GetName()),
		strings.TrimSpace(Annotation_W2.GetAbbreviation()):
		return Annotation_W2, nil
	case
		strings.TrimSpace(Annotation_U.GetName()),
		strings.TrimSpace(Annotation_U.GetAbbreviation()):
		return Annotation_U, nil
	case
		strings.TrimSpace(Annotation_ZH.GetName()),
		strings.TrimSpace(Annotation_ZH.GetAbbreviation()):
		return Annotation_ZH, nil
	case
		strings.TrimSpace(Annotation_ZG.GetName()),
		strings.TrimSpace(Annotation_ZG.GetAbbreviation()):
		return Annotation_ZG, nil

	default:
		return "error", errors.New("could not parse class type (unknown)")
	}
}
