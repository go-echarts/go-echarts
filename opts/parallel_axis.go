package opts

import "github.com/go-echarts/go-echarts/v2/types"

// ParallelAxis is the option set for a parallel axis.
type ParallelAxis struct {
	// Dimension index of coordinate axis.
	Dim int `json:"dim,omitempty"`

	// Name of axis.
	Name string `json:"name,omitempty"`

	// The maximum value of axis.
	// It can be set to a special value 'dataMax' so that the minimum value on this axis is set to be the maximum label.
	// It will be automatically computed to make sure axis tick is equally distributed when not set.
	// In category axis, it can also be set as the ordinal number.
	Max interface{} `json:"max,omitempty"`

	// Compulsively set segmentation interval for axis.
	Inverse types.Bool `json:"inverse,omitempty"`

	// Location of axis name. Options: "start", "middle", "center", "end"
	// default "end"
	NameLocation string `json:"nameLocation,omitempty"`

	// Type of axis.
	// Options：
	// "value"：Numerical axis, suitable for continuous data.
	// "category" Category axis, suitable for discrete category data. Category data can be auto retrieved from series.
	// "log"  Log axis, suitable for log data.
	Type string `json:"type,omitempty"`

	// Category data，works on (type: "category").
	Data interface{} `json:"data,omitempty"`
}
