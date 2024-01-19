package render

import (
	"bytes"
	"io"

	tpls "github.com/go-echarts/go-echarts/v2/templates"
)

type pageRender struct {
	c      interface{}
	before []func()
}

// NewPageRender returns a render implementation for Page.
func NewPageRender(c interface{}, before ...func()) Renderer {
	return &pageRender{c: c, before: before}
}

// Render renders the page into the given io.Writer.
func (r *pageRender) Render(w io.Writer) error {
	for _, fn := range r.before {
		fn()
	}

	contents := []string{tpls.HeaderTpl, tpls.BaseTpl, tpls.PageTpl}
	tpl := MustTemplate(ModPage, contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, ModPage, r.c); err != nil {
		panic(err)
	}

	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := w.Write(content)
	return err
}
