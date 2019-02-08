---
id: page
title: Page
sidebar_label: Page（顺序多图）
---

## API
```go
// 实例化图表
func NewPage(routers ...RouterOpts) *Page
// 新增 Page 图表，支持一次接收多个 Chart
func Add(charts ...charter) *Page 
// 负责渲染图表，支持传入多个实现了 io.Writer 接口的对象
func Render(w ...io.Writer) error
```

## Interface
```go
type charter interface {
    validateOpts()
    yieldAssets() ([]string, []string)
    chartType() string
}
```

## 预定义
> Note: 示例用到的一些变量及方法，部分重复的以后代码中不会再次列出
```go
func gaugeBase() *charts.Gauge {
    gauge := charts.NewGauge()
    gauge.SetGlobalOptions(charts.TitleOpts{Title: "Gauge-示例图"})
    m := make(map[string]interface{})
    m["工作进度"] = rand.Intn(50)
    gauge.Add("gauge", m)
    return gauge
}

func funnelBase() *charts.Funnel {
    funnel := charts.NewFunnel()
    funnel.SetGlobalOptions(charts.TitleOpts{Title: "Funnel-示例图"})
    funnel.Add("funnel", genKvData())
    return funnel
}

func heatMapBase() *charts.HeatMap {
    hm := charts.NewHeatMap()
    hm.SetGlobalOptions(charts.TitleOpts{Title: "HeatMap-示例图"})
    hm.AddXAxis(hours).AddYAxis("heatmap", genHeatMapData())
    hm.SetGlobalOptions(
        charts.YAxisOpts{Data: days, Type: "category", SplitArea: charts.SplitAreaOpts{Show: true}},
        charts.XAxisOpts{Type: "category", SplitArea: charts.SplitAreaOpts{Show: true}},
        charts.VisualMapOpts{Calculable: true, Max: 10, Min: 0,
            InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
    )
    return hm
}

func pieBase() *charts.Pie {
    pie := charts.NewPie()
    pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-示例图"})
    pie.Add("pie", genKvData())
    return pie
}

//...
```

## Demo

### Page-示例图
```go
// 结合 Overlap 的几个 xxBase() 示例函数
p := charts.NewPage()

p.Add(
    barBase(),
    esBase(),
    funnelBase(),
    gaugeBase(),
    geoBase(),
    heatMapBase(),
    klineBase(),
    lineBase(),
    liquidBase(),
    mapBase(),
    overlapBase(),
    pieBase(),
    scatterBase(),
    wcBase(),
)
f, err := os.Create("page.html")
if err != nil {
    log.Println(err)
}
p.Render(w, f)
```

![](https://user-images.githubusercontent.com/19553554/52349844-50118d80-2a62-11e9-83d5-790db128736a.png)