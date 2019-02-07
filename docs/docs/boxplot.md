---
id: boxPlot
title: BoxPlot
sidebar_label: BoxPlot（箱线图）
---

> 箱形图是一种用作显示一组数据分散情况资料的统计图。它能显示出一组数据的最大值、最小值、中位数、下四分位数及上四分位数

## API
```go
// 实例化图表
func NewBoxPlot(routers ...HTTPRouter) *BoxPlot
// 新增 X 轴数据
func AddXAxis(xAxis interface{}) *BoxPlot
// 新增 Y 轴数据及配置项
func AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *BoxPlot
// 结合不同类型图表叠加画在同张图上
// 只适用于 RectChart 图表，RectChart 图表包括 Bar/BoxPlot/Line/Scatter/EffectScatter/Kline/HeatMap
// 将 RectChart 图表的 Series 追加到调用者的 Series 里面，Series 是完全独立的
// 而全局配置使用的是调用者的配置项
func Overlap(a ...rectCharter)
// 新增扩展 X 轴
func ExtendXAxis(xAxis ...XAxisOpts)
// 新增扩展 Y 轴
func ExtendYAxis(yAxis ...YAxisOpts)
// 新增 JS 函数
func AddJSFuncs(fn ...string)
// 设置全局配置项
func SetGlobalOptions(options ...globalOptser)
// 设置 Series 配置项
func SetSeriesOptions(options ...seriesOptser)
// 负责渲染图表，支持传入多个实现了 io.Writer 接口的对象
func Render(w ...io.Writer)
```

## 预定义
> Note: 示例用到的一些变量及方法，部分重复的以后代码中不会再次列出
```go
var (
    bpX = [...]string{"expr1", "expr2", "expr3", "expr4", "expr5"}
    bpY = [][]int{
        {850, 740, 900, 1070, 930, 850, 950, 980, 980, 880,
            1000, 980, 930, 650, 760, 810, 1000, 1000, 960, 960},
        {960, 940, 960, 940, 880, 800, 850, 880, 900, 840,
            830, 790, 810, 880, 880, 830, 800, 790, 760, 800},
        {880, 880, 880, 860, 720, 720, 620, 860, 970, 950,
            880, 910, 850, 870, 840, 840, 850, 840, 840, 840},
        {890, 810, 810, 820, 800, 770, 760, 740, 750, 760,
            910, 920, 890, 860, 880, 720, 840, 850, 850, 780},
        {890, 840, 780, 810, 760, 810, 790, 810, 820, 850,
            870, 870, 810, 740, 810, 940, 950, 800, 810, 870},
    }
)
```

## Demo

### BoxPlot-示例图
```go
bp := charts.NewBoxPlot()
bp.SetGlobalOptions(charts.TitleOpts{Title: "BoxPlot-示例图"})
bp.AddXAxis(bpX).AddYAxis("boxplot", bpY)
```
![](https://user-images.githubusercontent.com/19553554/52360729-ad640980-2a77-11e9-84e2-feff7e11aea5.gif)


### BoxPlot-多 Series
```go
bp := charts.NewBoxPlot()
bp.SetGlobalOptions(charts.TitleOpts{Title: "BoxPlot-多 Series"})
bp.AddXAxis(bpX[:2]).
	AddYAxis("boxplot1", bpY[:2]).
	AddYAxis("boxplot2", bpY[2:])
```
![](https://user-images.githubusercontent.com/19553554/52392733-1e3e0c80-2ade-11e9-898e-720128e69007.png)
