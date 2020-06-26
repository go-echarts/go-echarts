package render

import (
	"bytes"
	"html/template"
	"io"
	"regexp"

	tpls "github.com/go-echarts/go-echarts/templates"
)

type Mod int

const (
	ModChart Mod = iota
	ModPage
)

// 渲染图表
func renderChart(chart interface{}, w io.Writer, mod Mod) error {
	contents := []string{tpls.HeaderTpl, tpls.RoutersTpl, tpls.BaseTpl}
	switch mod {
	case ModChart:
		contents = append(contents, tpls.ChartTpl)
	case ModPage:
		contents = append(contents, tpls.PageTpl)
	}
	tpl := template.Must(template.New("page").Parse(contents[0]))
	mustTpl(tpl, contents[1:]...)
	return tpl.ExecuteTemplate(w, "page", chart)
}

func mustTpl(tpl *template.Template, html ...string) {
	for i := 0; i < len(html); i++ {
		tpl = template.Must(tpl.Parse(html[i]))
	}
}

// 过滤替换渲染结果
func replaceRender(b bytes.Buffer, notReplace ...string) []byte {
	// __x__ 与模板占位符相匹配
	idPat, _ := regexp.Compile(`(__x__")|("__x__)`)
	// 替换并转为 []byte 类型
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
	// 移除无用的 JSON object
	// 另一种解决方案是使用 *struct
	var unusedPat string
	for i := 0; i < len(unusedObj); i++ {
		unusedPat += unusedObj[i] + "|"
	}
	p, _ := regexp.Compile(unusedPat)
	res := p.ReplaceAllString(content, "")
	return []byte(res)
}

// 针对某些图表移除
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

func render(chart interface{}, mod Mod, w io.Writer, removeStr ...string) error {
	var b bytes.Buffer
	if err := renderChart(chart, &b, mod); err != nil {
		return err
	}
	res := replaceRender(b, removeStr...)
	_, err := w.Write(res)
	return err
}

func ChartRender(chart interface{}, w io.Writer, removeStr ...string) error {
	return render(chart, ModChart, w, removeStr...)
}

func PageRender(chart interface{}, w io.Writer, removeStr ...string) error {
	return render(chart, ModPage, w, removeStr...)
}
