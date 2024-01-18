package opts

import "github.com/go-echarts/go-echarts/v2/types"

// Legend is the option set for a legend component.
// Legend component shows symbol, color and name of different series. You can click legends to toggle displaying series in the chart.
// https://echarts.apache.org/en/option.html#legend
type Legend struct {
	// Whether to show the Legend, default true.
	// Once you set other options, need to manually set it to true
	Show types.Bool `json:"show,omitempty" default:"true"`

	// Type of legend. Optional values:
	// "plain": Simple legend. (default)
	// "scroll": Scrollable legend. It helps when too many legend items needed to be shown.
	Type string `json:"type,omitempty"`

	// Distance between legend component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between legend component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between legend component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`

	// Distance between legend component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// Data array of legend. An array item is usually a name representing string.
	// set Data as []string{} if you wants to hide the legend.
	Data interface{} `json:"data,omitempty"`

	// The layout orientation of legend.
	// Options: 'horizontal', 'vertical'
	Orient string `json:"orient,omitempty"`

	// Legend color when not selected.
	InactiveColor string `json:"inactiveColor,omitempty"`

	// State table of selected legend.
	// example:
	// var selected = map[string]bool{}
	// selected["series1"] = true
	// selected["series2"] = false
	Selected map[string]bool `json:"selected,omitempty"`

	// Selected mode of legend, which controls whether series can be toggled displaying by clicking legends.
	// It is enabled by default, and you may set it to be false to disabled it.
	// Besides, it can be set to 'single' or 'multiple', for single selection and multiple selection.
	SelectedMode string `json:"selectedMode,omitempty"`

	// Legend space around content. The unit is px.
	// Default values for each position are 5.
	// And they can be set to different values with left, right, top, and bottom.
	// Examples:
	// 1. Set padding to be 5
	//    padding: 5
	// 2. Set the top and bottom paddings to be 5, and left and right paddings to be 10
	//    padding: [5, 10]
	// 3. Set each of the four paddings separately
	//    padding: [
	//      5,  // up
	//      10, // right
	//      5,  // down
	//      10, // left
	//    ]
	Padding interface{} `json:"padding,omitempty"`

	// Image width of legend symbol.
	ItemWidth int `json:"itemWidth,omitempty"`

	// Image height of legend symbol.
	ItemHeight int `json:"itemHeight,omitempty"`

	// Legend X position, right/left/center
	X string `json:"x,omitempty"`

	// Legend Y position, right/left/center
	Y string `json:"y,omitempty"`

	// Width of legend component. Adaptive by default.
	Width string `json:"width,omitempty"`

	// Height of legend component. Adaptive by default.
	Height string `json:"height,omitempty"`

	// Legend marker and text aligning.
	// By default, it automatically calculates from component location and orientation.
	// When left value of this component is 'right' and orient is 'vertical', it would be aligned to 'right'.
	// Options: auto/left/right
	Align string `json:"align,omitempty"`

	// Legend text style.
	TextStyle *TextStyle `json:"textStyle,omitempty"`
}
