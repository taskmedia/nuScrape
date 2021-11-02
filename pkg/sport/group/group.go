package group

import (
	"errors"
	"strconv"
)

// Group represents a id of a nuLiga group used to identify multiple groups in a league.
type Group int

// New creates a new Group and validate it
func New(groupString string) (Group, error) {
	g, err := strconv.Atoi(groupString)
	if err != nil {
		return -1, errors.New("group not an integer")
	}

	if g <= 0 {
		return -1, errors.New("group not a positive integer")
	}

	return Group(g), nil
}

// String converts a Group to a string
func (g Group) String() string {
	return string(g)
}
