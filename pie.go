package geocharts

type Pie struct {
	InitOptions
	BaseOptions
	Series []Series

	HasXYAxis bool
}

func NewPie(opt InitOptions) *Pie {
	pie := new(Pie)
	pie.InitOptions = opt
	pie.setDefault()
	pie.HasXYAxis = false
	return pie
}

func (pie *Pie) setDefault() {
	pie.InitOptions.SetDefault()
	pie.BaseOptions.SetDefault()
}

func (pie *Pie) Add(name string, data map[string]interface{}) {
	type pieData struct {
		Name  string      `json:"name"`
		Value interface{} `json:"value"`
	}
	pd := make([]pieData, 0)
	for k, v := range data {
		pd = append(pd, pieData{k, v})
	}
	pie.Series = append(pie.Series, Series{Name: name, Type: PIE, Data: pd})
}

func (pie *Pie) Render(saveFile string) {
	for i := 0; i < len(pie.Series); i++ {
		pie.Series[i].LabelOptions = pie.LabelOptions
	}
	RenderChart(pie, saveFile)
}
