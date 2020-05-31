---
id: funnel
title: Funnel
sidebar_label: Funnel（漏斗图）
---

> 漏斗图适用于业务流程比较规范、周期长、环节多的单流程单向分析，通过漏斗各环节业务数据的比较能够直观地发现和说明问题所在的环节，进而做出决策

## API
```go
// 实例化图表
func NewFunnel(routers ...RouterOpts) *Funnel
// 新增数据及配置项
func Add(name string, data map[string]interface{}, options ...seriesOptser) *Funnel
// 新增 JS 函数
func AddJSFuncs(fn ...string)
// 设置全局配置项
func SetGlobalOptions(options ...globalOptser)
// 设置全部 Series 配置项
func SetSeriesOptions(options ...seriesOptser)
// 负责渲染图表，支持传入多个实现了 io.Writer 接口的对象
func Render(w ...io.Writer) error
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
