package page

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/component"
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/templates"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/go-echarts/go-echarts/v2/util"
	"regexp"
)

const DefaultPageTitle = "Awesome go-echarts"
const DefaultEchartsAsset = "https://go-echarts.github.io/go-echarts-assets/assets/echarts.min.js"
const FuncMarker = "__f__"

var newlineTabPat = regexp.MustCompile(`\n|\t`)
var commentPat = regexp.MustCompile(`(//.*)\n`)

// Page represents a page which may contains one or more containers.
type Page struct {
	//Title the HTML title
	Title primitive.String

	JSAssets  *types.OrderedSet
	CSSAssets *types.OrderedSet

	Templates primitive.String

	Containers []*component.Container

	JSFunctions JSFunctions

	// Global styles, CSSAssets should be better
	// Style

}

type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, newlineTabPat.ReplaceAllString(fn[i], ""))
	}
}

// FuncOpts returns a string suitable for options expecting JavaScript code.
func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
}

// FuncStripCommentsOpts returns a string suitable for options expecting JavaScript code,
// stripping '//' comments.
func FuncStripCommentsOpts(fn string) string {
	fn = commentPat.ReplaceAllString(fn, "")
	return replaceJsFuncs(fn)
}

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	fn = newlineTabPat.ReplaceAllString(fn, "")
	return fmt.Sprintf("%s%s%s", FuncMarker, fn, FuncMarker)
}

type Config func(p *Page)

func NewPage(containers ...*component.Container) *Page {

	return &Page{
		Title:      DefaultPageTitle,
		JSAssets:   (&types.OrderedSet{}).Add(DefaultEchartsAsset),
		CSSAssets:  &types.OrderedSet{},
		Templates:  util.StrConv(templates.Tpl),
		Containers: containers,
	}
}

func (page *Page) AddCharts(charts ...component.Chart) *Page {

	for _, c := range charts {
		page.Containers = append(page.Containers, c.GetContainer())
	}
	return page
}

func (page *Page) Config(configs ...Config) *Page {

	for _, fn := range configs {
		fn(page)
	}
	return page
}

func New(config Config) *Page {
	page := NewPage()
	config(page)
	return page
}

func NewWithTemplate(tpl TemplateProvider, config Config) *Page {
	page := tpl.Provide()
	config(page)
	return page
}
