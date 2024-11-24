// Package book implements the Book type and methods.
package book

import (
	"strings"
	"unicode"

	"github.com/stvmln86/syrax/syrax/lang/unit"
	"github.com/stvmln86/syrax/syrax/lang/word"
)

// Book is a language container and translator.
type Book struct {
	Words []*word.Word
}

// New returns a new Book.
func New(words []*word.Word) *Book {
	return &Book{words}
}

// Parse returns a new Book from a string slice.
func Parse(texts []string) (*Book, error) {
	var words []*word.Word
	for _, text := range texts {
		word, err := word.Parse(text)
		if err != nil {
			return nil, err
		}

		words = append(words, word)
	}

	return New(words), nil
}

// Match returns a Unit slice from an atom slice.
func (b *Book) Match(atoms []string) []*unit.Unit {
	var umap = make(map[*unit.Unit]bool)
	for _, atom := range atoms {
		for _, word := range b.Words {
			if word.Match(atom) {
				umap[word.Unit] = true
			}
		}
	}

	var units []*unit.Unit
	for unit := range umap {
		units = append(units, unit)
	}

	return units
}

// Translate returns a Unit slice from a string.
func (b *Book) Translate(text string) []*unit.Unit {
	var chars []rune
	for _, char := range strings.ToLower(text) {
		if !unicode.IsPunct(char) {
			chars = append(chars, char)
		}
	}

	atoms := strings.Fields(string(chars))
	return b.Match(atoms)
}
