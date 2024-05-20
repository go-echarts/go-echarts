package render

import (
	"bytes"
	"html"
	"io"

	"github.com/go-echarts/go-echarts/v2/templates"
)

type chartRender struct {
	BaseRender
	c interface{}
	// before the pre-process functions for chart, it only calls once to support multi renders
	before []func()
}

// NewChartRender returns a render implementation for Chart.
func NewChartRender(c interface{}, before ...func()) Renderer {
	return &chartRender{c: c, before: before}
}

// Render renders the chart(s) into the given io.Writer.
func (r *chartRender) Render(w io.Writer) error {
	content := r.RenderContent()
	_, err := w.Write(content)
	return err
}

func (r *chartRender) RenderContent() []byte {
	for _, fn := range r.before {
		fn()
	}

	r.before = []func(){}

	contents := []string{templates.HeaderTpl, templates.BaseTpl, templates.ChartTpl}
	tpl := MustTemplate(ModChart, contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, ModChart, r.c); err != nil {
		panic(err)
	}

	return pat.ReplaceAll(buf.Bytes(), []byte(""))
}

func (r *chartRender) RenderSnippet() ChartSnippet {
	for _, fn := range r.before {
		fn()
	}

	r.before = []func(){}

	var snippet ChartSnippet

	elementTpl := MustTemplate(ModChart, []string{templates.BaseTpl, templates.BaseElementTpl})
	var elementBuf bytes.Buffer
	if err := elementTpl.ExecuteTemplate(&elementBuf, ModChart, r.c); err != nil {
		panic(err)
	}
	snippet.Element = string(pat.ReplaceAll(elementBuf.Bytes(), []byte("")))

	scriptTpl := MustTemplate(ModChart, []string{templates.BaseTpl, templates.BaseScriptTpl})
	var scriptBuf bytes.Buffer
	if err := scriptTpl.ExecuteTemplate(&scriptBuf, ModChart, r.c); err != nil {
		panic(err)
	}
	snippet.Script = string(pat.ReplaceAll(scriptBuf.Bytes(), []byte("")))

	optionTpl := MustTemplate(ModChart, []string{templates.BaseTpl, templates.BaseOptionTpl})
	var optionBuf bytes.Buffer
	if err := optionTpl.ExecuteTemplate(&optionBuf, ModChart, r.c); err != nil {
		panic(err)
	}
	snippet.Option = html.UnescapeString(string(pat.ReplaceAll(optionBuf.Bytes(), []byte(""))))

	return snippet
}
