package opts

import "github.com/go-echarts/go-echarts/v2/types"

type RadarChart struct {
	// ColorBy The policy to take color from option.color. Valid values:
	// 'series': assigns the colors in the palette by series, so that all data in the same series are in the same color;
	// 'data': assigns colors in the palette according to data items, with each data item using a different color.
	ColorBy string

	// RadarIndex Index of radar component that radar chart uses.
	RadarIndex int
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

// RadarData
// https://echarts.apache.org/en/option.html#series-radar
type RadarData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}
