package main

import (
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/opts"
)

var (
	itemCnt = 7
	weeks   = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
)

func generateBarItems() []opts.BarChartItem {
	items := make([]opts.BarChartItem, 0)
	for i := 0; i < itemCnt; i++ {
		items = append(items, opts.BarChartItem{Value: rand.Intn(300)})
	}
	return items
}

func barBase() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Bar-basic-example",
			Subtitle: "This is the subtitle.",
		}),
	)

	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barTitle() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Bar-with-title/legend-opts",
			Subtitle: "go-echarts is an awesome chart library written in Golang",
			Link:     "https://github.com/go-echarts/go-echarts",
			Right:    "40%",
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show: true,
		}),
		charts.WithLegendOpts(opts.Legend{
			Right: "80%",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barShowLabel() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bar-with-label-opts",
		}),
	)
	bar.SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:      true,
			Formatter: "{b}:{c}",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barXYName() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bar-display-axes-name",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "XAxisName",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "YAxisName",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

//func barColor() *charts.Bar {
//	bar := charts.NewBar()
//	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-设置系列颜色"})
//	bar.AddXAxis(nameItems).
//		AddYAxis("商家A", randInt(), charts.ColorOpts{"lightblue"}).
//		AddYAxis("商家B", randInt(), charts.ColorOpts{"pink"})
//	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
//	// 或者可以这样设置
//	//bar.SetGlobalOptions(charts.ColorOpts{"lightblue", "pink"})
//	return bar
//}
//
func barSplitLine() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bar-splitline",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "XAxisName",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "YAxisName",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
	)
	bar.SetXAxis(nameItems).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

func barGap() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bar-bargap",
		}),
	)
	bar.SetXAxis(nameItems).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	bar.SetSeriesOptions(
		charts.WithBarChartOpts(opts.BarChart{
			BarGap: "70%",
		}),
	)
	return bar
}

func barYAxis() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bar-yaxis-formatter",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: opts.Label{
				Formatter: "{value} 件/天",
			},
		}),
	)

	bar.SetXAxis(nameItems).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	return bar
}

//func barMultiYAxis() *charts.Bar {
//	bar := charts.NewBar()
//	bar.SetGlobalOptions(
//		charts.TitleOpts{Title: "Bar-多 Y 轴"},
//		charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}},
//	)
//	bar.AddXAxis(nameItems).
//		AddYAxis("商家A", randInt(), charts.BarOpts{YAxisIndex: 0}).
//		AddYAxis("商家B", randInt(), charts.BarOpts{YAxisIndex: 1})
//	bar.ExtendYAxis(charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/月"}})
//	return bar
//}
//
//func barMultiXAxis() *charts.Bar {
//	bar := charts.NewBar()
//	bar.SetGlobalOptions(
//		charts.TitleOpts{Title: "Bar-多 X 轴"},
//		charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}},
//	)
//	bar.AddXAxis(nameItems).
//		AddYAxis("商家A", randInt(), charts.BarOpts{XAxisIndex: 0}).
//		AddYAxis("商家B", randInt(), charts.BarOpts{XAxisIndex: 1})
//	bar.ExtendXAxis(charts.XAxisOpts{Data: foodItems})
//	return bar
//}
//
//func barDataZoom() *charts.Bar {
//	bar := charts.NewBar()
//	bar.SetGlobalOptions(
//		charts.TitleOpts{Title: "Bar-DataZoom"},
//		charts.ToolboxOpts{Show: true},
//		charts.DataZoomOpts{XAxisIndex: []int{0}, Start: 50, End: 100},
//	)
//	bar.AddXAxis(nameItems).
//		AddYAxis("商家A", randInt()).
//		AddYAxis("商家B", randInt())
//	return bar
//}

func barReverse() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bar-reverse-xy-axis",
		}),
	)

	bar.SetXAxis(nameItems).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	bar.XYReversal()
	return bar
}

//func barStack() *charts.Bar {
//	bar := charts.NewBar()
//	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-堆叠效果"})
//	bar.AddXAxis(nameItems).
//		AddYAxis("商家A", randInt(), charts.BarOpts{Stack: "stack"}).
//		AddYAxis("商家B", randInt(), charts.BarOpts{Stack: "stack"})
//	return bar
//
//}
//
//func barMark() *charts.Bar {
//	bar := charts.NewBar()
//	bar.SetGlobalOptions(
//		charts.TitleOpts{Title: "Bar-标记线&标记点"},
//	)
//	bar.AddXAxis(nameItems).
//		AddYAxis("商家A", randInt()).
//		AddYAxis("商家B", randInt())
//	bar.SetSeriesOptions(
//		charts.MLNameTypeItem{Name: "最大值", Type: "max"},
//		charts.MLStyleOpts{SymbolSize: 20, Label: charts.LabelTextOpts{Show: true}},
//		charts.MPNameTypeItem{Name: "最大值点", Type: "max"},
//		charts.MPStyleOpts{Label: charts.LabelTextOpts{Show: true}},
//	)
//	return bar
//}
//
//func barMarkCustom() *charts.Bar {
//	bar := charts.NewBar()
//	bar.SetGlobalOptions(
//		charts.TitleOpts{Title: "Bar-自定义标记+主题"},
//		charts.InitOpts{PageTitle: "Awesome", Theme: charts.ThemeType.Macarons},
//	)
//	bar.AddXAxis(nameItems).
//		AddYAxis("商家A", randInt()).
//		AddYAxis("商家B", randInt(),
//			charts.MLNameYAxisItem{Name: "Y 值为 5", YAxis: 20},
//			charts.MLNameCoordItem{Name: "自定义标记线",
//				Coord0: []interface{}{"衬衫", 10}, Coord1: []interface{}{"羊毛衫", 40}},
//			charts.MLStyleOpts{SymbolSize: 20, Label: charts.LabelTextOpts{Show: true}},
//		)
//	return bar
//}

func barSize() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Bar-display-axes-name",
		}),
		charts.WithInitializationOpts(opts.Initialization{
			Width:  "800px",
			Height: "600px",
		}),
	)
	bar.SetXAxis(weeks).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category A", generateBarItems())
	return bar
}

func barHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage()
	page.AddCharts(
		barBase(),
		barTitle(),
		barShowLabel(),
		barXYName(),
		//barColor(),
		barSplitLine(),
		barGap(),
		barYAxis(),
		//barMultiYAxis(),
		//barMultiXAxis(),
		//barDataZoom(),
		barReverse(),
		//barStack(),
		//barMark(),
		//barMarkCustom(),
		barSize(),
	)
	f, err := os.Create(getRenderPath("bar.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(io.MultiWriter(w, f))
}
