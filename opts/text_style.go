package opts

import "github.com/go-echarts/go-echarts/v2/types"

// TextStyle is the option set for a text style component.
type TextStyle struct {
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

	// Horizontal alignment of text, automatic by default.
	// Options: 'left', 'center', 'right'
	Align string `json:"align,omitempty"`

	// Vertical alignment of text, automatic by default.
	// Options: 'top', 'middle', 'bottom'
	VerticalAlign string `json:"verticalAlign,omitempty"`

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

	// compatible for WordCloud
	Normal *TextStyle `json:"normal,omitempty"`
}
