package main

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opt"
	"github.com/go-echarts/go-echarts/v2/series"
	"math"
)

var line3DColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genLine3dData() []opt.Chart3DData {
	data := make([][3]float64, 0)
	for i := 0; i < 25000; i++ {
		t := float64(i) / 1000
		data = append(data,
			[3]float64{(1 + 0.25*math.Cos(75*t)) * math.Cos(t), (1 + 0.25*math.Cos(75*t)) * math.Sin(t), t + 2.0*math.Sin(75.0*t)},
		)
	}

	ret := make([]opt.Chart3DData, 0, len(data))
	for _, d := range data {
		ret = append(ret, opt.Chart3DData{Value: []interface{}{d[0], d[1], d[2]}})
	}
	return ret

}

func NewLine3D() *charts.Line3D {
	line3d := charts.NewLine3D()
	// seems the echarts version not match somehow
	line3d.Page.JSAssets.Reset()
	line3d.Page.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts/dist/echarts.min.js")
	line3d.Page.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts-gl/dist/echarts-gl.min.js")

	line3d.VisualMaps = &opt.VisualMap{
		Min:       0,
		Max:       30,
		Dimension: "2",
		InRange:   &opt.VisualMapInRange{Color: line3DColor},
		Show:      false,
	}

	line3d.XAxis3D.Type = "value"
	line3d.YAxis3D.Type = "value"
	line3d.ZAxis3D.Type = "value"
	line3d.Grid3D = &opt.Grid3D{ViewControl: &opt.ViewControl{
		Projection: "orthographic",
	}}

	s := &series.Line3DSeries{
		Type:      "line3D",
		LineStyle: &opt.LineStyle{Width: 4},
		Data:      genLine3dData(),
	}
	line3d.AddSeries(s)
	return line3d

}

func NewLine3DChart() {

	NewLine3D().Render("line3d.html")

}
