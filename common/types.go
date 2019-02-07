package common

type chartType struct {
	Bar,
	Bar3D,
	BoxPlot,
	Cartesian3D,
	EffectScatter,
	Funnel,
	Gauge,
	Geo,
	HeatMap,
	Kline,
	Line,
	Liquid,
	Map,
	Pie,
	Scatter,
	WordCloud string
}

var ChartType = chartType{
	Bar:           "bar",
	Bar3D:         "bar3D",
	BoxPlot:       "boxplot",
	Cartesian3D:   "cartesian3D",
	EffectScatter: "effectScatter",
	Funnel:        "funnel",
	Gauge:         "gauge",
	Geo:           "geo",
	HeatMap:       "heatmap",
	Kline:         "candlestick",
	Line:          "line",
	Liquid:        "liquidFill",
	Map:           "map",
	Pie:           "pie",
	Scatter:       "scatter",
	WordCloud:     "wordCloud",
}

type themeType struct {
	Chalk,
	Essos,
	Halloween,
	Infographic,
	Macarons,
	PurplePassion,
	Roma,
	Romantic,
	Shine,
	Vintage,
	Walden string
}

var ThemeType = themeType{
	Chalk:         "chalk",
	Essos:         "essos",
	Halloween:     "halloween",
	Infographic:   "infographic",
	Macarons:      "macarons",
	PurplePassion: "purple-passion",
	Roma:          "roma",
	Romantic:      "romantic",
	Shine:         "shine",
	Vintage:       "vintage",
	Walden:        "walden",
}
