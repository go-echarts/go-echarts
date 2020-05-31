package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func overlapBase() *charts.Bar {
	bar := barBase()
	bar.Title = "Overlap-示例图"
	bar.Overlap(lineBase())
	return bar
}

func overlapBarScatter() *charts.Bar {
	bar := barBase()
	bar.Title = "Overlap-Bar+Scatter"
	bar.Overlap(scatterBase())
	return bar
}

func overlapLineScatter() *charts.Line {
	line := lineBase()
	line.Title = "Overlap-Line+Scatter"
	line.Overlap(scatterBase())
	return line
}

func overlapBarLineScatterEs() *charts.Bar {
	bar := barBase()
	bar.Title = "Overlap-All"
	bar.Overlap(lineBase(), scatterBase(), esBase())
	return bar
}

func overlapHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("overlap")...)
	page.Add(
		overlapBase(),
		overlapBarScatter(),
		overlapLineScatter(),
		overlapBarLineScatterEs(),
	)
	f, err := os.Create(getRenderPath("overlap.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
