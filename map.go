package goecharts

import (
	"io"
)

type Map struct {
	BaseOpts
	Series

	mapType string
}

// 工厂函数，生成 `Map` 实例
func NewMap(mapType string, routers ...HTTPRouter) *Map {
	chart := new(Map)
	chart.mapType = mapType
	chart.HasXYAxis = false
	chart.initBaseOpts(routers...)
	chart.initAssetsOpts()
	chart.JSAssets.Add("maps/" + MapFilenames[mapType] + ".js")
	return chart
}

func (c *Map) Add(name string, data map[string]interface{}, options ...interface{}) *Map {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: "map", MapType: c.mapType, Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Map) SetGlobalConfig(options ...interface{}) *Map {
	c.BaseOpts.setBaseGlobalConfig(options...)
	return c
}

func (c *Map) validateOpts() {
	c.validateInitOpt()
	c.validateAssets(c.AssetsHost)
}

func (c *Map) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", w...)
}
