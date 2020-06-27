package render

import (
	"bytes"
	"html/template"
	"io"
	"regexp"

	tpls "github.com/go-echarts/go-echarts/templates"
)

const (
	ModChart = "chart"
	ModPage  = "page"
)

func renderChart(chart interface{}, w io.Writer, mod string) error {
	contents := []string{tpls.HeaderTpl, tpls.RoutersTpl, tpls.BaseTpl}
	switch mod {
	case ModChart:
		contents = append(contents, tpls.ChartTpl)
	case ModPage:
		contents = append(contents, tpls.PageTpl)
	}

	tpl := template.Must(template.New(mod).Parse(contents[0]))
	for _, cont := range contents[1:] {
		tpl = template.Must(tpl.Parse(cont))
	}
	return tpl.ExecuteTemplate(w, mod, chart)
}

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

func render(chart interface{}, mod string, w io.Writer, removeStr ...string) error {
	var b bytes.Buffer
	if err := renderChart(chart, &b, mod); err != nil {
		return err
	}
	res := replaceRender(b, removeStr...)
	_, err := w.Write(res)
	return err
}

// ChartRender renders the Chart types.
func ChartRender(chart interface{}, w io.Writer, removeStr ...string) error {
	return render(chart, ModChart, w, removeStr...)
}

// PageRender renders the Page component.
func PageRender(chart interface{}, w io.Writer, removeStr ...string) error {
	return render(chart, ModPage, w, removeStr...)
}
