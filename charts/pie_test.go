package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPieAssetsBeforeRender(t *testing.T) {
	pie := NewPie()
	assert.Equal(t, []string{"echarts.min.js"}, pie.JSAssets.Values)
}

func TestPieAssetsAfterRender(t *testing.T) {
	pie := NewPie()
	err := pie.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, pie.JSAssets.Values)
}
