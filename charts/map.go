package charts

import (
	"github.com/go-echarts/go-echarts/datasets"
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Map represents a map chart.
type Map struct {
	BaseConfiguration
	MultiSeries

	mapType string
}

// Type returns the chart type.
func (Map) Type() string { return types.ChartMap }

// NewMap creates a new map chart.
func NewMap() *Map {
	chart := &Map{}
	chart.initBaseConfiguration()
	chart.Renderer = render.NewChartRender(chart, chart.Validate)
	return chart
}

func (c *Map) RegisterMapType(mapType string) {
	c.mapType = mapType
	c.JSAssets.Add("maps/" + datasets.MapFileNames[mapType] + ".js")
}

// Add adds new data sets.
func (c *Map) AddSeries(name string, data []opts.MapData, opts ...SeriesOpts) *Map {
	series := SingleSeries{Name: name, Type: types.ChartMap, MapType: c.mapType, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Map instance.
func (c *Map) SetGlobalOptions(opts ...GlobalOpts) *Map {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

// Validate
func (c *Map) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writer.
//func (c *Map) Render(w io.Writer) error {
//	c.Validate()
//	return render.ChartRender(c, w)
//}
