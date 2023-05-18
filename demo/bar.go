package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/series"
)

func NewBar() *charts.Bar {

	bar := charts.NewBar()
	bar.Title.Text = "Title-Title"
	bar.Title.SubText = "Subtitle-01"

	bar.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	s := &series.BarSeries{}
	s.Name = "Sale"
	s.Data = []int{150, 230, 224, 218, 135, 147, 260}

	s1 := &series.BarSeries{}
	s1.Name = "Out"
	s1.Data = []int{12, 11, 4, 32, 121, 12, 1}
	bar.AddSeries(s, s1)

	return bar
}

func NewBarChart() {
	NewBar().Render("bar.html")
}
