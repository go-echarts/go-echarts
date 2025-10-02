package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunnelAssetsBeforeRender(t *testing.T) {
	funnel := NewFunnel()
	assert.Equal(t, []string{"echarts.min.js"}, funnel.JSAssets.Values)
}

func TestFunnelAssetsAfterRender(t *testing.T) {
	funnel := NewFunnel()
	err := funnel.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, funnel.JSAssets.Values)
}
