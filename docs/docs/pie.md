---
id: pie
title: Pie
sidebar_label: Pie（饼图）
---

> 饼图主要用于表现不同类目的数据在总和中的占比。每个的弧度表示数据数量的比例

## API
```go
// 实例化图表
func NewPie(routers ...RouterOpts) *Pie
// 新增数据及配置项
func Add(name string, data map[string]interface{}, options ...seriesOptser) *Pie
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
type PieOpts struct {
    // 是否展示成南丁格尔图，通过半径区分数据大小。可选择两种模式：
    // 1."radius": 扇区圆心角展现数据的百分比，半径展现数据的大小。
    // 2."area": 所有扇区圆心角相同，仅通过半径展现数据大小。
    RoseType string
    // 饼图的中心（圆心）坐标，数组的第一项是横坐标，第二项是纵坐标。
    // 支持设置成百分比，设置成百分比时第一项是相对于容器宽度，第二项是相对于容器高度
    // 使用示例
    // 设置成绝对的像素值: center: [400, 300]
    // 设置成相对的百分比: center: ['50%', '50%']
    Center   interface{}
    // 饼图的半径。可以为如下类型：
    // 1.number：直接指定外半径值。
    // 2.string：例如，'20%'，表示外半径为可视区尺寸（容器高宽中较小一项）的 20% 长度。
    // 3.Array.<number|string>：数组的第一项是内半径，第二项是外半径。
    // 每一项遵从上述 number string 的描述。
    Radius   interface{}
}
```

## Demo

### Pie-示例图
```go
pie := charts.NewPie()
pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-示例图"})
pie.Add("pie", genKvData())
```
![](https://user-images.githubusercontent.com/19553554/52348202-bb596080-2a5e-11e9-84a7-60732be0743a.gif)


### Pie-显示 Label
```go
pie := charts.NewPie()
pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-显示 Label"})
pie.Add("pie", genKvData(), charts.LabelTextOpts{Show: true})
```
![](https://user-images.githubusercontent.com/19553554/52348226-c9a77c80-2a5e-11e9-8760-30d4b397b571.png)


### Pie-Label 格式
```go
pie := charts.NewPie()
pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-Label 格式"})
pie.Add("pie", genKvData(), charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"})
```
![](https://user-images.githubusercontent.com/19553554/52348239-d4faa800-2a5e-11e9-8640-487e6e7494fc.png)


### Pie-Radius
```go
pie := charts.NewPie()
pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-Radius"})
pie.Add("pie", genKvData(),
    charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
    charts.PieOpts{Radius: []string{"40%", "75%"}},
)
```
![](https://user-images.githubusercontent.com/19553554/52348239-d4faa800-2a5e-11e9-8640-487e6e7494fc.png)


### Pie-玫瑰图(Area)
```go
pie := charts.NewPie()
pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-玫瑰图(Area)"})
pie.Add("pie", genKvData(),
    charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
    charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "area"},
)
```
![](https://user-images.githubusercontent.com/19553554/52348273-ee035900-2a5e-11e9-81e5-38852b06ff28.png)


### Pie-玫瑰图(Radius)
```go
pie := charts.NewPie()
pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-玫瑰图(Radius)"})
pie.Add("pie", genKvData(),
    charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
    charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "radius"},
)
```
![](https://user-images.githubusercontent.com/19553554/52348322-fc517500-2a5e-11e9-85fc-08291d339f09.png)


### Pie-玫瑰图(Area/Radius)
```go
pie := charts.NewPie()
pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-玫瑰图(Area/Radius)"})
pie.Add("area", genKvData(),
    charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "area", Center: []string{"25%", "50%"}},
)
pie.Add("radius", genKvData(),
    charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
    charts.PieOpts{Radius: []string{"30%", "75%"}, RoseType: "radius", Center: []string{"75%", "50%"}},
)
```
![](https://user-images.githubusercontent.com/19553554/52348346-0bd0be00-2a5f-11e9-807f-e779c706fcf4.png)


### Pie-饼中饼
```go
pie := charts.NewPie()
pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-饼中饼"})
pie.Add("area", genKvData(),
    charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}"},
    charts.PieOpts{Radius: []string{"50%", "55%"}, RoseType: "area"},
)
pie.Add("radius", genKvData(),
    charts.PieOpts{Radius: []string{"0%", "45%"}, RoseType: "radius"},
)
```
![](https://user-images.githubusercontent.com/19553554/52348367-168b5300-2a5f-11e9-8098-2d5d6b2cb984.png)
