package goecharts

import (
	"io"
	"log"
	"strings"
)

type validator interface {
	validateOpts()
	yieldAssets() ([]string, []string)
}

type Page struct {
	InitOpts
	AssetsOpts
	Charts []interface{}

	HttpRouters
}

//工厂函数，生成 `Bar` 实例
func NewPage(routers ...HttpRouter) *Page {
	page := new(Page)
	for i := 0; i < len(routers); i++ {
		page.HttpRouters = append(page.HttpRouters, routers[i])
	}
	page.initAssetsOpts()
	return page
}

// 新增 Page 图表，支持一次接收多个 Chart
func (page *Page) Add(charts ...validator) *Page {
	if len(charts) < 1 {
		log.Println("Charts should length > 0")
	}
	for i := 0; i < len(charts); i++ {
		charts[i].validateOpts()
		jsList, cssList := charts[i].yieldAssets()
		page.extractJSAssets(jsList)
		page.extractCSSAssets(cssList)
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}

func (page *Page) extractJSAssets(jsList []string) {
	for i := 0; i < len(jsList); i++ {
		js := strings.Split(jsList[i], "/")
		if len(js) < 1 {
			continue
		}
		if !page.AssetsOpts.jsIn(js[len(js)-1]) {
			page.JSAssets = append(page.JSAssets, js[len(js)-1])
		}
	}
}

func (page *Page) extractCSSAssets(cssList []string) {
	for i := 0; i < len(cssList); i++ {
		css := strings.Split(cssList[i], "/")
		if len(css) < 1 {
			continue
		}
		if !page.AssetsOpts.cssIn(css[len(css)-1]) {
			page.CSSAssets = append(page.CSSAssets, css[len(css)-1])
		}
	}
}

// 渲染图表，支持多 io.Writer
func (page *Page) Render(w ...io.Writer) {
	page.InitOpts.setDefault()
	page.validateAssets(page.AssetsHost)
	renderToWriter(page, "page", w...)
}
