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
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// FunnelData
type FunnelData struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// GaugeData
type GaugeData struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// GraphOpts is the option set for graph chart.
//type GraphChart struct {
//	//图的布局。可选：
//	// "none" 不采用任何布局，使用节点中提供的 x， y 作为节点的位置。
//	// "circular" 采用环形布局
//	// "force" 采用力引导布局
//	Layout string
//	// "force", "circular" 布局详细配置项
//	Force GraphForce
//	// 是否开启鼠标缩放和平移漫游。默认不开启。
//	Roam bool
//	// 是否在鼠标移到节点上的时候突出显示节点以及节点的边和邻接节点
//	FocusNodeAdjacency bool
//	//
//	Categories []GraphCategory
//}

// HeatMapChart is the option set for a heatmap chart.
type HeatMapChart struct {
	//使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	//使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

// LineChart is the options set for a line chart.
type LineChart struct {
	// 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
	Stack string
	// 曲线是否平滑
	Smooth bool
	// 是否使用阶梯图
	Step bool
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
	// 是否连接空数据。
	ConnectNulls bool
}

// LineData
type LineData struct {
	// Name
	Name string `json:"name,omitempty"`
	// Value
	Value interface{} `json:"value,omitempty"`
	// Symbol
	Symbol string `json:"symbol,omitempty"`
	// SymbolSize
	SymbolSize int `json:"symbolSize,omitempty"`

	// Index of x axis to combine with, which is useful for multiple x axes in one chart.
	XAxisIndex int

	// Index of y axis to combine with, which is useful for multiple y axes in one chart.
	YAxisIndex int
}

// LiquidChart
type LiquidChart struct {
	// 水球图形状，可选
	// "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Shape string
	// 是否显示水球轮廓
	IsShowOutline bool
	// 是否停止动画
	IsWaveAnimation bool
}

// LiquidData
// reference https://github.com/ecomfe/echarts-liquidfill
type LiquidData struct {
	// Name
	Name string `json:"name,omitempty"`
	// Value
	Value interface{} `json:"value,omitempty"`
}

// PieChart is the option set for a pie chart.
type PieChart struct {
	// 是否展示成南丁格尔图，通过半径区分数据大小。可选择两种模式：
	// 1."radius": 扇区圆心角展现数据的百分比，半径展现数据的大小。
	// 2."area": 所有扇区圆心角相同，仅通过半径展现数据大小。
	RoseType string
	// 饼图的中心（圆心）坐标，数组的第一项是横坐标，第二项是纵坐标。
	// 支持设置成百分比，设置成百分比时第一项是相对于容器宽度，第二项是相对于容器高度
	// 使用示例
	// 设置成绝对的像素值: center: [400, 300]
	// 设置成相对的百分比: center: ['50%', '50%']
	// 默认 ["50%", "50%"]
	Center interface{}
	// 饼图的半径。可以为如下类型：
	// 1.number：直接指定外半径值。
	// 2.string：例如，'20%'，表示外半径为可视区尺寸（容器高宽中较小一项）的 20% 长度。
	// 3.Array.<number|string>：数组的第一项是内半径，第二项是外半径。
	// 每一项遵从上述 number string 的描述。
	// 默认 [0, "75%"]
	Radius interface{}
}

// PieData
type PieData struct {
	Name       string      `json:"name,omitempty"`
	Value      interface{} `json:"value,omitempty"`
	Selected   bool        `json:"selected,omitempty"`
	*Label     `json:"label,omitempty"`
	*ItemStyle `json:"itemStyle,omitempty"`
	*Tooltip   `json:"tooltip,omitempty"`
}

// ScatterChart is the option set for a scatter chart.
type ScatterChart struct {
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
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
	// 边的源节点名称的字符串，也支持使用数字表示源节点的索引
	Source interface{} `json:"source,omitempty"`
	// 边的目标节点名称的字符串，也支持使用数字表示源节点的索引
	Target interface{} `json:"target,omitempty"`
	// 边的数值，可以在力引导布局中用于映射到边的长度
	Value float32 `json:"value,omitempty"`
}

// SankeyNode represents a data node.
type SankeyNode struct {
	// 数据项名称
	Name string `json:"name,omitempty"`
	// 数据项值
	Value string `json:"value,omitempty"`
}

// ThemeRiverChartItem
type ThemeRiverData struct {
}

// RadarChartItem
type RadarData struct {
	// Name
	Name string `json:"name,omitempty"`
	// Value
	Value interface{} `json:"value,omitempty"`
}

// KlineData
type KlineData struct {
}

// ScatterData
type ScatterData struct {
	// Name
	Name string `json:"name,omitempty"`

	// Value
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
}

// HeatMapData
type HeatMapData struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// WordCloudData
type WordCloudData struct {
	Name  string      `json:"name,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// GeoData
type GeoData struct {
}

// ParallelData
type ParallelData struct {
}
