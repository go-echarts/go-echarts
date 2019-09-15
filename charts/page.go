package charts

import (
	"github.com/go-echarts/go-echarts/datatypes"
	"io"
	"log"
)

type charter interface {
	validateOpts()
	yieldAssets() ([]string, []string)
	chartType() string
}

// Page represents a page chart.
type Page struct {
	InitOpts
	AssetsOpts
	Charts []interface{}
	Routers

	unusedStr datatypes.OrderedSet
}

// NewPage creates a new page.
func NewPage(routers ...RouterOpts) *Page {
	page := new(Page)
	for i := 0; i < len(routers); i++ {
		page.Routers = append(page.Routers, routers[i])
	}
	page.AssetsOpts.initAssetsOptsWithoutArg()
	page.unusedStr.Init()
	return page
}

// Add adds new charts to the page.
func (page *Page) Add(charts ...charter) *Page {
	if len(charts) < 1 {
		log.Println("Charts length should > 0")
		return page
	}
	for i := 0; i < len(charts); i++ {
		charts[i].validateOpts()
		page.extractAssets(charts[i].yieldAssets())
		page.Charts = append(page.Charts, charts[i])

		if charts[i].chartType() == ChartType.Liquid {
			page.unusedStr.Add(`"outline":{"show":false},?`)
			page.unusedStr.Add(`"waveAnimation":false,?`)
		}
	}
	return page
}

func (page *Page) extractAssets(jsList, cssList []string) {
	for i := 0; i < len(jsList); i++ {
		page.JSAssets.Add(jsList[i])
	}
	for j := 0; j < len(cssList); j++ {
		page.CSSAssets.Add(cssList[j])
	}
}

// Render renders the chart and writes the output to given writers.
func (page *Page) Render(w ...io.Writer) error {
	page.InitOpts.setDefault()
	return renderToWriter(page, "page", page.unusedStr.Values, w...)
}
