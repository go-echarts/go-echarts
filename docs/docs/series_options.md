---
id: series_options
title: 系列配置项
sidebar_label: 系列配置项
---

以下所有 Opts/Item 结尾的配置项均可以作为 SetSeriesOptions 方法的参数，它们均实现了 seriesOptser 接口

## LabelTextOpts
> 图形上的文本标签配置项
```go
type LabelTextOpts struct {
    // 是否显示标签
    Show bool `json:"show"`
    // 文字的颜色
    Color string `json:"color,omitempty"`
    // 标签的位置
    // 通过相对的百分比或者绝对像素值表示标签相对于图形包围盒左上角的位置。示例：
    // 绝对的像素值    position: [10, 10],
    // 相对的百分比    position: ["50%", "50%"]
    // "top", "left", "right", "bottom"
    // "inside", "insideLeft", "insideRight", "insideTop", "insideBottom"
    // "insideTopLeft", "insideBottomLeft", "insideTopRight", "insideBottomRight"
    Position string `json:"position,omitempty"`
    // 标签内容格式器，支持字符串模板和回调函数两种形式，字符串模板与回调函数返回的字符串均支持用 \n 换行。
    // 1. 字符串模板 模板变量有：
    // {a}：系列名。
    // {b}：数据名。
    // {c}：数据值。
    // {@xxx}：数据中名为'xxx'的维度的值，如{@product}表示名为'product'`的维度的值。
    // {@[n]}：数据中维度 n 的值，如{@[3]}` 表示维度 3 的值，从 0 开始计数。
    // 2, 回调函数
    // 回调函数格式：
    // (params: Object|Array, ticket: string, callback: (ticket: string, html: string)) => string
    // 第一个参数 params 是 formatter 需要的数据集。格式如下：
    // {
    //    componentType: 'series',
    //    seriesType: string,    // 系列类型
    //    seriesIndex: number,    // 系列在传入的 option.series 中的 index
    //    seriesName: string,    // 系列名称
    //    name: string,            // 数据名，类目名
    //    dataIndex: number,    // 数据在传入的 data 数组中的 index
    //    data: Object,            // 传入的原始数据项
    //    value: number|Array,    // 传入的数据值
    //    color: string,        // 数据图形的颜色
    //    percent: number,        // 饼图的百分比
    // }
    Formatter string `json:"formatter,omitempty"`
}
```

## RippleEffectOpts
> 涟漪特效配置项
```go
type RippleEffectOpts struct {
    // 动画的周期，秒数
    // 默认 4(s)
    Period float32 `json:"period,omitempty"`
    // 动画中波纹的最大缩放比例
    // 默认 2.5
    Scale float32 `json:"scale,omitempty"`
    // 波纹的绘制方式，可选 "stroke" 和 "fill"
    // 默认 "fill"
    BrushType string `json:"brushType,omitempty"`
}
```

## LineStyleOpts
> 线风格配置项
```go
type LineStyleOpts struct {
    // 线的颜色
    Color string `json:"color,omitempty"`
    // 线的宽度
    // 默认 1
    Width float32 `json:"width,omitempty"`
    // 线的类型，可选 "solid", "dashed", "dotted"
    // 默认 "solid"
    Type string `json:"type,omitempty"`
    // 线的透明度。支持从 0 到 1 的数字，为 0 时不绘制线
    Opacity float32 `json:"opacity,omitempty"`
    // 线的曲度，支持从 0 到 1 的值，值越大曲度越大
    // 默认 0
    Curveness float32 `json:"curveness,omitempty"`
}
```

## AreaStyleOpts
> 区域风格配置项
```go
type AreaStyleOpts struct {
    // 填充区域的颜色
    Color string `json:"color,omitempty"`
    // 填充区域的透明度。支持从 0 到 1 的数字，为 0 时不填充区域
    Opacity float32 `json:"opacity,omitempty"`
}
```

## ItemStyleOpts
> 数据 Item 配置项
```go
type ItemStyleOpts struct {
    // 图形的颜色
    // Kline 图中为 阳线图形颜色
    Color string `json:"color,omitempty"`
    // Kline 图中为 阴线图形颜色
    Color0 string `json:"color0,omitempty"`
    // 图形的描边颜色
    // Kline 途中为 阳线图形的描边颜色
    BorderColor string `json:"borderColor,omitempty"`
    // Kline 途中为 阴线图形的描边颜色
    BorderColor0 string `json:"borderColor0,omitempty"`
    // 图形透明度。支持从 0 到 1 的数字，为 0 时不绘制该图形
    Opacity float32 `json:"opacity,omitempty"`
}
```

## TextStyleOpts
> 字体样式配置项
```go
type TextStyleOpts struct {
    // 文字字体颜色
    Color string `json:"color,omitempty"`
    // 文字字体的风格
    // 可选  'normal', 'italic', 'oblique'
    FontStyle string `json:"fontStyle,omitempty"`
    // 字体大小
    FontSize int `json:"fontSize,omitempty"`
    // 递归结构，为了兼容 wordCloud
    Normal *TextStyleOpts `json:"normal,omitempty"`
}
```

## MLNameTypeItem
> MarkLine 数据 Name-Type
```go
type MLNameTypeItem struct {
    // 标记线名称
    Name string `json:"name"`
    // 内置类型，可选 "average", "min", "max"
    Type string `json:"type"`
    // 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
    // 可以是维度的直接名称，例如折线图时可以是 x、angle 等
    // candlestick 图时可以是 open、close 等维度名称。
    ValueDim string `json:"valueDim,omitempty"`
}
```

## MLNameXAxisItem
> MarkLine 数据 Name-XAxis
```go
type MLNameXAxisItem struct {
    // 标记线名称
    Name string `json:"name"`
    // X 轴数据
    XAxis interface{} `json:"xAxis"`
    // 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
    // 可以是维度的直接名称，例如折线图时可以是 x、angle 等
    // candlestick 图时可以是 open、close 等维度名称。
    ValueDim string `json:"valueDim,omitempty"`
}
```

## MLNameYAxisItem
> MarkLine 数据 Name-YAxis
```go
type MLNameYAxisItem struct {
    // 标记线名称
    Name string `json:"name"`
    // Y 轴数据
    YAxis interface{} `json:"yAxis"`
    // 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
    // 可以是维度的直接名称，例如折线图时可以是 x、angle 等
    // candlestick 图时可以是 open、close 等维度名称。
    ValueDim string `json:"valueDim,omitempty"`
}
```

## MLNameCoordItem
> MarkLine 数据 Name-Coordinates
```go
type MLNameCoordItem struct {
    // 标记线名称
    Name string
    // 标记线起始坐标
    Coord0 []interface{}
    // 标记线结束坐标
    Coord1 []interface{}
    // 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
    // 可以是维度的直接名称，例如折线图时可以是 x、angle 等
    // candlestick 图时可以是 open、close 等维度名称。
    ValueDim string `json:"valueDim,omitempty"`
}
```

## MLStyleOpts
> MarkLine 风格配置项
```go
type MLStyleOpts struct {
    // 图元的图形类别
    // 可选 "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
    Symbol []string `json:"symbol,omitempty"`
    // 图元的大小
    SymbolSize float32 `json:"symbolSize,omitempty"`
    // 标线文本配置项
    Label LabelTextOpts `json:"label,omitempty"`
}
```

## MPNameTypeItem
> MarkPoint 数据 Name-Type
```go
type MPNameTypeItem struct {
    // 标记点名称
    Name string `json:"name"`
    // 内置类型，可选 "average", "min", "max"
    Type string `json:"type"`
    // 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
    // 可以是维度的直接名称，例如折线图时可以是 x、angle 等
    // candlestick 图时可以是 open、close 等维度名称。
    ValueDim string `json:"valueDim,omitempty"`
}
```

## MPNameCoordItem
> MarkPoint 数据 Name-Coordinates
```go
type MPNameCoordItem struct {
    // 标记点名称
    Name string `json:"name"`
    // 标记点坐标
    Coord []interface{} `json:"coord"`
    // 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
    // 可以是维度的直接名称，例如折线图时可以是 x、angle 等
    // candlestick 图时可以是 open、close 等维度名称。
    ValueDim string `json:"valueDim,omitempty"`
}
```

## MPStyleOpts
> MarkPoint 风格配置项
```go
type MPStyleOpts struct {
    // 图元的图形类别
    // 可选 "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
    Symbol string `json:"symbol,omitempty"`
    // 图元的大小
    SymbolSize float32 `json:"symbolSize,omitempty"`
    // 标注文本配置项
    Label LabelTextOpts `json:"label,omitempty"`
}
```

## BarOpts
> Bar series options
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

## HeatMapOpts
> heatMap series options
```go
type HeatMapOpts struct {
    //使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
    XAxisIndex int
    //使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
    YAxisIndex int
}
```

## LineOpts
> Line series options
```go
    // 数据堆叠，同个类目轴上系列配置相同的 stack 值可以堆叠放置
    Stack      string
    // 曲线是否平滑
    Smooth     bool
    // 是否使用阶梯图
    Step       bool
    // 使用的 x 轴的 index，在单个图表实例中存在多个 x 轴的时候有用
    XAxisIndex int
    // 使用的 y 轴的 index，在单个图表实例中存在多个 y 轴的时候有用
    YAxisIndex int
```

## LiquidOpts
> Liquid series options
```go
type LiquidOpts struct {
    // 水球图形状，可选
    // "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
    Shape           string
    // 是否显示水球轮廓
    IsShowOutline   bool
    // 是否停止动画
    IsWaveAnimation bool
}
```

## PieOpts
> Pie series options
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

## WordCLoudOpts
> WordCLoud series options
```go
type WordCLoudOpts struct {
    // 词云图形状，可选
    //"circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow"
    Shape         string
    // 字体大小范围
    SizeRange     []float32
    // 字体选择角度范围
    RotationRange []float32
}
```