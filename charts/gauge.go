package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

type Gauge struct {
	BaseOpts
	Series
}

func (Gauge) chartType() string { return common.GaugeType }

func NewGauge(routers ...HTTPRouter) *Gauge {
	chart := new(Gauge)
	chart.initBaseOpts(false, routers...)
	return chart
}

func (c *Gauge) Add(name string, data map[string]interface{}, options ...seriesOptser) *Gauge {
	nvs := make([]common.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, common.NameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: "gauge", Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	return c
}

func (c *Gauge) SetGlobalOptions(options ...globalOptser) *Gauge {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Gauge) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Gauge) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
