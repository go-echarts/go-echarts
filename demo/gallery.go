package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/core"
)

func main() {
	core.New().
		Page().
		UseTemplate(&MyPageTplProvider{}).
		PageConfig(func(p *core.Page) {
			p.Title = "My Page Title"
		}).Charts().
		AddCharts(
			NewBar(),
			NewBaseLine(),
			NewComplexLine(),
			NewLine3D(),
		).
		Render("gallery.html")

	// CustomRenderer(&MyMockPNGRenderer{}).
	// Render("gallery.png")

}

type MyMockPNGRenderer struct {
}

func (mcr *MyMockPNGRenderer) Render(p *core.Page, dest string) {
	// Process render PNG to dest
	fmt.Println("Process render PNG to " + dest)
}

func (mcr *MyMockPNGRenderer) GetRenderer() *core.Renderer {
	return nil
}

type MyPageTplProvider struct {
}

func (mptp *MyPageTplProvider) Provide() *core.Page {
	p := core.NewPage()
	p.JSAssets.Reset()
	p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts/dist/echarts.min.js")
	p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts-gl/dist/echarts-gl.min.js")
	p.JSAssets.Add("My.js")
	return p
}
