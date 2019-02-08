package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func pageHandler(w http.ResponseWriter, _ *http.Request) {
	p := charts.NewPage(orderRouters("page")...)

	p.Add(
		barBase(),
		bar3DBase(),
		boxPlotBase(),
		esBase(),
		funnelBase(),
		gaugeBase(),
		geoBase(),
		heatMapBase(),
		klineBase(),
		lineBase(),
		line3DBase(),
		liquidBase(),
		mapBase(),
		overlapBase(),
		pieBase(),
		scatterBase(),
		scatter3DBase(),
		surface3DBase(),
		wcBase(),
	)
	f, err := os.Create(getRenderPath("page.html"))
	if err != nil {
		log.Println(err)
	}
	p.Render(w, f)
}
