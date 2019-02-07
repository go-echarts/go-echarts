package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
	"github.com/chenjiandongx/go-echarts/datasets"
)

type Geo struct {
	BaseOpts
	Series
}

func (Geo) chartType() string { return common.ChartType.Geo }

var geoFormatter = `function (params) {
		return params.name + ' : ' + params.value[2];
}`

func NewGeo(mapType string, routers ...HTTPRouter) *Geo {
	chart := new(Geo)
	chart.initBaseOpts(false, routers...)
	chart.JSAssets.Add("maps/" + datasets.MapFileNames[mapType] + ".js")
	chart.GeoOpts.Map = mapType
	return chart
}

func (c *Geo) Add(name, geoType string, data map[string]float32, options ...seriesOptser) *Geo {
	nvs := make([]common.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, common.NameValueItem{k, c.extendValue(k, v)})
	}
	series := singleSeries{Name: name, Type: geoType, Data: nvs, CoordSystem: common.ChartType.Geo}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Geo) extendValue(region string, v float32) []float32 {
	tv := datasets.Coordinates[region]
	res := append(tv[:], v)
	return res
}

func (c *Geo) SetGlobalOptions(options ...globalOptser) *Geo {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Geo) validateGeoFormatter() {
	if c.TooltipOpts.Formatter == "" {
		c.TooltipOpts.Formatter = FuncOpts(geoFormatter)
	}
}

func (c *Geo) validateOpts() {
	c.validateAssets(c.AssetsHost)
	c.validateGeoFormatter()
}

func (c *Geo) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
