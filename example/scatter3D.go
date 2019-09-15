package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func genScatter3dData() [][3]int {
	data := make([][3]int, 0)
	for i := 0; i < 80; i++ {
		data = append(data, [3]int{
			int(seed.Int63()) % 100,
			int(seed.Int63()) % 100,
			int(seed.Int63()) % 100,
		})
	}
	return data
}

func scatter3DBase() *charts.Scatter3D {
	scatter3d := charts.NewScatter3D()
	scatter3d.SetGlobalOptions(
		charts.TitleOpts{Title: "Scatter3D-示例图"},
		charts.VisualMapOpts{
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        100,
		},
		charts.Grid3DOpts{BoxDepth: 80, BoxWidth: 200},
	)
	scatter3d.AddZAxis("scatter3d", genScatter3dData())
	return scatter3d
}

func scatter3DHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("scatter3D")...)
	page.Add(
		scatter3DBase(),
	)
	f, err := os.Create(getRenderPath("scatter3D.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
