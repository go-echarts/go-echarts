package render

import (
	"fmt"
	"html/template"
	"io"
	"regexp"
	"strings"

	"github.com/go-echarts/go-echarts/v2/types"
)

const (
	ModChart = "chart"
	ModPage  = "page"
	// EchartsInstancePrefix the default prefix for each echarts instance
	EchartsInstancePrefix = "goecharts_"
	// EchartsInstancePlaceholder a placeholder for types.FuncStr inject echarts instance
	EchartsInstancePlaceholder = "%MY_ECHARTS%"
)

var pat = regexp.MustCompile(`(__f__")|("__f__)|(__f__)`)

type ChartSnippet struct {
	Element string
	Script  string
	Option  string
}

// Renderer
// Any kinds of charts have their render implementation and
// you can define your own render logic easily.
type Renderer interface {
	Render(w io.Writer) error
	RenderContent() []byte
	RenderSnippet() ChartSnippet
}

// MustTemplate creates a new template with the given name and parsed contents.
func MustTemplate(name string, contents []string) *template.Template {
	tpl := template.New(name).Funcs(template.FuncMap{
		"safeHTML": func(s interface{}) template.HTML {
			return template.HTML(fmt.Sprint(s))
		},
		"safeJS": func(s interface{}) template.JS {
			return template.JS(fmt.Sprint(s))
		},
		"injectInstance": func(funcStr types.FuncStr, echartsInstancePlaceholder string, chartID string) string {
			instance := EchartsInstancePrefix + chartID
			return strings.Replace(string(funcStr), echartsInstancePlaceholder, instance, -1)
		},
	})
	tpl = template.Must(tpl.Parse(contents[0]))

	for _, cont := range contents[1:] {
		tpl = template.Must(tpl.Parse(cont))
	}
	return tpl
}
