package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func geoBase() *charts.Geo {
	geo := charts.NewGeo("china")
	geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-示例图(effectScatter)"})
	geo.Add("geo", charts.ChartType.EffectScatter, mapData,
		charts.RippleEffectOpts{Period: 4, Scale: 6, BrushType: "stroke"})
	return geo
}

func geoShowLabel() *charts.Geo {
	geo := charts.NewGeo("china")
	geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-显示 Label"})

	fn := `function (params) {
		return params.name + ' : ' + params.value[2];
}`
	geo.Add("geo", charts.ChartType.EffectScatter, mapData,
		charts.LabelTextOpts{Show: true, Formatter: charts.FuncOpts(fn), Color: "black", Position: "right"},
		charts.RippleEffectOpts{Period: 4, Scale: 6, BrushType: "stroke"},
	)
	return geo
}

func geoScatter() *charts.Geo {
	geo := charts.NewGeo("china")
	geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-Scatter"})
	geo.Add("geo", charts.ChartType.Scatter, mapData)
	return geo
}

func geoScatterVisualMap() *charts.Geo {
	geo := charts.NewGeo("china")
	geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-设置 VisualMap"})
	geo.Add("geo", charts.ChartType.Scatter, mapData)
	geo.SetGlobalOptions(charts.VisualMapOpts{Max: 60, Calculable: true})
	return geo
}

func geoHeatMap() *charts.Geo {
	geo := charts.NewGeo("china")
	geo.SetGlobalOptions(
		charts.TitleOpts{Title: "Geo-HeatMap"},
		charts.VisualMapOpts{Max: 60, Calculable: true},
	)
	geo.Add("geo", charts.ChartType.HeatMap, mapData)
	return geo
}

func geoEsHeatMap() *charts.Geo {
	geo := charts.NewGeo("china")
	geo.SetGlobalOptions(
		charts.TitleOpts{Title: "Geo-示例图(effectScatter+heatMap)"},
		charts.VisualMapOpts{Max: 60, Calculable: true},
	)
	geo.Add("es", charts.ChartType.EffectScatter, mapData,
		charts.RippleEffectOpts{Period: 4, Scale: 10, BrushType: "stroke"})
	geo.Add("heatmap", charts.ChartType.HeatMap, mapData)
	return geo
}

func geoGuangdong() *charts.Geo {
	geo := charts.NewGeo("广东")
	geo.SetGlobalOptions(
		charts.TitleOpts{Title: "Geo-广东地图"},
		charts.VisualMapOpts{Calculable: true,
			InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
	)
	geo.Add("geo", charts.ChartType.Scatter, guangdongData)
	return geo
}

func geoShantou() *charts.Geo {
	geo := charts.NewGeo("汕头")
	geo.SetGlobalOptions(
		charts.TitleOpts{Title: "Geo-汕头地图"},
		charts.VisualMapOpts{Calculable: true,
			InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
	)
	geo.Add("geo", charts.ChartType.HeatMap, shantouData)
	return geo
}

func geoHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("geo")...)
	page.Add(
		geoBase(),
		geoShowLabel(),
		geoScatter(),
		geoScatterVisualMap(),
		geoHeatMap(),
		geoEsHeatMap(),
		geoGuangdong(),
		geoShantou(),
	)
	f, err := os.Create(getRenderPath("geo.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
