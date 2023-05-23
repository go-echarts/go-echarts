package core

import (
	"bytes"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/util"
	"html/template"
	"reflect"
	"regexp"
)

var (
	pat = regexp.MustCompile(`(__f__")|("__f__)|(__f__)`)
)

const DefaultTplName = "__GO_ECHARTS__"

func (r *Renderer) Render(page *Page) []byte {

	tpl := MustTemplate(DefaultTplName, page.Templates)

	var buf bytes.Buffer

	if err := tpl.ExecuteTemplate(&buf, DefaultTplName, page); err != nil {
		panic(err)
	}

	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	return content
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

func prettier() func(chart interface{}) interface{} {
	return func(c interface{}) interface{} {
		util.Prettier(c)
		return c
	}
}

// MustTemplate creates a new template with the given name and parsed contents.
func MustTemplate(name string, contents primitive.String) *template.Template {
	tpl := template.Must(template.New(name).Funcs(template.FuncMap{
		"safeJS": func(s interface{}) template.JS {
			return template.JS(fmt.Sprint(s))
		},
		"isSet":    isSet,
		"prettier": prettier(),
	}).Parse(contents.StringVal()))

	return tpl
}
