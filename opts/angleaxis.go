package opts

type AngleAxis struct {
	PolarAxisBase
	Clockwise bool `json:"clockwise,omitempty"`
}

type PolarAxisBase struct {
	ID           string  `json:"id,omitempty"`
	PolarIndex   int     `json:"polarIndex,omitempty"`
	StartAngle   float64 `json:"startAngle,omitempty"`
	Type         string  `json:"type,omitempty"`
	BoundaryGap  bool    `json:"boundaryGap,omitempty"`
	Min          float64 `json:"min,omitempty"`
	Max          float64 `json:"max,omitempty"`
	Scale        bool    `json:"scale,omitempty"`
	SplitNumber  int     `json:"splitNumber,omitempty"`
	MinInterval  float64 `json:"minInterval,omitempty"`
	MaxInterval  float64 `json:"maxInterval,omitempty"`
	Interval     float64 `json:"interval,omitempty"`
	LogBase      float64 `json:"logBase,omitempty"`
	Silent       bool    `json:"silent,omitempty"`
	TriggerEvent bool    `json:"triggerEvent,omitempty"`
}
