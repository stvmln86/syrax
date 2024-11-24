// Package unit implements the Unit type and methods.
package unit

import (
	"fmt"
	"regexp"
)

// Unit is a stored language definition.
type Unit struct {
	Form string
	Item string
}

// Regex is a regular expression for matching Units.
var Regex = regexp.MustCompile(`(\w+)\.(\w+)`)

// New returns a new Unit.
func New(base, item string) *Unit {
	return &Unit{base, item}
}

// Parse returns a new Unit from a string.
func Parse(unit string) (*Unit, error) {
	elems := Regex.FindStringSubmatch(unit)
	if len(elems) < 3 {
		return nil, fmt.Errorf("cannot parse Unit %q", unit)
	}

	return New(elems[1], elems[2]), nil
}

// String returns the Unit as a string.
func (u *Unit) String() string {
	return fmt.Sprintf("%s.%s", u.Form, u.Item)
}
