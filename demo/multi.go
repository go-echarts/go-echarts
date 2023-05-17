package main

import (
	"github.com/go-echarts/go-echarts/v2/core"
)

func MultiCharts() {

	page := core.NewPage()
	page.Config(func(p *core.Page) {
		p.Title = "New Title"
		p.JSAssets.Add("New.js")

	}).AddCharts(
		NewBaseLine(),
		NewBar(),
		NewComplexLine(),
	).Render("multi-charts.html")

}
