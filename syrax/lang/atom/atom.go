// Package atom implements the Atom type and methods.
package atom

import (
	"fmt"
	"regexp"
)

// Atom is a stored language definition.
type Atom struct {
	Form string
	Item string
}

// Regex is a regular expression for matching Atoms.
var Regex = regexp.MustCompile(`(\w+)\.(\w+)`)

// New returns a new Atom.
func New(base, item string) *Atom {
	return &Atom{base, item}
}

// Parse returns a new Atom from a string.
func Parse(atom string) (*Atom, error) {
	elems := Regex.FindStringSubmatch(atom)
	if len(elems) < 3 {
		return nil, fmt.Errorf("cannot parse Atom %q", atom)
	}

	return New(elems[1], elems[2]), nil
}

// String returns the Atom as a string.
func (a *Atom) String() string {
	return fmt.Sprintf("%s.%s", a.Form, a.Item)
}
