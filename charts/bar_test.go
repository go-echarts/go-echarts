package charts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBarAssetsBeforeRender(t *testing.T) {
	bar := NewBar()
	assert.Equal(t, bar.JSAssets.Values, []string{"echarts.min.js"})
	assert.Equal(t, bar.CSSAssets.Values, []string{"bulma.min.css"})
}

func TestBarAssetsAfterRender(t *testing.T) {
	bar := NewBar()
	err := bar.Render()
	assert.NoError(t, err)
	var host = "http://chenjiandongx.com/go-echarts-assets/assets/"
	assert.Equal(t, bar.JSAssets.Values, []string{host + "echarts.min.js"})
	assert.Equal(t, bar.CSSAssets.Values, []string{host + "bulma.min.css"})
}

func TestBarDefaultValue(t *testing.T) {
	bar := NewBar()
	err := bar.Render()
	assert.NoError(t, err)
	assert.Equal(t, bar.Width, "900px")
	assert.Equal(t, bar.Height, "500px")
	assert.Equal(t, bar.PageTitle, "Awesome go-echarts")
	assert.Equal(t, bar.AssetsHost, "http://chenjiandongx.com/go-echarts-assets/assets/")
}
