package opts

import "github.com/go-echarts/go-echarts/v2/types"

type AngleAxis struct {
	PolarAxisBase
	Clockwise types.Bool `json:"clockwise,omitempty"`
}
