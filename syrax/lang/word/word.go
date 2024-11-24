// Package word implements the Word type and methods.
package word

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/stvmln86/syrax/syrax/lang/unit"
)

// Word is an atom set with a Unit definition.
type Word struct {
	Atoms []string
	Unit  *unit.Unit
}

// Regex is a regular expression for matching Words.
var Regex = regexp.MustCompile(`([\w\,]+)\:(\w+)\.(\w+)`)

// New returns a new Word.
func New(atoms []string, unit *unit.Unit) *Word {
	return &Word{atoms, unit}
}

// Parse returns a new Word from a string.
func Parse(text string) (*Word, error) {
	elems := Regex.FindStringSubmatch(text)
	if len(elems) < 4 {
		return nil, fmt.Errorf("cannot parse Word %q", text)
	}

	atoms := strings.Split(elems[1], ",")
	return New(atoms, unit.New(elems[2], elems[3])), nil
}

// Match returns true if the Word contains an atom.
func (w *Word) Match(atom string) bool {
	return slices.Contains(w.Atoms, atom)
}

// String returns the Word as a string.
func (w *Word) String() string {
	atoms := strings.Join(w.Atoms, ",")
	return fmt.Sprintf("%s:%s.%s", atoms, w.Unit.Form, w.Unit.Item)
}
