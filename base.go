package geocharts

import (
	"github.com/gobuffalo/packr"
	"html/template"
	"log"
	"os"
)

type Series struct {
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	Data         interface{} `json:"data"`
	LabelOptions `json:"label"`
}

type InitOptions struct {
	PageTitle   string `default:"Awesome go-echarts"`
	Width       string `default:"800px"`
	Height      string `default:"500px"`
	ContainerID string `default:"main"`
	TitleOptions
}

func (opt *InitOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type BaseOptions struct {
	InitOptions
	LabelOptions
	LegendOptions
	TooltipOptions
}

func (opt *BaseOptions) SetDefault() {
	opt.InitOptions.SetDefault()
	opt.LabelOptions.SetDefault()
	opt.LegendOptions.SetDefault()
	opt.TooltipOptions.SetDefault()
}

type XYOptions struct {
	XAxisOptions
	YAxisOptions
}

func (opt *XYOptions) SetDefault() {
	opt.XAxisOptions.SetDefault()
	opt.YAxisOptions.SetDefault()
}

type LegendOptions struct {
	Show   bool   `json:"show" default:"true"`
	Left   string `json:"left,omitempty"`
	Top    string `json:"top,omitempty"`
	Right  string `json:"right,omitempty"`
	Bottom string `json:"bottom,omitempty"`
}

func (opt *LegendOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type LabelOptions struct {
	Show     bool   `json:"show" default:"false"`
	Position string `json:"position" default:"top"`
	Color    string `json:"color"`
}

func (opt *LabelOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type TooltipOptions struct {
	Show    bool   `json:"show" default:"true" `
	Trigger string `json:"trigger" default:"item"`
}

func (opt *TooltipOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type XAxisOptions struct {
	Name string      `json:"name"`
	Show bool        `json:"show" default:"true"`
	Data interface{} `json:"data"`
}

func (opt *XAxisOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type YAxisOptions struct {
	Name string
	Show bool `default:"true"`
}

func (opt *YAxisOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

func RenderChart(chart interface{}, saveFile string) {
	box := packr.NewBox("./templates")
	htmlContent, err := box.FindString("index.html")
	t, err := template.New("").Parse(htmlContent)
	if err != nil {
		log.Println(err)
	}
	newFile, err := os.Create(saveFile)
	if err != nil {
		log.Println(err)
	}
	t.Execute(newFile, chart)
}

type RectOptions struct {
	BaseOptions
	XYOptions
}

func (rect *RectOptions) SetDefault() {
	rect.BaseOptions.SetDefault()
	rect.XAxisOptions.SetDefault()
}

func (rect *RectOptions) setRectGlobalConfig(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case InitOptions:
			rect.InitOptions = option.(InitOptions)
		case LegendOptions:
			rect.LegendOptions = option.(LegendOptions)
		case XAxisOptions:
			rect.XAxisOptions = option.(XAxisOptions)
		case YAxisOptions:
			rect.YAxisOptions = option.(YAxisOptions)
		}
	}
}

type TitleOptions struct {
	Title         string `json:"text"`
	TitleStyle    TextStyle
	Subtitle      string `json:"subtext"`
	SubtitleStyle TextStyle
	Link          string `json:"link"`
	Target        string `json:"target"`
	Top           string `json:"top"`
	Bottom        string `json:"bottom"`
	Left          string `json:"left"`
	Right         string `json:"right"`
}

type TextStyle struct {
	Color     string `json:"color"`
	FontStyle string `json:"fontStyle"`
	FontSize  string `json:"fontSize"`
}

type SeriesOptions struct {
	LabelOptions
}