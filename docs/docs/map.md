---
id: map
title: Map
sidebar_label: Map（地图）
---

> 地图主要用于地理区域数据的可视化

## API
```go
// 实例化图表
func NewMap(mapType string, routers ...RouterOpts) *Map
// 新增数据及配置项
func Add(name string, data map[string]float32, options ...seriesOptser) *Map
// 新增 JS 函数
func AddJSFuncs(fn ...string)
// 设置全局配置项
func SetGlobalOptions(options ...globalOptser)
// 设置全部 Series 配置项
func SetSeriesOptions(options ...seriesOptser)
// 负责渲染图表，支持传入多个实现了 io.Writer 接口的对象
func Render(w ...io.Writer) error
```

## 预定义
> Note: 示例用到的一些变量及方法，部分重复的以后代码中不会再次列出
```go
mapData = map[string]float32{
    "北京":   float32(rand.Intn(150)),
    "上海":   float32(rand.Intn(150)),
    "深圳":   float32(rand.Intn(150)),
    "辽宁":   float32(rand.Intn(150)),
    "青岛":   float32(rand.Intn(150)),
    "山西":   float32(rand.Intn(150)),
    "陕西":   float32(rand.Intn(150)),
    "乌鲁木齐": float32(rand.Intn(150)),
    "齐齐哈尔": float32(rand.Intn(150)),
}

guangdongData = map[string]float32{
    "深圳市": float32(rand.Intn(150)),
    "广州市": float32(rand.Intn(150)),
    "湛江市": float32(rand.Intn(150)),
    "汕头市": float32(rand.Intn(150)),
    "东莞市": float32(rand.Intn(150)),
    "佛山市": float32(rand.Intn(150)),
    "云浮市": float32(rand.Intn(150)),
    "肇庆市": float32(rand.Intn(150)),
    "梅州市": float32(rand.Intn(150)),
}

shantouData = map[string]float32{
    "澄海区": float32(rand.Intn(150)),
    "潮阳区": float32(rand.Intn(150)),
    "潮南区": float32(rand.Intn(150)),
    "南澳县": float32(rand.Intn(150)),
}
```

## Demo

### Map-示例图
```go
mc := charts.NewMap("china")
mc.SetGlobalOptions(charts.TitleOpts{Title: "Map-示例图"})
mc.Add("map", mapData)
```
![](https://user-images.githubusercontent.com/19553554/52347776-bf38b300-2a5d-11e9-85e9-2307a50ac692.gif)


### Map-展示 Label
```go
mc := charts.NewMap("china")
mc.SetGlobalOptions(charts.TitleOpts{Title: "Map-展示 Label"})
mc.Add("map", mapData, charts.LabelTextOpts{Show: true})
```
![](https://user-images.githubusercontent.com/19553554/52347807-ce1f6580-2a5d-11e9-8b2d-84dfc244d160.png)


### Map-设置 VisualMap
```go
mc := charts.NewMap("china")
mc.SetGlobalOptions(
    charts.TitleOpts{Title: "Map-设置 VisualMap"},
    charts.VisualMapOpts{Calculable: true},
)
mc.Add("map", mapData)
```
![](https://user-images.githubusercontent.com/19553554/52347856-e8594380-2a5d-11e9-800f-f5b3f991fb7a.gif)


### Map-广东地图
```go
mc := charts.NewMap("广东")
mc.SetGlobalOptions(
    charts.TitleOpts{Title: "Map-广东地图"},
    charts.VisualMapOpts{Calculable: true,
        InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
)
mc.Add("map", guangdongData)
```
![](https://user-images.githubusercontent.com/19553554/52347915-0a52c600-2a5e-11e9-8039-41268238576c.gif)


### Map-汕头地图
```go
mc := charts.NewMap("汕头")
mc.SetGlobalOptions(
    charts.TitleOpts{Title: "Map-汕头地图"},
    charts.VisualMapOpts{Calculable: true,
        InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
)
mc.Add("map", shantouData)
```
![](https://user-images.githubusercontent.com/19553554/52347939-1cccff80-2a5e-11e9-9c27-d58704c43805.png)


### Map-设置风格
```go
import "github.com/go-echarts/go-echarts/common"

mc := charts.NewMap("china")
mc.SetGlobalOptions(
    charts.InitOpts{Theme: common.ThemeType.Macarons},
    charts.TitleOpts{Title: "Map-设置风格"},
    charts.VisualMapOpts{Calculable: true, Max: 150},
)
mc.Add("map", mapData)
```
![](https://user-images.githubusercontent.com/19553554/52347966-2a828500-2a5e-11e9-8732-84ed1ad30deb.png)
