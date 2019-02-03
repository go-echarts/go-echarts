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
		esBase(),
		funnelBase(),
		gaugeBase(),
		geoBase(),
		heatMapBase(),
		klineBase(),
		lineBase(),
		liquidBase(),
		mapBase(),
		overlapBase(),
		pieBase(),
		scatterBase(),
		wcBase(),
	)
	f, err := os.Create(getRenderPath("page.html"))
	if err != nil {
		log.Println(err)
	}
	p.Render(w, f)
}
