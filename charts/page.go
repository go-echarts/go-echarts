package charts

import (
	"io"
	"log"
)

// 校验器接口
type validator interface {
	validateOpts()
	yieldAssets() ([]string, []string)
}

type Page struct {
	InitOpts
	AssetsOpts
	Charts []interface{}

	HTTPRouters
}

// 工厂函数，生成 `Bar` 实例
func NewPage(routers ...HTTPRouter) *Page {
	page := new(Page)
	for i := 0; i < len(routers); i++ {
		page.HTTPRouters = append(page.HTTPRouters, routers[i])
	}
	page.AssetsOpts.initAssetsOptsWithoutArg()
	return page
}

// 新增 Page 图表，支持一次接收多个 Chart
func (page *Page) Add(charts ...validator) *Page {
	if len(charts) < 1 {
		log.Println("Charts should length > 0")
		return page
	}
	for i := 0; i < len(charts); i++ {
		charts[i].validateOpts()
		page.extractAssets(charts[i].yieldAssets())
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}

// 提取引用资源
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
	return renderToWriter(page, "page", w...)
}
