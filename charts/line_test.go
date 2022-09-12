package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineAssetsBeforeRender(t *testing.T) {
	line := NewLine()
	assert.Equal(t, []string{"echarts.min.js"}, line.JSAssets.Values)
}

func TestLineAssetsAfterRender(t *testing.T) {
	line := NewLine()
	err := line.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, line.JSAssets.Values)
}
