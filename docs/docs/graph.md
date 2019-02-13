---
id: graph
title: Graph
sidebar_label: Graph（关系图）
---

> 用于展现节点以及节点之间的关系数据

## API
```go
// 实例化图表
func NewGraph(routers ...RouterOpts) *Graph
// 新增数据及配置项
func Add(name string, nodes []GraphNode, links []GraphLink, options ...seriesOptser) *Graph
// 新增 JS 函数
func AddJSFuncs(fn ...string)
// 设置全局配置项
func SetGlobalOptions(options ...globalOptser)
// 设置全部 Series 配置项
func SetSeriesOptions(options ...seriesOptser)
// 负责渲染图表，支持传入多个实现了 io.Writer 接口的对象
func Render(w ...io.Writer) error
```

## 专属 Options
> 在 `SetSeriesOptions` 中设置
### GraphNode
> 关系图节点数据
```go
type GraphNode struct {
    // 数据项名称
    Name string `json:"name,omitempty"`
    // 节点的初始 x 值。在布局为 "none" 时生效
    X float32 `json:"x,omitempty"`
    // 节点的初始 y 值。在布局为 "none" 时生效
    Y float32 `json:"y,omitempty"`
    // 数据项值
    Value float32 `json:"value,omitempty"`
    // 节点在力引导布局中是否固定
    Fixed bool `json:"fixed,omitempty"`
    // 数据项所在类目的 index
    Category int `json:"category,omitempty"`
    // 该类目节点标记的图形，可选
    // 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
    Symbol string `json:"symbol,omitempty"`
    // 该类目节点标记的大小，可以设置成诸如 10 这样单一的数字，
    // 也可以用数组分开表示宽和高，例如 [20, 10] 表示标记宽为20，高为10
    SymbolSize interface{}   `json:"symbolSize,omitempty"`
    // 该节点的样式
    ItemStyle  ItemStyleOpts `json:"itemStyle,omitempty"`
}
```

### GraphLink
> 关系图节点间的关系数据
```go
type GraphLink struct {
    // 边的源节点名称的字符串，也支持使用数字表示源节点的索引
    Source interface{} `json:"source,omitempty"`
    // 边的目标节点名称的字符串，也支持使用数字表示源节点的索引
    Target interface{} `json:"target,omitempty"`
    // 边的数值，可以在力引导布局中用于映射到边的长度
    Value float32 `json:"value,omitempty"`
}
```

### GraphCategory
> 节点分类的类目
```go
type GraphCategory struct {
    // 类目名称，用于和 legend 对应以及格式化 tooltip 的内容
    Name string `json:"name"`
    // 该类目节点标签的样式
    Label LabelTextOpts `json:"label"`
}
```

### GraphForce
> 力引导布局相关的配置项
```go
type GraphForce struct {
    // 进行力引导布局前的初始化布局，初始化布局会影响到力引导的效果
    //InitLayout string `json:"initLayout,omitempty"`
    // 节点之间的斥力因子。
    // 支持设置成数组表达斥力的范围，此时不同大小的值会线性映射到不同的斥力。
    // 值越大则斥力越大。默认为 50
    Repulsion float32 `json:"repulsion,omitempty"`
    // 节点受到的向中心的引力因子。该值越大节点越往中心点靠拢
    // 默认为 0.1
    Gravity float32 `json:"gravity,omitempty"`
    // 边的两个节点之间的距离，这个距离也会受 Repulsion 影响。
    // 支持设置成数组表达边长的范围，此时不同大小的值会线性映射到不同的长度。值越小则长度越长。
    // 如下示例：值最大的边长度会趋向于 10，值最小的边长度会趋向于 50
    // edgeLength: [10, 50]
    EdgeLength float32 `json:"edgeLength,omitempty"`
}
```

### GraphOpts
```go
type GraphOpts struct {
    //图的布局。可选：
    // "none" 不采用任何布局，使用节点中提供的 x， y 作为节点的位置。
    // "circular" 采用环形布局
    // "force" 采用力引导布局
    Layout string
    // "force", "circular" 布局详细配置项
    Force GraphForce
    // 是否开启鼠标缩放和平移漫游。默认不开启。
    Roam bool
    // 是否在鼠标移到节点上的时候突出显示节点以及节点的边和邻接节点
    FocusNodeAdjacency bool
    //
    Categories []GraphCategory
}
```

## 预定义
> Note: 示例用到的一些变量及方法，部分重复的以后代码中不会再次列出
```go
var graphNodes = []charts.GraphNode{
    {Name: "节点1"},
    {Name: "节点2"},
    {Name: "节点3"},
    {Name: "节点4"},
    {Name: "节点5"},
    {Name: "节点6"},
    {Name: "节点7"},
    {Name: "节点8"},
}

func genLinks() []charts.GraphLink {
    links := make([]charts.GraphLink, 0)
    for i := 0; i < len(graphNodes); i++ {
        for j := 0; j < len(graphNodes); j++ {
            links = append(links,
                charts.GraphLink{Source: graphNodes[i].Name, Target: graphNodes[j].Name})
        }
    }
    return links
}
```

## Demo

### Graph-示例图
```go
graph := charts.NewGraph()
graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-示例图"})
graph.Add("graph", graphNodes, genLinks(),
    charts.GraphOpts{Force: charts.GraphForce{Repulsion: 8000}},
)
```
![](https://user-images.githubusercontent.com/19553554/52727715-c8db9100-2ff0-11e9-8b58-f4a224a138fd.png)


### Graph-布局(Circular)
```go
graph := charts.NewGraph()
graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-布局(Circular)"})
graph.Add("graph", graphNodes, genLinks(),
    charts.GraphOpts{Layout: "circular", Force: charts.GraphForce{Repulsion: 8000}},
    charts.LabelTextOpts{Show: true, Position: "right"},
)
```
![](https://user-images.githubusercontent.com/19553554/52727731-d42ebc80-2ff0-11e9-8a49-041a8f2d730c.png)


### Graph-npm package 依赖关系
=> [npmdepgraph.json](https://github.com/chenjiandongx/go-echarts/tree/master/example/fixtures/npmdepgraph.json) <=>
```go
import "github.com/gobuffalo/packr"

graph := charts.NewGraph()
graph.SetGlobalOptions(charts.TitleOpts{Title: "Graph-npm package 依赖关系"})
box := packr.NewBox(path.Join(".", "fixtures"))
f, err := box.Find("npmdepgraph.json")
if err != nil {
    log.Fatal(err)
}
type Data struct {
    Nodes []charts.GraphNode
    Links []charts.GraphLink
}

var data Data
if err := json.Unmarshal(f, &data); err != nil {
    fmt.Println(err)
}
graph.Add("graph", data.Nodes, data.Links,
    charts.GraphOpts{Layout: "none", Roam: true, FocusNodeAdjacency: true},
    charts.EmphasisOpts{Label: charts.LabelTextOpts{Show: true, Position:"left", Color:"black"}},
    charts.LineStyleOpts{Curveness: 0.3},
)
```
![](https://user-images.githubusercontent.com/19553554/52727805-f7f20280-2ff0-11e9-91ab-cd99848e3127.gif)
