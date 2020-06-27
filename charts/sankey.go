package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Sankey represents a sankey chart.
type Sankey struct {
	BaseConfiguration
	MultiSeries
}

// Type returns the chart type.
func (Sankey) Type() string { return types.ChartSankey }

// NewSankey creates a new sankey chart.
func NewSankey() *Sankey {
	chart := &Sankey{}
	chart.initBaseConfiguration()
	return chart
}

// Add adds new data sets.
func (c *Sankey) AddSeries(name string, nodes []opts.SankeyNode, links []opts.SankeyLink, opts ...SeriesOpts) *Sankey {
	series := SingleSeries{Name: name, Type: types.ChartSankey, Data: nodes, Links: links}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Sankey instance.
func (c *Sankey) SetGlobalOptions(opts ...GlobalOpts) *Sankey {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Sankey) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Sankey) Render(w io.Writer) error {
	c.Validate()
	return render.ChartRender(c, w)
}
