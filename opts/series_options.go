package opts

// Label contains options for a label text.
type Label struct {
	// 是否显示标签
	Show bool `json:"show,omitempty"`
	// 文字的颜色
	Color string `json:"color,omitempty"`
	// 标签的位置
	// 通过相对的百分比或者绝对像素值表示标签相对于图形包围盒左上角的位置。示例：
	// 绝对的像素值	position: [10, 10],
	// 相对的百分比	position: ["50%", "50%"]
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
	//    seriesType: string,	// 系列类型
	//    seriesIndex: number,	// 系列在传入的 option.series 中的 index
	//    seriesName: string,	// 系列名称
	//    name: string,			// 数据名，类目名
	//    dataIndex: number,	// 数据在传入的 data 数组中的 index
	//    data: Object,			// 传入的原始数据项
	//    value: number|Array,	// 传入的数据值
	//    color: string,		// 数据图形的颜色
	//    percent: number,		// 饼图的百分比
	// }
	Formatter string `json:"formatter,omitempty"`
}

type Emphasis struct {
	// 高亮的标签样式
	Label `json:"label,omitempty"`
	// 高亮的图形样式
	ItemStyle `json:"itemStyle,omitempty"`
}

// ItemStyleOpts contains styling options for a MarkLine.
type ItemStyle struct {
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

// MLStyleOpts contains styling options for a MarkLine.
type MarkLineStyle struct {
	// 图元的图形类别
	// 可选 "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol []string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
	// 标线文本配置项
	Label `json:"label,omitempty"`
}

// MPStyleOpts contains styling options for a MarkPoint.
type MarkPointStyle struct {
	// 图元的图形类别
	// 可选 "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
	// 标注文本配置项
	Label `json:"label,omitempty"`
}

// RippleEffectOpts is the option set for the ripple effect.
type RippleEffect struct {
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

// LineStyleOpts is the option set for a link style component.
type LineStyle struct {
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

// AreaStyleOpts is the option set for an area style component.
type AreaStyle struct {
	// 填充区域的颜色
	Color string `json:"color,omitempty"`
	// 填充区域的透明度。支持从 0 到 1 的数字，为 0 时不填充区域
	Opacity float32 `json:"opacity,omitempty"`
}

// MLNameTypeItem represents type for a MarkLine.
type MarkLineNameTypeItem struct {
	// 标记线名称
	Name string `json:"name,omitempty"`
	// 内置类型，可选 "average", "min", "max"
	Type string `json:"type,omitempty"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkLine represents a mark line.
type MarkLines struct {
	Data []interface{} `json:"data,omitempty"`
	MarkLineStyle
}

// MarkPoint represents a mark point.
type MarkPoints struct {
	Data []interface{} `json:"data,omitempty"`
	MarkPointStyle
}

// MPNameTypeItem represents type for a MarkPoint.
type MarkPointNameTypeItem struct {
	// 标记点名称
	Name string `json:"name"`
	// 内置类型，可选 "average", "min", "max"
	Type string `json:"type"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MPNameCoordItem represents coordinates for a MarkPoint.
type MarkPointNameCoordItem struct {
	// 标记点名称
	Name string `json:"name,omitempty"`
	// 标记点坐标
	Coordinate []interface{} `json:"coord,omitempty"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

type MarkLineNameYAxisItem struct {
	// 标记线名称
	Name string `json:"name,omitempty"`
	// Y 轴数据
	YAxis interface{} `json:"yAxis,omitempty"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

type MarkLineNameXAxisItem struct {
	// 标记线名称
	Name string `json:"name,omitempty"`
	// X 轴数据
	XAxis interface{} `json:"xAxis,omitempty"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MLNameCoordItem represents coordinates for a MarkLine.
type MarkLineNameCoordItem struct {
	// 标记线名称
	Name string
	// 标记线起始坐标
	Coordinate0 []interface{}
	// 标记线结束坐标
	Coordinate1 []interface{}
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}
