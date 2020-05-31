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

var graphNodes = []charts.GraphNode{
	{Name: "节点1"},
	{Name: "节点2"},
	{Name: "节点3"},
	{Name: "节点4"},
	{Name: "节点5"},
	{Name: "节点6"},
	{Name: "节点7"},
	{Name: "节点8"},
}

func genLinks() []charts.GraphLink {
	links := make([]charts.GraphLink, 0)
	for i := 0; i < len(graphNodes); i++ {
		for j := 0; j < len(graphNodes); j++ {
			links = append(links,
				charts.GraphLink{Source: graphNodes[i].Name, Target: graphNodes[j].Name})
		}
	}
	return links
}

func graphBase() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-示例图"})
	graph.Add("graph", graphNodes, genLinks(),
		charts.GraphOpts{Force: charts.GraphForce{Repulsion: 8000}},
	)
	return graph
}

func graphCircle() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-布局(Circular)"})
	graph.Add("graph", graphNodes, genLinks(),
		charts.GraphOpts{Layout: "circular", Force: charts.GraphForce{Repulsion: 8000}},
		charts.LabelTextOpts{Show: true, Position: "right"},
	)
	return graph
}

func graphNpmDep() *charts.Graph {
	graph := charts.NewGraph()
	graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-npm package 依赖关系"})
	box := packr.NewBox(path.Join(".", "fixtures"))
	f, err := box.Find("npmdepgraph.json")
	if err != nil {
		log.Fatal(err)
	}
	type Data struct {
		Nodes []charts.GraphNode
		Links []charts.GraphLink
	}

	var data Data
	if err := json.Unmarshal(f, &data); err != nil {
		fmt.Println(err)
	}
	graph.Add("graph", data.Nodes, data.Links,
		charts.GraphOpts{Layout: "none", Roam: true, FocusNodeAdjacency: true},
		charts.EmphasisOpts{Label: charts.LabelTextOpts{Show: true, Position: "left", Color: "black"}},
		charts.LineStyleOpts{Curveness: 0.3},
	)
	return graph
}

func graphHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("graph")...)
	page.Add(
		graphBase(),
		graphCircle(),
		graphNpmDep(),
	)
	f, err := os.Create(getRenderPath("graph.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
