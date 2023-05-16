package render

import (
	"bytes"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/components"
	"html/template"
	"os"
	"reflect"
	"regexp"
)

var (
	pat = regexp.MustCompile(`(__f__")|("__f__)|(__f__)`)
)

type RenderV3 struct {
}

func (r *RenderV3) Render(file string, page *components.PageV3) error {

	f, _ := os.Create(file)
	tpl := MustTemplate("chart", page.Templates)

	var buf bytes.Buffer

	if err := tpl.ExecuteTemplate(&buf, "chart", page); err != nil {
		return err
	}

	content := pat.ReplaceAll(buf.Bytes(), []byte(""))

	_, err := f.Write(content)
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
	tpl := template.Must(template.New(name).Funcs(template.FuncMap{
		"safeJS": func(s interface{}) template.JS {
			return template.JS(fmt.Sprint(s))
		},
		"isSet": isSet,
	}).Parse(contents[0]))

	//for _, cont := range contents[1:] {
	//	tpl = template.Must(tpl.Parse(cont))
	//}
	return tpl
}
