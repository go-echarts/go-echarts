---
id: scatter
title: Scatter
sidebar_label: Scatter（散点图）
---

> 直角坐标系上的散点图可以用来展现数据的 x，y 之间的关系

## API
```go
// // 实例化图表
func NewScatter(routers ...RouterOpts) *Scatter
// 新增 X 轴数据
func AddXAxis(xAxis interface{}) *Scatter
// 新增 Y 轴数据及配置项
func AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Scatter
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
func Render(w ...io.Writer) error
```

## 专属 Options
> 在 `SetSeriesOptions` 中设置
```go
type ScatterOpts struct {
    // 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
    XAxisIndex int
    // 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
    YAxisIndex int
}
```

## Demo

### Scatter-示例图
```go
scatter := charts.NewScatter()
scatter.SetGlobalOptions(charts.TitleOpts{Title: "Scatter-示例图"})
scatter.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
```
![](https://user-images.githubusercontent.com/19553554/52348431-420e3d80-2a5f-11e9-8cab-7b415592dc77.gif)


### Scatter-显示 Label
```go
scatter := charts.NewScatter()
scatter.SetGlobalOptions(charts.TitleOpts{Title: "Scatter-显示 Label"})
scatter.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
scatter.SetSeriesOptions(charts.LabelTextOpts{Show: true, Position: "right"})
```
![](https://user-images.githubusercontent.com/19553554/52348566-94e7f500-2a5f-11e9-9ca3-ecd73bd85fd3.png)


### Scatter-显示分割线
```go
scatter := charts.NewScatter()
scatter.SetGlobalOptions(
    charts.TitleOpts{Title: "Scatter-显示分割线"},
    charts.XAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
    charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}},
)
scatter.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
```
![](https://user-images.githubusercontent.com/19553554/52348578-9f09f380-2a5f-11e9-8ea9-8133e1e4c644.png)
