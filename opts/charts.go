package opts

import "github.com/go-echarts/go-echarts/v2/types"

// SunburstChart
// https://echarts.apache.org/en/option.html#series-sunburst
type SunburstChart struct {
	// The action of clicking a sector
	NodeClick string `json:"nodeClick,omitempty"`
	// Sorting method that sectors use based on value
	Sort string `json:"sort,omitempty"`
	// If there is no name, whether need to render it.
	RenderLabelForZeroData types.Bool `json:"renderLabelForZeroData,omitempty"`
	// Selected mode
	SelectedMode types.Bool `json:"selectedMode,omitempty"`
	// Whether to enable animation.
	Animation types.Bool `json:"animation,omitempty"`
	// Whether to set graphic number threshold to animation
	AnimationThreshold int `json:"animationThreshold,omitempty"`
	// Duration of the first animation
	AnimationDuration int `json:"animationDuration,omitempty"`
	// Easing method used for the first animation
	AnimationEasing string `json:"animationEasing,omitempty"`
	// Delay before updating the first animation
	AnimationDelay int `json:"animationDelay,omitempty"`
	// Time for animation to complete
	AnimationDurationUpdate int `json:"animationDurationUpdate,omitempty"`
	// Easing method used for animation.
	AnimationEasingUpdate string `json:"animationEasingUpdate,omitempty"`
	// Delay before updating animation
	AnimationDelayUpdate int `json:"animationDelayUpdate,omitempty"`
}

// Bar3DChart is the option set for a 3D bar chart.
type Bar3DChart struct {
	// Shading is the coloring effect of 3D graphics in 3D Bar.
	// The following three coloring methods are supported in echarts-gl:
	// Options:
	//
	// * "color": Only display colors, not affected by other factors such as lighting.
	// * "lambert": Through the classic [lambert] coloring, can express the light and dark that the light shows.
	// * "realistic": Realistic rendering, combined with light.ambientCubemap and postEffect,
	//   can improve the quality and texture of the display.
	//   [Physical Based Rendering (PBR)] (https://www.marmoset.co/posts/physically-based-rendering-and-you-can-too/)
	//   is used in ECharts GL to represent realistic materials.
	Shading string
}

// BoxPlotData
// https://echarts.apache.org/en/option.html#series-boxplot.data
type BoxPlotData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// The style setting of the text label in a single bar.
	Label *Label `json:"label,omitempty"`

	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Emphasis settings in this series data.
	Emphasis *Emphasis `json:"emphasis,omitempty"`

	// Tooltip settings in this series data.
	Tooltip *Tooltip `json:"tooltip,omitempty"`
}

// FunnelData
// https://echarts.apache.org/en/option.html#series-funnel.data
type FunnelData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GeoData
type GeoData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GaugeData
// https://echarts.apache.org/en/option.html#series-gauge.data
type GaugeData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GraphChart is the option set for graph chart.
// https://echarts.apache.org/en/option.html#series-graph
type GraphChart struct {
	// Graph layout.
	// * 'none' No layout, use x, y provided in node as the position of node.
	// * 'circular' Adopt circular layout, see the example Les Miserables.
	// * 'force' Adopt force-directed layout, see the example Force, the
	// detail about layout configurations are in graph.force
	Layout string

	// Force is the option set for graph force layout.
	Force *GraphForce

	// Whether to enable mouse zooming and translating. false by default.
	// If either zooming or translating is wanted, it can be set to 'scale' or 'move'.
	// Otherwise, set it to be true to enable both.
	Roam types.Bool

	// EdgeSymbol is the symbols of two ends of edge line.
	// * 'circle'
	// * 'arrow'
	// * 'none'
	// example: ["circle", "arrow"] or "circle"
	EdgeSymbol interface{}

	// EdgeSymbolSize is size of symbol of two ends of edge line. Can be an array or a single number
	// example: [5,10] or 5
	EdgeSymbolSize interface{}

	// Draggable allows you to move the nodes with the mouse if they are not fixed.
	Draggable types.Bool

	// Whether to focus/highlight the hover node and it's adjacencies.
	FocusNodeAdjacency types.Bool

	// The categories of node, which is optional. If there is a classification of nodes,
	// the category of each node can be assigned through data[i].category.
	// And the style of category will also be applied to the style of nodes. categories can also be used in legend.
	Categories []*GraphCategory

	// EdgeLabel is the properties of an label of edge.
	EdgeLabel *EdgeLabel `json:"edgeLabel"`

	// SymbolKeepAspect is whether to keep aspect for symbols in the form of path://.
	SymbolKeepAspect types.Bool

	// Animation determines whether to enable animation.
	Animation types.Bool

	// Emphasis contains configurations of emphasis state.
	Emphasis *Emphasis

	// LineStyle controls the style of edge line. lineStyle.color can be 'source' or 'target',
	// which will use the color of source node or target node.
	LineStyle *LineStyle
}

// GraphNode represents a data node in graph chart.
// https://echarts.apache.org/en/option.html#series-graph.data
type GraphNode struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// x value of node position.
	X float32 `json:"x,omitempty"`

	// y value of node position.
	Y float32 `json:"y,omitempty"`

	// Value of data item.
	Value float32 `json:"value,omitempty"`

	// If node are fixed when doing force directed layout.
	Fixed types.Bool `json:"fixed,omitempty"`

	// Index of category which the data item belongs to.
	Category interface{} `json:"category,omitempty"`

	// Symbol of node of this category.
	// Icon types provided by ECharts includes
	// 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Symbol string `json:"symbol,omitempty"`

	// node of this category symbol size. It can be set to single numbers like 10,
	// or use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10.
	SymbolSize interface{} `json:"symbolSize,omitempty"`

	// The style of this node.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// The tooltip of this node.
	Tooltip *Tooltip `json:"tooltip,omitempty"`
}

// GraphLink represents relationship between two data nodes.
// https://echarts.apache.org/en/option.html#series-graph.links
type GraphLink struct {
	// A string representing the name of source node on edge. Can also be a number representing the node index.
	Source interface{} `json:"source,omitempty"`

	// A string representing the name of target node on edge. Can also be a number representing node index.
	Target interface{} `json:"target,omitempty"`

	// value of edge, can be mapped to edge length in force graph.
	Value float32 `json:"value,omitempty"`

	// Label for this link.
	Label *EdgeLabel `json:"label,omitempty"`

	// LineStyle settings in this series data.
	LineStyle *LineStyle `json:"lineStyle,omitempty"`
}

// GraphCategory represents a category for data nodes.
// The categories of node, which is optional. If there is a classification of nodes,
// the category of each node can be assigned through data[i].category.
// And the style of category will also be applied to the style of nodes. categories can also be used in legend.
// https://echarts.apache.org/en/option.html#series-graph.categories
type GraphCategory struct {
	// Name of category, which is used to correspond with legend and the content of tooltip.
	Name string `json:"name"`

	// The style of this node.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// The label style of node in this category.
	Label *Label `json:"label,omitempty"`
}

// HeatMapChart is the option set for a heatmap chart.
// https://echarts.apache.org/en/option.html#series-heatmap
type HeatMapChart struct {
	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int
}

// HeatMapData
// https://echarts.apache.org/en/option.html#series-heatmap.data
type HeatMapData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// KlineData
// https://echarts.apache.org/en/option.html#series-candlestick.data
type KlineData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// KlineChart is the options set for a chandlestick chart.
// https://echarts.apache.org/en/option.html#series-candlestick
type KlineChart struct {
	// Specify bar width. Absolute value (like 10) or percentage (like '20%', according to band width) can be used. Auto adapt by default.
	BarWidth string

	// Specify bar min width. Absolute value (like 10) or percentage (like '20%', according to band width) can be used. Auto adapt by default.
	BarMinWidth string

	// Specify bar max width. Absolute value (like 10) or percentage (like '20%', according to band width) can be used. Auto adapt by default.
	BarMaxWidth string
}

// LiquidChart
// reference https://github.com/ecomfe/echarts-liquidfill
type LiquidChart struct {
	// Shape of single data.
	// Icon types provided by ECharts includes 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Shape string

	// Whether to show outline
	IsShowOutline types.Bool

	// Whether to stop animation
	IsWaveAnimation types.Bool
}

// LiquidData
// reference https://github.com/ecomfe/echarts-liquidfill
type LiquidData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// MapData
// https://echarts.apache.org/en/option.html#series-map.data
type MapData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// ParallelData
// https://echarts.apache.org/en/option.html#series-parallel.data
type ParallelData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// SankeyLink represents relationship between two data nodes.
// https://echarts.apache.org/en/option.html#series-sankey.links
type SankeyLink struct {
	// The name of source node of edge
	Source interface{} `json:"source,omitempty"`

	// The name of target node of edge
	Target interface{} `json:"target,omitempty"`

	// The value of edge, which decides the width of edge.
	Value float32 `json:"value,omitempty"`
}

// SankeyNode represents a data node.
// https://echarts.apache.org/en/option.html#series-sankey.nodes
type SankeyNode struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value string `json:"value,omitempty"`

	// Depth of the node within the chart
	Depth *int `json:"depth,omitempty"`

	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
}

// ThemeRiverData
// https://echarts.apache.org/en/option.html#series-themeRiver
type ThemeRiverData struct {
	// the time attribute of time and theme.
	Date string `json:"date,omitempty"`

	// the value of an event or theme at a time point.
	Value float64 `json:"value,omitempty"`

	// the name of an event or theme.
	Name string `json:"name,omitempty"`
}

// ToList converts the themeriver data to a list
func (trd ThemeRiverData) ToList() [3]interface{} {
	return [3]interface{}{trd.Date, trd.Value, trd.Name}
}

// WordCloudChart is the option set for a word cloud chart.
type WordCloudChart struct {
	// Shape of WordCloud
	// Optional: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow"
	Shape string

	// range of font size
	SizeRange []float32

	// range of font rotation angle
	RotationRange []float32
}

type WordCloudData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

type Chart3DData struct {
	// Name of the data item.
	Name string `json:"name,omitempty"`

	// Value of the data item.
	// []interface{}{1, 2, 3}
	Value []interface{} `json:"value,omitempty"`

	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// The style setting of the text label in a single bar.
	Label *Label `json:"label,omitempty"`
}

type TreeChart struct {
	// The layout of the tree, which can be orthogonal and radial.
	// * 'orthogonal' refer to the horizontal and vertical direction.
	// * 'radial' refers to the view that the root node as the center and each layer of nodes as the ring.
	Layout string

	// The direction of the orthogonal layout in the tree diagram.
	// * 'from left to right' or 'LR'
	// * 'from right to left' or 'RL'
	// * 'from top to bottom' or 'TB'
	// * 'from bottom to top' or 'BT'
	Orient string `json:"orient,omitempty"`

	// Whether to enable mouse zooming and translating. false by default.
	// If either zooming or translating is wanted, it can be set to 'scale' or 'move'.
	// Otherwise, set it to be true to enable both.
	Roam types.Bool `json:"roam,omitempty"`

	// Subtree collapses and expands interaction, default true.
	ExpandAndCollapse types.Bool `json:"expandAndCollapse,omitempty"`

	// The initial level (depth) of the tree. The root node is the 0th layer, then the first layer, the second layer, ... , until the leaf node.
	// This configuration item is primarily used in conjunction with collapsing and expansion interactions.
	// The purpose is to prevent the nodes from obscuring each other. If set as -1 or null or undefined, all nodes are expanded.
	InitialTreeDepth int `json:"initialTreeDepth,omitempty"`

	// The style setting of the text label in a single bar.
	Label *Label `json:"label,omitempty"`

	// Leaf node special configuration, the leaf node and non-leaf node label location is different.
	Leaves *TreeLeaves `json:"leaves,omitempty"`

	// Distance between tree component and the sides of the container.
	// value can be instant pixel value like 20;
	// It can also be a percentage value relative to container width like '20%';
	Left   string `json:"left,omitempty"`
	Right  string `json:"right,omitempty"`
	Top    string `json:"top,omitempty"`
	Bottom string `json:"bottom,omitempty"`

	// SymbolKeepAspect is whether to keep aspect for symbols in the form of path://.
	SymbolKeepAspect types.Bool
}

type TreeData struct {
	// Name of the data item.
	Name string `json:"name,omitempty"`

	// Value of the data item.
	Value interface{} `json:"value,omitempty"`

	Children []*TreeData `json:"children,omitempty"`

	// Symbol of node of this category.
	// Icon types provided by ECharts includes
	// 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Symbol string `json:"symbol,omitempty"`

	// node of this category symbol size. It can be set to single numbers like 10,
	// or use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10.
	SymbolSize interface{} `json:"symbolSize,omitempty"`

	// If set as `true`, the node is collapsed in the initialization.
	Collapsed types.Bool `json:"collapsed,omitempty"`

	// LineStyle settings in this series data.
	LineStyle *LineStyle `json:"lineStyle,omitempty"`

	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
}

type TreeMapChart struct {
	// Whether to enable animation.
	Animation types.Bool `json:"animation,omitempty"`

	// leafDepth represents how many levels are shown at most. For example, when leafDepth is set to 1, only one level will be shown.
	// leafDepth is null/undefined by default, which means that "drill down" is disabled.
	LeafDepth int `json:"leafDeapth,omitempty"`

	// Roam describes whether to enable mouse zooming and translating. false by default.
	Roam types.Bool `json:"roam,omitempty"`

	// Label decribes the style of the label in each node.
	Label *Label `json:"label,omitempty"`

	// UpperLabel is used to specify whether show label when the treemap node has children.
	UpperLabel *UpperLabel `json:"upperLabel,omitempty"`

	// ColorMappingBy specifies the rule according to which each node obtain color from color list.
	ColorMappingBy string `json:"colorMappingBy,omitempty"`

	// Levels provide configration for each node level
	Levels *[]TreeMapLevel `json:"levels,omitempty"`

	// Distance between treemap component and the sides of the container.
	// value can be instant pixel value like 20;
	// It can also be a percentage value relative to container width like '20%';
	Left   string `json:"left,omitempty"`
	Right  string `json:"right,omitempty"`
	Top    string `json:"top,omitempty"`
	Bottom string `json:"bottom,omitempty"`
}

type TreeMapNode struct {
	// Name of the tree node item.
	Name string `json:"name"`

	// Value of the tree node item.
	Value int `json:"value,omitempty"`

	Children []TreeMapNode `json:"children,omitempty"`
}

// SunBurstData holds the data structure of series-sunburst.data is like tree.
// https://echarts.apache.org/en/option.html#series-sunburst.data
type SunBurstData struct {
	// Name displayed in each sector.
	Name string `json:"name,omitempty"`
	// Value for each item. If contains children, value can be left unset, and sum of children values will be used in this case.
	Value float64 `json:"value,omitempty"`
	// ItemStyle specifies the style of the sector of the sunburst chart.
	// You can specify the style of all sectors with series.itemStyle, or specify the style of each level of sectors with
	// series.levels.itemStyle, or specify a specific style for each sector with series.data.itemStyle. The priority is
	// from low to high, that is, if series.data.itemStyle is defined, it will override series.itemStyle and series.levels.itemStyle.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
	// Tooltip configures the tool-tip settings in this series data.
	Tooltip *Tooltip `json:"tooltip,omitempty"`
	// Label configures the style of the label of the sector.
	Label *Label `json:"label,omitempty"`
	// Emphasis configures the emphasis state.
	Emphasis *Emphasis `json:"emphasis,omitempty"`
	// Children are the children nodes defined recursively.
	Children []*SunBurstData `json:"children,omitempty"`
}

// CustomChart is the options set for a custom chart.
// https://echarts.apache.org/en/option.html#series-custom
type CustomChart struct {
	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int

	// Custom series requires developers to write a render logic by themselves in JavaScript.
	// This render logic is called RenderItem. Use opts.FuncOpts to embed JavaScript.
	RenderItem types.FuncStr
}

// Progress is the options set for progress.
type Progress struct {
	// Whether to show the progress, default is false.
	Show types.Bool `json:"show,omitempty"`

	// Whether the progress overlaps when there are multiple groups of data, default is true.
	Overlap types.Bool `json:"overlap,omitempty"`

	// Width of the progress in px
	Width int `json:"width,omitempty"`

	// Whether to add round caps at the end, default is false.
	RoundCap types.Bool `json:"roundCap,omitempty"`

	// Whether to clip overflow, default is false.
	Clip types.Bool `json:"clip,omitempty"`

	// The style of progress.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
}

// Detail is the options set for detail (e.g. on a gauge).
type Detail struct {
	// Whether to show the details, default is true.
	Show types.Bool `json:"show,omitempty"`

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

	// Font size of the value in px
	FontSize int `json:"fontSize,omitempty"`

	// Line height of the text fragment.
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

	// Width of text block.
	Width int `json:"width,omitempty"`

	// Height of text block.
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

	// Determine how to display the text when it's overflow. Available when width is set.
	//
	// 'truncate' Truncate the text and trailing with ellipsis.
	// 'break' Break by word
	// 'breakAll' Break by character.
	Overflow string `json:"overflow,omitempty"`

	// Ellipsis
	Ellipsis types.Bool `json:"ellipsis,omitempty"`

	// The content formatter of value
	//
	// 1. String template
	// The template variable is {value}.
	//
	// 2. Callback function
	// The format of callback function:
	// (value: number) => string
	Formatter types.FuncStr `json:"formatter,omitempty"`

	// Value position relative to the center of chart
	// OffceCenter is provided as [x, y] where x and y are either a number (px, provided
	// as string) or a percentage.
	// Positive values move the chart value to [right, bottom], negative values vice
	// versa.
	OffsetCenter []string `json:"offsetCenter,omitempty"`
}

// Pointer is the options set for Pointer (e.g. on a gauge).
type Pointer struct {
	// Whether to show the pointer, default true.
	Show types.Bool `json:"show,omitempty"`

	// Whether to show the pointer above detail and title, default true.
	ShowAbove types.Bool `json:"ShowAbove,omitempty"`

	// Icon of the legend items.
	// Icon types provided by ECharts includes
	// 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	// An image URL example:
	//   'image://http://example.website/a/b.png'
	// A dataURI example:
	//
	// 'image://data:image/gif;base64,KOY......'
	// Icons can be set to arbitrary vector path via 'path://' in ECharts.
	// As compared with a raster image, vector paths prevent jagging and blurring when scaled, and have better control over changing colors.
	// For example:
	//
	// 'path://M30.9,53.2C16.8,...'
	Icon string `json:"icon,omitempty"`

	// Gauge
	// Value position relative to the center of chart
	// OffsetCenter is provided as [x, y] where x and y are either a number (px, provided
	// as string) or a percentage.
	// Positive values move the chart value to [right, bottom], negative values vice
	// versa.
	OffsetCenter []string `json:"offsetCenter,omitempty"`

	// The length of pointer which could be absolute value and also the percentage relative to radius, e.g. '60' or '60%'.
	Length string `json:"length,omitempty"`

	// The style of pointer.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`
}

// CustomData
// https://echarts.apache.org/en/option.html#series-custom.data
type CustomData struct {
	// Name of data item.
	Name string `json:"name,omitempty"`

	// Value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// ItemStyle settings in this series data.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Emphasis settings in this series data.
	Emphasis *Emphasis `json:"emphasis,omitempty"`

	// Tooltip settings in this series data.
	Tooltip *Tooltip `json:"tooltip,omitempty"`
}
