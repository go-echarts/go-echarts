---
id: wordCloud
title: WordCloud
sidebar_label: WordCloud（词云图）
---

> 词云图

## API
```go
// 实例化图表
func NewWordCloud(routers ...RouterOpts) *WordCloud
// 新增数据及配置项
func Add(name string, data map[string]interface{}, options ...seriesOptser) *WordCloud
// 新增 JS 函数
func AddJSFuncs(fn ...string)
// 设置全局配置项
func SetGlobalOptions(options ...globalOptser)
// 设置全部 Series 配置项
func SetSeriesOptions(options ...seriesOptser)
// 负责渲染图表，支持传入多个实现了 io.Writer 接口的对象
func Render(w ...io.Writer) error
```

## 专属 Options
> 在 `SetSeriesOptions` 中设置
```go
type WordCloudOpts struct {
    // 词云图形状，可选
    //"circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow"
    Shape         string
    // 字体大小范围
    SizeRange     []float32
    // 字体选择角度范围
    RotationRange []float32
}
```

## 预定义
> Note: 示例用到的一些变量及方法，部分重复的以后代码中不会再次列出
```go
var wcData = map[string]interface{}{
    "Sam S Club":               10000,
    "Macys":                    6181,
    "Amy Schumer":              4386,
    "Jurassic World":           4055,
    "Charter Communications":   2467,
    "Chick Fil A":              2244,
    "Planet Fitness":           1898,
    "Pitch Perfect":            1484,
    "Express":                  1689,
    "Home":                     1112,
    "Johnny Depp":              985,
    "Lena Dunham":              847,
    "Lewis Hamilton":           582,
    "KXAN":                     555,
    "Mary Ellen Mark":          550,
    "Farrah Abraham":           462,
    "Rita Ora":                 366,
    "Serena Williams":          282,
    "NCAA baseball tournament": 273,
    "Point Break":              265,
}
```

## Demo

### WordCloud-示例图
```go
wc := charts.NewWordCloud()
wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-示例图"})
wc.Add("wordcloud", wcData, charts.WordCloudOpts{SizeRange: []float32{14, 80}})
```
![](https://user-images.githubusercontent.com/19553554/52348737-01fb8a80-2a60-11e9-94ac-dacbd7b58811.png)


### WordCloud-形状(cardioid)
```go
wc := charts.NewWordCloud()
wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-形状(cardioid)"})
wc.Add("wordcloud", wcData,
    charts.WordCloudOpts{SizeRange: []float32{14, 80}}, charts.WordCloudOpts{Shape: "cardioid"})
```
![](https://user-images.githubusercontent.com/19553554/52348901-5bfc5000-2a60-11e9-94f5-fbdce2f2ec46.png)


### WordCloud-形状(star)
```go
wc := charts.NewWordCloud()
wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-形状(star)"})
wc.Add("wordcloud", wcData,
    charts.WordCloudOpts{SizeRange: []float32{14, 80}}, charts.WordCloudOpts{Shape: "cardioid"})
```
![](https://user-images.githubusercontent.com/19553554/52349093-bf867d80-2a60-11e9-81d7-2c45ddcce0cc.png)
