package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeMapAssetsBeforeRender(t *testing.T) {
	treeMap := NewTreeMap()
	assert.Equal(t, []string{"echarts.min.js"}, treeMap.JSAssets.Values)
}

func TestTreeMapAssetsAfterRender(t *testing.T) {
	treeMap := NewTreeMap()
	err := treeMap.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, treeMap.JSAssets.Values)
}
