package charts

type seriesOptser interface {
	markSeries()
}

// 图形上的文本标签配置项
type LabelTextOpts struct {
	// 是否显示标签
	Show bool `json:"show"`
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

func (LabelTextOpts) markSeries() {}

// MarkLine 配置项
type MarkPoint struct {
	Data []interface{} `json:"data,omitempty"`
	MPStyleOpts
}

// MarkPoint 风格配置项
type MPStyleOpts struct {
	// 图元的图形类别
	// 可选 "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
	// 标注文本配置项
	Label LabelTextOpts `json:"label,omitempty"`
}

func (MPStyleOpts) markSeries() {}

// MarkPoint 数据 Name-Type
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

func (MPNameTypeItem) markSeries() {}

// MarkPoint 数据 Name-Coordinates
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

func (MPNameCoordItem) markSeries() {}

// MarkLine 配置项
type MarkLine struct {
	Data []interface{} `json:"data,omitempty"`
	MLStyleOpts
}

// MarkLine 风格配置项
type MLStyleOpts struct {
	// 图元的图形类别
	// 可选 "circle", "rect", "roundRect", "triangle", "diamond", "pin", "arrow", "none"
	Symbol []string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
	// 标线文本配置项
	Label LabelTextOpts `json:"label,omitempty"`
}

func (MLStyleOpts) markSeries() {}

// MarkLine 数据 Name-Type
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

func (MLNameTypeItem) markSeries() {}

// MarkLine 数据 Name-YAxis
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

func (MLNameYAxisItem) markSeries() {}

// MarkLine 数据 Name-XAxis
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

func (MLNameXAxisItem) markSeries() {}

// MarkLine 数据 Name-Coordinates
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

func (MLNameCoordItem) markSeries() {}

// 数据 Item 配置项
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

func (ItemStyleOpts) markSeries() {}

type singleSeries struct {
	// 系列名称
	Name string `json:"name,omitempty"`
	// 系列类型
	Type string `json:"type"`

	// Rectangular charts
	Stack      string `json:"stack,omitempty"`
	XAxisIndex int    `json:"xAxisIndex,omitempty"`
	YAxisIndex int    `json:"yAxisIndex,omitempty"`

	// Bar chart
	BarGap         string `json:"barGap,omitempty"`
	BarCategoryGap string `json:"barCategoryGap,omitempty"`

	// Map chart
	MapType     string `json:"map,omitempty"`
	CoordSystem string `json:"coordinateSystem,omitempty"`

	// Line chart
	Step   bool `json:"step,omitempty"`
	Smooth bool `json:"smooth,omitempty"`

	// Pie chart
	RoseType interface{} `json:"roseType,omitempty"`
	Center   interface{} `json:"center,omitempty"`
	Radius   interface{} `json:"radius,omitempty"`

	// wordCloud chart
	Shape         string    `json:"shape,omitempty"`
	SizeRange     []float32 `json:"sizeRange,omitempty"`
	RotationRange []float32 `json:"rotationRange,omitempty"`

	// Liquid chart
	LiquidOutlineOpts `json:"outline,omitempty"`
	IsWaveAnimation   bool `json:"waveAnimation"`

	// bar3D chart
	Shading string `json:"shading,omitempty"`

	// 系列数据项
	Data interface{} `json:"data"`

	// 系列配置项
	ItemStyleOpts    `json:"itemStyle,omitempty"`
	LabelTextOpts    `json:"label,omitempty"`
	MarkLine         `json:"markLine,omitempty"`
	MarkPoint        `json:"markPoint,omitempty"`
	RippleEffectOpts `json:"rippleEffect,omitempty"`
	LineStyleOpts    `json:"lineStyle,omitempty"`
	AreaStyleOpts    `json:"areaStyle,omitempty"`
	TextStyleOpts    `json:"textStyle,omitempty"`
}

// 设置 singleSeries 配置项
func (s *singleSeries) switchSeriesOpts(options ...seriesOptser) {
	// 实际 MarkLevel Name Coordinates 结构
	type MLNameCoord struct {
		Name  string        `json:"name,omitempty"`
		Coord []interface{} `json:"coord"`
	}

	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case LabelTextOpts:
			s.LabelTextOpts = option.(LabelTextOpts)
		case RippleEffectOpts:
			s.RippleEffectOpts = option.(RippleEffectOpts)
		case LineStyleOpts:
			s.LineStyleOpts = option.(LineStyleOpts)
		case AreaStyleOpts:
			s.AreaStyleOpts = option.(AreaStyleOpts)
		case ItemStyleOpts:
			s.ItemStyleOpts = option.(ItemStyleOpts)
		case TextStyleOpts:
			s.TextStyleOpts = option.(TextStyleOpts)
			s.TextStyleOpts.Normal = &TextStyleOpts{
				Color:     s.TextStyleOpts.Color,
				FontSize:  s.TextStyleOpts.FontSize,
				FontStyle: s.TextStyleOpts.FontStyle,
			}

			// MarkLine 配置项
		case MLNameTypeItem:
			s.MarkLine.Data = append(s.MarkLine.Data, option.(MLNameTypeItem))
		case MLNameXAxisItem:
			s.MarkLine.Data = append(s.MarkLine.Data, option.(MLNameXAxisItem))
		case MLNameYAxisItem:
			s.MarkLine.Data = append(s.MarkLine.Data, option.(MLNameYAxisItem))
		case MLNameCoordItem:
			m := option.(MLNameCoordItem)
			s.MarkLine.Data = append(
				s.MarkLine.Data, []MLNameCoord{{Name: m.Name, Coord: m.Coord0}, {Coord: m.Coord1}})
		case MLStyleOpts:
			s.MarkLine.MLStyleOpts = option.(MLStyleOpts)

			// MarkPoint 配置项
		case MPNameTypeItem:
			s.MarkPoint.Data = append(s.MarkPoint.Data, option.(MPNameTypeItem))
		case MPNameCoordItem:
			s.MarkPoint.Data = append(s.MarkPoint.Data, option.(MPNameCoordItem))
		case MPStyleOpts:
			s.MarkPoint.MPStyleOpts = option.(MPStyleOpts)

		case BarOpts:
			opt := option.(BarOpts)
			opt.setChartOpt(s)
		case Bar3DOpts:
			opt := option.(Bar3DOpts)
			opt.setChartOpt(s)
		case HeatMapOpts:
			opt := option.(HeatMapOpts)
			opt.setChartOpt(s)
		case LineOpts:
			opt := option.(LineOpts)
			opt.setChartOpt(s)
		case LiquidOpts:
			opt := option.(LiquidOpts)
			opt.setChartOpt(s)
		case PieOpts:
			opt := option.(PieOpts)
			opt.setChartOpt(s)
		case ScatterOpts:
			opt := option.(ScatterOpts)
			opt.setChartOpt(s)
		case WordCLoudOpts:
			opt := option.(WordCLoudOpts)
			opt.setChartOpt(s)
		}
	}
}

func (s *singleSeries) setSingleSeriesOpts(options ...seriesOptser) {
	s.switchSeriesOpts(options...)
}

type Series []singleSeries

func (series *Series) exportSeries() Series {
	return *series
}

func (series *Series) SetSeriesOptions(options ...seriesOptser) {
	tsl := *series
	for i := 0; i < len(tsl); i++ {
		tsl[i].switchSeriesOpts(options...)
	}
}
