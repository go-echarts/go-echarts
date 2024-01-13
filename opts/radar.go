package opts

// RadarComponent is the option set for a radar component.
// https://echarts.apache.org/en/option.html#radar
type RadarComponent struct {
	// Indicator of radar chart, which is used to assign multiple variables(dimensions) in radar chart.
	Indicator []*Indicator `json:"indicator,omitempty"`

	// Radar render type, in which 'polygon' and 'circle' are supported.
	Shape string `json:"shape,omitempty"`

	// Segments of indicator axis.
	// default 5
	SplitNumber int `json:"splitNumber,omitempty"`

	// Center position of , the first of which is the horizontal position, and the second is the vertical position.
	// Percentage is supported. When set in percentage, the item is relative to the container width and height.
	Center interface{} `json:"center,omitempty"`

	// Split area of axis in grid area.
	SplitArea *SplitArea `json:"splitArea,omitempty"`

	// Split line of axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`
}
