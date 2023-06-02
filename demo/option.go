package main

import (
	"github.com/go-echarts/go-echarts/v2/page"
	"github.com/go-echarts/go-echarts/v2/render"
)

func NewOptionGallery() {

	p := page.New(func(p *page.Page) {
		p.Title = "My Option Gallery Title"
		p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts/dist/echarts.min.js")
		p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts-gl/dist/echarts-gl.min.js")
	})

	p.AddCharts(NewBar(), NewBaseLine(), NewComplexLine(), NewLine3D())

	render.Render(p, "gallery-option.html")

}

func NewOptionGalleryWithCustom() {

	p := page.NewWithTemplate(&MyPageTplProvider{}, func(p *page.Page) {
		p.Title = "My Option Custom Gallery Title"
	})

	p.AddCharts(NewBar())

	render.WithConfig(&render.Config{
		CustomerRenderer: &MyMockPNGRenderer{},
	}, p, "gallery-png.png")

}
