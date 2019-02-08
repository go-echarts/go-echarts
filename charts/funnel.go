package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

type Funnel struct {
	BaseOpts
	Series
}

func (Funnel) chartType() string { return common.ChartType.Funnel }

func NewFunnel(routers ...RouterOpts) *Funnel {
	chart := new(Funnel)
	chart.initBaseOpts(false, routers...)
	return chart
}

func (c *Funnel) Add(name string, data map[string]interface{}, options ...seriesOptser) *Funnel {
	nvs := make([]common.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, common.NameValueItem{Name: k, Value: v})
	}
	series := singleSeries{Name: name, Type: common.ChartType.Funnel, Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Funnel) SetGlobalOptions(options ...globalOptser) *Funnel {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Funnel) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Funnel) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
