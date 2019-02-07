package common

type chartType struct {
	BarType,
	Bar3DType,
	BoxPlotType,
	EffectScatterType,
	FunnelType,
	GaugeType,
	GeoType,
	HeatMapType,
	KlineType,
	LineType,
	LiquidType,
	MapType,
	PieType,
	ScatterType,
	WordCloudType string
}

var ChartType = chartType{
	BarType:           "bar",
	Bar3DType:         "bar3D",
	BoxPlotType:       "boxPlot",
	EffectScatterType: "effectScatter",
	FunnelType:        "funnel",
	GaugeType:         "gauge",
	GeoType:           "geo",
	HeatMapType:       "heatMap",
	KlineType:         "kline",
	LineType:          "line",
	LiquidType:        "liquid",
	MapType:           "map",
	PieType:           "pie",
	ScatterType:       "scatter",
	WordCloudType:     "wordCloud",
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
