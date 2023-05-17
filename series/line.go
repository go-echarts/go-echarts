package series

import (
	"github.com/go-echarts/go-echarts/v2/opt"
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/types"
)

type LineSeries0 []*LineSeries

type LineSeries struct {
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
	*opt.Encode        `json:"encode,omitempty"`
	*opt.ItemStyle     `json:"itemStyle,omitempty"`
	*opt.Label         `json:"label,omitempty"`
	*opt.LabelLine     `json:"labelLine,omitempty"`
	*opt.Emphasis      `json:"emphasis,omitempty"`
	*opt.MarkLines     `json:"markLine,omitempty"`
	*opt.MarkAreas     `json:"markArea,omitempty"`
	*opt.MarkPoint     `json:"markPoint,omitempty"`
	*opt.RippleEffect  `json:"rippleEffect,omitempty"`
	*opt.LineStyle     `json:"lineStyle,omitempty"`
	*opt.AreaStyle     `json:"areaStyle,omitempty"`
	*opt.TextStyle     `json:"textStyle,omitempty"`
	*opt.CircularStyle `json:"circular,omitempty"`
}

func (ls LineSeries) New() *LineSeries {
	return &LineSeries{
		Type: types.ChartLine,
	}

}
