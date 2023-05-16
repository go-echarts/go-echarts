package components

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

	Templates []string
	// Containers, one - one
	Containers []*Container
}

func NewDefaultPage(containers ...*Container) *Page {

	return &Page{
		Title:      DefaultPageTitle,
		JSAssets:   (&types.OrderedSet{}).Add(DefaultEchartsAsset),
		CSSAssets:  &types.OrderedSet{},
		Templates:  []string{templates.Tpl},
		Containers: containers,
	}

}

func (page *Page) AddCharts(chart ...interface{}) *Page {
	// TODO set page to chart in same page
	return page
}
