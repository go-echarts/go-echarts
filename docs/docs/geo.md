---
id: geo
title: Geo（地理坐标系）
sidebar_label: Geo（地理坐标系）
---

## API
```go
// 实例化图表
func NewGeo(mapType string, routers ...HTTPRouter) *Geo {}
// 新增数据及配置项
func Add(name, geoType string, data map[string]interface{}, options ...seriesOptser) *Geo {}
```

## 预定义
```go
mapData = map[string]float32{
    "北京":   float32(rand.Intn(150)),
    "上海":   float32(rand.Intn(150)),
    "深圳":   float32(rand.Intn(150)),
    "辽宁":   float32(rand.Intn(150)),
    "青岛":   float32(rand.Intn(150)),
    "山西":   float32(rand.Intn(150)),
    "陕西":   float32(rand.Intn(150)),
    "乌鲁木齐": float32(rand.Intn(150)),
    "齐齐哈尔": float32(rand.Intn(150)),
}

guangdongData = map[string]float32{
    "深圳市": float32(rand.Intn(150)),
    "广州市": float32(rand.Intn(150)),
    "湛江市": float32(rand.Intn(150)),
    "汕头市": float32(rand.Intn(150)),
    "东莞市": float32(rand.Intn(150)),
    "佛山市": float32(rand.Intn(150)),
    "云浮市": float32(rand.Intn(150)),
    "肇庆市": float32(rand.Intn(150)),
    "梅州市": float32(rand.Intn(150)),
}

shantouData = map[string]float32{
    "澄海区": float32(rand.Intn(150)),
    "潮阳区": float32(rand.Intn(150)),
    "潮南区": float32(rand.Intn(150)),
    "南澳县": float32(rand.Intn(150)),
}
```

## Demo

### Geo-示例图(effectScatter)
```go
geo := charts.NewGeo("china")
geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-示例图(effectScatter)"})
geo.Add("geo", "effectScatter", mapData,
    charts.RippleEffectOpts{Period: 4, Scale: 6, BrushType: "stroke"})
```
![](https://user-images.githubusercontent.com/19553554/52333496-60147780-2a38-11e9-86b4-967bee2ae954.gif)


### Geo-显示 Label
```go
geo := charts.NewGeo("china")
geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-显示 Label"})

fn := `function (params) {
    return params.name + ' : ' + params.value[2];
}`
geo.Add("geo", "effectScatter", mapData,
    charts.LabelTextOpts{
        Show: true, 
        Formatter: charts.FuncOpts(fn), Color: "black", Position: "right"},
    charts.RippleEffectOpts{Period: 4, Scale: 6, BrushType: "stroke"},
)
```
![](https://user-images.githubusercontent.com/19553554/52344575-133f9980-2a56-11e9-93e0-568e484936ce.gif)


### Geo-Scatter
```go
geo := charts.NewGeo("china")
geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-Scatter"})
geo.Add("geo", "scatter", mapData)
```
![](https://user-images.githubusercontent.com/19553554/52344634-3ec28400-2a56-11e9-9020-e2eb18c95826.png)


### Geo-设置 VisualMap
```go
geo := charts.NewGeo("china")
geo.SetGlobalOptions(charts.TitleOpts{Title: "Geo-设置 VisualMap"})
geo.Add("geo", "scatter", mapData)
geo.SetGlobalOptions(charts.VisualMapOpts{Max: 60, Calculable: true})
```
![](https://user-images.githubusercontent.com/19553554/52344700-63b6f700-2a56-11e9-9075-efc97a088889.gif)


### Geo-HeatMap
```go
geo := charts.NewGeo("china")
geo.SetGlobalOptions(
    charts.TitleOpts{Title: "Geo-HeatMap"},
    charts.VisualMapOpts{Max: 60, Calculable: true},
)
geo.Add("geo", "heatmap", mapData)
```
![](https://user-images.githubusercontent.com/19553554/52344749-834e1f80-2a56-11e9-92b3-bbc8618f6ea2.gif)


### Geo-示例图(effectScatter+heatMap)
```go
geo := charts.NewGeo("china")
geo.SetGlobalOptions(
    charts.TitleOpts{Title: "Geo-示例图(effectScatter+heatMap)"},
    charts.VisualMapOpts{Max: 60, Calculable: true},
)
geo.Add("es", "effectScatter", mapData,
    charts.RippleEffectOpts{Period: 4, Scale: 10, BrushType: "stroke"})
geo.Add("heatmap", "heatmap", mapData)
```
![](https://user-images.githubusercontent.com/19553554/52344796-a547a200-2a56-11e9-9dbc-9b7f4ed6431c.gif)


### Geo-广东地图
```go
geo := charts.NewGeo("广东")
geo.SetGlobalOptions(
    charts.TitleOpts{Title: "Geo-广东地图"},
    charts.VisualMapOpts{Calculable: true,
        InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
)
geo.Add("geo", "scatter", guangdongData)
```
![](https://user-images.githubusercontent.com/19553554/52344868-d922c780-2a56-11e9-9114-f7dc06f9c7e5.gif)


### Geo-汕头地图
```go
geo := charts.NewGeo("汕头")
geo.SetGlobalOptions(
    charts.TitleOpts{Title: "Geo-汕头地图"},
    charts.VisualMapOpts{Calculable: true,
        InRange: charts.VMInRange{Color: []string{"#50a3ba", "#eac736", "#d94e5d"}}},
)
geo.Add("geo", "heatmap", shantouData)
```
![](https://user-images.githubusercontent.com/19553554/52344892-e93aa700-2a56-11e9-92e9-faaba7c79fc6.png)
