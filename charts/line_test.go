package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineAssetsBeforeRender(t *testing.T) {
	line := NewLine()
	assert.Equal(t, []string{"echarts.min.js"}, line.JSAssets.Values)
}

func TestLineAssetsAfterRender(t *testing.T) {
	line := NewLine()
	err := line.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, line.JSAssets.Values)
}
