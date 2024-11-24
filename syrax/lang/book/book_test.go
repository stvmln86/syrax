package book

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/syrax/syrax/util/test"
)

func TestNew(t *testing.T) {
	// success
	book := New(test.Words("one,uno:test.one"))
	assert.Equal(t, test.Words("one,uno:test.one"), book.Words)
}

func TestParse(t *testing.T) {
	// success
	book, err := Parse([]string{"one,uno:test.one"})
	assert.Equal(t, test.Words("one,uno:test.one"), book.Words)
	assert.NoError(t, err)
}

func TestMatch(t *testing.T) {
	// setup
	book := New(test.Words("one,uno:test.one"))

	// success
	units := book.Match([]string{"one", "uno", "nope"})
	assert.ElementsMatch(t, test.Units("test.one"), units)
}

func TestTranslate(t *testing.T) {
	// setup
	book := New(test.Words("one,uno:test.one"))

	// success
	units := book.Translate("One! Uno! Nope!\n")
	assert.ElementsMatch(t, test.Units("test.one"), units)
}
