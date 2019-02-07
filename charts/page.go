package charts

import (
	"io"
	"log"

	"github.com/chenjiandongx/go-echarts/common"
)

type charter interface {
	validateOpts()
	yieldAssets() ([]string, []string)
	chartType() string
}

type Page struct {
	InitOpts
	AssetsOpts
	Charts []interface{}
	HTTPRouters

	unusedStr common.OrderedSet
}

func NewPage(routers ...HTTPRouter) *Page {
	page := new(Page)
	for i := 0; i < len(routers); i++ {
		page.HTTPRouters = append(page.HTTPRouters, routers[i])
	}
	page.AssetsOpts.initAssetsOptsWithoutArg()
	page.unusedStr.InitWithoutArg()
	return page
}

// 新增 Page 图表，支持一次接收多个 Chart
func (page *Page) Add(charts ...charter) *Page {
	if len(charts) < 1 {
		log.Println("Charts length should > 0")
		return page
	}
	for i := 0; i < len(charts); i++ {
		charts[i].validateOpts()
		page.extractAssets(charts[i].yieldAssets())
		page.Charts = append(page.Charts, charts[i])

		if charts[i].chartType() == common.ChartType.Liquid {
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

func (page *Page) Render(w ...io.Writer) error {
	page.InitOpts.setDefault()
	return renderToWriter(page, "page", page.unusedStr.Values, w...)
}
