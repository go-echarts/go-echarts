package opt

import "github.com/go-echarts/go-echarts/v2/primitive"

// Title is the option set for a title component.
// https://echarts.apache.org/en/option.html#title
type Title struct {
	Id primitive.String `json:"id,omitempty"`

	// Default is true, Set this to false to prevent the title from showing
	Show primitive.Bool `json:"show,omitempty"`

	// The main title text, supporting \n for newlines.
	Text primitive.String `json:"text,omitempty"`

	// The hyper link of main title text.
	Link primitive.String `json:"link,omitempty"`

	// Open the hyper link of main title in specified tab.
	// options:
	// 'self' opening it in current tab
	// 'blank' opening it in a new tab
	Target primitive.String `json:"target,omitempty"`

	// TextStyle of the main title.
	TextStyle *TextStyle `json:"textStyle,omitempty"`

	// Subtitle text, supporting \n for newlines.
	SubText primitive.String `json:"subtext,omitempty"`

	// TextStyle of the sub title.
	SubTextStyle *TextStyle `json:"subtextStyle,omitempty"`

	// The hyper link of sub title text.
	SubLink primitive.String `json:"sublink,omitempty"`

	// SubTarget

	// Distance between title component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top primitive.String `json:"top,omitempty"`

	// Distance between title component and the bottom side of the container.
	// bottom value can be instant pixel value like 20;
	// it can also be a percentage value relative to container width like '20%'.
	// Adaptive by default.
	Bottom primitive.String `json:"bottom,omitempty"`

	// Distance between title component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left primitive.String `json:"left,omitempty"`

	// Distance between title component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right primitive.String `json:"right,omitempty"`
}

func (title *Title) NewDefault() *Title {
	return &Title{}
}
