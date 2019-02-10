package charts

import (
	"fmt"
	"io"
	"log"

	"github.com/chenjiandongx/go-echarts/common"
	"github.com/chenjiandongx/go-echarts/datasets"
)

// 地理坐标系组件配置项
type GeoComponentOpts struct {
	Map string `json:"map,omitempty"`
}

type Geo struct {
	BaseOpts
	Series
	GeoComponentOpts
}

func (Geo) chartType() string { return common.ChartType.Geo }

var geoFormatter = `function (params) {
		return params.name + ' : ' + params.value[2];
}`

func NewGeo(mapType string, routers ...RouterOpts) *Geo {
	chart := new(Geo)
	chart.initBaseOpts(false, routers...)
	chart.HasGeo = true
	chart.JSAssets.Add("maps/" + datasets.MapFileNames[mapType] + ".js")
	chart.GeoComponentOpts.Map = mapType
	return chart
}

// geoType 是 Geo 图形的种类，有以下三种类型可选
// common.ChartType.Scatter
// common.ChartType.EffectScatter
// common.ChartType.HeatMap
func (c *Geo) Add(name, geoType string, data map[string]float32, options ...seriesOptser) *Geo {
	nvs := make([]common.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, common.NameValueItem{Name: k, Value: c.extendValue(k, v)})
	}
	series := singleSeries{Name: name, Type: geoType, Data: nvs, CoordSystem: common.ChartType.Geo}
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

func (c *Geo) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
