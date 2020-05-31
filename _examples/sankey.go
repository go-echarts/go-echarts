package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/go-echarts/go-echarts/charts"
	"github.com/gobuffalo/packr"
)

var sankeyNode = []charts.SankeyNode{
	{Name: "category1"},
	{Name: "category2"},
	{Name: "category3"},
	{Name: "category4"},
	{Name: "category5"},
	{Name: "category6"},
}

var sankeyLink = []charts.SankeyLink{
	{Source: "category1", Target: "category2", Value: 10},
	{Source: "category2", Target: "category3", Value: 15},
	{Source: "category3", Target: "category4", Value: 20},
	{Source: "category5", Target: "category6", Value: 25},
}

func sankeyBase() *charts.Sankey {
	sankey := charts.NewSankey()
	sankey.SetGlobalOptions(charts.TitleOpts{Title: "Sankey-示例图"})
	sankey.Add("sankey", sankeyNode, sankeyLink, charts.LabelTextOpts{Show: true})
	return sankey
}

func graphEnergy() *charts.Sankey {
	sankey := charts.NewSankey()
	sankey.SetGlobalOptions(charts.TitleOpts{Title: "Sankey-官方示例"})
	box := packr.NewBox(path.Join(".", "fixtures"))
	f, err := box.Find("energy.json")
	if err != nil {
		log.Fatal(err)
	}
	type Data struct {
		Nodes []charts.SankeyNode
		Links []charts.SankeyLink
	}

	var data Data
	if err := json.Unmarshal(f, &data); err != nil {
		fmt.Println(err)
	}
	sankey.Add("sankey", data.Nodes, data.Links,
		charts.LineStyleOpts{Curveness: 0.5, Color: "source"},
		charts.LabelTextOpts{Show: true},
	)
	return sankey
}

func sankeyHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("sankey")...)
	page.Add(
		sankeyBase(),
		graphEnergy(),
	)
	f, err := os.Create(getRenderPath("sankey.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
