---
id: pie
title: Pie（饼图）
sidebar_label: Pie（饼图）
---

## API
```go
// 实例化图表
func NewPie(routers ...HTTPRouter) *Pie {}
// 新增数据及配置项
func Add(name string, data map[string]interface{}, options ...seriesOptser) *Pie {}
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
