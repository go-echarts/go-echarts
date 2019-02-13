package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

type WordCloud struct {
	BaseOpts
	Series
}

func (WordCloud) chartType() string { return common.ChartType.WordCloud }

// WordCLoud series options
type WordCLoudOpts struct {
	// 词云图形状，可选
	// "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow"
	Shape string
	// 字体大小范围
	SizeRange []float32
	// 字体选择角度范围
	RotationRange []float32
}

func (WordCLoudOpts) markSeries() {}

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

func NewWordCloud(routers ...RouterOpts) *WordCloud {
	chart := new(WordCloud)
	chart.initBaseOpts(routers...)
	chart.JSAssets.Add("echarts-wordcloud.min.js")
	return chart
}

func (c *WordCloud) Add(name string, data map[string]interface{}, options ...seriesOptser) *WordCloud {
	nvs := make([]common.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, common.NameValueItem{Name: k, Value: v})
	}
	series := singleSeries{Name: name, Type: common.ChartType.WordCloud, Data: nvs}
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

func (c *WordCloud) SetGlobalOptions(options ...globalOptser) *WordCloud {
	c.BaseOpts.setBaseGlobalOptions(options...)
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
