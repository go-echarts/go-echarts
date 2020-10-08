package opts

// BarChart
type BarChart struct {
	// Name of stack. On the same category axis, the series with the
	// same stack name would be put on top of each other.
	Stack string

	// The gap between bars between different series, is a percent value like '30%',
	// which means 30% of the bar width.
	// Set barGap as '-100%' can overlap bars that belong to different series,
	// which is useful when making a series of bar be background.
	// In a single coordinate system, this attribute is shared by multiple 'bar' series.
	// This attribute should be set on the last 'bar' series in the coordinate system,
	// then it will be adopted by all 'bar' series in the coordinate system.
	BarGap string

	// The bar gap of a single series, defaults to be 20% of the category gap,
	// can be set as a fixed value.
	// In a single coordinate system, this attribute is shared by multiple 'bar' series.
	// This attribute should be set on the last 'bar' series in the coordinate system,
	// then it will be adopted by all 'bar' series in the coordinate system.
	BarCategoryGap string

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int
}

// BarData
type BarData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// The style setting of the text label in a single bar.
	*Label `json:"label,omitempty"`

	// ItemStyle settings in this series data.
	*ItemStyle `json:"itemStyle,omitempty"`

	// Tooltip settings in this series data.
	*Tooltip `json:"tooltip,omitempty"`
}

// BoxPlotData
type BoxPlotData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// The style setting of the text label in a single bar.
	*Label `json:"label,omitempty"`

	// ItemStyle settings in this series data.
	*ItemStyle `json:"itemStyle,omitempty"`

	// Emphasis settings in this series data.
	*Emphasis `json:"emphasis,omitempty"`

	// Tooltip settings in this series data.
	*Tooltip `json:"tooltip,omitempty"`
}

// EffectScatterData
type EffectScatterData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// FunnelData
type FunnelData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GaugeData
type GaugeData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GraphChart is the option set for graph chart.
type GraphChart struct {
	// Graph layout.
	// * 'none' No any layout, use x, y provided in node as the position of node.
	// * 'circular' Adopt circular layout, see the example Les Miserables.
	// * 'force' Adopt force-directed layout, see the example Force, the
	// detail about configrations of layout are in graph.force
	Layout string

	// Force is the option set for graph force layout.
	Force GraphForce

	// Whether to enable mouse zooming and translating. false by default.
	// If either zooming or translating is wanted, it can be set to 'scale' or 'move'.
	// Otherwise, set it to be true to enable both.
	Roam bool

	// Whether to focus/highlight the hover node and it's adjacencies.
	FocusNodeAdjacency bool

	// The categories of node, which is optional. If there is a classification of nodes,
	// the category of each node can be assigned through data[i].category.
	// And the style of category will also be applied to the style of nodes. categories can also be used in legend.
	Categories []GraphCategory
}

// HeatMapChart is the option set for a heatmap chart.
type HeatMapChart struct {
	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int
}

// LineChart is the options set for a line chart.
type LineChart struct {
	// If stack the value. On the same category axis, the series with the same stack name would be put on top of each other.
	// The effect of the below example could be seen through stack switching of toolbox on the top right corner:
	Stack string

	// Whether to show as smooth curve.
	// If is typed in boolean, then it means whether to enable smoothing. If is
	// typed in number, valued from 0 to 1, then it means smoothness. A smaller value makes it less smooth.
	Smooth bool

	// Whether to show as a step line. It can be true, false. Or 'start', 'middle', 'end'.
	// Which will configure the turn point of step line.
	Step bool

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int

	// Whether to connect the line across null points.
	ConnectNulls bool
}

// LineData
type LineData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// Symbol of single data.
	// Icon types provided by ECharts includes 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Symbol string `json:"symbol,omitempty"`

	// single data symbol size. It can be set to single numbers like 10, or
	// use an array to represent width and height. For example, [20, 10] means symbol width is 20, and height is10
	SymbolSize int `json:"symbolSize,omitempty"`

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int
}

// LiquidChart
type LiquidChart struct {
	// Shape of single data.
	// Icon types provided by ECharts includes 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	// It can be set to an image with 'image://url' , in which URL is the link to an image, or dataURI of an image.
	Shape string

	// Whether to show outline
	IsShowOutline bool

	// Whether to stop animation
	IsWaveAnimation bool
}

// LiquidData
// reference https://github.com/ecomfe/echarts-liquidfill
type LiquidData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// PieChart is the option set for a pie chart.
type PieChart struct {
	// Whether to show as Nightingale chart, which distinguishs data through radius. There are 2 optional modes:
	// * 'radius' Use central angle to show the percentage of data, radius to show data size.
	// * 'area' All the sectors will share the same central angle, the data size is shown only through radiuses.
	RoseType string

	// Center position of Pie chart, the first of which is the horizontal position, and the second is the vertical position.
	// Percentage is supported. When set in percentage, the item is relative to the container width,
	// and the second item to the height.
	//
	// Example:
	//
	// Set to absolute pixel values ->> center: [400, 300]
	// Set to relative percent ->> center: ['50%', '50%']
	Center interface{}

	// Radius of Pie chart. Value can be:
	// * number: Specify outside radius directly.
	// * string: For example, '20%', means that the outside radius is 20% of the viewport
	// size (the little one between width and height of the chart container).
	//
	// Array.<number|string>: The first item specifies the inside radius, and the
	// second item specifies the outside radius. Each item follows the definitions above.
	Radius interface{}
}

// PieData
type PieData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// Whether the data item is selected.
	Selected bool `json:"selected,omitempty"`

	// The label configuration of a single sector.
	*Label `json:"label,omitempty"`

	// Graphic style of , emphasis is the style when it is highlighted, like being hovered by mouse, or highlighted via legend connect.
	*ItemStyle `json:"itemStyle,omitempty"`

	// tooltip settings in this series data.
	*Tooltip `json:"tooltip,omitempty"`
}

// ScatterChart is the option set for a scatter chart.
type ScatterChart struct {
	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of x axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int
}

// Bar3DChart is the option set for a 3D bar chart.
type Bar3DChart struct {
	Shading string
}

// WordCloudChart is the option set for a word cloud chart.
type WordCloudChart struct {
	// shape of WordCloud
	// Optional: "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow"
	Shape string

	// range of font size
	SizeRange []float32

	// range of font rotation angle
	RotationRange []float32
}

// SankeyLink represents relationship between two data nodes.
type SankeyLink struct {
	// The name of source node of edge
	Source interface{} `json:"source,omitempty"`

	// The name of target node of edge
	Target interface{} `json:"target,omitempty"`

	// The value of edge, which decides the width of edge.
	Value float32 `json:"value,omitempty"`
}

// SankeyNode represents a data node.
type SankeyNode struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value string `json:"value,omitempty"`
}

// ThemeRiverChartItem
type ThemeRiverData struct {
	// the time attribute of time and theme.
	Date string

	// the value of an event or theme at a time point.
	Value float64

	// the name of an event or theme.
	Name string
}

func (trd ThemeRiverData) ToList() [3]interface{} {
	return [3]interface{}{trd.Date, trd.Value, trd.Name}
}

// RadarData
type RadarData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// KlineData
type KlineData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// ScatterData
type ScatterData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`

	// Symbol
	Symbol string `json:"symbol,omitempty"`

	// SymbolSize
	SymbolSize int `json:"symbolSize,omitempty"`

	// SymbolRotate
	SymbolRotate int `json:"symbolRotate,omitempty"`

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int
}

// MapData
type MapData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// HeatMapData
type HeatMapData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// WordCloudData
type WordCloudData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GeoData
type GeoData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// ParallelData
type ParallelData struct {
	// The name of data item.
	Name string `json:"name,omitempty"`

	// The value of a single data item.
	Value interface{} `json:"value,omitempty"`
}

// GraphNode represents a data node in graph chart.
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
	Fixed bool `json:"fixed,omitempty"`

	// Index of category which the data item belongs to.
	Category int `json:"category,omitempty"`

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
}

// GraphLink represents relationship between two data nodes.
type GraphLink struct {
	// A string representing the name of source node on edge. Can also be a number representing the node index.
	Source interface{} `json:"source,omitempty"`

	// A string representing the name of target node on edge. Can also be a number representing node index.
	Target interface{} `json:"target,omitempty"`

	// value of edge, can be mapped to edge length in force graph.
	Value float32 `json:"value,omitempty"`
}

// GraphCategory represents a category for data nodes.
type GraphCategory struct {
	// Name of category, which is used to correspond with legend and the content of tooltip.
	Name string `json:"name"`

	// The label style of node in this category.
	Label *Label `json:"label,omitempty"`
}
