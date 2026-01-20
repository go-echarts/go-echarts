package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThemeRiverAssetsBeforeRender(t *testing.T) {
	themeRiver := NewThemeRiver()
	assert.Equal(t, []string{"echarts.min.js"}, themeRiver.JSAssets.Values)
}

func TestThemeRiverAssetsAfterRender(t *testing.T) {
	themeRiver := NewThemeRiver()
	err := themeRiver.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, themeRiver.JSAssets.Values)
}
