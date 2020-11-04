package charts

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCloudAssetsBeforeRender(t *testing.T) {
	wc := NewWordCloud()
	assert.Equal(t, []string{"echarts.min.js", "echarts-wordcloud.min.js"}, wc.JSAssets.Values)
}

func TestWordCloudAssetsAfterRender(t *testing.T) {
	wc := NewWordCloud()
	err := wc.Render(os.Stdout)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js", host + "echarts-wordcloud.min.js"}, wc.JSAssets.Values)
}
