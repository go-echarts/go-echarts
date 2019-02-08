---
id: line3d
title: Line3D
sidebar_label: Line3D（3D 折线图）
---

> 3D 折线图

## API
```go
// 实例化图表
func NewLine3D(routers ...HTTPRouter) *Line3D
// 新增 XY 轴数据
func AddXYAxis(xAxis, yAxis interface{}) *Line3D
// 新增 Z 轴数据及配置项
func AddZAxis(name string, zAxis interface{}, options ...seriesOptser) *Line3D
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

func genLine3dData() [][3]float64 {
    data := make([][3]float64, 0)
    for i := 0; i < 25000; i++ {
        t := float64(i) / 1000
        data = append(data,
            [3]float64{
                (1 + 0.25*math.Cos(75*float64(t))) * math.Cos(float64(t)),
                (1 + 0.25*math.Cos(75*float64(t))) * math.Sin(float64(t)),
                float64(t) + 2.0*math.Sin(75.0*float64(t)),
            },
        )
    }
    return data
}

```

## Demo

### Line3D-示例图
```go
line3d := charts.NewLine3D()
line3d.SetGlobalOptions(
    charts.TitleOpts{Title: "Line3D-示例图"},
    charts.VisualMapOpts{
        Calculable: true,
        InRange:    charts.VMInRange{Color: rangeColor},
        Max:        30,
    },
)
line3d.AddZAxis("line3D", genLine3dData())
```
![](https://user-images.githubusercontent.com/19553554/52464826-4baab900-2bb7-11e9-8299-776f5ee43670.gif)


### Line3D-旋转的弹簧
```go
line3d := charts.NewLine3D()
line3d.SetGlobalOptions(
    charts.TitleOpts{Title: "Line3D-旋转的弹簧"},
    charts.VisualMapOpts{
        Calculable: true,
        InRange:    charts.VMInRange{Color: rangeColor},
        Max:        30,
    },
    charts.Grid3DOpts{
        ViewControl: charts.ViewControlOpts{AutoRotate: true, AutoRotateSpeed: 30},
    },
)
line3d.AddZAxis("line3D", genLine3dData())
```
![](https://user-images.githubusercontent.com/19553554/52465010-0935ac00-2bb8-11e9-97eb-de531df19b16.gif)
