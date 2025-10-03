package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGaugeAssetsBeforeRender(t *testing.T) {
	gauge := NewGauge()
	assert.Equal(t, []string{"echarts.min.js"}, gauge.JSAssets.Values)
}

func TestGaugeAssetsAfterRender(t *testing.T) {
	gauge := NewGauge()
	err := gauge.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, gauge.JSAssets.Values)
}
