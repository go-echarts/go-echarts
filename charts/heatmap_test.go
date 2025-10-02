package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeatMapAssetsBeforeRender(t *testing.T) {
	heatMap := NewHeatMap()
	assert.Equal(t, []string{"echarts.min.js"}, heatMap.JSAssets.Values)
}

func TestHeatMapAssetsAfterRender(t *testing.T) {
	heatMap := NewHeatMap()
	err := heatMap.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, heatMap.JSAssets.Values)
}
