package series

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/primitive"
)

type LineSeries struct {
	Id   primitive.String `json:"id,omitempty"`
	Name primitive.String `json:"name,omitempty"`
	Type primitive.String `json:"type,omitempty"`

	Step         primitive.Mixed  `json:"step,omitempty"`
	Smooth       primitive.Bool   `json:"smooth,omitempty"`
	ConnectNulls primitive.Bool   `json:"connectNulls,omitempty"`
	ShowSymbol   primitive.Bool   `json:"showSymbol,omitempty"`
	Symbol       primitive.String `json:"symbol,omitempty"`
	Color        primitive.String `json:"color,omitempty"`

	// series data
	Data primitive.Mixed `json:"data"`

	// series options
	*opts.Encode        `json:"encode,omitempty"`
	*opts.ItemStyle     `json:"itemStyle,omitempty"`
	*opts.Label         `json:"label,omitempty"`
	*opts.LabelLine     `json:"labelLine,omitempty"`
	*opts.Emphasis      `json:"emphasis,omitempty"`
	*opts.MarkLines     `json:"markLine,omitempty"`
	*opts.MarkAreas     `json:"markArea,omitempty"`
	*opts.MarkPoints    `json:"markPoint,omitempty"`
	*opts.RippleEffect  `json:"rippleEffect,omitempty"`
	*opts.LineStyle     `json:"lineStyle,omitempty"`
	*opts.AreaStyle     `json:"areaStyle,omitempty"`
	*opts.TextStyle     `json:"textStyle,omitempty"`
	*opts.CircularStyle `json:"circular,omitempty"`
}

func (ls LineSeries) New() *LineSeries {
	return &LineSeries{
		Type: ls.Type,
	}

}
