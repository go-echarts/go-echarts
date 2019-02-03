package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func pieBase() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-示例图"})
	pie.Add("pie", genKvData())
	return pie
}

func pieShowLabel() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-显示 Label"})
	pie.Add("pie", genKvData(), charts.LabelTextOpts{Show: true})
	return pie
}

func pieHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("pie")...)
	page.Add(
		pieBase(),
		pieShowLabel(),
	)
	f, err := os.Create(getRenderPath("pie.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
