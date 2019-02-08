---
id: scatter3d
title: Scatter3D
sidebar_label: Scatter3D（3D 散点图）
---

> 3D 散点图

## API
```go
// 实例化图表
func NewScatter3D(routers ...RouterOpts) *Scatter3D
// 新增 XY 轴数据
func AddXYAxis(xAxis, yAxis interface{}) *Scatter3D
// 新增 Z 轴数据及配置项
func AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Scatter3D
// 新增 JS 函数
func AddJSFuncs(fn ...string)
// 设置全局配置项
func SetGlobalOptions(options ...globalOptser)
// 设置 Series 配置项
func SetSeriesOptions(options ...seriesOptser)
// 负责渲染图表，支持传入多个实现了 io.Writer 接口的对象
func Render(w ...io.Writer) error
```

## 预定义
> Note: 示例用到的一些变量及方法，部分重复的以后代码中不会再次列出
```go
var rangeColor = []string{
    "#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
    "#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func genScatter3dData() [][3]int {
    data := make([][3]int, 0)
    for i := 0; i < 80; i++ {
        data = append(data, [3]int{
            int(seed.Int63()) % 100,
            int(seed.Int63()) % 100,
            int(seed.Int63()) % 100,
        })
    }
    return data
}
```

## Demo

### Scatter3D-示例图
```go
scatter3d := charts.NewScatter3D()
scatter3d.SetGlobalOptions(
	charts.TitleOpts{Title: "Scatter3D-示例图"},
	charts.VisualMapOpts{
		Calculable: true,
		InRange:    charts.VMInRange{Color: rangeColor},
		Max:        100,
	},
	charts.Grid3DOpts{BoxDepth: 80, BoxWidth: 200},
)
scatter3d.AddZAxis("scatter3d", genScatter3dData())
```
![](https://user-images.githubusercontent.com/19553554/52464647-aee81b80-2bb6-11e9-864e-c544392e523a.gif)
