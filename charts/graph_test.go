package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphAssetsBeforeRender(t *testing.T) {
	graph := NewGraph()
	assert.Equal(t, []string{"echarts.min.js"}, graph.JSAssets.Values)
}

func TestGraphAssetsAfterRender(t *testing.T) {
	graph := NewGraph()
	err := graph.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, graph.JSAssets.Values)
}
