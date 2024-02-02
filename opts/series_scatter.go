package opts

import "github.com/go-echarts/go-echarts/v2/types"

// ScatterChart is the option set for a scatter chart.
// https://echarts.apache.org/en/option.html#series-scatter
type ScatterChart struct {
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

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of x axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int

	// Icon types provided by ECharts includes
	//  'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// Full documentation: https://echarts.apache.org/en/option.html#series-line.symbol
	Symbol string

	// symbol size. It can be set to single numbers like 10, or use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10.
	// Full documentation: https://echarts.apache.org/en/option.html#series-line.symbolSize
	SymbolSize interface{}

	// SymbolKeepAspect is whether to keep aspect for symbols in the form of path://.
	SymbolKeepAspect types.Bool
}

// ScatterData
// https://echarts.apache.org/en/option.html#series-scatter.data
type ScatterData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// Symbol
	Symbol string `json:"symbol,omitempty"`

	// SymbolSize
	SymbolSize int `json:"symbolSize,omitempty"`

	// SymbolRotate
	SymbolRotate int `json:"symbolRotate,omitempty"`

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int `json:"xAxisIndex,omitempty"`

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int `json:"yAxisIndex,omitempty"`
}
