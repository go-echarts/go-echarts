package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/primitive"
)

func main() {

	line := charts.NewLine()
	line.TitleV3.Text = "Title-Title"
	line.TitleV3.SubText = "Subtitle-01"

	line.XAxis.Type = "category"
	line.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	line.YAxis.Type = "value"
	line.Series.Data = []int{150, 230, 224, 218, 135, 147, 260}

	// show Toolbox but no other empty structs, test
	line.Toolbox.Show = primitive.True()

	line.Render("line.html")
}
