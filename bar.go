package geocharts

import (
	"github.com/gobuffalo/packr"
	"html/template"
	"log"
	"os"
)

type BarOptions struct {
	BaseOptions
	RectOptions
}

func (opt *BarOptions) SetDefault() {
	opt.BaseOptions.SetDefault()
	opt.RectOptions.SetDefault()
}

type Bar struct {
	InitOptions
	BarOptions
	XValue interface{}
	Series []Series
}

func NewBar(opt InitOptions) *Bar {
	bar := new(Bar)
	bar.InitOptions = opt
	bar.InitOptions.SetDefault()
	return bar
}

func (bar *Bar) AddXAxis(xAxis interface{}) *Bar {
	bar.XValue = xAxis
	return bar
}

func (bar *Bar) AddYAxis(name string, yAxis interface{}) *Bar {
	bar.Series = append(bar.Series, Series{name, "bar", yAxis})
	return bar
}

func (bar *Bar) SetConfig(options ...interface{}) *Bar {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case LabelOptions:
			bar.LabelOptions = option.(LabelOptions)
		case LegendOptions:
			bar.LegendOptions = option.(LegendOptions)
		case XAxisOptions:
			bar.XAxisOptions = option.(XAxisOptions)
		case YAxisOptions:
			bar.YAxisOptions = option.(YAxisOptions)
		}
	}
	return bar
}

func (bar *Bar) Render(saveFile string) {
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
	t.Execute(newFile, bar)
}
