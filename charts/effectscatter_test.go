package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEffectScatterAssetsBeforeRender(t *testing.T) {
	effectScatter := NewEffectScatter()
	assert.Equal(t, []string{"echarts.min.js"}, effectScatter.JSAssets.Values)
}

func TestEffectScatterAssetsAfterRender(t *testing.T) {
	effectScatter := NewEffectScatter()
	err := effectScatter.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, effectScatter.JSAssets.Values)
}
