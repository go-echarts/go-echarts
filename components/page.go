package components

import (
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/templates"
	"github.com/go-echarts/go-echarts/v2/types"
)

const DefaultPageTitle = "Awesome go-echarts"
const DefaultEchartsAsset = "https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"

// PageV3 represents a page chart.
type PageV3 struct {
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

func NewDefaultPage(containers ...*Container) *PageV3 {

	return &PageV3{
		Title:      DefaultPageTitle,
		JSAssets:   (&types.OrderedSet{}).Add(DefaultEchartsAsset),
		CSSAssets:  &types.OrderedSet{},
		Templates:  []string{templates.Tpl},
		Containers: containers,
	}

}

func (page *PageV3) AddCharts(chart ...interface{}) *PageV3 {
	// set page to chart
	return page
}

//func (page *PageV3) Render(file string) error {
//	// set page to chart
//	return page.RenderEngineer.Render(file, page)
//}
