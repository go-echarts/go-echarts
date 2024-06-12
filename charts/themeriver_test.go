package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThemeRiverAssetsBeforeRender(t *testing.T) {
	themeRiver := NewThemeRiver()
	assert.Equal(t, []string{"echarts.min.js"}, themeRiver.JSAssets.Values)
}

func TestThemeRiverAssetsAfterRender(t *testing.T) {
	themeRiver := NewThemeRiver()
	err := themeRiver.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, themeRiver.JSAssets.Values)
}
