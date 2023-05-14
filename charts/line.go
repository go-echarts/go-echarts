package charts

import (
	"github.com/go-echarts/go-echarts/v3/opts"
	"github.com/go-echarts/go-echarts/v3/types"
)

// Line represents a line chart.
type Line struct {
	commonConf CommonConfig

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

// Type returns the chart type.
func (c *Line) Type() string { return types.ChartLine }

func (c *Line) SetCommonConfig(commonConf CommonConfig) {
	c.commonConf = commonConf
}

// AddSeries adds the new series.
func (c *Line) AddSeries(data ...opts.LineData) *Line {
	return c
}
