package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
)

func NewBar() *charts.Bar {

	bar := charts.NewBar()
	bar.Title.Text = "Title-Title"
	bar.Title.SubText = "Subtitle-01"

	bar.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	bar.Series.Name = "Sale"
	bar.Series.Data = []int{150, 230, 224, 218, 135, 147, 260}

	bar.Render("bar.html")
	return bar
}

func NewBarChart() {
	NewBar().Render("bar.html")
}
