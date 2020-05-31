package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

var wcData = map[string]interface{}{
	"Sam S Club":               10000,
	"Macys":                    6181,
	"Amy Schumer":              4386,
	"Jurassic World":           4055,
	"Charter Communications":   2467,
	"Chick Fil A":              2244,
	"Planet Fitness":           1898,
	"Pitch Perfect":            1484,
	"Express":                  1689,
	"Home":                     1112,
	"Johnny Depp":              985,
	"Lena Dunham":              847,
	"Lewis Hamilton":           582,
	"KXAN":                     555,
	"Mary Ellen Mark":          550,
	"Farrah Abraham":           462,
	"Rita Ora":                 366,
	"Serena Williams":          282,
	"NCAA baseball tournament": 273,
	"Point Break":              265,
}

func wcBase() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-示例图"})
	wc.Add("wordcloud", wcData, charts.WordCloudOpts{SizeRange: []float32{14, 80}})
	return wc
}

func wcCardioid() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-形状(cardioid)"})
	wc.Add("wordcloud", wcData,
		charts.WordCloudOpts{SizeRange: []float32{14, 80}}, charts.WordCloudOpts{Shape: "cardioid"})
	return wc
}

func wcStar() *charts.WordCloud {
	wc := charts.NewWordCloud()
	wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-形状(star)"})
	wc.Add("wordcloud", wcData,
		charts.WordCloudOpts{SizeRange: []float32{14, 80}}, charts.WordCloudOpts{Shape: "cardioid"})
	return wc
}

func wcHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("wordCloud")...)
	page.Add(
		wcBase(),
		wcCardioid(),
		wcStar(),
	)
	f, err := os.Create(getRenderPath("wordCloud.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
