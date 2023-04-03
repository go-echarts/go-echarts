package components

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

type Layout string

const (
	PageNoneLayout   Layout = "none"
	PageCenterLayout Layout = "center"
	PageFlexLayout   Layout = "flex"
)

// Charter represents a chart value which provides its type, assets and can be validated.
type Charter interface {
	Type() string
	GetAssets() opts.Assets
	FillDefaultValues()
	Validate()
	RegisterMux(mux ...*http.ServeMux)
}

// Page represents a page chart.
type Page struct {
	render.Renderer
	opts.Initialization
	opts.Assets

	Charts []Charter
	Layout Layout
	Mux    *http.ServeMux
}

// NewPage creates a new page.
func NewPage() *Page {
	page := &Page{}
	page.Assets.InitAssets()
	page.Renderer = render.NewPageRender(page, page.Validate,page.ValidateMux)
	page.Layout = PageCenterLayout
	return page
}

// SetLayout sets the layout of the Page.
func (page *Page) SetLayout(layout Layout) *Page {
	page.Layout = layout
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

// Validate validates the given configuration.
func (page *Page) Validate() {
	page.Initialization.Validate()
	page.Assets.Validate(page.AssetsHost)
}

func (page *Page) ValidateMux() {
	for _, v := range page.Charts {
		v.RegisterMux(page.Mux)
	}
}
