package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

func liquidBase() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-示例图"})
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LiquidOpts{IsWaveAnimation: true},
	)
	return liquid
}

func liquidShowLabel() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-展示 Label"})
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LabelTextOpts{Show: true},
		charts.LiquidOpts{IsWaveAnimation: true},
	)
	return liquid
}

func liquidOutline() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-显示轮廓"})
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LabelTextOpts{Show: true},
		charts.LiquidOpts{IsShowOutline: true, IsWaveAnimation: true},
	)
	return liquid
}

func liquidWaveAnimation() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-关闭动画"})
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LabelTextOpts{Show: true},
		charts.LiquidOpts{IsShowOutline: true},
	)
	return liquid
}

func liquidDiamond() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-形状(diamond)"})
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LiquidOpts{Shape: "diamond", IsWaveAnimation: true})
	return liquid
}

func liquidPin() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-形状(pin)"})
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LiquidOpts{Shape: "pin", IsWaveAnimation: true})
	return liquid
}

func liquidArrow() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-形状(arrow)"})
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LiquidOpts{Shape: "arrow", IsWaveAnimation: true})
	return liquid
}

func liquidTriangle() *charts.Liquid {
	liquid := charts.NewLiquid()
	liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-形状(triangle)"})
	liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
		charts.LiquidOpts{Shape: "triangle", IsWaveAnimation: true})
	return liquid
}

func liquidHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("liquid")...)
	page.Add(
		liquidBase(),
		liquidShowLabel(),
		liquidOutline(),
		liquidWaveAnimation(),
		liquidDiamond(),
		liquidPin(),
		liquidArrow(),
		liquidTriangle(),
	)
	f, err := os.Create(getRenderPath("liquid.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
