package relayName

import (
	"errors"
	"reflect"
	"strings"
)

type RelayName string

const (
	N  = RelayName("Nord")
	NO = RelayName("Nord-Ost")
	NW = RelayName("Nord-West")

	O = RelayName("Ost")

	S  = RelayName("Süd")
	SO = RelayName("Süd-Ost")
	SW = RelayName("Süd-West")

	W = RelayName("West")

	M = RelayName("Mitte")

	A = RelayName("A")
	B = RelayName("B")
	C = RelayName("C")
	D = RelayName("D")
	E = RelayName("E")
	F = RelayName("F")
)

// replaces dash and space
var replacer = strings.NewReplacer("-", "", " ", "")

// func GetAbbreviation returns short name of a relayname
func (r RelayName) GetAbbreviation() string {
	switch r {
	case N:
		return "N"
	case NO:
		return "NO"
	case NW:
		return "NW"
	case O:
		return "O"
	case S:
		return "S"
	case SO:
		return "SO"
	case SW:
		return "SW"
	case W:
		return "W"
	case M:
		return "M"
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	case F:
		return "F"
	default:
		return "invalid relay (abbreviation)"
	}
}

// func GetName returns the full name of a relayname
func (r RelayName) GetName() string {
	return reflect.ValueOf(r).String()
}

// func Parse converts a given string to a RelayName
// it tries to convert different styles of relaynamess to a RelayName type
func Parse(str string) (RelayName, error) {
	switch unifyString(str) {
	case
		unifyString(N.GetName()),
		unifyString(N.GetAbbreviation()):
		return N, nil

	case
		unifyString(NO.GetName()),
		unifyString(NO.GetAbbreviation()):
		return NO, nil

	case
		unifyString(NW.GetName()),
		unifyString(NW.GetAbbreviation()):
		return NW, nil

	case
		unifyString(O.GetName()),
		unifyString(O.GetAbbreviation()):
		return O, nil

	case
		unifyString(S.GetName()),
		unifyString(S.GetAbbreviation()):
		return S, nil

	case
		unifyString(SO.GetName()),
		unifyString(SO.GetAbbreviation()):
		return SO, nil

	case
		unifyString(SW.GetName()),
		unifyString(SW.GetAbbreviation()):
		return SW, nil

	case
		unifyString(W.GetName()),
		unifyString(W.GetAbbreviation()):
		return W, nil

	case
		unifyString(M.GetName()),
		unifyString(M.GetAbbreviation()):
		return M, nil

	case
		unifyString(A.GetName()),
		unifyString(A.GetAbbreviation()):
		return A, nil

	case
		unifyString(B.GetName()),
		unifyString(B.GetAbbreviation()):
		return B, nil

	case
		unifyString(C.GetName()),
		unifyString(C.GetAbbreviation()):
		return C, nil

	case
		unifyString(D.GetName()),
		unifyString(D.GetAbbreviation()):
		return D, nil

	case
		unifyString(E.GetName()),
		unifyString(E.GetAbbreviation()):
		return E, nil

	case
		unifyString(F.GetName()),
		unifyString(F.GetAbbreviation()):
		return F, nil

	default:
		return "", errors.New("could not parse relayName type (unknown)")
	}
}

// func unifyString returns the value removed from dash or spaces in lowercase
// this will be used to compare strings with each other
func unifyString(s string) string {
	return strings.ToLower(replacer.Replace(s))
}
