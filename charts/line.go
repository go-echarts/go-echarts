package charts

import (
	"github.com/go-echarts/go-echarts/v3/opts"
	"github.com/go-echarts/go-echarts/v3/types"
)

type LineBaseConfiguration struct {
	*BaseConfiguration
	//XAxis *opts.XAxis
	*opts.XAxis
}

// Line represents a line chart.
type Line struct {
	commonConf *LineBaseConfiguration

	ID         string      `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	XAxisIndex int         `json:"xAxisIndex,omitempty"`
	YAxisIndex int         `json:"yAxisIndex,omitempty"`
	PolarIndex int         `json:"polarIndex,omitempty"`
	Label      *opts.Label `json:"label,omitempty"`
}

// NewLine creates a new line chart.
func NewLine(name string) *Line {
	c := &Line{
		Name: name,
	}
	return c
}

func DefaultXAxis() *opts.XAxis {
	return &opts.XAxis{}
}

func DefaultBaseConfiguration() *BaseConfiguration {
}

// Type returns the chart type.
func (c *Line) Type() string { return types.ChartLine }

func (c *Line) SetBaseConfig(commonConf *LineBaseConfiguration) {
	c.commonConf = commonConf

	bc := DefaultBaseConfiguration()
	//xaxis.AxisLabel.Align = "xx"

	conf := &LineBaseConfiguration{
		BaseConfiguration: bc,
		XAxis:             &opts.XAxis{},
	}

}

func (c *Line) SetXAxis(x *opts.XAxis) {
	c.commonConf.XAxis = x
}

func (c *Line) SetYAxis(y *opts.YAxis) {
	//c.commonConf.Y
}

// AddSeries adds the new series.
func (c *Line) AddSeries(data ...opts.LineData) *Line {
	return c
}

func (c *Line) AddDataset() {

}
