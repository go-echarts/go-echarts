package opts

import "github.com/go-echarts/go-echarts/v2/types"

// Title is the option set for a title component.
// https://echarts.apache.org/en/option.html#title
type Title struct {
	// Set this to false to prevent the title from showing
	Show types.Bool `json:"show,omitempty"`

	// The main title text, supporting \n for newlines.
	Title string `json:"text,omitempty"`

	// The hyper link of main title text.
	Link string `json:"link,omitempty"`

	// Open the hyper link of main title in specified tab.
	// options:
	// 'self' opening it in current tab
	// 'blank' opening it in a new tab
	Target string `json:"target,omitempty"`

	// TextStyle of the main title.
	TitleStyle *TextStyle `json:"textStyle,omitempty"`

	// Subtitle text, supporting \n for newlines.
	Subtitle string `json:"subtext,omitempty"`

	// The hyper link of sub title text.
	SubLink string `json:"sublink,omitempty"`

	// Open the hyper link of subtitle in specified tab.
	// options:
	// 'self' opening it in current tab
	// 'blank' opening it in a new tab
	SubTarget string `json:"subtarget,omitempty"`

	// TextStyle of the sub title.
	SubtitleStyle *TextStyle `json:"subtextStyle,omitempty"`

	// The horizontal align of the component (including "text" and "subtext").
	// Optional values: 'auto', 'left', 'right', 'center'.
	TextAlign string `json:"textAlign,omitempty"`
	// The vertical align of the component (including "text" and "subtext").
	// Optional values: 'auto', 'top', 'bottom', 'middle'.
	TextVerticalAlign string `json:"textVerticalAlign,omitempty"`

	// Distance between title component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between title component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between title component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`

	// Distance between title component and the bottom side of the container.
	// bottom value can be instant pixel value like 20;
	// it can also be a percentage value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// TitleBackgroundColor Background color of title, which is transparent by default.
	TitleBackgroundColor string `json:"backgroundColor,omitempty"`

	// BorderColor Border color of title. Support the same color format as backgroundColor.
	BorderColor string `json:"borderColor,omitempty"`

	// BorderWidth Border width of title.
	BorderWidth int `json:"borderWidth,omitempty"`

	// BorderRadius The radius of rounded corner. Its unit is px. And it supports use array to respectively specify the 4 corner radiuses.
	// For example:
	// borderRadius: 5, // consistently set the size of 4 rounded corners
	// borderRadius: [5, 5, 0, 0] // (clockwise upper left, upper right, bottom right and bottom left)
	// FYI, use [1] is same to number 1
	BorderRadius []int `json:"borderRadius,omitempty"`

	// Gauge
	// Value position relative to the center of chart
	// OffsetCenter is provided as [x, y] where x and y are either a number (px, provided
	// as string) or a percentage.
	// Positive values move the chart value to [right, bottom], negative values vice
	// versa.
	OffsetCenter []string `json:"offsetCenter,omitempty"`
}
