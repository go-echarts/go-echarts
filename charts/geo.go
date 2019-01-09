package charts

import (
	"io"
)

type Geo struct {
	BaseOpts
	Series
}

var geoFormatter = `function (params) {
		return params.name + ' : ' + params.value[2];
}`

// 工厂函数，生成 `Geo` 实例
func NewGeo(mapType string, routers ...HTTPRouter) *Geo {
	chart := new(Geo)
	chart.initBaseOpts(false, routers...)
	chart.JSAssets.Add("maps/" + MapFilenames[mapType] + ".js")
	chart.GeoOpts.Map = mapType
	return chart
}

func (c *Geo) Add(name, geoType string, data map[string]float32, options ...interface{}) *Geo {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, c.extendValue(k, v)})
	}
	series := singleSeries{Name: name, Type: geoType, Data: nvs, CoordSystem: "geo"}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Geo) extendValue(region string, v float32) []float32 {
	tv := Coordinates[region]
	res := append(tv[:], v)
	return res
}

func (c *Geo) SetGlobalConfig(options ...interface{}) *Geo {
	c.BaseOpts.setBaseGlobalConfig(options...)
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
