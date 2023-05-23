package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/canvas"
	"github.com/go-echarts/go-echarts/v2/core"
	"strings"
)

func NewFlowGallery() {
	canvas.New().
		Page().
		PageConfig(func(p *core.Page) {
			p.Title = "My Flow Page Title"
			p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts/dist/echarts.min.js")
			p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts-gl/dist/echarts-gl.min.js")
		}).Charts().
		AddCharts(
			NewBar(),
			NewBaseLine(),
			NewComplexLine(),
			NewLine3D()).
		RendererConfig(nil, func(defaultWriter core.Writer) core.Writer {
			// allow to rewriter Writer
			return defaultWriter

		}).
		Render("gallery-flow.html")

}

func NewFlowGalleryWithCustom() {
	canvas.New().
		Page().
		UseTemplate(&MyPageTplProvider{}).
		PageConfig(func(p *core.Page) {
			p.Title = "My Flow Custom Page Title"
		}).Charts().
		AddCharts(
			NewBar(),
		).
		CustomRenderer(&MyMockPNGRenderer{}).
		Render("gallery.png")
}

type MyMockPNGRenderer struct {
}

func (mcr *MyMockPNGRenderer) Resolve(p *core.Page, dest string) {
	title := p.Title.StringVal()
	suffix := "Option Way"
	if strings.Contains(title, "Flow") {
		suffix = "Flow Way"
	}
	// Process render PNG to dest
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	fmt.Print("\n")
	fmt.Println("\n Process render PNG to " + dest + " in " + suffix)
	fmt.Print("\n")
}

func (mcr *MyMockPNGRenderer) GetRenderer() *core.Renderer {
	return nil
}

func (mcr *MyMockPNGRenderer) SetRender(r core.Render) {
	panic("implement me")
}

func (mcr *MyMockPNGRenderer) SetWriter(w core.Writer) {
	panic("implement me")
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
