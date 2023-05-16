package charts

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/render"
)

type Chart interface {
	GetChart() interface{}
	GetChartName() string
	GetContainer() *components.Container
	GetPage() *components.Page
	Render(file string)
	// GetChartToJson()
	// GetChartToMap()
}

func doRender(file string, page *components.Page) {
	_ = (&render.Render{}).Render(file, page)
}
