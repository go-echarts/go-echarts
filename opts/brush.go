package opts

// Brush is an area-selecting component, with which user can select part of data from a chart to display in detail, or do calculations with them.
// https://echarts.apache.org/en/option.html#brush
type Brush struct {
	// XAxisIndex Assigns which of the xAxisIndex can use brush selecting.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`

	// Brushlink is a mapping of dataIndex. So data of every series with brushLink should be guaranteed to correspond to the other.
	Brushlink interface{} `json:"brushlink,omitempty"`

	// OutOfBrush Defines visual effects of items out of selection
	OutOfBrush *BrushOutOfBrush `json:"outOfBrush,omitempty"`
}

// BrushOutOfBrush
// https://echarts.apache.org/en/option.html#brush.outOfBrush
type BrushOutOfBrush struct {
	ColorAlpha float32 `json:"colorAlpha,omitempty"`
}
