package geocharts

type Bar struct {
	InitOptions
	RectOptions
	Series []Series

	HasXYAxis bool
	xAxisData interface{}
}

func NewBar() *Bar {
	bar := new(Bar)
	bar.setDefault()
	bar.HasXYAxis = true
	return bar
}

func (bar *Bar) setDefault() {
	bar.InitOptions.SetDefault()
	bar.RectOptions.SetDefault()
}

func (bar *Bar) AddXAxis(xAxis interface{}) *Bar {
	bar.xAxisData = xAxis
	return bar
}

func (bar *Bar) AddYAxis(name string, yAxis interface{}, options ...interface{}) *Bar {
	series := Series{Name: name, Type: BAR, Data: yAxis}
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case LabelOptions:
			series.LabelOptions = option.(LabelOptions)
		}
	}
	bar.Series = append(bar.Series, series)
	return bar
}

func (bar *Bar) SetGlobalConfig(options ...interface{}) *Bar {
	bar.RectOptions.setRectGlobalConfig(options...)
	return bar
}

func (bar *Bar) SetSeriesConfig(options ...interface{}) *Bar {
	for i := 0; i < len(bar.Series); i++ {
		for j := 0; j < len(options); j++ {
			option := options[j]
			switch option.(type) {
			case LabelOptions:
				bar.Series[i].LabelOptions = option.(LabelOptions)
			}
		}
	}
	return bar
}

func (bar *Bar) Render(saveFile string) {
	bar.XAxisOptions.Data = bar.xAxisData
	RenderChart(bar, saveFile)
}
