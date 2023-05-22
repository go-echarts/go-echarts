package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/series"
)

func NewBar() *charts.Bar {

	bar := charts.NewBar()
	bar.Title.Text = "Title-Title"
	bar.Title.SubText = "Subtitle-01"

	bar.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	s := &series.BarSingleSeries{}
	s.Name = "Sale"
	s.Data = []int{150, 230, 224, 218, 135, 147, 260}

	s1 := &series.BarSingleSeries{}
	s1.Name = "Out"
	s1.Data = []int{12, 11, 4, 32, 121, 12, 1}
	bar.AddSeries(s, s1)

	bar.Container.AddEvent(&core.Event{
		EventBinder: "on",
		EventType:   "click",
		Action:      "function(params) {\n   console.log(params.name);\n}",
	})

	return bar
}

func NewBarChart() {

	p := core.NewPage(NewBar().GetContainer())
	core.NewDefaultRenderer().Render(p, "bar.html")
}
