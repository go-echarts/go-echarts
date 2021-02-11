package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Tree represents a Tree chart.
type Tree struct {
	RectChart
}

// Type returns the chart type.
func (Tree) Type() string { return types.Tree }

// NewTree creates a new Tree chart instance.
func NewTree() *Tree {
	c := &Tree{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

func (c *Tree) AddSeries(name string, data []opts.TreeData, options ...SeriesOpts) *Tree {
	series := SingleSeries{Name: name, Type: types.Tree, Data: data}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// Validate validates the given configuration.
func (c *Tree) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
