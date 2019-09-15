---
id: sankey
title: Sankey
sidebar_label: Sankey（桑基图）
---

> 是一种特殊的流图（可以看作是有向无环图）。 它主要用来表示原材料、能量等如何从最初形式经过中间过程的加工或转化达到最终状态

## API
```go
// 实例化图表
func NewSankey(routers ...RouterOpts) *Sankey
// 新增数据及配置项
func Add(name string, nodes []SankeyNode, links []SankeyLink, options ...seriesOptser) *Sankey
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

### SankeyLink
> 桑基图节点间的关系数据
```go
type SankeyLink struct {
    // 边的源节点名称的字符串，也支持使用数字表示源节点的索引
    Source interface{} `json:"source,omitempty"`
    // 边的目标节点名称的字符串，也支持使用数字表示源节点的索引
    Target interface{} `json:"target,omitempty"`
    // 边的数值，可以在力引导布局中用于映射到边的长度
    Value float32 `json:"value,omitempty"`
}
```

### SankeyNode
> 桑基图节点数据
```go
type SankeyNode struct {
    // 数据项名称
    Name string `json:"name,omitempty"`
    // 数据项值
    Value string `json:"value,omitempty"`
}
```

## 预定义
> Note: 示例用到的一些变量及方法，部分重复的以后代码中不会再次列出
```go
var sankeyNode = []charts.SankeyNode{
    {Name: "category1"},
    {Name: "category2"},
    {Name: "category3"},
    {Name: "category4"},
    {Name: "category5"},
    {Name: "category6"},
}

var sankeyLink = []charts.SankeyLink{
    {Source: "category1", Target: "category2", Value: 10},
    {Source: "category2", Target: "category3", Value: 15},
    {Source: "category3", Target: "category4", Value: 20},
    {Source: "category5", Target: "category6", Value: 25},
}
```

## Demo

### Sankey-示例图
```go
sankey := charts.NewSankey()
sankey.SetGlobalOptions(charts.TitleOpts{Title: "Sankey-示例图"})
sankey.Add("sankey", sankeyNode, sankeyLink, charts.LabelTextOpts{Show: true})
```
![](https://user-images.githubusercontent.com/19553554/52802192-6cdd3f00-30ba-11e9-8ada-e0b0145381cc.png)


### Sankey-官方示例
=> [energy.json](https://github.com/go-echarts/go-echarts/tree/master/example/fixtures/energy.json) <=>
```go
import "github.com/gobuffalo/packr"

sankey := charts.NewSankey()
sankey.SetGlobalOptions(charts.TitleOpts{Title: "Sankey-官方示例"})
box := packr.NewBox(path.Join(".", "fixtures"))
f, err := box.Find("energy.json")
if err != nil {
    log.Fatal(err)
}
type Data struct {
    Nodes []charts.SankeyNode
    Links []charts.SankeyLink
}

var data Data
if err := json.Unmarshal(f, &data); err != nil {
    fmt.Println(err)
}
sankey.Add("sankey", data.Nodes, data.Links,
    charts.LineStyleOpts{Curveness: 0.5, Color: "source"},
    charts.LabelTextOpts{Show: true},
)
```
![](https://user-images.githubusercontent.com/19553554/52802261-8d0cfe00-30ba-11e9-8ae7-ae0773770a59.gif)
