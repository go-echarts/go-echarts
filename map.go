package goecharts

import (
	"io"
)

type Map struct {
	BaseOpts
	Series MapSeries

	mapType string
}

// 工厂函数，生成 `Map` 实例
func NewMap(mapType string, routers ...HttpRouter) *Map {
	mapChart := new(Map)
	mapChart.mapType = mapType
	mapChart.HasXYAxis = false
	mapChart.initBaseOpts(routers...)
	mapChart.initAssetsOpts()
	mapChart.JSAssets = append(mapChart.JSAssets, "maps/"+MapFilenames[mapType]+".js")
	return mapChart
}

type singleMapSeries struct {
	singleSeries
	MapType string `json:"mapType"`
	//Roam    bool   `json:"roam"`
}

type MapSeries []singleMapSeries

func (c *Map) Add(name string, data map[string]interface{}, options ...interface{}) *Map {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleMapSeries{
		singleSeries: singleSeries{Name: name, Type: "map", Data: nvs},
		MapType:      c.mapType,
	}
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
	if err := renderToWriter(c, "chart", w...); err != nil {
		return err
	}
	return nil
}
