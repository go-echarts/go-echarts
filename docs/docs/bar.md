---
id: bar
title: Bar
sidebar_label: Bar（柱状图）
---

> 柱状/条形图，通过柱形的高度/条形的宽度来表现数据的大小

## API
```go
// 实例化图表
func NewBar(routers ...HTTPRouter) *Bar
// 新增 X 轴数据
func AddXAxis(xAxis interface{}) *Bar
// 新增 Y 轴数据及配置项
func AddYAxis(name string, yAxis interface{}, options ...seriesOptser) *Bar
// 是否翻转 XY 坐标轴
func XYReversal()
// 结合不同类型图表叠加画在同张图上
// 只适用于 RectChart 图表，RectChart 图表包括 Bar/BoxPlot/Line/Scatter/EffectScatter/Kline/HeatMap
// 将 RectChart 图表的 Series 追加到调用者的 Series 里面，Series 是完全独立的
// 而全局配置使用的是调用者的配置项
func Overlap(a ...rectCharter)
// 新增扩展 X 轴
func ExtendXAxis(xAxis ...XAxisOpts)
// 新增扩展 Y 轴
func ExtendYAxis(yAxis ...YAxisOpts)
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
type BarOpts struct {
    // 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
    Stack      string
    // 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
    XAxisIndex int
    // 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
    YAxisIndex int
}
```

## 预定义
> Note: 示例用到的一些变量及方法，部分重复的以后代码中不会再次列出
```go
var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var foodItems = []string{"面包", "牛奶", "奶茶", "棒棒糖", "加多宝", "可口可乐"}

var seed = rand.NewSource(time.Now().UnixNano())
func randInt() []int {
    cnt := len(nameItems)
    r := make([]int, 0)
    for i := 0; i < cnt; i++ {
        r = append(r, int(seed.Int63())%50)
    }
    return r
}
```

## Demo

### Bar-示例图
```go
import (
    "..."   // 其他内置库，不一一列出，后面示例代码也是一样

    "github.com/chenjiandongx/go-echarts/charts"
)

bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
f, err := os.Create("bar.html")
    if err != nil {
        log.Println(err)
    }
bar.Render(f)

// 或者使用 net/http，同上，后面也不就列出
func handler(w http.ResponseWriter, _ *http.Request) {
    bar := charts.NewBar()
    bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
    bar.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
        AddYAxis("商家B", randInt())
    f, err := os.Create("bar.html")
    if err != nil {
        log.Println(err)
    }
    bar.Render(w, f) // Render 可接收多个 io.writer 接口
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8081", nil)
}
```
![](https://user-images.githubusercontent.com/19553554/52197440-843a5200-289a-11e9-8601-3ce8d945b04a.gif)


### Bar-标题
```go
bar := charts.NewBar()
bar.SetGlobalOptions(
    charts.TitleOpts{Title: "Bar-标题", Subtitle: "我是副标题，相对来讲我会长一点", Right: "40%"},
    charts.LegendOpts{Right: "80%"},
    charts.ToolboxOpts{Show: true},
)
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
```
![](https://user-images.githubusercontent.com/19553554/52395508-6ca4d880-2ae9-11e9-853d-4ced2e5db822.png)


### Bar-显示 Label
```go
bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-显示 Label"})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
```
![](https://user-images.githubusercontent.com/19553554/52331824-1c1f7380-2a34-11e9-928b-70b84c179fb0.png)


### Bar-XY 轴名称
```go
bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-XY 轴名称"})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
bar.SetGlobalOptions(charts.XAxisOpts{Name: "商品名称"}, charts.YAxisOpts{Name: "销售量"})
```
![](https://user-images.githubusercontent.com/19553554/52331865-2d688000-2a34-11e9-850c-a15adab4d775.png)


### Bar-设置系列颜色
```go
bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-设置系列颜色"})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt(), charts.ColorOpts{"lightblue"}).
    AddYAxis("商家B", randInt(), charts.ColorOpts{"pink"})
bar.SetSeriesOptions(charts.LabelTextOpts{Show: true})
// 或者可以这样设置
//bar.SetGlobalOptions(charts.ColorOpts{"lightblue", "pink"})
```
![](https://user-images.githubusercontent.com/19553554/52331995-74567580-2a34-11e9-9050-3e4f202b8810.png)


### Bar-显示分割线
```go
bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-显示分割线"})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
bar.SetGlobalOptions(charts.YAxisOpts{SplitLine: charts.SplitLineOpts{Show: true}})
```
![](https://user-images.githubusercontent.com/19553554/52332006-82a49180-2a34-11e9-8cc4-ddff94b06426.png)


### Bar-Y 轴格式
```go
bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-Y 轴格式"})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
bar.SetGlobalOptions(
    charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}},
)
```
![](https://user-images.githubusercontent.com/19553554/52332025-8fc18080-2a34-11e9-8961-7cd606628be8.png)


### Bar-多 Y 轴
```go
bar := charts.NewBar()
bar.SetGlobalOptions(
    charts.TitleOpts{Title: "Bar-多 Y 轴"},
    charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}},
)
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt(), charts.BarOpts{YAxisIndex: 0}).
    AddYAxis("商家B", randInt(), charts.BarOpts{YAxisIndex: 1})
bar.ExtendYAxis(charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/月"}})
```
![](https://user-images.githubusercontent.com/19553554/52392892-cfdd3d80-2ade-11e9-985d-d80ae110c1e3.png)


### Bar-多 X 轴
```go
bar := charts.NewBar()
bar.SetGlobalOptions(
    charts.TitleOpts{Title: "Bar-多 X 轴"},
    charts.YAxisOpts{AxisLabel: charts.LabelTextOpts{Formatter: "{value} 件/天"}},
)
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt(), charts.BarOpts{XAxisIndex: 0}).
    AddYAxis("商家B", randInt(), charts.BarOpts{XAxisIndex: 1})
bar.ExtendXAxis(charts.XAxisOpts{Data: foodItems})
```
![](https://user-images.githubusercontent.com/19553554/52393164-ea63e680-2adf-11e9-99fa-dc1d6c5dd61e.png)


### Bar-DataZoom
```go
bar := charts.NewBar()
bar.SetGlobalOptions(
    charts.TitleOpts{Title: "Bar-DataZoom"},
    charts.ToolboxOpts{Show: true},
    charts.DataZoomOpts{XAxisIndex: []int{0}, Start: 50, End: 100},
)
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
```
![](https://user-images.githubusercontent.com/19553554/52393958-9824c480-2ae3-11e9-8418-52e71fd33ccb.gif)


### Bar-翻转 XY 轴
```go
bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-翻转 XY 轴"})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
bar.XYReversal()
```
![](https://user-images.githubusercontent.com/19553554/52332106-bb446b00-2a34-11e9-8a79-a03db79c74a0.png)


### Bar-堆叠效果
```go
bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-堆叠效果"})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt(), charts.BarOpts{Stack: "stack"}).
    AddYAxis("商家B", randInt(), charts.BarOpts{Stack: "stack"})
```
![](https://user-images.githubusercontent.com/19553554/52332121-c7302d00-2a34-11e9-86f1-a6c363e8b808.png)


### Bar-标记线&标记点
```go
bar := charts.NewBar()
bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-标记线&标记点"})
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt())
bar.SetSeriesOptions(
    charts.MLNameTypeItem{Name: "最大值", Type: "max"},
    charts.MLStyleOpts{SymbolSize: 20, Label: charts.LabelTextOpts{Show: true}},
    charts.MPNameTypeItem{Name: "最大值点", Type: "max"},
    charts.MPStyleOpts{Label: charts.LabelTextOpts{Show: true}},
)
```
![](https://user-images.githubusercontent.com/19553554/52332133-d2835880-2a34-11e9-8e4f-792fb2ea540a.png)


### Bar-自定义标记+主题
```go
import "github.com/chenjiandongx/go-echarts/common"

bar := charts.NewBar()
bar.SetGlobalOptions(
    charts.TitleOpts{Title: "Bar-自定义标记+主题"},
    charts.InitOpts{PageTitle: "Awesome", Theme: common.ThemeType.Macarons},
)
bar.AddXAxis(nameItems).
    AddYAxis("商家A", randInt()).
    AddYAxis("商家B", randInt(),
        charts.MLNameYAxisItem{Name: "Y 值为 5", YAxis: 20},
        charts.MLNameCoordItem{Name: "自定义标记线",
            Coord0: []interface{}{"衬衫", 10}, Coord1: []interface{}{"羊毛衫", 40}},
        charts.MLStyleOpts{SymbolSize: 20, Label: charts.LabelTextOpts{Show: true}},
    )
```
![](https://user-images.githubusercontent.com/19553554/52332160-e29b3800-2a34-11e9-8218-66a0b7a24622.png)
