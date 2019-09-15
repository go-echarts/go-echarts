package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func scatterBase() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.TitleOpts{Title: "Scatter-示例图"})
	scatter.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	return scatter
}

func scatterShowLabel() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.TitleOpts{Title: "Scatter-显示 Label"})
	scatter.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	scatter.SetSeriesOptions(charts.LabelTextOpts{Show: true, Position: "right"})
	return scatter
}

func scatterSplitLine() *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.TitleOpts{Title: "Scatter-显示分割线"},
		charts.XAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
		charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
	)
	scatter.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	return scatter
}

func scatterHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("scatter")...)
	page.Add(
		scatterBase(),
		scatterShowLabel(),
		scatterSplitLine(),
	)
	f, err := os.Create(getRenderPath("scatter.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
