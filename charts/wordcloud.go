package charts

import (
	"io"
)

type WordCloud struct {
	BaseOpts
	Series
}

// WordCLoud series options
type WordCLoudOpts struct {
	Shape         string
	SizeRange     []float32
	RotationRange []float32
}

func (opt *WordCLoudOpts) setChartOpt(s *singleSeries) {
	s.Shape = opt.Shape
	s.SizeRange = opt.SizeRange
	s.RotationRange = opt.RotationRange
}

var wcTextColor = `function () {
	return 'rgb(' + [
		Math.round(Math.random() * 160),
		Math.round(Math.random() * 160),
		Math.round(Math.random() * 160)].join(',') + ')';
}`

// 工厂函数，生成 `WordCloud` 实例
func NewWordCloud(routers ...HTTPRouter) *WordCloud {
	chart := new(WordCloud)
	chart.initBaseOpts(false, routers...)
	chart.JSAssets.Add("echarts-wordcloud.min.js")
	return chart
}

func (c *WordCloud) Add(name string, data map[string]interface{}, options ...interface{}) *WordCloud {
	nvs := make([]nameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, nameValueItem{k, v})
	}
	series := singleSeries{Name: name, Type: "wordCloud", Data: nvs}
	series.setSingleSeriesOpts(options...)

	// 处理词云图默认随机颜色
	if series.TextStyleOpts.Normal == nil {
		series.TextStyleOpts.Normal = &TextStyleOpts{Color: FuncOpts(wcTextColor)}
	} else {
		if series.TextStyleOpts.Normal.Color == "" {
			series.TextStyleOpts.Normal.Color = FuncOpts(wcTextColor)
		}
	}

	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *WordCloud) SetGlobalConfig(options ...interface{}) *WordCloud {
	c.BaseOpts.setBaseGlobalConfig(options...)
	return c
}

func (c *WordCloud) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *WordCloud) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
