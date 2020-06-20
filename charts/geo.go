package charts

import (
	"fmt"
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/types"
	"io"
	"log"

	"github.com/go-echarts/go-echarts/datasets"
)

// GeoComponentOpts is the option set for geo component.
type GeoComponentOpts struct {
	Map string `json:"map,omitempty"`
}

// Geo represents a geo chart.
type Geo struct {
	BaseConfiguration
	MultiSeries
	GeoComponentOpts
}

func (Geo) Type() string { return types.ChartGeo }

var geoFormatter = `function (params) {
		return params.name + ' : ' + params.value[2];
}`

// NewGeo creates a new geo chart.
func NewGeo(mapType string) *Geo {
	chart := &Geo{}
	chart.initBaseConfiguration()
	chart.HasGeo = true
	chart.JSAssets.Add("maps/" + datasets.MapFileNames[mapType] + ".js")
	chart.GeoComponentOpts.Map = mapType
	return chart
}

// Add adds new data sets.
// geoType 是 Geo 图形的种类，有以下三种类型可选
// common.ChartType.Scatter
// common.ChartType.EffectScatter
// common.ChartType.HeatMap
func (c *Geo) Add(name, geoType string, data map[string]float32, opts ...SeriesOpts) *Geo {
	nvs := make([]types.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, types.NameValueItem{Name: k, Value: c.extendValue(k, v)})
	}
	series := SingleSeries{Name: name, Type: geoType, Data: nvs, CoordSystem: types.ChartGeo}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

func (c *Geo) extendValue(region string, v float32) []float32 {
	res := make([]float32, 0)
	tv := datasets.Coordinates[region]
	if tv == [2]float32{0, 0} {
		log.Println(fmt.Sprintf("No coordinate is specified for %s", region))
	} else {
		res = append(tv[:], v)
	}
	return res
}

// SetGlobalOptions sets options for the Geo instance.
func (c *Geo) SetGlobalOptions(opts ...GlobalOpts) *Geo {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Geo) Validate() {
	if c.Tooltip.Formatter == "" {
		c.Tooltip.Formatter = opts.FuncOpts(geoFormatter)
	}
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Geo) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.Validate()
	return renderToWriter(c, "chart", []string{}, w...)
}
