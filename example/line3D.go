package main

import (
	"log"
	"math"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func genLine3dData() [][3]float64 {
	data := make([][3]float64, 0)
	for i := 0; i < 25000; i++ {
		t := float64(i) / 1000
		data = append(data,
			[3]float64{
				(1 + 0.25*math.Cos(75*float64(t))) * math.Cos(float64(t)),
				(1 + 0.25*math.Cos(75*float64(t))) * math.Sin(float64(t)),
				float64(t) + 2.0*math.Sin(75.0*float64(t)),
			},
		)
	}
	return data
}

func line3DBase() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.TitleOpts{Title: "Line3D-示例图"},
		charts.VisualMapOpts{
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        30,
		},
	)
	line3d.AddZAxis("line3D", genLine3dData())
	return line3d
}

func line3DAutoRotate() *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.TitleOpts{Title: "Line3D-旋转的弹簧"},
		charts.VisualMapOpts{
			Calculable: true,
			InRange:    charts.VMInRange{Color: rangeColor},
			Max:        30,
		},
		charts.Grid3DOpts{
			ViewControl: charts.ViewControlOpts{AutoRotate: true},
		},
	)
	line3d.AddZAxis("line3D", genLine3dData())
	return line3d
}

func line3DHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("line3D")...)
	page.Add(
		line3DBase(),
		line3DAutoRotate(),
	)
	f, err := os.Create(getRenderPath("line3D.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
