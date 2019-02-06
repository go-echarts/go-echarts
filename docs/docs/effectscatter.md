---
id: effectScatter
title: Effectscatter
sidebar_label: Effectscatter（动态散点图）
---

## API
```go
// // 实例化图表
func NewEffectScatter(routers ...HTTPRouter) *EffectScatter {}
// 新增 X 轴数据
func AddXAxis(xAxis interface{}) *EffectScatter {}
// 新增 Y 轴数据及配置项
func AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *EffectScatter {}
// 结合不同类型图表叠加画在同张图上
// 只适用于 RectChart 图表，RectChart 图表包括 Bar/Line/Scatter/EffectScatter/Kline/HeatMap
// 将 RectChart 图表的 Series 追加到调用者的 Series 里面，Series 是完全独立的
// 而全局配置使用的是调用者的配置项
func Overlap(a ...serieser)
// 扩展新增 X 轴
func ExtendXAxis(xAxis ...XAxisOpts) {}
// 扩展新增 Y 轴
func ExtendYAxis(yAxis ...YAxisOpts) {}
```

## Demo

### EffectScatter-示例图
```go
es := charts.NewEffectScatter()
es.SetGlobalOptions(charts.TitleOpts{Title: "EffectScatter-示例图"})
es.AddXAxis(nameItems).AddYAxis("es1", randInt())
```
![](https://user-images.githubusercontent.com/19553554/52332299-28580080-2a35-11e9-95b8-c23b06c78ea0.gif)


### EffectScatter-涟漪效果
```go
es := charts.NewEffectScatter()
es.SetGlobalOptions(charts.TitleOpts{Title: "EffectScatter-涟漪效果"})
es.AddXAxis(nameItems).
    AddYAxis("es1", randInt(), 
        charts.RippleEffectOpts{Period: 4, Scale: 10, BrushType: "stroke"}).
    AddYAxis("es2", randInt(), 
        charts.RippleEffectOpts{Period: 3, Scale: 6, BrushType: "fill"})
```
![](https://user-images.githubusercontent.com/19553554/52332354-4f163700-2a35-11e9-837e-b45f8ed1b371.gif)
