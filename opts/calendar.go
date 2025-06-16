package opts

import "github.com/go-echarts/go-echarts/v2/types"

// Calendar is the option set for a calendar component.
// This works with the calendar coordinate system, and it is a heatmap calendar.
type Calendar struct {
	ID     string `json:"id,omitempty"`
	Zlevel int    `json:"zlevel,omitempty"`
	Z      int    `json:"z,omitempty"`
	// Distance between grid component and the left side of the container.
	Left string `json:"left,omitempty"`

	// Distance between grid component and the right side of the container.
	Right string `json:"right,omitempty"`

	// Distance between grid component and the top side of the container.
	Top string `json:"top,omitempty"`

	// Distance between grid component and the bottom side of the container.
	Bottom string `json:"bottom,omitempty"`

	// Width of grid component.
	Width string `json:"width,omitempty"`

	// Height of grid component. Adaptive by default.
	Height string `json:"height,omitempty"`

	// Required, range of Calendar coordinates, support multiple formats.
	Range []string `json:"range,omitempty"`

	// The size of each rect of calendar coordinates.
	CellSize string `json:"cellSize,omitempty"`

	// The layout orientation of legend.
	// Options: 'horizontal', 'vertical'
	Orient string `json:"orient,omitempty"`

	// Split line of X axis in grid area.
	SplitLine *SplitLine `json:"splitLine,omitempty"`

	// Graphic style of Map Area Border, emphasis is the style when it is highlighted,
	// like being hovered by mouse, or highlighted via legend connect.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Day Label
	DayLabel *CalendarLabel `json:"dayLabel,omitempty"`

	// Month Label
	MonthLabel *CalendarLabel `json:"monthLabel,omitempty"`

	// Year Label
	YearLabel *CalendarLabel `json:"yearLabel,omitempty"`

	// Whether to ignore mouse events. Default value is false, for triggering and responding to mouse events.
	Silent types.Bool `json:"silent,omitempty"`
}

// CalendarLabel is the option set for a calendar label: DayLabel, MonthLabel, YearLabel.
type CalendarLabel struct {
	// Whether to show the label.
	Show types.Bool `json:"show,omitempty"`

	FirstDay int `json:"firstDay,omitempty"`

	// The margin between the month label and the axis line.
	Margin float64 `json:"margin,omitempty"`

	// Position of year.
	// Default: when orient is set as horizontal, position is left when orient is set as vertical, position is top.
	// Options: 'left', 'right', 'top', 'bottom'
	Position string `json:"position,omitempty"`

	// Text color.
	Color string `json:"color,omitempty"`

	// Font style.
	// Options: 'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`

	// Font weight.
	// Options: 'normal', 'bold', 'bolder', 'lighter', 100 | 200 | 300 | 400...
	FontWeight string `json:"fontWeight,omitempty"`

	// Font family.
	FontFamily string `json:"fontFamily,omitempty"`

	// Font size.
	FontSize int `json:"fontSize,omitempty"`

	// Horizontal alignment of text, automatic by default.
	// Options: 'left', 'center', 'right'
	Align string `json:"align,omitempty"`

	// Vertical alignment of text, automatic by default.
	// Options: 'top', 'middle', 'bottom'
	VerticalAlign string `json:"verticalAlign,omitempty"`

	// Line height of text.
	LineHeight int `json:"lineHeight,omitempty"`

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

	// Padding
	Padding int `json:"padding,omitempty"`

	// Shadow blur of text block.
	ShadowBlur int `json:"shadowBlur,omitempty"`

	// Shadow color of text block.
	ShadowColor string `json:"shadowColor,omitempty"`

	// Shadow X offset of text block.
	ShadowOffsetX int `json:"shadowOffsetX,omitempty"`

	// Shadow Y offset of text block.
	ShadowOffsetY int `json:"shadowOffsetY,omitempty"`

	// Width
	Width int `json:"width,omitempty"`

	// Height
	Height int `json:"height,omitempty"`

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

	// Overflow
	Overflow string `json:"overflow,omitempty"`

	// Ellipsis
	Ellipsis types.Bool `json:"ellipsis,omitempty"`

	// Silent
	Silent types.Bool `json:"silent,omitempty"`
}
