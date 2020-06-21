package charts

//
//// WordCloud represents a word cloud chart.
//type WordCloud struct {
//	BaseConfiguration
//	MultiSeries
//}
//
//func (WordCloud) Type() string { return types.ChartWordCloud }
//
//var wcTextColor = `function () {
//	return 'rgb(' + [
//		Math.round(Math.random() * 160),
//		Math.round(Math.random() * 160),
//		Math.round(Math.random() * 160)].join(',') + ')';
//}`
//
//// NewWordCloud creates a new word cloud chart.
//func NewWordCloud() *WordCloud {
//	chart := &WordCloud{}
//	chart.initBaseConfiguration()
//	chart.JSAssets.Add("echarts-wordcloud.min.js")
//	return chart
//}
//
//// Add adds new data sets.
//func (c *WordCloud) Add(name string, data map[string]interface{}, opts ...SeriesOpts) *WordCloud {
//	nvs := make([]types.NameValueItem, 0)
//	for k, v := range data {
//		nvs = append(nvs, types.NameValueItem{Name: k, Value: v})
//	}
//	series := SingleSeries{Name: name, Type: types.ChartWordCloud, Data: nvs}
//	series.configureSeriesOpts(opts...)
//
//	// set default random color for WordCloud chart
//	if series.TextStyle.Normal == nil {
//		series.TextStyle.Normal = &opts.TextStyle{Color: FuncOpts(wcTextColor)}
//	} else {
//		if series.TextStyleOpts.Normal.Color == "" {
//			series.TextStyleOpts.Normal.Color = FuncOpts(wcTextColor)
//		}
//	}
//
//	c.MultiSeries = append(c.MultiSeries, series)
//	return c
//}
//
//// SetGlobalOptions sets options for the WordCloud instance.
//func (c *WordCloud) SetGlobalOptions(opts ...GlobalOpts) *WordCloud {
//	c.BaseConfiguration.setBaseGlobalOptions(opts...)
//	return c
//}
//
//func (c *WordCloud) Validate() {
//	c.Assets.Validate(c.AssetsHost)
//}
//
//// Render renders the chart and writes the output to given writers.
//func (c *WordCloud) Render(w ...io.Writer) error {
//	c.insertSeriesColors(c.appendColor)
//	c.Validate()
//	return renderToWriter(c, "chart", []string{}, w...)
//}
