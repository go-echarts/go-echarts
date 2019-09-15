package main

import (
	"log"
	"math"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func genSurface3dData0() [][3]interface{} {
	data := make([][3]interface{}, 0)

	for i := -60; i < 60; i++ {
		y := float64(i) / 60
		for j := -60; j < 60; j++ {
			x := float64(j) / 60
			z := math.Sin(x*math.Pi) * math.Sin(y*math.Pi)
			data = append(data, [3]interface{}{x, y, z})
		}
	}
	return data
}

func genSurface3dData1() [][3]interface{} {
	data := make([][3]interface{}, 0)
	for i := -30; i < 30; i++ {
		y := float64(i) / 10
		for j := -30; j < 30; j++ {
			x := float64(j) / 10
			z := math.Sin(x*x+y*y) * x / math.Pi
			data = append(data, [3]interface{}{x, y, z})
		}
	}
	return data
}

func surface3DBase() *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.TitleOpts{Title: "surface3D-示例图"},
		charts.VisualMapOpts{
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        3,
			Min:        -3,
		},
	)
	surface3d.AddZAxis("surface3d", genSurface3dData0())
	return surface3d
}

func surface3DRose() *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.TitleOpts{Title: "surface3D-一朵玫瑰"},
		charts.VisualMapOpts{
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        3,
			Min:        -3,
		},
	)
	surface3d.AddZAxis("surface3d", genSurface3dData1())
	return surface3d
}

func surface3DHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("surface3D")...)
	page.Add(
		surface3DBase(),
		surface3DRose(),
	)
	f, err := os.Create(getRenderPath("surface3D.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
