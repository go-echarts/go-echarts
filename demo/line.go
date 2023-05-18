package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opt"
	"github.com/go-echarts/go-echarts/v2/series"
)

func NewBaseLine() *charts.Line {
	line := charts.NewLine()

	line.JSAssets.Add("My.js")

	line.Title.Text = "Title-Title"
	line.Title.SubText = "Subtitle-01"

	line.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	s := &series.LineSeries{}
	s.Data = []int{150, 230, 224, 218, 135, 147, 260}

	line.AddSeries(
		s,
	)

	// show Toolbox
	line.Toolbox.Feature = &opt.ToolBoxFeature{
		SaveAsImage: &opt.ToolBoxFeatureSaveAsImage{},
		Brush:       &opt.ToolBoxFeatureBrush{},
		DataZoom:    &opt.ToolBoxFeatureDataZoom{},
		DataView:    &opt.ToolBoxFeatureDataView{},
		Restore:     &opt.ToolBoxFeatureRestore{},
	}

	// change container
	line.Container.ChartID = "customId"
	line.Container.Theme = "dark"
	// change page title
	line.Page.Title = "My go-echarts title"
	return line

}

func NewComplexLine() *charts.Line {
	line := charts.NewLine()

	line.Title.Text = "Line Complex"
	line.Title.SubText = "Subtitle-Complex"

	line.Tooltip.Trigger = "axis"

	line.Toolbox.Feature = &opt.ToolBoxFeature{
		SaveAsImage: &opt.ToolBoxFeatureSaveAsImage{},
		Brush:       &opt.ToolBoxFeatureBrush{},
		DataZoom:    &opt.ToolBoxFeatureDataZoom{},
		DataView:    &opt.ToolBoxFeatureDataView{},
		Restore:     &opt.ToolBoxFeatureRestore{},
	}

	line.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	line.YAxis.AxisLabel = &opt.AxisLabel{Formatter: "{value} c"}

	s := &series.LineSeries{}
	s.Name = "Highest"
	s.Data = []int{10, 11, 13, 11, 12, 12, 9}

	d1 := &opt.MarkPointNameTypeItem{
		Name: "Max",
		Type: "max",
	}

	d2 := &opt.MarkPointNameTypeItem{
		Name: "Min",
		Type: "min",
	}

	s.MarkPoint = &opt.MarkPoint{
		Data: []interface{}{d1, d2},
	}

	d3 := &opt.MarkPointNameTypeItem{
		Name: "Avg",
		Type: "average",
	}

	s.MarkLines = &opt.MarkLines{
		Data: []interface{}{d3},
	}

	s1 := &series.LineSeries{}
	s1.Name = "Lowest"
	s1.Data = []int{1, -2, 2, 5, 3, 2, 0}

	d4 := &opt.MarkPointNameTypeItem{
		Name:  "Week Lowest",
		Value: -2,
		XAxis: 1,
		YAxis: -1.5,
	}

	s1.MarkPoint = &opt.MarkPoint{
		Data: []interface{}{d4},
	}

	d5 := &opt.MarkPointNameTypeItem{
		Name: "Avg",
		Type: "average",
	}

	s1.MarkLines = &opt.MarkLines{
		Data: []interface{}{d5},
	}

	line.AddSeries(s, s1)

	return line

}

func NewLineChart() {
	NewBaseLine().Render("line.html")
}

func NewComplexLineChart() {
	NewComplexLine().Render("line-complex.html")
}
