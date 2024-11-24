package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/syrax/syrax/lang/unit"
	"github.com/stvmln86/syrax/syrax/lang/word"
)

func TestUnits(t *testing.T) {
	// setup
	comps := []*unit.Unit{unit.New("test", "one")}

	// success
	units := Units("test.one")
	assert.ElementsMatch(t, units, comps)
}

func TestWords(t *testing.T) {
	// setup
	comps := []*word.Word{word.New([]string{"one", "uno"}, unit.New("test", "one"))}

	// success
	words := Words("one,uno:test.one")
	assert.ElementsMatch(t, words, comps)
}
