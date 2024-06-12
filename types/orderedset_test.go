package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderedSet(t *testing.T) {
	target := &OrderedSet{}
	target.Init("foo", "bar")

	assert.Equal(t, 2, target.Size())

	target.Add("chart")
	assert.Equal(t, 3, target.Size())

	// add same
	target.Add("chart")
	assert.Equal(t, 3, target.Size())

	target.Remove("foo")
	assert.Equal(t, 2, target.Size())

	assert.Equal(t, true, target.Contains("bar"))
	assert.Equal(t, false, target.Contains("foo"))
}
