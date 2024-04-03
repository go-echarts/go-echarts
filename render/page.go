package render

import (
	"bytes"
	"io"

	"github.com/go-echarts/go-echarts/v2/templates"
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

func (r *pageRender) RenderSnippet() ChartSnippet {
	for _, fn := range r.before {
		fn()
	}

	elementTpl := MustTemplate(ModPage, []string{templates.BaseTpl, templates.BaseElementTpl, templates.BaseElementsTpl})

	var snippet ChartSnippet
	var elementBuf bytes.Buffer

	if err := elementTpl.ExecuteTemplate(&elementBuf, ModPage, r.c); err != nil {
		panic(err)
	}
	snippet.Element = string(pat.ReplaceAll(elementBuf.Bytes(), []byte("")))

	scriptTpl := MustTemplate(ModPage, []string{templates.BaseTpl, templates.BaseScriptTpl, templates.BaseScriptsTpl})
	var scriptBuf bytes.Buffer

	if err := scriptTpl.ExecuteTemplate(&scriptBuf, ModPage, r.c); err != nil {
		panic(err)
	}
	snippet.Script = string(pat.ReplaceAll(scriptBuf.Bytes(), []byte("")))

	return snippet
}
