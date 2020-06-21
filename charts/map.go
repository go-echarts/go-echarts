package charts

import (
	"github.com/go-echarts/go-echarts/types"
	"io"

	"github.com/go-echarts/go-echarts/datasets"
)

// Map represents a map chart.
type Map struct {
	BaseConfiguration
	MultiSeries

	mapType string
}

func (Map) Type() string { return types.ChartMap }

// NewMap creates a new map chart.
func NewMap(mapType string) *Map {
	chart := &Map{}
	chart.mapType = mapType
	chart.initBaseConfiguration()
	chart.JSAssets.Add("maps/" + datasets.MapFileNames[mapType] + ".js")
	return chart
}

// Add adds new data sets.
func (c *Map) Add(name string, data map[string]float32, opts ...SeriesOpts) *Map {
	nvs := make([]types.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, types.NameValueItem{Name: k, Value: v})
	}
	series := SingleSeries{Name: name, Type: types.ChartMap, MapType: c.mapType, Data: nvs}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Map instance.
func (c *Map) SetGlobalOptions(opts ...GlobalOpts) *Map {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Map) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Map) Render(w io.Writer) error {
	c.Validate()
	return renderToWriter(c, "chart", []string{}, w)
}
