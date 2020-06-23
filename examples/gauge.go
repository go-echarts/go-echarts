package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func gaugeBase() *charts.Gauge {
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(charts.TitleOpts{Title: "Gauge-示例图"})
	m := make(map[string]interface{})
	m["工作进度"] = rand.Intn(50)
	gauge.AddSeries("gauge", m)
	return gauge
}

func gaugeTimer() *charts.Gauge {
	gauge := charts.NewGauge()

	m := make(map[string]interface{})
	m["工作进度"] = rand.Intn(50)
	gauge.AddSeries("gauge1", m)
	gauge.SetGlobalOptions(charts.TitleOpts{Title: "Gauge-定时器"})
	fn := fmt.Sprintf(`setInterval(function () {
			option_%s.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
			myChart_%s.setOption(option_%s, true);
		}, 2000);`, gauge.ChartID, gauge.ChartID, gauge.ChartID)
	gauge.AddJSFuncs(fn)
	return gauge
}

func gaugeHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("gauge")...)
	page.AddCharts(
		gaugeBase(),
		gaugeTimer(),
	)
	f, err := os.Create(getRenderPath("gauge.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
