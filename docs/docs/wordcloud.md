---
id: wordCloud
title: WordCloud（词云图）
sidebar_label: WordCloud（词云图）
---

## API
```go
// 实例化图表
func NewWordCloud(routers ...HTTPRouter) *WordCloud {}
// 新增数据及配置项
func Add(name string, data map[string]interface{}, options ...seriesOptser) *WordCloud {}
```

## 预定义
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
wc.Add("wordcloud", wcData, charts.WordCLoudOpts{SizeRange: []float32{14, 80}})
```
![](https://user-images.githubusercontent.com/19553554/52348737-01fb8a80-2a60-11e9-94ac-dacbd7b58811.png)


### WordCloud-形状(cardioid)
```go
wc := charts.NewWordCloud()
wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-形状(cardioid)"})
wc.Add("wordcloud", wcData,
    charts.WordCLoudOpts{SizeRange: []float32{14, 80}}, charts.WordCLoudOpts{Shape: "cardioid"})
```
![](https://user-images.githubusercontent.com/19553554/52348901-5bfc5000-2a60-11e9-94f5-fbdce2f2ec46.png)


### WordCloud-形状(star)
```go
wc := charts.NewWordCloud()
wc.SetGlobalOptions(charts.TitleOpts{Title: "WordCloud-形状(star)"})
wc.Add("wordcloud", wcData,
    charts.WordCLoudOpts{SizeRange: []float32{14, 80}}, charts.WordCLoudOpts{Shape: "cardioid"})
```
![](https://user-images.githubusercontent.com/19553554/52349093-bf867d80-2a60-11e9-81d7-2c45ddcce0cc.png)
