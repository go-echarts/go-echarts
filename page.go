package goecharts

import (
	"bytes"
	"github.com/gobuffalo/packr"
	"html/template"
	"io"
	"log"
)

type verifier interface {
	Validate()
}

type Page struct {
	InitOptions
	Charts []interface{}
}

//工厂函数，生成 `Bar` 实例
func NewPage() *Page {
	page := new(Page)
	page.InitOptions.setDefault()
	return page
}

// 新增 Page 图表，支持一次接收多个
func (page *Page) Add(charts ...verifier) *Page {
	if len(charts) < 1 {
		log.Println("Please make sure len(charts) > 0")
	}
	for i := 0; i < len(charts); i++ {
		charts[i].Validate()
		page.Charts = append(page.Charts, charts[i])

	}
	return page
}

func (page *Page) Render(w ...io.Writer) {
	page.setDefault()
	var b bytes.Buffer
	renderChart1(page, &b)
	res := replaceRender(b)
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}

// 渲染图表
func renderChart1(chart interface{}, w io.Writer) {
	box := packr.NewBox("./templates")
	htmlContent, err := box.FindString("page.html")
	t, err := template.New("").Parse(htmlContent)
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, chart)
}
