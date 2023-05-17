package core

import (
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/templates"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/go-echarts/go-echarts/v2/util"
)

const DefaultPageTitle = "Awesome go-echarts"
const DefaultEchartsAsset = "https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"

// Page represents a page which may contains one or more containers.
type Page struct {
	//Title the HTML title
	Title primitive.String

	JSAssets  *types.OrderedSet
	CSSAssets *types.OrderedSet

	Templates primitive.String

	Containers []*Container
}

type PageConfig func(p *Page)

// Render Default render
func (page *Page) Render(file string) {
	_ = (&DefaultRender{}).Render(file, page)
}

// CustomRender Render by the CustomRender
func (page *Page) CustomRender(file string, customRender Render) {
	_ = customRender.Render(file, page)
}

func NewPage(containers ...*Container) *Page {

	return &Page{
		Title:      DefaultPageTitle,
		JSAssets:   (&types.OrderedSet{}).Add(DefaultEchartsAsset),
		CSSAssets:  &types.OrderedSet{},
		Templates:  util.StrConv(templates.Tpl),
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
