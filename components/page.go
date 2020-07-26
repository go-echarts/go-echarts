package components

import (
	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/render"
	"github.com/go-echarts/go-echarts/types"
)

// Charter
type Charter interface {
	Type() string
	GetAssets() opts.Assets
	Validate()
}

// Page represents a page chart.
type Page struct {
	render.Renderer

	opts.Initialization
	opts.Assets

	Charts []interface{}

	unusedStr types.OrderedSet
}

// NewPage creates a new page.
func NewPage() *Page {
	page := &Page{}
	page.Assets.InitAssets()
	page.unusedStr.Init()
	page.Renderer = render.NewPageRender(page, page.Validate)
	return page
}

// AddCharts adds new charts to the page.
func (page *Page) AddCharts(charts ...Charter) *Page {
	for i := 0; i < len(charts); i++ {
		assets := charts[i].GetAssets()
		for _, v := range assets.JSAssets.Values {
			page.JSAssets.Add(v)
		}

		for _, v := range assets.CSSAssets.Values {
			page.CSSAssets.Add(v)
		}
		charts[i].Validate()
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}

// Render renders the chart and writes the output to given writer.
func (page *Page) Validate() {
	page.Initialization.Validate()
	page.Assets.Validate(page.AssetsHost)
}
