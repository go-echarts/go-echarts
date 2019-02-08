---
id: api
title: API 设计
sidebar_label: API 设计
---

go-echarts 秉承着 API 设计要简洁的原则，对所有图表只提供了下面的接口

## Common API

*func Add(arg0 Type0, arg1 Type1, ...)*

新增 series，参数由具体图表决定

*func SetGlobalOptions(options ...globalOptser)*

设置全局配置项

*func SetSeriesOptions(options ...seriesOptser)*

设置 Series 配置项

*func AddJSFuncs(fn ...string)*

新增 JS 函数

*func Render(w ...io.Writer) error*

负责渲染图表，支持传入多个实现了 io.Writer 接口的对象

## RectChart API

> RectChart 包括 Bar/BoxPlot/Line/Scatter/EffectScatter/Kline/HeatMap

*func AddXAxis(arg0 Type0, arg1 Type1, ...)*

新增 X 轴数据，参数由具体图表决定

*func AddYAxis(arg0 Type0, arg1 Type1, ...)*

新增 X 轴数据，参数由具体图表决定

*func Overlap(a ...rectCharter)*

新增重叠 Series，对象必须实现 rectCharter 接口

*func ExtendXAxis(xAxis ...XAxisOpts)*

新增扩展 X 轴

*func ExtendYAxis(yAxis ...YAxisOpts)*

新增扩展 Y 轴

## 3D Chart API

> 3D charts 包括 Bar3D/Line3D/Scatter3D/Surface3D

*func AddXYAxis(xAxis, yAxis interface{})*

新增 XY 轴数据

*func AddZAxis(name string, zAxis interface{}, options ...seriesOptser)*

新增 Z 轴数据及配置项

## Interface

*interface globalOptser*
```go
type globalOptser interface {
    markGlobal()
}
```

*interface seriesOptser*
```go
type seriesOptser interface {
    markSeries()
}
```

*interface rectCharter*
```go
type rectCharter interface {
    markRectChart()
    exportSeries() Series
}
```