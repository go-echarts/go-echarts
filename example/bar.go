package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chenjiandongx/go-echarts/charts"
)

func barBase() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	return bar
}

func barShowLabel() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-显示 Label"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
	return bar
}

func barXYName() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-XY 轴名称"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
	bar.SetGlobalOptions(charts.XAxisOpts{Name: "商品名称"}, charts.YAxisOpts{Name: "销售量"})
	return bar
}

func barColor() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-设置系列颜色"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt(), charts.ColorOpts{"lightblue"}).
		AddYAxis("商家B", randInt(), charts.ColorOpts{"pink"})
	bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
	// 或者可以这样设置
	//bar.SetGlobalOptions(charts.ColorOpts{"lightblue", "pink"})
	return bar
}

func barSplitLine() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-显示分割线"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetGlobalOptions(charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}})
	return bar
}

func barYAxis() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-Y 轴格式"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetGlobalOptions(charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}})
	return bar
}

func barReverse() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-翻转 XY 轴"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.XYReversal()
	return bar
}

func barStack() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-堆叠效果"})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt(), charts.BarOpts{Stack: "stack"}).
		AddYAxis("商家B", randInt(), charts.BarOpts{Stack: "stack"})
	return bar

}

func barMark() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-标记线&标记点"},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	bar.SetSeriesOptions(
		charts.MLNameTypeItem{Name: "最大值", Type: "max"},
		charts.MLStyleOpts{SymbolSize: 20, Label: charts.LabelTextOpts{Show: true}},
		charts.MPNameTypeItem{Name: "最大值点", Type: "max"},
		charts.MPStyleOpts{Label: charts.LabelTextOpts{Show: true}},
	)
	return bar
}

func barMarkCustom() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.TitleOpts{Title: "Bar-自定义标记+主题"},
		charts.InitOpts{PageTitle: "Awesome", Theme: "macarons"},
	)
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt(),
			charts.MLNameYAxisItem{Name: "Y 值为 5", YAxis: 20},
			charts.MLNameCoordItem{Name: "自定义标记线",
				Coord0: []interface{}{"衬衫", 10}, Coord1: []interface{}{"羊毛衫", 40}},
			charts.MLStyleOpts{SymbolSize: 20, Label: charts.LabelTextOpts{Show: true}},
		)
	return bar
}

func BarHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("bar")...)
	page.Add(
		barBase(),
		barShowLabel(),
		barXYName(),
		barColor(),
		barSplitLine(),
		barYAxis(),
		barReverse(),
		barStack(),
		barMark(),
		barMarkCustom(),
	)
	f, err := os.Create(getRenderPath("bar.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
