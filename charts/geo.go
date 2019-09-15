package charts

import (
	"fmt"
	"github.com/go-echarts/go-echarts/datatypes"
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
	BaseOpts
	Series
	GeoComponentOpts
}

func (Geo) chartType() string { return ChartType.Geo }

var geoFormatter = `function (params) {
		return params.name + ' : ' + params.value[2];
}`

// NewGeo creates a new geo chart.
func NewGeo(mapType string, routers ...RouterOpts) *Geo {
	chart := new(Geo)
	chart.initBaseOpts(routers...)
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
func (c *Geo) Add(name, geoType string, data map[string]float32, options ...seriesOptser) *Geo {
	nvs := make([]datatypes.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, datatypes.NameValueItem{Name: k, Value: c.extendValue(k, v)})
	}
	series := singleSeries{Name: name, Type: geoType, Data: nvs, CoordSystem: ChartType.Geo}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
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
func (c *Geo) SetGlobalOptions(options ...globalOptser) *Geo {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Geo) validateOpts() {
	if c.TooltipOpts.Formatter == "" {
		c.TooltipOpts.Formatter = FuncOpts(geoFormatter)
	}
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Geo) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
