---
id: liquid
title: Liquid
sidebar_label: Liquid（水球图）
---

> 主要用来突出数据的百分比

## API
```go
// 实例化图表
func NewLiquid(routers ...HTTPRouter) *Liquid
// 新增数据及配置项
func Add(name string, data map[string]interface{}, options ...seriesOptser) *Liquid
// 新增 JS 函数
func AddJSFuncs(fn ...string)
// 设置全局配置项
func SetGlobalOptions(options ...globalOptser)
// 设置 Series 配置项
func SetSeriesOptions(options ...seriesOptser)
// 负责渲染图表，支持传入多个实现了 io.Writer 接口的对象
func Render(w ...io.Writer)
```

## 专属 Options
> 在 `SetSeriesOptions` 中设置
```go
type LiquidOpts struct {
    // 水球图形状，可选
    // "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
    Shape           string
    // 是否显示水球轮廓
    IsShowOutline   bool
    // 是否停止动画
    IsWaveAnimation bool
}
```

## Demo

### Liquid-示例图
```go
liquid := charts.NewLiquid()
liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-示例图"})
liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
    charts.LiquidOpts{IsWaveAnimation: true},
)
```
![](https://user-images.githubusercontent.com/19553554/52347061-03c34f00-2a5c-11e9-9bd3-5176f9ea24a6.gif)


### Liquid-展示 Label
```go
liquid := charts.NewLiquid()
liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-展示 Label"})
liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
    charts.LabelTextOpts{Show: true},
    charts.LiquidOpts{IsWaveAnimation: true},
)
```
![](https://user-images.githubusercontent.com/19553554/52347117-248ba480-2a5c-11e9-8402-5a94054dca50.gif)


### Liquid-显示轮廓
```go
liquid := charts.NewLiquid()
liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-显示轮廓"})
liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
    charts.LabelTextOpts{Show: true},
    charts.LiquidOpts{IsShowOutline: true, IsWaveAnimation: true},
)
```
![](https://user-images.githubusercontent.com/19553554/52347152-3bca9200-2a5c-11e9-9d76-9fdd634eee48.gif)


### Liquid-关闭动画
```go
liquid := charts.NewLiquid()
liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-关闭动画"})
liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
    charts.LabelTextOpts{Show: true},
    charts.LiquidOpts{IsShowOutline: true},
)
```
![](https://user-images.githubusercontent.com/19553554/52347196-5270e900-2a5c-11e9-9352-1b5d06237b1e.png)


### Liquid-形状(diamond)
```go
liquid := charts.NewLiquid()
liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-形状(diamond)"})
liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
    charts.LiquidOpts{Shape: "diamond", IsWaveAnimation: true})
```
![](https://user-images.githubusercontent.com/19553554/52347239-69afd680-2a5c-11e9-811a-464bf795206b.gif)


### Liquid-形状(pin)
```go
liquid := charts.NewLiquid()
liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-形状(pin)"})
liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
    charts.LiquidOpts{Shape: "pin", IsWaveAnimation: true})
```
![](https://user-images.githubusercontent.com/19553554/52347274-80562d80-2a5c-11e9-87d3-a83b15117eb7.gif)


### Liquid-形状(arrow)
```go
liquid := charts.NewLiquid()
liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-形状(arrow)"})
liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
    charts.LiquidOpts{Shape: "arrow", IsWaveAnimation: true})
```
![](https://user-images.githubusercontent.com/19553554/52347342-a380dd00-2a5c-11e9-855c-ea751bc95482.gif)


### Liquid-形状(triangle)
```go
liquid := charts.NewLiquid()
liquid.SetGlobalOptions(charts.TitleOpts{Title: "Liquid-形状(triangle)"})
liquid.Add("liquid", []float32{0.3, 0.4, 0.5},
    charts.LiquidOpts{Shape: "triangle", IsWaveAnimation: true})
```
![](https://user-images.githubusercontent.com/19553554/52347361-af6c9f00-2a5c-11e9-8b70-0025c1a60b09.gif)
