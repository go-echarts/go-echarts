package opts

import "github.com/go-echarts/go-echarts/v2/types"

type RadiusAxis struct {
	PolarAxisBase

	Name string `json:"name,omitempty"`
	// NameLocation Location of axis name.
	// Options:
	// 'start'
	// 'middle' or 'center'
	// 'end'
	NameLocation string `json:"nameLocation,omitempty"`

	// NameTextStyle Text style of axis name.
	NameTextStyle TextStyle `json:"nameTextStyle,omitempty"`

	// NameGap Gap between axis name and axis line.
	NameGap float64 `json:"nameGap,omitempty"`

	// NameRadius Rotation of axis name.
	NameRadius float64 `json:"nameRotate,omitempty"`

	// Inverse Set this to true to invert the axis. This is a new option available from Echarts 3 and newer.
	Inverse types.Bool `json:"inverse,omitempty"`
}
