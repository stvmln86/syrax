package atom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	atom := New("test", "alpha")
	assert.Equal(t, "test", atom.Form)
	assert.Equal(t, "alpha", atom.Item)
}

func TestParse(t *testing.T) {
	// success
	atom, err := Parse("test.alpha")
	assert.Equal(t, "test", atom.Form)
	assert.Equal(t, "alpha", atom.Item)
	assert.NoError(t, err)

	// error - cannot parse Atom
	atom, err = Parse("nope")
	assert.Nil(t, atom)
	assert.EqualError(t, err, `cannot parse Atom "nope"`)
}

func TestString(t *testing.T) {
	// success
	text := New("test", "alpha").String()
	assert.Equal(t, "test.alpha", text)
}
