package opts

import "github.com/go-echarts/go-echarts/v2/types"

// Polar coordinate can be used in scatter and line chart. Every polar coordinate has an angleAxis and a radiusAxis.
type Polar struct {
	ID      string    `json:"id,omitempty"`
	Zlevel  int       `json:"zlevel,omitempty"`
	Z       int       `json:"z,omitempty"`
	Center  [2]string `json:"center,omitempty"`
	Radius  [2]string `json:"radius,omitempty"`
	Tooltip Tooltip   `json:"tooltip,omitempty"`
}

type PolarAxisBase struct {
	ID           string     `json:"id,omitempty"`
	PolarIndex   int        `json:"polarIndex,omitempty"`
	StartAngle   float64    `json:"startAngle,omitempty"`
	Type         string     `json:"type,omitempty"`
	BoundaryGap  types.Bool `json:"boundaryGap,omitempty"`
	Min          float64    `json:"min,omitempty"`
	Max          float64    `json:"max,omitempty"`
	Scale        types.Bool `json:"scale,omitempty"`
	SplitNumber  int        `json:"splitNumber,omitempty"`
	MinInterval  float64    `json:"minInterval,omitempty"`
	MaxInterval  float64    `json:"maxInterval,omitempty"`
	Interval     float64    `json:"interval,omitempty"`
	LogBase      float64    `json:"logBase,omitempty"`
	Silent       types.Bool `json:"silent,omitempty"`
	TriggerEvent types.Bool `json:"triggerEvent,omitempty"`
}
