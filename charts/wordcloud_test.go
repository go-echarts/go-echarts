package charts

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCloudAssetsBeforeRender(t *testing.T) {
	bar := NewBar()
	assert.Equal(t, bar.JSAssets.Values, []string{"echarts.min.js"})
}

func TestWordCloudAssetsAfterRender(t *testing.T) {
	bar := NewBar()
	err := bar.Render(os.Stdout)
	assert.NoError(t, err)
	assert.Equal(t, bar.JSAssets.Values, []string{host + "echarts.min.js"})
}
