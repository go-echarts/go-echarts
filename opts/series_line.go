package opts

import "github.com/go-echarts/go-echarts/v2/types"

// LineChart is the options set for a line chart.
// https://echarts.apache.org/en/option.html#series-line
type LineChart struct {
	// ColorBy The policy to take color from option.color. Valid values:
	// 'series': assigns the colors in the palette by series, so that all data in the same series are in the same color;
	// 'data': assigns colors in the palette according to data items, with each data item using a different color.
	ColorBy string

	// CoordinateSystem The coordinate used in the series, whose options are:
	//'cartesian2d' Use a two-dimensional rectangular coordinate (also known as Cartesian coordinate), with xAxisIndex and
	//  yAxisIndex to assign the corresponding axis component.
	//
	// 'polar' Use polar coordinates, with polarIndex to assign the corresponding polar coordinate component.
	CoordSystem string

	// Index of x-axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y-axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int

	// Index of polar coordinate to combine with, which is useful for multiple polar axes in one chart.
	PolarIndex int

	// Icon types provided by ECharts includes
	//  'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// Full documentation: https://echarts.apache.org/en/option.html#series-line.symbol
	Symbol string

	// symbol size. It can be set to single numbers like 10, or use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10.
	// Full documentation: https://echarts.apache.org/en/option.html#series-line.symbolSize
	SymbolSize interface{}

	// SymbolKeepAspect is whether to keep aspect for symbols in the form of path://.
	SymbolKeepAspect types.Bool

	// Whether to show symbol. It would be shown during tooltip hover.
	ShowSymbol types.Bool

	// If stack the value. On the same category axis, the series with the same stack name would be put on top of each other.
	// The effect of the below example could be seen through stack switching of toolbox on the top right corner:
	Stack string

	// Whether to show as smooth curve.
	Smooth types.Bool

	// Whether to connect the line across null points.
	ConnectNulls types.Bool

	// Whether to show as a step line. It can be true, false. Or 'start', 'middle', 'end'.
	// Which will configure the turn point of step line.
	Step interface{}
}

// LineData
// https://echarts.apache.org/en/option.html#series-line.data
type LineData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// Symbol of single data.
	// Icon types provided by ECharts includes 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Symbol string `json:"symbol,omitempty"`

	// single data symbol size. It can be set to single numbers like 10, or
	// use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10
	SymbolSize int `json:"symbolSize,omitempty"`

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int `json:"XAxisIndex,omitempty"`

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int `json:"YAxisIndex,omitempty"`
}
