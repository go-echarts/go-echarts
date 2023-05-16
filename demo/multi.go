package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/primitive"
)

func MultiCharts() {

	line := charts.NewLine()
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

	line1 := charts.NewLine()
	line1.Title.Text = "Title-Title-1"
	line1.Title.SubText = "Subtitle-01-1"

	line1.XAxis.Type = "category"
	line1.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	line1.YAxis.Type = "value"
	line1.Series.Data = []int{150, 230, 224, 218, 135, 147, 260}

	core.NewDefaultPage().AddCharts(
		line,
		line1,
	).Render("new.html")

}
