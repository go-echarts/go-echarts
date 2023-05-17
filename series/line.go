package series

import (
	"github.com/go-echarts/go-echarts/v2/components"
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
	*components.Encode        `json:"encode,omitempty"`
	*components.ItemStyle     `json:"itemStyle,omitempty"`
	*components.Label         `json:"label,omitempty"`
	*components.LabelLine     `json:"labelLine,omitempty"`
	*components.Emphasis      `json:"emphasis,omitempty"`
	*components.MarkLines     `json:"markLine,omitempty"`
	*components.MarkAreas     `json:"markArea,omitempty"`
	*components.MarkPoint     `json:"markPoint,omitempty"`
	*components.RippleEffect  `json:"rippleEffect,omitempty"`
	*components.LineStyle     `json:"lineStyle,omitempty"`
	*components.AreaStyle     `json:"areaStyle,omitempty"`
	*components.TextStyle     `json:"textStyle,omitempty"`
	*components.CircularStyle `json:"circular,omitempty"`
}

func (ls LineSeries) New() *LineSeries {
	return &LineSeries{
		Type: types.ChartLine,
	}

}
