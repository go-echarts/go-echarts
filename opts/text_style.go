package opts

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

	// Padding title space around content. See legend.textStyle.padding
	// The unit is px. Default values for each position are 5.
	// And they can be set to different values with left, right, top, and bottom.
	Padding interface{} `json:"padding,omitempty"`

	// compatible for WordCloud
	Normal *TextStyle `json:"normal,omitempty"`
}
