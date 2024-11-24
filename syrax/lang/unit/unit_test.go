package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// success
	unit := New("test", "one")
	assert.Equal(t, "test", unit.Form)
	assert.Equal(t, "one", unit.Item)
}

func TestParse(t *testing.T) {
	// success
	unit, err := Parse("test.one")
	assert.Equal(t, "test", unit.Form)
	assert.Equal(t, "one", unit.Item)
	assert.NoError(t, err)

	// error - cannot parse Unit
	unit, err = Parse("nope")
	assert.Nil(t, unit)
	assert.EqualError(t, err, `cannot parse Unit "nope"`)
}

func TestString(t *testing.T) {
	// setup
	unit := New("test", "one")

	// success
	text := unit.String()
	assert.Equal(t, "test.one", text)
}
