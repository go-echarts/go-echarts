package charts

import (
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/opt"
	"github.com/go-echarts/go-echarts/v2/series"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Line3DConfiguration struct {
	*opt.BaseConfiguration
	Series  *series.Line3DSeries0 `json:"series,omitempty,reserved"`
	XAxis3D *opt.XAxis3D          `json:"xAxis3D,omitempty,reserved"`
	YAxis3D *opt.YAxis3D          `json:"yAxis3D,omitempty,reserved"`
	ZAxis3D *opt.ZAxis3D          `json:"zAxis3D,omitempty,reserved"`
	Grid3D  *opt.Grid3D           `json:"grid3D,omitempty,reserved"`
}

// Line3D represents a line3D chart.
type Line3D struct {
	*Line3DConfiguration
}

func (line3D *Line3D) GetChartName() string {
	return types.ChartLine3D
}

func (line3D *Line3D) GetChart() interface{} {
	return line3D
}

func (line3D *Line3D) GetContainer() *core.Container {
	if line3D.Container != nil {
		return line3D.Container
	}

	line3D.Container = core.NewContainer(line3D)
	return line3D.Container

}

func (line3D *Line3D) GetPage() *core.Page {
	if line3D.Page != nil {
		return line3D.Page
	}

	line3D.Page = core.NewPage(line3D.GetContainer())
	return line3D.Page
}

func NewLine3D() *Line3D {
	line3D := &Line3D{}

	line3D.Line3DConfiguration = &Line3DConfiguration{
		BaseConfiguration: opt.NewBaseConfiguration(),
		XAxis3D:           &opt.XAxis3D{},
		YAxis3D:           &opt.YAxis3D{},
		ZAxis3D:           &opt.ZAxis3D{},
		Series:            &series.Line3DSeries0{},
	}

	line3D.Container = core.NewContainer(line3D)
	line3D.Page = core.NewPage(line3D.Container)
	line3D.Page.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts-gl/dist/echarts-gl.min.js")

	return line3D
}

func (line3D *Line3D) AddSeries(series ...*series.Line3DSeries) {
	for _, s := range series {
		s.Type = types.ChartLine3D
		*line3D.Series = append(*line3D.Series, s)
	}
}
