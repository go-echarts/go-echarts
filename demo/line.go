package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/series"
)

func NewBaseLine() *charts.Line {
	line := charts.NewLine()

	line.JSAssets.Add("My.js")

	line.Title.Text = "Title-Title"
	line.Title.SubText = "Subtitle-01"

	line.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	s := series.LineSeries{}.New()
	s.Data = []int{150, 230, 224, 218, 135, 147, 260}

	line.AddSeries(
		s,
	)

	// show Toolbox
	line.Toolbox.Feature = &opts.ToolBoxFeature{
		SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{},
		Brush:       &opts.ToolBoxFeatureBrush{},
		DataZoom:    &opts.ToolBoxFeatureDataZoom{},
		DataView:    &opts.ToolBoxFeatureDataView{},
		Restore:     &opts.ToolBoxFeatureRestore{},
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

	line.Toolbox.Feature = &opts.ToolBoxFeature{
		SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{},
		Brush:       &opts.ToolBoxFeatureBrush{},
		DataZoom:    &opts.ToolBoxFeatureDataZoom{},
		DataView:    &opts.ToolBoxFeatureDataView{},
		Restore:     &opts.ToolBoxFeatureRestore{},
	}

	line.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	line.YAxis.AxisLabel = &opts.AxisLabel{Formatter: "{value} c"}

	s := series.LineSeries{}.New()
	s.Name = "Highest"
	s.Data = []int{10, 11, 13, 11, 12, 12, 9}

	d1 := &opts.MarkPointNameTypeItem{
		Name: "Max",
		Type: "max",
	}

	d2 := &opts.MarkPointNameTypeItem{
		Name: "Min",
		Type: "min",
	}

	s.MarkPoint = &opts.MarkPoint{
		Data: []interface{}{d1, d2},
	}

	d3 := &opts.MarkPointNameTypeItem{
		Name: "Avg",
		Type: "average",
	}

	s.MarkLines = &opts.MarkLines{
		Data: []interface{}{d3},
	}

	s1 := series.LineSeries{}.New()
	s1.Name = "Lowest"
	s1.Data = []int{1, -2, 2, 5, 3, 2, 0}

	d4 := &opts.MarkPointNameTypeItem{
		Name:  "Week Lowest",
		Value: -2,
		XAxis: 1,
		YAxis: -1.5,
	}

	s1.MarkPoint = &opts.MarkPoint{
		Data: []interface{}{d4},
	}

	d5 := &opts.MarkPointNameTypeItem{
		Name: "Avg",
		Type: "average",
	}

	s1.MarkLines = &opts.MarkLines{
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
