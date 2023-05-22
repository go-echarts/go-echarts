package charts

import (
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/opt"
	"github.com/go-echarts/go-echarts/v2/series"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Line3DConfiguration struct {
	*opt.BaseConfiguration
	Series  *series.Line3DSeries `json:"series,omitempty,reserved"`
	XAxis3D *opt.XAxis3D         `json:"xAxis3D,omitempty,reserved"`
	YAxis3D *opt.YAxis3D         `json:"yAxis3D,omitempty,reserved"`
	ZAxis3D *opt.ZAxis3D         `json:"zAxis3D,omitempty,reserved"`
	Grid3D  *opt.Grid3D          `json:"grid3D,omitempty,reserved"`
}

// Line3D represents a line3D chart.
type Line3D struct {
	*Line3DConfiguration
}

func (line3D *Line3D) GetChartType() string {
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

func NewLine3D() *Line3D {
	line3D := &Line3D{}

	line3D.Line3DConfiguration = &Line3DConfiguration{
		BaseConfiguration: opt.NewBaseConfiguration(),
		XAxis3D:           &opt.XAxis3D{},
		YAxis3D:           &opt.YAxis3D{},
		ZAxis3D:           &opt.ZAxis3D{},
		Series:            &series.Line3DSeries{},
	}

	line3D.Container = core.NewContainer(line3D)

	return line3D
}

func (line3D *Line3D) AddSeries(series ...*series.Line3DSingleSeries) {
	for _, s := range series {
		s.Type = types.ChartLine3D
		*line3D.Series = append(*line3D.Series, s)
	}
}
