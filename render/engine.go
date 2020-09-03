package render

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"regexp"

	tpls "github.com/go-echarts/go-echarts/templates"
)

type Renderer interface {
	Render(w io.Writer) error
}

const (
	ModChart = "chart"
	ModPage  = "page"
)

type pageRender struct {
	c      interface{}
	before []func()
}

// NewPageRender
func NewPageRender(c interface{}, before ...func()) Renderer {
	return &pageRender{c: c, before: before}
}

// Render
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

	pat, _ := regexp.Compile(`(__x__")|("__x__)|(__x__)`)
	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := w.Write(content)
	return err
}

type chartRender struct {
	c      interface{}
	before []func()
}

// NewChartRender
func NewChartRender(c interface{}, before ...func()) Renderer {
	return &chartRender{c: c, before: before}
}

// Render
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

	pat, _ := regexp.Compile(`(__x__")|("__x__)`)
	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := w.Write(content)
	return err
}

// MustTemplate
func MustTemplate(name string, contents []string) *template.Template {
	tpl := template.Must(template.New(name).Parse(contents[0])).Funcs(template.FuncMap{
		"safeJS": func(s interface{}) template.JS {
			return template.JS(fmt.Sprint(s))
		},
	})
	for _, cont := range contents[1:] {
		tpl = template.Must(tpl.Parse(cont))
	}
	return tpl
}

// todo: Is it necessary?
func replaceRender(b bytes.Buffer, notReplace ...string) []byte {
	// set `__x__` as placeholder.
	idPat, _ := regexp.Compile(`(__x__")|("__x__)`)
	content := idPat.ReplaceAllString(b.String(), "")
	unusedObj := []string{
		`geo: {},?`,
		`"normal":{},?`,
		`"textStyle":{},?`,
		`"subtextStyle":{},?`,
		`"inRange":{},?`,
		`"label":{},?`,
		`"markLine":{},?`,
		`"markPoint":{},?`,
		`"itemStyle":{},?`,
		`"areaStyle":{},?`,
		`"lineStyle":{},?`,
		`"rippleEffect":{},?`,
		`"splitArea":{},?`,
		`"outline":{"show":false},?`,
		`"waveAnimation":false,?`,
		`"viewControl":{},?`,
		`"force":{},?`,
	}
	unusedObj = removeNotReplace(unusedObj, notReplace...)
	// remove unused JSON object.
	var unusedPat string
	for i := 0; i < len(unusedObj); i++ {
		unusedPat += unusedObj[i] + "|"
	}
	p, _ := regexp.Compile(unusedPat)
	res := p.ReplaceAllString(content, "")
	return []byte(res)
}

// todo: Is it necessary?
func removeNotReplace(unusedObj []string, removeStr ...string) []string {
	res := make([]string, 0)
	for i := 0; i < len(unusedObj); i++ {
		isRemove := false
		for j := 0; j < len(removeStr); j++ {
			if unusedObj[i] == removeStr[j] {
				isRemove = true
				break
			}
		}
		if !isRemove {
			res = append(res, unusedObj[i])
		}
	}
	return res
}
