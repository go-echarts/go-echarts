package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/types"
)

// WordCloud represents a word cloud chart.
type WordCloud struct {
	BaseConfiguration
	BaseActions
}

// Type returns the chart type.
func (WordCloud) Type() string { return types.ChartWordCloud }

var wcTextColor = `function () {
	return 'rgb(' + [
		Math.round(Math.random() * 160),
		Math.round(Math.random() * 160),
		Math.round(Math.random() * 160)].join(',') + ')';
}`

// NewWordCloud creates a new word cloud chart.
func NewWordCloud() *WordCloud {
	c := &WordCloud{}
	c.initBaseConfiguration()
	c.Renderer = render.NewChartRender(c, c.Validate)
	c.JSAssets.Add("echarts-wordcloud.min.js")
	return c
}

// AddSeries adds new data sets.
func (c *WordCloud) AddSeries(name string, data []opts.WordCloudData, options ...SeriesOpts) *WordCloud {
	series := SingleSeries{Name: name, Type: types.ChartWordCloud, Data: data}
	series.ConfigureSeriesOpts(options...)

	// set default random color for WordCloud chart
	if series.TextStyle == nil {
		series.TextStyle = &opts.TextStyle{Normal: &opts.TextStyle{}}
	}
	if series.TextStyle.Normal.Color == "" {
		series.TextStyle.Normal.Color = opts.FuncOpts(wcTextColor)
	}

	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the WordCloud instance.
func (c *WordCloud) SetGlobalOptions(options ...GlobalOpts) *WordCloud {
	c.BaseConfiguration.setBaseGlobalOptions(options...)
	return c
}

// SetDispatchActions sets actions for the WordCloud instance.
func (c *WordCloud) SetDispatchActions(actions ...GlobalActions) *WordCloud {
	c.BaseActions.setBaseGlobalActions(actions...)
	return c
}

// Validate validates the given configuration.
func (c *WordCloud) Validate() {
	c.Assets.Validate(c.AssetsHost)
}
