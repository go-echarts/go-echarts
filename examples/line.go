package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func lineBase() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-示例图"})
	line.SetXAxis(nameItems).AddSeries("商家A", randInt())
	return line
}

func lineShowLabel() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-显示 Label"})
	line.SetXAxis(nameItems).AddSeries("商家A", randInt(), charts.LabelTextOpts{Show: true})
	return line
}

func lineMarkPoint() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-标记点"})

	var markpoints = []charts.SeriesOptser{
		charts.MPNameTypeItem{Name: "最大值", Type: "max"},
		charts.MPNameTypeItem{Name: "平均值", Type: "average"},
		charts.MPNameTypeItem{Name: "最小值", Type: "min"},
		charts.MPStyleOpts{Label: charts.LabelTextOpts{Show: true}},
	}
	line.SetXAxis(nameItems).AddSeries("商家A", randInt(), markpoints...)
	return line
}

func lineSplitLine() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-显示分割线"})
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LabelTextOpts{Show: true})
	line.SetGlobalOptions(charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}})
	return line
}

func lineStep() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-阶梯图"})
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LineOpts{Step: true})
	return line
}

func lineSmooth() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-平滑曲线"})
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LineOpts{Smooth: true})
	return line
}

func lineArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-填充区域"})
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt(),
		charts.LabelTextOpts{Show: true},
		charts.AreaStyleOpts{Opacity: 0.2},
	)
	return line
}

func lineSmoothArea() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-平滑区域"})
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt(),
		charts.LabelTextOpts{Show: true},
		charts.AreaStyleOpts{Opacity: 0.2},
		charts.LineOpts{Smooth: true},
	)
	return line
}

func lineMulti() *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-多线"}, charts.InitOpts{Theme: "shine"})
	line.AddXAxis(nameItems).
		AddYAxis("商家 A", randInt()).
		AddYAxis("商家 B", randInt()).
		AddYAxis("商家 C", randInt()).
		AddYAxis("商家 D", randInt())
	return line
}

func lineDemo() *charts.Line {
	line := charts.NewLine()
	line.AddXAxis([]string{"10e1", "10e2", "10e3", "10e4", "10e5", "10e6", "10e7"}).
		AddYAxis("map", []float32{19.9, 16.8, 19.9, 29.4, 61.3, 77.3, 93.0},
			charts.LabelTextOpts{Show: true, Position: "bottom"}).
		AddYAxis("slice", []float32{24.9, 34.9, 48.1, 58.3, 69.7, 123, 131},
			charts.LabelTextOpts{Show: true, Position: "top"})
	line.SetSeriesOptions(
		charts.MLNameTypeItem{Name: "平均值", Type: "average"},
		charts.LineOpts{Smooth: true},
		charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{a}: {b}"}},
	)
	line.SetGlobalOptions(
		charts.TitleOpts{Title: "查询时间对比 哈希表 vs 二分查找"},
		charts.YAxisOpts{Name: "搜索时间(ns)", SplitLine: charts.SplitLineOpts{Show: false}},
		charts.XAxisOpts{Name: "元素数量"})
	return line
}

func lineHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("line")...)
	page.Add(
		lineBase(),
		lineShowLabel(),
		lineMarkPoint(),
		lineSplitLine(),
		lineStep(),
		lineSmooth(),
		lineArea(),
		lineSmoothArea(),
		lineMulti(),
		lineDemo(),
	)
	f, err := os.Create(getRenderPath("line.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
