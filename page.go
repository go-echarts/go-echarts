package goecharts

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

	HttpRouters
}

//工厂函数，生成 `Bar` 实例
func NewPage(routers ...HttpRouter) *Page {
	page := new(Page)
	for i := 0; i < len(routers); i++ {
		page.HttpRouters = append(page.HttpRouters, routers[i])
	}
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
		jsList, cssList := charts[i].yieldAssets()
		page.extractJSAssets(jsList)
		page.extractCSSAssets(cssList)
		page.Charts = append(page.Charts, charts[i])
	}
	return page
}

// 提取 js 引用
func (page *Page) extractJSAssets(jsList []string) {
	for i := 0; i < len(jsList); i++ {
		if !page.AssetsOpts.jsIn(jsList[i]) {
			page.JSAssets = append(page.JSAssets, jsList[i])
		}
	}
}

// 提取 css 引用
func (page *Page) extractCSSAssets(cssList []string) {
	for i := 0; i < len(cssList); i++ {
		if !page.AssetsOpts.cssIn(cssList[i]) {
			page.CSSAssets = append(page.CSSAssets, cssList[i])
		}
	}
}

func (page *Page) Render(w ...io.Writer) error {
	page.InitOpts.setDefault()
	if err := renderToWriter(page, "page", w...); err != nil {
		return err
	}
	return nil
}
