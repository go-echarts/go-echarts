package opts

import "github.com/go-echarts/go-echarts/v2/types"

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

	// AxisLine controls settings related to axis line.
	AxisLine *AxisLine `json:"axisLine,omitempty"`

	// The start angle of coordinate, which is the angle of the first indicator axis.
	StartAngle float64 `json:"startAngle,omitempty"`

	// Name options for radar indicators.
	AxisName *AxisName `json:"axisName,omitempty"`
}

// Indicator is the option set for a radar chart.
type Indicator struct {
	// Indicator name
	Name string `json:"name,omitempty"`

	// The maximum value of indicator. It is an optional configuration, but we recommend to set it manually.
	Max float32 `json:"max,omitempty"`

	// The minimum value of indicator. It is an optional configuration, with default value of 0.
	Min float32 `json:"min,omitempty"`

	// Specify a color the indicator.
	Color string `json:"color,omitempty"`
}

// Name options for radar indicators.
type AxisName struct {
	// Whether to display the indicator's name.
	Show types.Bool `json:"show,omitempty"`

	// The formatter of indicator's name, using string template, the template variable should be the indicator's name {value}
	Formatter types.FuncStr `json:"formatter,omitempty"`

	// Font color
	Color string `json:"color,omitempty"`

	// Font style
	// Options: 'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`

	// FontWeight main title font thick weight.
	// Options are:
	// 'normal'
	// 'bold'
	// 'bolder'
	// 'lighter'
	// 100 | 200 | 300 | 400...
	FontWeight string `json:"fontWeight,omitempty"`

	// Font family the main title font family.
	// Options: "sans-serif", 'serif' , 'monospace', 'Arial', 'Courier New', 'Microsoft YaHei', ...
	FontFamily string `json:"fontFamily,omitempty"`

	// Font size
	FontSize int `json:"fontSize,omitempty"`

	// Width of text block.
	Width int `json:"width,omitempty"`

	// Height of text block.
	Height int `json:"height,omitempty"`

	// Line height of the text fragment.
	LineHeight int `json:"lineHeight,omitempty"`

	// Text border color.
	TextBorderColor string `json:"textBorderColor,omitempty"`

	// Text border width.
	TextBorderWidth int `json:"textBorderWidth,omitempty"`

	// Text border type
	// Options: 'solid', 'dashed', 'dotted'
	TextBorderType string `json:"textBorderType,omitempty"`

	// Text border dash offset.
	TextBorderDashOffset int `json:"textBorderDashOffset,omitempty"`

	// Text shadow color.
	TextShadowColor string `json:"textShadowColor,omitempty"`

	// Text shadow blur.
	TextShadowBlur int `json:"textShadowBlur,omitempty"`

	// Text shadow X offset.
	TextShadowOffsetX int `json:"textShadowOffsetX,omitempty"`

	// Text shadow Y offset.
	TextShadowOffsetY int `json:"textShadowOffsetY,omitempty"`

	// Determine how to display the text when it's overflow. Available when width is set.
	//
	// 'truncate' Truncate the text and trailing with ellipsis.
	// 'break' Break by word
	// 'breakAll' Break by character.
	Overflow string `json:"overflow,omitempty"`

	// Ellipsis
	Ellipsis types.Bool `json:"ellipsis,omitempty"`

	// Background color of label, which is transparent by default.
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Border color of label.
	BorderColor string `json:"borderColor,omitempty"`

	// Border width of label.
	BorderWidth int `json:"borderWidth,omitempty"`

	// Border radius of label.
	BorderRadius int `json:"borderRadius,omitempty"`

	// Border type of label.
	// Options: 'solid', 'dashed', 'dotted'
	BorderType string `json:"borderType,omitempty"`

	// Border dash offset of label.
	BorderDashOffset int `json:"borderDashOffset,omitempty"`

	// Shadow blur of text block.
	ShadowBlur int `json:"shadowBlur,omitempty"`

	// Shadow color of text block.
	ShadowColor string `json:"shadowColor,omitempty"`

	// Shadow X offset of text block.
	ShadowOffsetX int `json:"shadowOffsetX,omitempty"`

	// Shadow Y offset of text block.
	ShadowOffsetY int `json:"shadowOffsetY,omitempty"`

	// Padding title space around content. See legend.textStyle.padding
	// The unit is px. Default values for each position are 5.
	// And they can be set to different values with left, right, top, and bottom.
	Padding interface{} `json:"padding,omitempty"`
}
