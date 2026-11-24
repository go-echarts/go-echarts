package charts

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-echarts/go-echarts/v2/opts"
)

func TestTreeMapAssetsBeforeRender(t *testing.T) {
	treeMap := NewTreeMap()
	assert.Equal(t, []string{"echarts.min.js"}, treeMap.JSAssets.Values)
}

func TestTreeMapAssetsAfterRender(t *testing.T) {
	treeMap := NewTreeMap()
	err := treeMap.Render(io.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, treeMap.JSAssets.Values)
}

func TestTreeMapSeriesOptions(t *testing.T) {
	chart := NewTreeMap().AddSeries("sales", []opts.TreeMapNode{{Name: "Books", Value: 3}},
		WithTreeMapOpts(opts.TreeMapChart{
			Label:          &opts.Label{Color: "red"},
			ColorMappingBy: "id",
		}))

	options := string(chart.JSONNotEscaped())
	assert.Contains(t, options, `"label":{"color":"red"}`)
	assert.Contains(t, options, `"colorMappingBy":"id"`)
}

func TestTreeMapOptsPreserveExistingLabel(t *testing.T) {
	chart := NewTreeMap().AddSeries("sales", nil,
		WithLabelOpts(opts.Label{Color: "red"}),
		WithTreeMapOpts(opts.TreeMapChart{LeafDepth: 1}))

	assert.Contains(t, string(chart.JSONNotEscaped()), `"label":{"color":"red"}`)
}
