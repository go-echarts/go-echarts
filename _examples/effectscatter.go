package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func esBase() *charts.EffectScatter {
	es := charts.NewEffectScatter()
	es.SetGlobalOptions(charts.TitleOpts{Title: "EffectScatter-示例图"})
	es.AddXAxis(nameItems).AddYAxis("es1", randInt())
	return es
}

func esEffectStyle() *charts.EffectScatter {
	es := charts.NewEffectScatter()
	es.SetGlobalOptions(charts.TitleOpts{Title: "EffectScatter-涟漪效果"})
	es.AddXAxis(nameItems).
		AddYAxis("es1", randInt(),
			charts.RippleEffectOpts{Period: 4, Scale: 10, BrushType: "stroke"}).
		AddYAxis("es2", randInt(),
			charts.RippleEffectOpts{Period: 3, Scale: 6, BrushType: "fill"})
	return es
}

func esHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("effectScatter")...)
	page.Add(
		esBase(),
		esEffectStyle(),
	)
	f, err := os.Create(getRenderPath("effectscatter.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
