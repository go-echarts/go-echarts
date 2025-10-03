package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScatterAssetsBeforeRender(t *testing.T) {
	scatter := NewScatter()
	assert.Equal(t, []string{"echarts.min.js"}, scatter.JSAssets.Values)
}

func TestScatterAssetsAfterRender(t *testing.T) {
	scatter := NewScatter()
	err := scatter.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, scatter.JSAssets.Values)
}
