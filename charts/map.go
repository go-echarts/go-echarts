package charts

import (
	"github.com/go-echarts/go-echarts/datatypes"
	"io"

	"github.com/go-echarts/go-echarts/datasets"
)

// Map represents a map chart.
type Map struct {
	BaseOpts
	MultiSeries

	mapType string
}

func (Map) chartType() string { return ChartType.Map }

// NewMap creates a new map chart.
func NewMap(mapType string, routers ...RouterOpts) *Map {
	chart := new(Map)
	chart.mapType = mapType
	chart.initBaseOpts(routers...)
	chart.JSAssets.Add("maps/" + datasets.MapFileNames[mapType] + ".js")
	return chart
}

// Add adds new data sets.
func (c *Map) Add(name string, data map[string]float32, fns ...SeriesOptser) *Map {
	nvs := make([]datatypes.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, datatypes.NameValueItem{Name: k, Value: v})
	}
	series := SingleSeries{Name: name, Type: ChartType.Map, MapType: c.mapType, Data: nvs}
	series.configureSeriesFns(fns...)
	c.MultiSeries = append(c.MultiSeries, series)
	c.setColor(options...)
	return c
}

// SetGlobalOptions sets options for the Map instance.
func (c *Map) SetGlobalOptions(options ...GlobalOptser) *Map {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Map) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Map) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
