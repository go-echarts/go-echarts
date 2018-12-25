package goecharts

import (
	"io"
)

type Map struct {
	BaseOpts
	Series MapSeries

	HasXYAxis bool
	mapType   string
}

// 工厂函数，生成 `Map` 实例
func NewMap(mapType string, routers ...HttpRouter) *Map {
	m := new(Map)
	m.mapType = mapType
	m.HasXYAxis = false
	m.initBaseOpts(routers...)
	m.initAssetsOpts()
	m.JSAssets = append(m.JSAssets, "maps/"+mapType+".js")
	return m
}

type singleMapSeries struct {
	singleSeries
	MapType string `json:"mapType"`
	//Roam    bool   `json:"roam"`
}

type MapSeries []singleMapSeries

func (m *Map) Add(name string, data map[string]interface{}, options ...interface{}) *Map {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleMapSeries{
		singleSeries: singleSeries{Name: name, Type: mapType, Data: nvs},
		MapType:      m.mapType}
	series.setSingleSeriesOpts(options...)
	m.Series = append(m.Series, series)
	m.setColor(options...)
	return m
}

func (m *Map) SetGlobalConfig(options ...interface{}) *Map {
	m.BaseOpts.setBaseGlobalConfig(options...)
	return m
}

func (m *Map) validateOpts() {
	m.validateInitOpt()
	m.validateAssets(m.AssetsHost)
}

// 渲染图表，支持多 io.Writer
func (m *Map) Render(w ...io.Writer) {
	m.insertSeriesColors(m.appendColor)
	m.validateOpts()
	renderToWriter(m, "chart", w...)
}
