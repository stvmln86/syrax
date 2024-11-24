// Package test implements unit-testing helpers and values.
package test

import (
	"github.com/stvmln86/syrax/syrax/lang/unit"
	"github.com/stvmln86/syrax/syrax/lang/word"
)

// Units returns a Unit slice from parsed strings.
func Units(texts ...string) []*unit.Unit {
	var units []*unit.Unit
	for _, text := range texts {
		unit, _ := unit.Parse(text)
		units = append(units, unit)
	}

	return units
}

// Words returns a Word slice from parsed strings.
func Words(texts ...string) []*word.Word {
	var words []*word.Word
	for _, text := range texts {
		word, _ := word.Parse(text)
		words = append(words, word)
	}

	return words
}
