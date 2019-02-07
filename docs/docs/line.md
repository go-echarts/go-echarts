---
id: line
title: Line
sidebar_label: Line（折线图）
---

> 折线图是用折线将各个数据点标志连接起来的图表，用于展现数据的变化趋势

## API
```go
// // 实例化图表
func NewLine(routers ...HTTPRouter) *Line
// 新增 X 轴数据
func AddXAxis(xAxis interface{}) *Line
// 新增 Y 轴数据及配置项
func AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Line
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
type LineOpts struct {
    // 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
    Stack      string
    // 曲线是否平滑
    Smooth     bool
    // 是否使用阶梯图
    Step       bool
    // 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
    XAxisIndex int
    // 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
    YAxisIndex int
}
```


## Demo

### Line-示例图
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-示例图"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt())
```
![](https://user-images.githubusercontent.com/19553554/52346064-b7770f80-2a59-11e9-9e03-6dae3a8c637d.gif)


### Line-显示 Label
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-显示 Label"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LabelTextOpts{Show: true})
```
![](https://user-images.githubusercontent.com/19553554/52346108-d2e21a80-2a59-11e9-8786-2164cbb23e89.png)


### Line-标记点
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-标记点"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt(),
    charts.MPNameTypeItem{Name:"最大值", Type:"max"},
    charts.MPNameTypeItem{Name:"平均值", Type:"average"},
    charts.MPNameTypeItem{Name:"最小值", Type:"min"},
    charts.MPStyleOpts{Label:charts.LabelTextOpts{Show:true}})
```
![](https://user-images.githubusercontent.com/19553554/52346684-1db06200-2a5b-11e9-9488-4a51db0806d4.gif)


### Line-显示分割线
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-显示分割线"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LabelTextOpts{Show: true})
line.SetGlobalOptions(charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}})
```
![](https://user-images.githubusercontent.com/19553554/52346154-f6a56080-2a59-11e9-9c24-9679c6ebc8d5.png)


### Line-阶梯图
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-阶梯图"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LineOpts{Step: true})
```
![](https://user-images.githubusercontent.com/19553554/52346208-0d4bb780-2a5a-11e9-821d-bf7f1a9b26c6.png)


### Line-平滑曲线
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-平滑曲线"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LineOpts{Smooth: true})
```
![](https://user-images.githubusercontent.com/19553554/52346235-1d639700-2a5a-11e9-8bd0-5f85ea24194d.png)


### Line-填充区域
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-填充区域"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt(),
    charts.LabelTextOpts{Show: true},
    charts.AreaStyleOpts{Opacity: 0.2},
)
```
![](https://user-images.githubusercontent.com/19553554/52346258-2f453a00-2a5a-11e9-845e-89a64a7fe317.png)


### Line-平滑区域
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-平滑区域"})
line.AddXAxis(nameItems).AddYAxis("商家A", randInt(),
    charts.LabelTextOpts{Show: true},
    charts.AreaStyleOpts{Opacity: 0.2},
    charts.LineOpts{Smooth: true},
)
```
![](https://user-images.githubusercontent.com/19553554/52346289-3f5d1980-2a5a-11e9-9b7f-939f90d0c685.png)


### Line-多线
```go
line := charts.NewLine()
line.SetGlobalOptions(charts.TitleOpts{Title: "Line-多线"}, charts.InitOpts{Theme: "shine"})
line.AddXAxis(nameItems).
    AddYAxis("商家 A", randInt()).
    AddYAxis("商家 B", randInt()).
    AddYAxis("商家 C", randInt()).
    AddYAxis("商家 D", randInt())
```
![](https://user-images.githubusercontent.com/19553554/52346315-51d75300-2a5a-11e9-83ce-49f9f73308aa.png)


### 查询时间对比 哈希表 vs 二分查找
```go
line := charts.NewLine()
line.AddXAxis([]string{"10e1", "10e2", "10e3", "10e4", "10e5", "10e6", "10e7"}).
    AddYAxis("map", []float32{19.9, 16.8, 19.9, 29.4, 61.3, 77.3, 93.0},
        charts.LabelTextOpts{Show: true, Position: "bottom"}).
    AddYAxis("slice", []float32{24.9, 34.9, 48.1, 58.3, 69.7, 123, 131},
        charts.LabelTextOpts{Show: true, Position: "top"})
line.SetSeriesOptions(
    charts.MLNameTypeItem{Name: "平均值", Type: "average"},
    charts.LineOpts{Smooth: true},
    charts.MLStyleOpts{Label: charts.LabelTextOpts{Show: true, Formatter: "{a}: {b}"}},
)
line.SetGlobalOptions(
    charts.TitleOpts{Title: "查询时间对比 哈希表 vs 二分查找"},
    charts.YAxisOpts{Name: "搜索时间(ns)", SplitLine: charts.SplitLineOpts{Show: false}},
    charts.XAxisOpts{Name: "元素数量"})
```
![](https://user-images.githubusercontent.com/19553554/52346349-60256f00-2a5a-11e9-85f5-df627c655844.png)
