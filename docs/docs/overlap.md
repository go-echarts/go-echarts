---
id: overlap
title: Overlap（重叠图）
sidebar_label: Overlap（重叠图）
---

## API

## 预定义
```go
func barBase() *charts.Bar {
    bar := charts.NewBar()
    bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
    bar.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
        AddYAxis("商家B", randInt())
    return bar
}

func lineBase() *charts.Line {
    line := charts.NewLine()
    line.AddXAxis(nameItems).AddYAxis("商家A", randInt())
    return line
}

func scatterBase() *charts.Scatter {
    scatter := charts.NewScatter()
    scatter.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
        AddYAxis("商家B", randInt())
    return scatter
}

func esBase() *charts.EffectScatter {
    es := charts.NewEffectScatter()
    es.AddXAxis(nameItems).AddYAxis("es1", randInt())
    return es
}
```

## Demo

### Overlap-示例图
```go
bar := barBase()
bar.Title = "Overlap-示例图"
bar.Overlap(lineBase())
```
![](https://user-images.githubusercontent.com/19553554/52349544-c2ce3900-2a61-11e9-82af-28aaaaae0d67.gif)


### Overlap-Bar+Scatte
```go
bar := barBase()
bar.Title = "Overlap-Bar+Scatter"
bar.Overlap(scatterBase())
```
![](https://user-images.githubusercontent.com/19553554/52349578-d2e61880-2a61-11e9-9c79-b8d2925f92cd.png)


### Overlap-Line+Scatter
```go
line := lineBase()
line.Title = "Overlap-Line+Scatter"
line.Overlap(scatterBase())
```
![](https://user-images.githubusercontent.com/19553554/52349601-e09b9e00-2a61-11e9-84e8-69c729280e7a.png)


### Overlap-All
```go
bar := barBase()
bar.Title = "Overlap-All"
bar.Overlap(lineBase(), scatterBase(), esBase())
```
![](https://user-images.githubusercontent.com/19553554/52349628-edb88d00-2a61-11e9-925b-62d8d9e4ea3f.png)
