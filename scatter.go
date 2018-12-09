package geocharts

type Scatter struct {
	InitOptions
	RectOptions
	Series []Series

	HasXYAxis bool
}

func NewScatter(opt InitOptions) *Scatter {
	scatter := new(Scatter)
	scatter.InitOptions = opt
	scatter.InitOptions.SetDefault()
	scatter.HasXYAxis = true
	return scatter
}

func (scatter *Scatter) setDefault() {
	scatter.InitOptions.SetDefault()
	scatter.BaseOptions.SetDefault()
}

func (scatter *Scatter) AddXAxis(xAxis interface{}) *Scatter {
	scatter.XAxisOptions.Data = xAxis
	return scatter
}

func (scatter *Scatter) AddYAxis(name string, yAxis interface{}) *Scatter {
	scatter.Series = append(scatter.Series, Series{Name: name, Type: SCATTER, Data: yAxis})
	return scatter
}

func (scatter *Scatter) SetGlobalConfig(options ...interface{}) *Scatter {
	scatter.RectOptions.setRectGlobalConfig(options...)
	return scatter
}

func (scatter *Scatter) Render(saveFile string) {
	for i := 0; i < len(scatter.Series); i++ {
		scatter.Series[i].LabelOptions = scatter.LabelOptions
	}
	RenderChart(scatter, saveFile)
}
