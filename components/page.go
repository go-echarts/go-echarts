package components

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
	"github.com/go-echarts/go-echarts/v2/util"
)

type Layout string

const (
	PageNoneLayout   Layout = "none"
	PageCenterLayout Layout = "center"
	PageFlexLayout   Layout = "flex"
	PageFullLayout   Layout = "full"
)

// Charter represents a chart value which provides its type, assets and can be validated.
type Charter interface {
	Type() string
	GetAssets() opts.Assets
	// FillDefaultValues fill default values and would be overridden if any struct fields has been manually set
	FillDefaultValues()
	// Validate a validator as well as a post processor before render
	Validate()
}

// Page represents a page chart.
type Page struct {
	render.Renderer
	opts.PageConfiguration
	opts.Assets

	Charts []Charter
	Layout Layout
}

// NewPage creates a new page.
func NewPage() *Page {
	page := &Page{}
	page.Assets.InitAssets()
	page.Renderer = render.NewPageRender(page, page.Validate)
	page.Layout = PageCenterLayout
	return page
}

func (page *Page) SetPageTitle(title string) *Page {
	page.PageConfiguration.PageTitle = title
	return page
}

func (page *Page) SetAssetsHost(assetsHost string) *Page {
	page.PageConfiguration.AssetsHost = assetsHost
	return page
}

// SetLayout sets the layout of the Page.
func (page *Page) SetLayout(layout Layout) *Page {
	page.Layout = layout
	return page
}

// AddCharts adds new charts to the page and merge assets.
func (page *Page) AddCharts(charts ...Charter) *Page {
	for i := 0; i < len(charts); i++ {
		assets := charts[i].GetAssets()
		for _, v := range assets.JSAssets.Values {
			page.JSAssets.Add(v)
		}

		for _, v := range assets.CSSAssets.Values {
			page.CSSAssets.Add(v)
		}

		for _, v := range assets.CustomizedJSAssets.Values {
			page.CustomizedJSAssets.Add(v)
		}

		for _, v := range assets.CustomizedCSSAssets.Values {
			page.CustomizedCSSAssets.Add(v)
		}

		for _, v := range assets.CustomizedHeaders.Values {
			page.CustomizedHeaders.Add(v)
		}

		charts[i].Validate()
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}

// Validate validates the given configuration.
func (page *Page) Validate() {
	util.SetDefaultValue(page)
	page.Assets.Validate(page.AssetsHost)
}
