package main

import "github.com/go-echarts/go-echarts/v2/core"

func main() {
	core.NewCanvas().
		Page().UseTemplate(&MyPageTplProvider{}).
		PageConfig(func(p *core.Page) {
			p.Title = "My Page Title"
		}).Charts().
		AddCharts(
			NewBar(),
			NewBaseLine(),
			NewComplexLine(),
			NewLine3D(),
		).
		Render("MyPage.html")

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
