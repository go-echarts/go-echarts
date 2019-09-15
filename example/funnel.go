package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func funnelBase() *charts.Funnel {
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(charts.TitleOpts{Title: "Funnel-示例图"})
	funnel.Add("funnel", genKvData())
	return funnel
}

func funnelShowLabel() *charts.Funnel {
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(charts.TitleOpts{Title: "Funnel-显示 Label"})
	funnel.Add("funnel", genKvData(), charts.LabelTextOpts{Show: true})
	return funnel
}

func funnelHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("funnel")...)
	page.Add(
		funnelBase(),
		funnelShowLabel(),
	)
	f, err := os.Create(getRenderPath("funnel.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
