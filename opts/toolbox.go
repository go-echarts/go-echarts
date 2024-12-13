package opts

import (
	"bytes"
	"encoding/json"

	"github.com/go-echarts/go-echarts/v2/types"
)

// Toolbox is the option set for a toolbox component.
// https://echarts.apache.org/en/option.html#toolbox
type Toolbox struct {
	// Whether to show toolbox component.
	Show types.Bool `json:"show,omitempty"`

	// The layout orientation of toolbox's icon.
	// Options: 'horizontal','vertical'
	Orient string `json:"orient,omitempty"`

	// Distance between toolbox component and the left side of the container.
	// left value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right', then the component
	// will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between toolbox component and the top side of the container.
	// top value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'; and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom', then the component
	// will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between toolbox component and the right side of the container.
	// right value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Right string `json:"right,omitempty"`

	// Distance between toolbox component and the bottom side of the container.
	// bottom value can be instant pixel value like 20; it can also be a percentage
	// value relative to container width like '20%'.
	// Adaptive by default.
	Bottom string `json:"bottom,omitempty"`

	// The configuration item for each tool.
	// Besides the tools we provide, user-defined toolbox is also supported.
	Feature *ToolBoxFeature `json:"feature,omitempty"`
}

// ToolBoxFeature is a feature component under toolbox.
// https://echarts.apache.org/en/option.html#toolbox
type ToolBoxFeature struct {
	// Save as image tool
	SaveAsImage *ToolBoxFeatureSaveAsImage `json:"saveAsImage,omitempty"`

	// Data brush
	Brush *ToolBoxFeatureBrush `json:"brush,omitempty"`

	// Data area zooming, which only supports rectangular coordinate by now.
	DataZoom *ToolBoxFeatureDataZoom `json:"dataZoom,omitempty"`

	// Data view tool, which could display data in current chart and updates chart after being edited.
	DataView *ToolBoxFeatureDataView `json:"dataView,omitempty"`

	// Restore configuration item.
	Restore *ToolBoxFeatureRestore `json:"restore,omitempty"`

	// User-defined tools. They have to start with "my".
	UserDefined map[string]ToolBoxFeatureUserDefined `json:"-"`
}

func (f ToolBoxFeature) MarshalJSON() ([]byte, error) {
	type ToolBoxFeature_ ToolBoxFeature

	buff := new(bytes.Buffer)
	enc := json.NewEncoder(buff)
	enc.SetEscapeHTML(false)

	err := enc.Encode(ToolBoxFeature_(f))
	if err != nil {
		return nil, err
	}
	// Remove trailing newline.
	buff.Truncate(buff.Len() - 1)

	if len(f.UserDefined) == 0 {
		return buff.Bytes(), nil
	}

	user := new(bytes.Buffer)
	enc = json.NewEncoder(user)
	enc.SetEscapeHTML(false)

	err = enc.Encode(f.UserDefined)
	if err != nil {
		return nil, err
	}
	// Remove trailing newline.
	user.Truncate(user.Len() - 1)

	// Remove trailing "}".
	buff.Truncate(buff.Len() - 1)
	_, _ = buff.WriteString(",")
	// Remove prefix "{".
	_, _ = user.ReadByte()
	// Copy user-defined tools over.
	_, _ = buff.ReadFrom(user)
	return buff.Bytes(), nil
}

// ToolBoxFeatureUserDefined is the option fro user-defined tools.
// https://echarts.apache.org/en/option.html#toolbox.feature
type ToolBoxFeatureUserDefined struct {
	// Whether to show the tool.
	Show types.Bool `json:"show,omitempty"`

	// Title for the tool.
	Title string `json:"title,omitempty"`

	// Icon for the tool.
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	// Icons can be also set to arbitrary vector path via 'path://', like ''path://M30.9,53.2C16.8,53.2,5.3,41......'
	Icon string `json:"icon,omitempty"`

	// On click handler in JavaScript. Use opts.FuncOpts to embed JavaScript.
	OnClick types.FuncStr `json:"onclick,omitempty"`
}

// ToolBoxFeatureSaveAsImage is the option for saving chart as image.
// https://echarts.apache.org/en/option.html#toolbox.feature.saveAsImage
type ToolBoxFeatureSaveAsImage struct {
	// Whether to show the tool.
	Show types.Bool `json:"show,omitempty"`

	// toolbox.feature.saveAsImage. type = 'png'
	// File suffix of the image saved.
	// If the renderer is set to be 'canvas' when chart initialized (default), t
	// hen 'png' (default) and 'jpeg' are supported.
	// If the renderer is set to be 'svg' when when chart initialized, then only 'svg' is supported
	// for type ('svg' type is supported since v4.8.0).
	Type string `json:"type,omitempty" default:"png"`

	// Name to save the image, whose default value is title.text.
	Name string `json:"name,omitempty"`

	// Title for the tool.
	Title string `json:"title,omitempty"`

	// Icon for the tool.
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	// Icons can be also set to arbitrary vector path via 'path://', like ''path://M30.9,53.2C16.8,53.2,5.3,41......'
	Icon string `json:"icon,omitempty"`
}

// ToolBoxFeatureBrush  brush-selecting icon.
// https://echarts.apache.org/en/option.html#toolbox.feature.brush
type ToolBoxFeatureBrush struct {
	// Icons used, whose values are:
	// 'rect': Enabling selecting with rectangle area.
	// 'polygon': Enabling selecting with any shape.
	// 'lineX': Enabling horizontal selecting.
	// 'lineY': Enabling vertical selecting.
	// 'keep': Switching between single selecting and multiple selecting. The latter one can select multiple areas, while the former one cancels previous selection.
	// 'clear': Clearing all selection.
	Type []string `json:"type,omitempty"`
}

// ToolBoxFeatureDataZoom
// https://echarts.apache.org/en/option.html#toolbox.feature.dataZoom
type ToolBoxFeatureDataZoom struct {
	// Whether to show the tool.
	Show types.Bool `json:"show,omitempty"`

	// Defines which yAxis should be controlled. By default, it controls all y axes.
	// If it is set to be false, then no y axis is controlled.
	// If it is set to be then it controls axis with axisIndex of 3.
	// If it is set to be [0, 3], it controls the x-axes with axisIndex of 0 and 3.
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`

	// Icon for the tool.
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	// Icons can be also set to arbitrary vector path via 'path://', like ''path://M30.9,53.2C16.8,53.2,5.3,41......'
	// m["zoom"] = "image://url.."
	// m["back"] = "path://M78..."
	Icon map[string]string `json:"icon"`

	// Restored and zoomed title text.
	// m["zoom"] = "area zooming"
	// m["back"] = "restore area zooming"
	Title map[string]string `json:"title"`
}

// ToolBoxFeatureDataView
// https://echarts.apache.org/en/option.html#toolbox.feature.dataView
type ToolBoxFeatureDataView struct {
	// Whether to show the tool.
	Show types.Bool `json:"show,omitempty"`

	// title for the tool.
	Title string `json:"title,omitempty"`

	// Icon for the tool.
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	// Icons can be also set to arbitrary vector path via 'path://', like ''path://M30.9,53.2C16.8,53.2,5.3,41......'
	Icon string `json:"icon,omitempty"`

	// There are 3 names in data view
	// you could set them like this: []string["data view", "turn off", "refresh"]
	Lang []string `json:"lang"`

	// Background color of the floating layer in data view.
	BackgroundColor string `json:"backgroundColor"`
}

// ToolBoxFeatureRestore
// https://echarts.apache.org/en/option.html#toolbox.feature.restore
type ToolBoxFeatureRestore struct {
	// Whether to show the tool.
	Show types.Bool `json:"show,omitempty"`

	// title for the tool.
	Title string `json:"title,omitempty"`

	// Icon for the tool.
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	// Icons can be also set to arbitrary vector path via 'path://', like ''path://M30.9,53.2C16.8,53.2,5.3,41......'
	Icon string `json:"icon,omitempty"`
}
