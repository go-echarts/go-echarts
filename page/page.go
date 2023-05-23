package page

import "github.com/go-echarts/go-echarts/v2/core"

func NewWithTemplate(tpl core.PageTemplateProvider, config core.PageConfig) *core.Page {
	page := tpl.Provide()
	config(page)
	return page
}

func New(config core.PageConfig) *core.Page {
	page := core.NewPage()
	config(page)
	return page

}
