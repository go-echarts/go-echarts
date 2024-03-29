package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapAssetsBeforeRender(t *testing.T) {
	Map := NewMap()
	assert.Equal(t, []string{"echarts.min.js"}, Map.JSAssets.Values)
}

func TestMapAssetsAfterRender(t *testing.T) {
	Map := NewMap()
	err := Map.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, Map.JSAssets.Values)
}
