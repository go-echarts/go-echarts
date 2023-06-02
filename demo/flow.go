package main

import (
	"github.com/go-echarts/go-echarts/v2/canvas"
	"github.com/go-echarts/go-echarts/v2/page"
	"github.com/go-echarts/go-echarts/v2/render"
)

func NewFlowGallery() {
	canvas.New().
		Page().
		PageConfig(func(p *page.Page) {
			p.Title = "My Flow Page Title"
			p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts/dist/echarts.min.js")
			p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts-gl/dist/echarts-gl.min.js")
		}).Charts().
		AddCharts(
			NewBar(),
			NewBaseLine(),
			NewComplexLine(),
			NewLine3D()).
		RendererConfig(nil, func(defaultWriter render.IWriter) (newWriter render.IWriter) {
			// allow to rewriter IWriter/IRender
			return defaultWriter
		}).
		Render("gallery-flow.html")

}

func NewFlowGalleryWithCustom() {
	canvas.New().
		Page().
		UseTemplate(&MyPageTplProvider{}).
		PageConfig(func(p *page.Page) {
			p.Title = "My Flow Custom Page Title"
		}).Charts().
		AddCharts(
			NewBar(),
		).
		CustomRenderer(&MyMockPNGRenderer{}).
		Render("gallery.png")
}
