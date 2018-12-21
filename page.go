package goecharts

import (
	"bytes"
	"io"
	"log"
)

type verifier interface {
	verifyOpts()
}

type Page struct {
	InitOpts
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

// 新增 Page 图表，支持一次接收多个
func (page *Page) Add(charts ...verifier) *Page {
	if len(charts) < 1 {
		log.Println("Please make sure len(charts) > 0")
	}
	for i := 0; i < len(charts); i++ {
		charts[i].verifyOpts()
		page.Charts = append(page.Charts, charts[i])

	}
	return page
}

// 渲染图表，支持多 io.Writer
func (page *Page) Render(w ...io.Writer) {
	page.InitOpts.setDefault()
	var b bytes.Buffer
	renderChart(page, &b, "page")
	res := replaceRender(b)
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}
