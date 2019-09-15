package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func mapBase() *charts.Map {
	mc := charts.NewMap("china")
	mc.SetGlobalOptions(charts.TitleOpts{Title: "Map-示例图"})
	mc.Add("map", mapData)
	return mc
}

func mapShowLabel() *charts.Map {
	mc := charts.NewMap("china")
	mc.SetGlobalOptions(charts.TitleOpts{Title: "Map-展示 Label"})
	mc.Add("map", mapData, charts.LabelTextOpts{Show: true})
	return mc
}

func mapVisualMap() *charts.Map {
	mc := charts.NewMap("china")
	mc.SetGlobalOptions(
		charts.TitleOpts{Title: "Map-设置 VisualMap"},
		charts.VisualMapOpts{Calculable: true},
	)
	mc.Add("map", mapData)
	return mc
}

func mapGuangdong() *charts.Map {
	mc := charts.NewMap("广东")
	mc.SetGlobalOptions(
		charts.TitleOpts{Title: "Map-广东地图"},
		charts.VisualMapOpts{Calculable: true,
			InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
	)
	mc.Add("map", guangdongData)
	return mc
}

func mapShantou() *charts.Map {
	mc := charts.NewMap("汕头")
	mc.SetGlobalOptions(
		charts.TitleOpts{Title: "Map-汕头地图"},
		charts.VisualMapOpts{Calculable: true,
			InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
	)
	mc.Add("map", shantouData)
	return mc
}

func mapTheme() *charts.Map {
	mc := charts.NewMap("china")
	mc.SetGlobalOptions(
		charts.InitOpts{Theme: charts.ThemeType.Macarons},
		charts.TitleOpts{Title: "Map-设置风格"},
		charts.VisualMapOpts{Calculable: true, Max: 150},
	)
	mc.Add("map", mapData)
	return mc
}

func mapHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("map")...)
	page.Add(
		mapBase(),
		mapShowLabel(),
		mapVisualMap(),
		mapGuangdong(),
		mapShantou(),
		mapTheme(),
	)
	f, err := os.Create(getRenderPath("map.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
