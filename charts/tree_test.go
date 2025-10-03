package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeAssetsBeforeRender(t *testing.T) {
	tree := NewTree()
	assert.Equal(t, []string{"echarts.min.js"}, tree.JSAssets.Values)
}

func TestTreeAssetsAfterRender(t *testing.T) {
	tree := NewTree()
	err := tree.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, tree.JSAssets.Values)
}
