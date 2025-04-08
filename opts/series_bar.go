package opts

import "github.com/go-echarts/go-echarts/v2/types"

// BarChart
// https://echarts.apache.org/en/option.html#series-bar
type BarChart struct {
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

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int

	// Index of polar coordinate to combine with, which is useful for multiple polar axes in one chart.
	PolarIndex int

	// Whether to add round caps at the end of the bar sectors. Valid only for bar series on polar coordinates.
	RoundCap types.Bool

	// Whether to show background behind each bar. Use backgroundStyle to set background style.
	ShowBackground types.Bool

	// Name of stack. On the same category axis, the series with the
	// same stack name would be put on top of each other.
	Stack string

	// The gap between bars between different series, is a percent value like '30%',
	// which means 30% of the bar width.
	// Set barGap as '-100%' can overlap bars that belong to different series,
	// which is useful when putting a series of bar as background.
	// In a single coordinate system, this attribute is shared by multiple 'bar' series.
	// This attribute should be set on the last 'bar' series in the coordinate system,
	// then it will be adopted by all 'bar' series in the coordinate system.
	BarGap string

	// The bar gap of a single series, defaults to be 20% of the category gap,
	// can be set as a fixed value.
	// In a single coordinate system, this attribute is shared by multiple 'bar' series.
	// This attribute should be set on the last 'bar' series in the coordinate system,
	// then it will be adopted by all 'bar' series in the coordinate system.
	BarCategoryGap string

	// The width of the bar. Adaptive when not specified.
	// Can be an absolute value like 40 or a percent value like '60%'. The percent is based on the calculated category width.
	// In a single coordinate system, this attribute is shared by multiple 'bar' series.
	// This attribute should be set on the last 'bar' series in the coordinate system, then it will be adopted by all 'bar' series in the coordinate system.
	BarWidth string
}

// BarData
// https://echarts.apache.org/en/option.html#series-bar.data
type BarData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// The style setting of the text label in a single bar.
	Label *Label `json:"label,omitempty"`

	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Tooltip settings in this series data.
	Tooltip *Tooltip `json:"tooltip,omitempty"`
}
