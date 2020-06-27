package opts

// todo: replace ' with " overall

// Label contains options for a label text.
type Label struct {
	// Whether to show label.
	Show bool `json:"show,omitempty"`

	// text color.
	// If set as 'auto', the color will assigned as visual color, such as series color.
	Color string `json:"color,omitempty"`

	// Label position. Followings are the options:
	//
	// [x, y]
	// Use relative percentage, or absolute pixel values to represent position of label
	// relative to top-left corner of bounding box. For example:
	//
	// Absolute pixel values: position: [10, 10],
	// Relative percentage: position: ['50%', '50%']
	//
	// 'top'
	// 'left'
	// 'right'
	// 'bottom'
	// 'inside'
	// 'insideLeft'
	// 'insideRight'
	// 'insideTop'
	// 'insideBottom'
	// 'insideTopLeft'
	// 'insideBottomLeft'
	// 'insideTopRight'
	// 'insideBottomRight'
	Position string `json:"position,omitempty"`

	// Data label formatter, which supports string template and callback function.
	// In either form, \n is supported to represent a new line.
	// String template, Model variation includes:
	//
	// {a}: series name.
	// {b}: the name of a data item.
	// {c}: the value of a data item.
	// {@xxx}: the value of a dimension named'xxx', for example,{@product}refers the value of'product'` dimension.
	// {@[n]}: the value of a dimension at the index ofn, for example,{@[3]}` refers the value at dimensions[3].
	Formatter string `json:"formatter,omitempty"`
}

// Emphasis
type Emphasis struct {
	// 高亮的标签样式
	Label `json:"label,omitempty"`
	// 高亮的图形样式
	ItemStyle `json:"itemStyle,omitempty"`
}

// ItemStyle
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

// MarkLine represents a mark line.
type MarkLines struct {
	Data []interface{} `json:"data,omitempty"`
	MarkLineStyle
}

// MarkLineStyle contains styling options for a MarkLine.
type MarkLineStyle struct {
	// 图元的图形类别
	// 可选 "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol []string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
	// 标线文本配置项
	*Label `json:"label,omitempty"`
}

// MarkLineNameTypeItem represents type for a MarkLine.
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

// MarkLineNameYAxisItem
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

// MarkLineNameXAxisItem
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

// MarkLineNameCoordItem represents coordinates for a MarkLine.
type MarkLineNameCoordItem struct {
	// 标记线名称
	Name string `json:"name,omitempty"`
	// 标记线起始坐标
	Coordinate0 []interface{}
	// 标记线结束坐标
	Coordinate1 []interface{}
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkPoint represents a mark point.
type MarkPoints struct {
	Data []interface{} `json:"data,omitempty"`
	MarkPointStyle
}

// MarkPointStyle contains styling options for a MarkPoint.
type MarkPointStyle struct {
	// 图元的图形类别
	// 可选 "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
	// 标注文本配置项
	*Label `json:"label,omitempty"`
}

// MarkPointNameTypeItem represents type for a MarkPoint.
type MarkPointNameTypeItem struct {
	// 标记点名称
	Name string `json:"name,omitempty"`
	// 内置类型，可选 "average", "min", "max"
	Type string `json:"type,omitempty"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkPointNameCoordItem represents coordinates for a MarkPoint.
type MarkPointNameCoordItem struct {
	// 标记点名称
	Name string `json:"name,omitempty"`
	// 标记点坐标
	Coordinate []interface{} `json:"coord,omitempty"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等
	// candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
	*Label   `json:"label,omitempty"`
}

// RippleEffect is the option set for the ripple effect.
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

// LineStyle is the option set for a link style component.
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

// AreaStyle is the option set for an area style component.
type AreaStyle struct {
	// 填充区域的颜色
	Color string `json:"color,omitempty"`
	// 填充区域的透明度。支持从 0 到 1 的数字，为 0 时不填充区域
	Opacity float32 `json:"opacity,omitempty"`
}
