package opts

import "github.com/go-echarts/go-echarts/v2/types"

type RadiusAxis struct {
	PolarAxisBase
	Name          string     `json:"name,omitempty"`
	NameLocation  string     `json:"nameLocation,omitempty"`
	NameTextStyle TextStyle  `json:"nameTextStyle,omitempty"`
	NameGap       float64    `json:"nameGap,omitempty"`
	NameRadius    float64    `json:"nameRotate,omitempty"`
	Inverse       types.Bool `json:"inverse,omitempty"`
}
