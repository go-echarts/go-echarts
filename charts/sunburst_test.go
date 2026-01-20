package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSunburstAssetsBeforeRender(t *testing.T) {
	sunburst := NewSunburst()
	assert.Equal(t, []string{"echarts.min.js"}, sunburst.JSAssets.Values)
}

func TestSunburstAssetsAfterRender(t *testing.T) {
	sunburst := NewSunburst()
	err := sunburst.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, sunburst.JSAssets.Values)
}
