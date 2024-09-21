package opts

import "github.com/go-echarts/go-echarts/v2/types"

// PieChart is the option set for a pie chart.
// https://echarts.apache.org/en/option.html#series-pie
type PieChart struct {
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

	// Whether to show as Nightingale chart, which distinguishes data through radius. There are 2 optional modes:
	// * 'radius' Use central angle to show the percentage of data, radius to show data size.
	// * 'area' All the sectors will share the same central angle, the data size is shown only through radiuses.
	RoseType string

	// Center position of Pie chart, the first of which is the horizontal position, and the second is the vertical position.
	// Percentage is supported. When set in percentage, the item is relative to the container width,
	// and the second item to the height.
	//
	// Example:
	//
	// Set to absolute pixel values ->> center: [400, 300]
	// Set to relative percent ->> center: ['50%', '50%']
	Center interface{}

	// Radius of Pie chart. Value can be:
	// * number: Specify outside radius directly.
	// * string: For example, '20%', means that the outside radius is 20% of the viewport
	// size (the little one between width and height of the chart container).
	//
	// Array.<number|string>: The first item specifies the inside radius, and the
	// second item specifies the outside radius. Each item follows the definitions above.
	Radius interface{}
}

// PieData
// https://echarts.apache.org/en/option.html#series-pie.data
type PieData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// Whether the data item is selected.
	Selected types.Bool `json:"selected,omitempty"`

	// The label configuration of a single sector.
	Label *Label `json:"label,omitempty"`

	// Graphic style of , emphasis is the style when it is highlighted, like being hovered by mouse, or highlighted via legend connect.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// tooltip settings in this series data.
	Tooltip *Tooltip `json:"tooltip,omitempty"`
}
