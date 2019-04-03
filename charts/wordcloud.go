package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

// WordCloud represents a word cloud chart.
type WordCloud struct {
	BaseOpts
	Series
}

func (WordCloud) chartType() string { return common.ChartType.WordCloud }

// WordCloudOpts is the option set for a word cloud chart.
type WordCloudOpts struct {
	// 词云图形状，可选
	// "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow"
	Shape string
	// 字体大小范围
	SizeRange []float32
	// 字体选择角度范围
	RotationRange []float32
}

func (WordCloudOpts) markSeries() {}

func (opt *WordCloudOpts) setChartOpt(s *singleSeries) {
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

// NewWordCloud creates a new word cloud chart.
func NewWordCloud(routers ...RouterOpts) *WordCloud {
	chart := new(WordCloud)
	chart.initBaseOpts(routers...)
	chart.JSAssets.Add("echarts-wordcloud.min.js")
	return chart
}

// Add adds new data sets.
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

// SetGlobalOptions sets options for the WordCloud instance.
func (c *WordCloud) SetGlobalOptions(options ...globalOptser) *WordCloud {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *WordCloud) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *WordCloud) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
