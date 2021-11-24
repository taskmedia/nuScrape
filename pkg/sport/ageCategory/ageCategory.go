package ageCategory

import (
	"errors"
	"fmt"
	"strings"
)

type AgeCategory struct {
	Sex string
	Age string
}

// replaces dash, space and period
var replacer = strings.NewReplacer("-", "", " ", "", ".", "")

// func GetName returns the full name of a ageCategory
func (c *AgeCategory) GetName() string {
	if c.Age == "" {
		if c.Sex == "m" {
			return "Männer"
		} else {
			return "Frauen"
		}
	}

	ac := ""

	if c.Sex == "m" {
		ac = "männliche"
	} else {
		ac = "weibliche"
	}

	ac += fmt.Sprintf(" %s-Jugend", strings.ToUpper(c.Age))

	return ac
}

// func GetAbbreviation returns short name of a ageCategory
func (c *AgeCategory) GetAbbreviation() string {
	if c.Age == "" {
		if c.Sex == "m" {
			return "M"
		} else {
			return "F"
		}
	}
	return strings.ToLower(c.Sex) + strings.ToUpper(c.Age)
}

// func Parse converts a given string to a ageCategory
// it tries to convert different styles of ageCategory to a ageCategory struct
func Parse(s string) (AgeCategory, error) {
	switch unifyString(s) {
	case "männer":
		return AgeCategory{Sex: "m"}, nil

	case "frauen":
		return AgeCategory{Sex: "w"}, nil

	case "ma", "majugend", "majgd", "männlicheajugend", "männlicheajgd":
		return AgeCategory{Sex: "m", Age: "A"}, nil

	case "mb", "mbjugend", "mbjgd", "männlichebjugend", "männlichebjgd":
		return AgeCategory{Sex: "m", Age: "B"}, nil

	case "mc", "mcjugend", "mcjgd", "männlichecjugend", "männlichecjgd":
		return AgeCategory{Sex: "m", Age: "C"}, nil

	case "md", "mdjugend", "mdjgd", "männlichedjugend", "männlichedjgd":
		return AgeCategory{Sex: "m", Age: "D"}, nil

	case "wa", "wajugend", "wajgd", "weiblicheajugend", "weiblicheajgd":
		return AgeCategory{Sex: "w", Age: "A"}, nil

	case "wb", "wbjugend", "wbjgd", "weiblichebjugend", "weiblichebjgd":
		return AgeCategory{Sex: "w", Age: "B"}, nil

	case "wc", "wcjugend", "wcjgd", "weiblichecjugend", "weiblichecjgd":
		return AgeCategory{Sex: "w", Age: "C"}, nil

	case "wd", "wdjugend", "wdjgd", "weiblichedjugend", "weiblichedjgd":
		return AgeCategory{Sex: "w", Age: "D"}, nil

	default:
		return AgeCategory{}, errors.New("could not parse ageCategory")
	}
}

// func unifyString returns the value removed from dash or spaces in lowercase
// this will be used to compare strings with each other
func unifyString(s string) string {
	return strings.ToLower(replacer.Replace(s))
}
