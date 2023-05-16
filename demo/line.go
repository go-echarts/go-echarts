package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/primitive"
)

func NewLine() *charts.Line {
	line := charts.NewLine()
	line.JSAssets.Add("My.js")
	line.Title.Text = "Title-Title"
	line.Title.SubText = "Subtitle-01"

	line.XAxis.Type = "category"
	line.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	line.YAxis.Type = "value"
	line.Series.Data = []int{150, 230, 224, 218, 135, 147, 260}

	// show Toolbox but no other empty structs, test
	line.Toolbox.Show = primitive.True()

	// change container
	line.Container.ChartID = "customId"
	line.Container.Theme = "dark"
	// change page title
	line.Page.Title = "My go-echarts title"
	return line

}

func NewLineChart() {
	NewLine().Render("line.html")
}
