package opts

// Toolbox is the option set for a toolbox component.
// https://echarts.apache.org/en/option.html#toolbox
type Toolbox struct {
	// Whether to show toolbox component.
	Show bool `json:"show"`

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
	Brush *ToolBoxFeatureBrush `json:"brush"`

	// Data area zooming, which only supports rectangular coordinate by now.
	DataZoom *ToolBoxFeatureDataZoom `json:"dataZoom,omitempty"`

	// Data view tool, which could display data in current chart and updates chart after being edited.
	DataView *ToolBoxFeatureDataView `json:"dataView,omitempty"`

	// Restore configuration item.
	Restore *ToolBoxFeatureRestore `json:"restore,omitempty"`
}

// ToolBoxFeatureSaveAsImage is the option for saving chart as image.
// https://echarts.apache.org/en/option.html#toolbox.feature.saveAsImage
type ToolBoxFeatureSaveAsImage struct {
	// Whether to show the tool.
	Show bool `json:"show"`

	// toolbox.feature.saveAsImage. type = 'png'
	// File suffix of the image saved.
	// If the renderer is set to be 'canvas' when chart initialized (default), t
	// hen 'png' (default) and 'jpeg' are supported.
	// If the renderer is set to be 'svg' when when chart initialized, then only 'svg' is supported
	// for type ('svg' type is supported since v4.8.0).
	Type string `json:"png,omitempty"`

	// Name to save the image, whose default value is title.text.
	Name string `json:"name,omitempty"`

	// Title for the tool.
	Title string `json:"title,omitempty"`
}

// ToolBoxFeatureBrush  brush-selecting icon.
// https://echarts.apache.org/en/option.html#toolbox.feature.brush
type ToolBoxFeatureBrush struct {

	//Icons used, whose values are:
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
	Show bool `json:"show"`

	//Defines which yAxis should be controlled. By default, it controls all y axes.
	//If it is set to be false, then no y axis is controlled.
	//If it is set to be then it controls axis with axisIndex of 3.
	//If it is set to be [0, 3], it controls the x-axes with axisIndex of 0 and 3.
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`

	// Restored and zoomed title text.
	// m["zoom"] = "area zooming"
	// m["back"] = "restore area zooming"
	Title map[string]string `json:"title"`
}

// ToolBoxFeatureDataView
// https://echarts.apache.org/en/option.html#toolbox.feature.dataView
type ToolBoxFeatureDataView struct {
	// Whether to show the tool.
	Show bool `json:"show"`

	// title for the tool.
	Title string `json:"title,omitempty"`

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
	Show bool `json:"show"`

	// title for the tool.
	Title string `json:"title,omitempty"`
}
