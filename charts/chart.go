package charts

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/render"
)

type Chart interface {
	GetChart() interface{}
	GetChartName() string
	GetContainer() *components.Container
	GetPage() *components.PageV3
	Render(file string)
	// GetChartToJson() string
	// GetChartToMap() string
}

func doRender(file string, page *components.PageV3) {
	_ = (&render.RenderV3{}).Render(file, page)
}
