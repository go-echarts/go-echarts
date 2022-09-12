package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelAssetsBeforeRender(t *testing.T) {
	parallel := NewParallel()
	assert.Equal(t, []string{"echarts.min.js"}, parallel.JSAssets.Values)
}

func TestParallelAssetsAfterRender(t *testing.T) {
	parallel := NewParallel()
	err := parallel.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, parallel.JSAssets.Values)
}
