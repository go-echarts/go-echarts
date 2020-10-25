package charts

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const host = "https://go-echarts.github.io/go-echarts-assets/assets/"

func TestBarAssetsBeforeRender(t *testing.T) {
	bar := NewBar()
	assert.Equal(t, bar.JSAssets.Values, []string{"echarts.min.js"})
}

func TestBarAssetsAfterRender(t *testing.T) {
	bar := NewBar()
	err := bar.Render(os.Stdout)
	assert.NoError(t, err)
	assert.Equal(t, bar.JSAssets.Values, []string{host + "echarts.min.js"})
}

func TestBarDefaultValue(t *testing.T) {
	bar := NewBar()
	err := bar.Render(os.Stdout)
	assert.NoError(t, err)
	assert.Equal(t, bar.Initialization.Width, "900px")
	assert.Equal(t, bar.Initialization.Height, "500px")
	assert.Equal(t, bar.PageTitle, "Awesome go-echarts")
	assert.Equal(t, bar.AssetsHost, host)
}
