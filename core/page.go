package core

import (
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/templates"
	"github.com/go-echarts/go-echarts/v2/types"
)

const DefaultPageTitle = "Awesome go-echarts"
const DefaultEchartsAsset = "https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"

// Page represents a page chart.
type Page struct {
	//PageTitle string `default:"Awesome go-echarts"`
	Title primitive.String

	// Assets host
	//EchartsJsAsset string `default:"https://go-echarts.github.io/go-echarts-assets/assets/"`
	JSAssets  *types.OrderedSet
	CSSAssets *types.OrderedSet

	Templates []primitive.String
	// Containers, one - one
	Containers []*Container
}

type PageConfig func(p *Page)

func (page *Page) Render(file string) {
	_ = (&DefaultRender{}).Render(file, page)
}

func (page *Page) RenderWithCustomRender(file string, renderer Render) {
	_ = renderer.Render(file, page)
}

func NewPage(containers ...*Container) *Page {

	return &Page{
		Title:      DefaultPageTitle,
		JSAssets:   (&types.OrderedSet{}).Add(DefaultEchartsAsset),
		CSSAssets:  &types.OrderedSet{},
		Templates:  []primitive.String{primitive.StringOf(templates.Tpl)},
		Containers: containers,
	}

}

func (page *Page) AddCharts(charts ...Chart) *Page {
	for _, c := range charts {
		page.Containers = append(page.Containers, c.GetContainer())
	}
	return page
}

func (page *Page) Config(configs ...PageConfig) *Page {

	for _, fn := range configs {
		fn(page)
	}
	return page
}
