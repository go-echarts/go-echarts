package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Graph represents a graph chart.
type Graph struct {
	BaseConfiguration
	MultiSeries
}

// Type returns the chart type.
func (Graph) Type() string { return types.ChartGraph }

// NewGraph creates a new graph chart.
func NewGraph() *Graph {
	chart := new(Graph)
	chart.initBaseConfiguration()
	return chart
}

// AddSeries adds the new series.
func (c *Graph) AddSeries(name string, nodes []opts.GraphNode, links []opts.GraphLink, options ...SeriesOpts) *Graph {
	series := SingleSeries{Name: name, Type: types.ChartGraph, Links: links, Data: nodes}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Graph instance.
func (c *Graph) SetGlobalOptions(options ...GlobalOpts) *Graph {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate validates the given configuration.
func (c *Graph) Validate() {
	// If there is no setting of layout, default layout "force".
	for i := 0; i < len(c.MultiSeries); i++ {
		if c.MultiSeries[i].Layout == "" {
			c.MultiSeries[i].Layout = "force"
		}
	}
	c.Assets.Validate(c.AssetsHost)
}
