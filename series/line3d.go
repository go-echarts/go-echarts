package series

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/types"
)

type Line3DSeries0 []*Line3DSeries

type Line3DSeries struct {
	Id   primitive.String `json:"id,omitempty"`
	Name primitive.String `json:"name,omitempty"`
	Type primitive.String `json:"type,omitempty"`

	Step         primitive.Mixed    `json:"step,omitempty"`
	Smooth       primitive.Bool     `json:"smooth,omitempty"`
	ConnectNulls primitive.Bool     `json:"connectNulls,omitempty"`
	ShowSymbol   primitive.Bool     `json:"showSymbol,omitempty"`
	Symbol       primitive.String   `json:"symbol,omitempty"`
	Color        []primitive.String `json:"color,omitempty"`

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
	*opts.MarkPoint     `json:"markPoint,omitempty"`
	*opts.RippleEffect  `json:"rippleEffect,omitempty"`
	*opts.LineStyle     `json:"lineStyle,omitempty"`
	*opts.AreaStyle     `json:"areaStyle,omitempty"`
	*opts.TextStyle     `json:"textStyle,omitempty"`
	*opts.CircularStyle `json:"circular,omitempty"`
}

func (ls Line3DSeries) New() *Line3DSeries {
	return &Line3DSeries{
		Type: types.ChartLine,
	}
}
