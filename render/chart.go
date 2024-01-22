package render

import (
	"bytes"
	"io"

	"github.com/go-echarts/go-echarts/v2/templates"
)

type chartRender struct {
	c      interface{}
	before []func()
}

// NewChartRender returns a render implementation for Chart.
func NewChartRender(c interface{}, before ...func()) Renderer {
	return &chartRender{c: c, before: before}
}

// Render renders the chart into the given io.Writer.
func (r *chartRender) Render(w io.Writer) error {

	content := r.RenderContent()
	_, err := w.Write(content)
	return err
}

func (r *chartRender) RenderContent() []byte {
	for _, fn := range r.before {
		fn()
	}

	contents := []string{templates.HeaderTpl, templates.BaseTpl, templates.ChartTpl}
	tpl := MustTemplate(ModChart, contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, ModChart, r.c); err != nil {
		panic(err)
	}

	return pat.ReplaceAll(buf.Bytes(), []byte(""))
}
