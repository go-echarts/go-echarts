package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func pieBase() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-示例图"})
	pie.AddSeries("pie", genKvData())
	return pie
}

func pieShowLabel() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-显示 Label"})
	pie.AddSeries("pie", genKvData(), charts.LabelTextOpts{Show: true})
	return pie
}

func pieLabelFormatter() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-Label 格式"})
	pie.AddSeries("pie", genKvData(), charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"})
	return pie
}

func pieRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-Radius"})
	pie.AddSeries("pie", genKvData(),
		charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
		charts.PieOpts{Radius: []string{"40%", "75%"}},
	)
	return pie
}

func pieRoseArea() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-玫瑰图(Area)"})
	pie.AddSeries("pie", genKvData(),
		charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
		charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "area"},
	)
	return pie
}

func pieRoseRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-玫瑰图(Radius)"})
	pie.AddSeries("pie", genKvData(),
		charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
		charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "radius"},
	)
	return pie
}

func pieRoseAreaRadius() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-玫瑰图(Area/Radius)"})
	pie.AddSeries("area", genKvData(),
		charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "area", Center: []string{"25%", "50%"}},
	)
	pie.AddSeries("radius", genKvData(),
		charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
		charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "radius", Center: []string{"75%", "50%"}},
	)
	return pie
}

func pieInPie() *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-饼中饼"})
	pie.AddSeries("area", genKvData(),
		charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
		charts.PieOpts{Radius: []string{"50%", "55%"}, RoseType: "area"},
	)
	pie.AddSeries("radius", genKvData(),
		charts.PieOpts{Radius: []string{"0%", "45%"}, RoseType: "radius"},
	)
	return pie
}

func pieHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("pie")...)
	page.AddCharts(
		pieBase(),
		pieShowLabel(),
		pieLabelFormatter(),
		pieRadius(),
		pieRoseArea(),
		pieRoseRadius(),
		pieRoseAreaRadius(),
		pieInPie(),
	)
	f, err := os.Create(getRenderPath("pie.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
