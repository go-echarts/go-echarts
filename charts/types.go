package charts

// ChartType contains string representations of chart types.
// TODO: should use enum.
var ChartType = struct {
	Bar, Bar3D, BoxPlot,
	Cartesian3D,
	EffectScatter,
	Funnel,
	Gauge, Geo, Graph,
	HeatMap,
	Kline,
	Line, Line3D, Liquid,
	Map,
	Parallel, Pie,
	Radar,
	Sankey, Scatter, Scatter3D, Surface3D,
	ThemeRiver,
	WordCloud string
}{
	Bar:           "bar",
	Bar3D:         "bar3D",
	BoxPlot:       "boxplot",
	Cartesian3D:   "cartesian3D",
	EffectScatter: "effectScatter",
	Funnel:        "funnel",
	Gauge:         "gauge",
	Geo:           "geo",
	Graph:         "graph",
	HeatMap:       "heatmap",
	Kline:         "candlestick",
	Line:          "line",
	Line3D:        "line3D",
	Liquid:        "liquidFill",
	Map:           "map",
	Parallel:      "parallel",
	Pie:           "pie",
	Radar:         "radar",
	Sankey:        "sankey",
	Scatter:       "scatter",
	Scatter3D:     "scatter3D",
	Surface3D:     "surface",
	ThemeRiver:    "themeRiver",
	WordCloud:     "wordCloud",
}

// ThemeType contains string representations of them types.
// TODO: should use enum.
var ThemeType = struct {
	Chalk,
	Essos,
	Infographic,
	Macarons,
	PurplePassion, Roma,
	Romantic,
	Shine,
	Vintage,
	Walden, Westeros, Wonderland string
}{
	Chalk:         "chalk",
	Essos:         "essos",
	Infographic:   "infographic",
	Macarons:      "macarons",
	PurplePassion: "purple-passion",
	Roma:          "roma",
	Romantic:      "romantic",
	Shine:         "shine",
	Vintage:       "vintage",
	Walden:        "walden",
	Westeros:      "westeros",
	Wonderland:    "wonderland",
}
