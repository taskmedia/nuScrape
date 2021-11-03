package group

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test func String
// This test will create a Group from a string and convert it back.
func TestString(t *testing.T) {
	expected := "281103"

	g, _ := New(expected)

	assert.Equal(t, expected, g.String(), "Expect a string as return from method")

}
