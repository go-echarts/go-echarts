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

	d1 := map[string]string{
		"type": "max",
		"name": "Max",
	}
	d2 := map[string]string{
		"type": "min",
		"name": "Min",
	}

	s.MarkPoints = &opts.MarkPoints{
		Data: []interface{}{d1, d2},
	}

	d3 := map[string]string{
		"type": "average",
		"name": "Avg",
	}

	s.MarkLines = &opts.MarkLines{
		Data: []interface{}{d3},
	}

	line.AddSeries(s)

	return line

}

func NewLineChart() {
	NewComplexLine().Render("line-complex.html")
}

func NewComplexLineChart() {
	NewBaseLine().Render("line.html")
}
