package render

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"reflect"
	"regexp"

	tpls "github.com/go-echarts/go-echarts/v2/templates"
)

// Renderer
// Any kinds of charts have their render implementation and
// you can define your own render logic easily.
type Renderer interface {
	Render(w io.Writer) error
}

const (
	ModChart = "chart"
	ModPage  = "page"
)

var (
	pat = regexp.MustCompile(`(__f__")|("__f__)|(__f__)`)
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
		return err
	}

	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := w.Write(content)
	return err
}

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
	for _, fn := range r.before {
		fn()
	}

	contents := []string{tpls.HeaderTpl, tpls.BaseTpl, tpls.ChartTpl}
	tpl := MustTemplate(ModChart, contents)

	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, ModChart, r.c); err != nil {
		return err
	}

	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := w.Write(content)
	return err
}

// isSet check if the field exist in the chart instance
// Shamed copy from https://stackoverflow.com/questions/44675087/golang-template-variable-isset
func isSet(name string, data interface{}) bool {
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return false
	}

	return v.FieldByName(name).IsValid()
}

// MustTemplate creates a new template with the given name and parsed contents.
func MustTemplate(name string, contents []string) *template.Template {
	tpl := template.Must(template.New(name).Parse(contents[0])).Funcs(template.FuncMap{
		"safeJS": func(s interface{}) template.JS {
			return template.JS(fmt.Sprint(s))
		},
		"isSet": isSet,
	})

	for _, cont := range contents[1:] {
		tpl = template.Must(tpl.Parse(cont))
	}
	return tpl
}
