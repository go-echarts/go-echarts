package geocharts

type Line struct {
	InitOptions
	RectOptions
	Series []Series

	HasXYAxis bool
}

func NewLine(opt InitOptions) *Line {
	line := new(Line)
	line.InitOptions = opt
	line.setDefault()
	line.HasXYAxis = true
	return line
}

func (line *Line) setDefault() {
	line.InitOptions.SetDefault()
	line.BaseOptions.SetDefault()
}

func (line *Line) AddXAxis(xAxis interface{}) *Line {
	line.XAxisOptions.Data = xAxis
	return line
}

func (line *Line) AddYAxis(name string, yAxis interface{}) *Line {
	line.Series = append(line.Series, Series{Name:name, Type:LINE, Data:yAxis})
	return line
}

func (line *Line) SetGlobalConfig(options ...interface{}) *Line{
	line.RectOptions.setRectGlobalConfig(options...)
	return line
}

func (line *Line) Render(saveFile string) {
	for i := 0; i < len(line.Series); i++ {
		line.Series[i].LabelOptions = line.LabelOptions
	}
	RenderChart(line, saveFile)
}
