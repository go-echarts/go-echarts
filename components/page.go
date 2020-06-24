package components

import (
	"io"

	"github.com/go-echarts/go-echarts/opts"
	"github.com/go-echarts/go-echarts/types"
)

type Charter interface {
	Type() string
	GetAssets() opts.Assets
	Validate()
}

// Page represents a page chart.
type Page struct {
	opts.Initialization
	opts.Assets

	Routers []opts.Router
	Charts  []interface{}

	unusedStr types.OrderedSet
}

// NewPage creates a new page.
func NewPage() *Page {
	page := &Page{}
	page.Assets.InitAssets()
	page.unusedStr.Init()
	return page
}

// Add adds new charts to the page.
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

// Render renders the chart and writes the output to given writers.
func (page *Page) Render(w io.Writer) error {
	page.Initialization.Validate()
	page.Assets.Validate(page.AssetsHost)
	return renderToWriter(page, ModPage, w, page.unusedStr.Values...)
}
