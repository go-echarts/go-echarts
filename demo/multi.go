package main

import (
	"github.com/go-echarts/go-echarts/v2/core"
)

func MultiCharts() {

	page := core.NewPage()
	page.Config(func(p *core.Page) {
		p.Title = "New Title"
		p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts/dist/echarts.min.js")
		p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts-gl/dist/echarts-gl.min.js")

	}).AddCharts(
		NewBaseLine(),
		NewBar(),
		NewComplexLine(),
		NewLine3D(),
	).CustomRender("multi-charts.html", &core.DefaultRender{})

}
