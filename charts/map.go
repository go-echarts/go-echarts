package charts

import (
	"github.com/go-echarts/go-echarts/v2/datasets"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// Map represents a map chart.
type Map struct {
	BaseConfiguration

	mapType string
}

// Type returns the chart type.
func (Map) Type() string { return types.ChartMap }

// NewMap creates a new map chart.
func NewMap() *Map {
	c := &Map{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	return c
}

// RegisterMapType
func (c *Map) RegisterMapType(mapType string) {
	c.mapType = mapType
	c.JSAssets.Add("maps/" + datasets.MapFileNames[mapType] + ".js")
}

// AddSeries adds new data sets.
func (c *Map) AddSeries(name string, data []opts.MapData, options ...SeriesOpts) *Map {
	series := SingleSeries{Name: name, Type: types.ChartMap, MapType: c.mapType, Data: data}
	series.configureSeriesOpts(options...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Map instance.
func (c *Map) SetGlobalOptions(options ...GlobalOpts) *Map {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// Validate
func (c *Map) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
