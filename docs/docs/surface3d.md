---
id: surface3d
title: surface3D
sidebar_label: surface3D（3D 曲面图）
---

> 3D 曲面图

## API
```go
// 实例化图表
func NewSurface3D(routers ...RouterOpts) *Surface3D
// 新增 XY 轴数据
func AddXYAxis(xAxis, yAxis interface{}) *Surface3D
// 新增 Z 轴数据及配置项
func AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Surface3D
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

func genSurface3dData0() [][3]interface{} {
    data := make([][3]interface{}, 0)

    for i := -60; i < 60; i++ {
        y := float64(i) / 60
        for j := -60; j < 60; j++ {
            x := float64(j) / 60
            z := math.Sin(x*math.Pi) * math.Sin(y*math.Pi)
            data = append(data, [3]interface{}{x, y, z})
        }
    }
    return data
}

func genSurface3dData1() [][3]interface{} {
    data := make([][3]interface{}, 0)
    for i := -30; i < 30; i++ {
        y := float64(i) / 10
        for j := -30; j < 30; j++ {
            x := float64(j) / 10
            z := math.Sin(x*x+y*y) * x / math.Pi
            data = append(data, [3]interface{}{x, y, z})
        }
    }
    return data
}
```

## Demo

###
```go
surface3d := charts.NewSurface3D()
surface3d.SetGlobalOptions(
	charts.TitleOpts{Title: "surface3D-示例图"},
	charts.VisualMapOpts{
		Calculable: true,
		InRange:    charts.VMInRange{Color: rangeColor},
		Max:        3,
		Min:        -3,
	},
)
surface3d.AddZAxis("surface3d", genSurface3dData0())
```
![](https://user-images.githubusercontent.com/19553554/52465183-a55fb300-2bb8-11e9-8c10-4519c4e3f758.gif)


###
```go
surface3d := charts.NewSurface3D()
surface3d.SetGlobalOptions(
	charts.TitleOpts{Title: "surface3D-一朵玫瑰"},
	charts.VisualMapOpts{
		Calculable: true,
		InRange:    charts.VMInRange{Color: rangeColor},
		Max:        3,
		Min:        -3,
	},
)
surface3d.AddZAxis("surface3d", genSurface3dData1())
```
![](https://user-images.githubusercontent.com/19553554/52465232-d3dd8e00-2bb8-11e9-841b-d0e82683551e.png)
