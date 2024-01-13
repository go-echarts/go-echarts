package opts

// Title is the option set for a title component.
// https://echarts.apache.org/en/option.html#title
type Title struct {
	// The main title text, supporting \n for newlines.
	Title string `json:"text,omitempty"`

	// TextStyle of the main title.
	TitleStyle *TextStyle `json:"textStyle,omitempty"`

	// The hyper link of main title text.
	Link string `json:"link,omitempty"`

	// Subtitle text, supporting \n for newlines.
	Subtitle string `json:"subtext,omitempty"`

	// TextStyle of the sub title.
	SubtitleStyle *TextStyle `json:"subtextStyle,omitempty"`

	// The hyper link of sub title text.
	SubLink string `json:"sublink,omitempty"`

	// Open the hyper link of main title in specified tab.
	// options:
	// 'self' opening it in current tab
	// 'blank' opening it in a new tab
	Target string `json:"target,omitempty"`

	// Distance between title component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between title component and the bottom side of the container.
	// bottom value can be instant pixel value like 20;
	// it can also be a percentage value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// Distance between title component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between title component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`
}
