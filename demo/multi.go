package main

import (
	"github.com/go-echarts/go-echarts/v2/core"
)

func MultiCharts() {

	core.NewDefaultPage().AddCharts(
		NewLine(),
		NewBar(),
	).Render("multi-charts.html")

}
