package charts

// Graph represents a graph chart.
type Graph struct {
	BaseConfiguration
	MultiSeries
}

//
//// GraphNode represents a data node in graph chart.
//type GraphNode struct {
//	// 数据项名称
//	Name string `json:"name,omitempty"`
//	// 节点的初始 x 值。在布局为 "none" 时生效
//	X float32 `json:"x,omitempty"`
//	// 节点的初始 y 值。在布局为 "none" 时生效
//	Y float32 `json:"y,omitempty"`
//	// 数据项值
//	Value float32 `json:"value,omitempty"`
//	// 节点在力引导布局中是否固定
//	Fixed bool `json:"fixed,omitempty"`
//	// 数据项所在类目的 index
//	Category int `json:"category,omitempty"`
//	// 该类目节点标记的图形，可选
//	// 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
//	Symbol string `json:"symbol,omitempty"`
//	// 该类目节点标记的大小，可以设置成诸如 10 这样单一的数字，
//	// 也可以用数组分开表示宽和高，例如 [20, 10] 表示标记宽为20，高为10
//	SymbolSize interface{} `json:"symbolSize,omitempty"`
//	// 该节点的样式
//	ItemStyle ItemStyleOpts `json:"itemStyle,omitempty"`
//}
//
//// GraphLink represents relationship between two data nodes.
//type GraphLink struct {
//	// 边的源节点名称的字符串，也支持使用数字表示源节点的索引
//	Source interface{} `json:"source,omitempty"`
//	// 边的目标节点名称的字符串，也支持使用数字表示源节点的索引
//	Target interface{} `json:"target,omitempty"`
//	// 边的数值，可以在力引导布局中用于映射到边的长度
//	Value float32 `json:"value,omitempty"`
//}
//
//// GraphCategory represents a category for data nodes.
//type GraphCategory struct {
//	// 类目名称，用于和 legend 对应以及格式化 tooltip 的内容
//	Name string `json:"name"`
//	// 该类目节点标签的样式
//	Label LabelTextOpts `json:"label"`
//}
//
//// GraphForce is the option set for graph force layout.
//type GraphForce struct {
//	// 进行力引导布局前的初始化布局，初始化布局会影响到力引导的效果
//	//InitLayout string `json:"initLayout,omitempty"`
//	// 节点之间的斥力因子。
//	// 支持设置成数组表达斥力的范围，此时不同大小的值会线性映射到不同的斥力。
//	// 值越大则斥力越大。默认为 50
//	Repulsion float32 `json:"repulsion,omitempty"`
//	// 节点受到的向中心的引力因子。该值越大节点越往中心点靠拢
//	// 默认为 0.1
//	Gravity float32 `json:"gravity,omitempty"`
//	// 边的两个节点之间的距离，这个距离也会受 Repulsion 影响。
//	// 支持设置成数组表达边长的范围，此时不同大小的值会线性映射到不同的长度。值越小则长度越长。
//	// 如下示例：值最大的边长度会趋向于 10，值最小的边长度会趋向于 50
//	// edgeLength: [10, 50]
//	EdgeLength float32 `json:"edgeLength,omitempty"`
//}
//
//func (Graph) Type() string { return types.ChartGraph }
//
//// NewGraph creates a new graph chart.
//func NewGraph() *Graph {
//	chart := new(Graph)
//	chart.initBaseConfiguration()
//	return chart
//}
//
//// Add adds new data sets.
//func (c *Graph) Add(name string, nodes []GraphNode, links []GraphLink, opts ...SeriesOpts) *Graph {
//	series := SingleSeries{Name: name, Type: types.ChartGraph, Links: links, Data: nodes}
//	series.configureSeriesOpts(opts...)
//	c.MultiSeries = append(c.MultiSeries, series)
//	return c
//}
//
//// SetGlobalOptions sets options for the Graph instance.
//func (c *Graph) SetGlobalOptions(opts ...GlobalOpts) *Graph {
//	c.BaseConfiguration.setBaseGlobalOptions(opts...)
//	return c
//}
//
//func (c *Graph) validateOpts() {
//	// 如果没有设置布局 则默认为 "force"
//	for i := 0; i < len(c.Series); i++ {
//		if c.Series[i].Layout == "" {
//			c.Series[i].Layout = "force"
//		}
//	}
//	c.validateAssets(c.AssetsHost)
//}
//
//// Render renders the chart and writes the output to given writers.
//func (c *Graph) Render(w ...io.Writer) error {
//	c.insertSeriesColors(c.appendColor)
//	c.validateOpts()
//	return renderToWriter(c, "chart", []string{`"force":{},?`}, w...)
//}
