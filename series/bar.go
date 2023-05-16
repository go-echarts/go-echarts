package series

import (
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/types"
)

type BarSeries struct {
	Id   primitive.String `json:"id,omitempty"`
	Name primitive.String `json:"name,omitempty"`
	Type primitive.String `json:"type,omitempty"`

	// series data
	Data primitive.Mixed `json:"data"`
}

func (bs BarSeries) New() *BarSeries {
	return &BarSeries{
		Type: types.ChartBar,
	}

}
