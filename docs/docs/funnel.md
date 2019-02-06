---
id: funnel
title: Funnel（漏斗图）
sidebar_label: Funnel（漏斗图）
---

## API
```go
// 实例化图表
func NewFunnel(routers ...HTTPRouter) *Funnel {}
// 新增数据及配置项
func Add(name string, data map[string]interface{}, options ...seriesOptser) *Funnel {}
```

## Demo

### Funnel-示例图
```go
funnel := charts.NewFunnel()
funnel.SetGlobalOptions(charts.TitleOpts{Title: "Funnel-示例图"})
funnel.Add("funnel", genKvData())
```
![](https://user-images.githubusercontent.com/19553554/52332816-ac5eb800-2a36-11e9-8227-3538976f447d.gif)


### Funnel-显示 Label
```go
funnel := charts.NewFunnel()
funnel.SetGlobalOptions(charts.TitleOpts{Title: "Funnel-显示 Label"})
funnel.Add("funnel", genKvData(), charts.LabelTextOpts{Show: true})
```
![](https://user-images.githubusercontent.com/19553554/52332845-baacd400-2a36-11e9-86e7-8072a3efae23.png)
