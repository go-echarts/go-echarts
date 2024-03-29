package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRadarAssetsBeforeRender(t *testing.T) {
	radar := NewPie()
	assert.Equal(t, []string{"echarts.min.js"}, radar.JSAssets.Values)
}

func TestRadarAssetsAfterRender(t *testing.T) {
	radar := NewRadar()
	err := radar.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, radar.JSAssets.Values)
}
