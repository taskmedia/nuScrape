package relay

import (
	"errors"
	"strings"
)

type Relay string

const (
	N  = Relay("Nord")
	NO = Relay("Nord-Ost")
	NW = Relay("Nord-West")

	O = Relay("Ost")

	S  = Relay("Süd")
	SO = Relay("Süd-Ost")
	SW = Relay("Süd-West")

	W = Relay("West")

	M = Relay("Mitte")

	A = Relay("A")
	B = Relay("B")
	C = Relay("C")
	D = Relay("D")
	E = Relay("E")
	F = Relay("F")
)

// replaces dash and space
var replacer = strings.NewReplacer("-", "", " ", "")

func (r Relay) GetAbbreviation() string {
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

func (r Relay) GetName() string {
	switch r {
	case N:
		return "Nord"
	case NO:
		return "Nord-Ost"
	case NW:
		return "Nord-West"
	case O:
		return "Ost"
	case S:
		return "Süd"
	case SO:
		return "Süd-Ost"
	case SW:
		return "Süd-West"
	case W:
		return "West"
	case M:
		return "Mitte"
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
		return "invalid relay (name)"
	}
}

// func Parse converts a given string to a Relay
// it tries to convert different styles of relays to a Relay type
func Parse(s string) (Relay, error) {
	switch unifyString(s) {
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
		return "", errors.New("todo")
	}
}

// func unifyString returns the value removed from dash or spaces in lowercase
// this will be used to compare strings with each other
func unifyString(s string) string {
	return strings.ToLower(replacer.Replace(s))
}
