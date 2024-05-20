package render

import (
	"bytes"
	"io"

	"github.com/go-echarts/go-echarts/v2/templates"
)

type pageRender struct {
	BaseRender
	c      interface{}
	before []func()
}

// NewPageRender returns a render implementation for Page.
func NewPageRender(c interface{}, before ...func()) Renderer {
	return &pageRender{c: c, before: before}
}

// Render renders the chart(s) into the given io.Writer.
func (r *pageRender) Render(w io.Writer) error {
	content := r.RenderContent()
	_, err := w.Write(content)
	return err
}

func (r *pageRender) RenderContent() []byte {
	for _, fn := range r.before {
		fn()
	}

	contents := []string{templates.HeaderTpl, templates.BaseTpl, templates.PageTpl}
	tpl := MustTemplate(ModPage, contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, ModPage, r.c); err != nil {
		panic(err)
	}

	return pat.ReplaceAll(buf.Bytes(), []byte(""))
}
