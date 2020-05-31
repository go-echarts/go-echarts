package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
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
		graphBase(),
		heatMapBase(),
		klineBase(),
		lineBase(),
		line3DBase(),
		liquidBase(),
		mapBase(),
		overlapBase(),
		parallelBase(),
		pieBase(),
		radarBase(),
		sankeyBase(),
		scatterBase(),
		scatter3DBase(),
		surface3DBase(),
		themeRiverBase(),
		wcBase(),
	)
	f, err := os.Create(getRenderPath("page.html"))
	if err != nil {
		log.Println(err)
	}
	p.Render(w, f)
}
