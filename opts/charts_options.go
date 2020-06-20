package opts

type BarChart struct {
	// 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
	Stack string
	// 不同系列的柱间距离，为百分比（如 "30%"，表示柱子宽度的 30%）。
	// 如果想要两个系列的柱子重叠，可以设置 barGap 为 "-100%"。这在用柱子做背景的时候有用。
	// 在同一坐标系上，此属性会被多个 "bar" 系列共享。
	// 此属性应设置于此坐标系中最后一个 "bar" 系列上才会生效，并且是对此坐标系中所有 "bar" 系列生效
	// 默认 "30%"
	BarGap string
	// 同一系列的柱间距离，默认为类目间距的 20%，可设固定值
	// 在同一坐标系上，此属性会被多个 "bar" 系列共享。
	// 此属性应设置于此坐标系中最后一个 "bar" 系列上才会生效，并且是对此坐标系中所有 "bar" 系列生效
	BarCategoryGap string
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

type BarChartItem struct {
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

// HeatMapOpts is the option set for a heatmap chart.
type HeatMapChart struct {
	//使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	//使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

// LineOpts is the options set for a line chart.
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

type LiquidChart struct {
	// 水球图形状，可选
	// "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Shape string
	// 是否显示水球轮廓
	IsShowOutline bool
	// 是否停止动画
	IsWaveAnimation bool
}

// PieOpts is the option set for a pie chart.
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

// ScatterOpts is the option set for a scatter chart.
type ScatterChart struct {
	// 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
	XAxisIndex int
	// 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
	YAxisIndex int
}

// Bar3DOpts is the option set for a 3D bar chart.
type Bar3DChart struct {
	Shading string
}

// WordCloudOpts is the option set for a word cloud chart.
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
