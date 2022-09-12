package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTreeMapAssetsBeforeRender(t *testing.T) {
	treeMap := NewTreeMap()
	assert.Equal(t, []string{"echarts.min.js"}, treeMap.JSAssets.Values)
}

func TestTreeMapAssetsAfterRender(t *testing.T) {
	treeMap := NewTreeMap()
	err := treeMap.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, treeMap.JSAssets.Values)
}
