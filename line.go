package geocharts

import (
	"github.com/gobuffalo/packr"
	"html/template"
	"log"
	"os"
)

type LineOptions struct {
	BaseOptions
	RectOptions
}

func (opt *LineOptions) SetDefault() {
	opt.BaseOptions.SetDefault()
	opt.RectOptions.SetDefault()
}

type Line struct {
	InitOptions
	BarOptions
	XValue interface{}
	Series []Series
}

func NewLine(opt InitOptions) *Line {
	line := new(Line)
	line.InitOptions = opt
	line.InitOptions.SetDefault()
	return line
}

func (line *Line) AddXAxis(xAxis interface{}) *Line {
	line.XValue = xAxis
	return line
}

func (line *Line) AddYAxis(name string, yAxis interface{}) *Line {
	line.Series = append(line.Series, Series{name, "line", yAxis})
	return line
}

func (line *Line) SetConfig(options ...interface{}) *Line {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case LabelOptions:
			line.LabelOptions = option.(LabelOptions)
		case LegendOptions:
			line.LegendOptions = option.(LegendOptions)
		case XAxisOptions:
			line.XAxisOptions = option.(XAxisOptions)
		case YAxisOptions:
			line.YAxisOptions = option.(YAxisOptions)
		}
	}
	return line
}

func (line *Line) Render(saveFile string) {
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
	t.Execute(newFile, line)
}
