package word

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/syrax/syrax/lang/unit"
)

func TestNew(t *testing.T) {
	// success
	word := New([]string{"one", "uno"}, unit.New("test", "one"))
	assert.Equal(t, []string{"one", "uno"}, word.Atoms)
	assert.Equal(t, "test", word.Unit.Form)
	assert.Equal(t, "one", word.Unit.Item)
}

func TestParse(t *testing.T) {
	// success
	word, err := Parse("one,uno:test.one")
	assert.Equal(t, []string{"one", "uno"}, word.Atoms)
	assert.Equal(t, "test", word.Unit.Form)
	assert.Equal(t, "one", word.Unit.Item)
	assert.NoError(t, err)

	// error - cannot parse Word
	word, err = Parse("nope")
	assert.Nil(t, word)
	assert.EqualError(t, err, `cannot parse Word "nope"`)
}

func TestMatch(t *testing.T) {
	// setup
	word := New([]string{"one", "uno"}, unit.New("test", "one"))

	// success - true
	okay := word.Match("one")
	assert.True(t, okay)

	// success - false
	okay = word.Match("nope")
	assert.False(t, okay)
}

func TestString(t *testing.T) {
	// setup
	word := New([]string{"one", "uno"}, unit.New("test", "one"))

	// success
	text := word.String()
	assert.Equal(t, "one,uno:test.one", text)
}
