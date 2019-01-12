package charts

// 图形上的文本标签配置项
type LabelTextOpts struct {
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
}

// MarkLine 配置项
type MarkPointOpts struct {
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
}

// MarkPoint 数据 Name-Type
type MPNameTypeItem struct {
	// 标记点名称
	Name string `json:"name"`
	// 内置类型，可选 "average", "min", "max"
	Type string `json:"type"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等、candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkPoint 数据 Name-Coordinates
type MPNameCoordItem struct {
	// 标记点名称
	Name string `json:"name"`
	// 标记点坐标
	Coord []interface{} `json:"coord"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等、candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkLine 配置项
type MarkLineOpts struct {
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
}

// MarkLine 数据 Name-Type
type MLNameTypeItem struct {
	// 标记线名称
	Name string `json:"name"`
	// 内置类型，可选 "average", "min", "max"
	Type string `json:"type"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等、candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkLine 数据 Name-YAxis
type MLNameYAxisItem struct {
	// 标记线名称
	Name string `json:"name"`
	// Y 轴数据
	YAxis interface{} `json:"yAxis"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等、candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkLine 数据 Name-XAxis
type MLNameXAxisItem struct {
	// 标记线名称
	Name string `json:"name"`
	// X 轴数据
	XAxis interface{} `json:"xAxis"`
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等、candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

// MarkLine 数据 Name-Coordinates
type MLNameCoordItem struct {
	// 标记线名称
	Name string
	// 标记线起始坐标
	Coord0 []interface{}
	// 标记线结束坐标
	Coord1 []interface{}
	// 在使用 type 时有效，用于指定在哪个维度上指定最大值最小值。
	// 可以是维度的直接名称，例如折线图时可以是 x、angle 等、candlestick 图时可以是 open、close 等维度名称。
	ValueDim string `json:"valueDim,omitempty"`
}

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

type singleSeries struct {
	// 系列名称
	Name string `json:"name,omitempty"`
	// 系列类型
	Type string `json:"type"`

	// Rectangular charts
	// 数据堆叠，同个类目轴上系列配置相同的 stack 值后，后一个系列的值会在前一个系列的值上相加
	Stack string `json:"stack,omitempty"`
	// 使用的 X 轴的 index，在单个图表实例中存在多个 X 轴的时候有用
	XAxisIndex int `json:"xAxisIndex,omitempty"`
	// 使用的 Y 轴的 index，在单个图表实例中存在多个 Y 轴的时候有用
	YAxisIndex int `json:"yAxisIndex,omitempty"`

	// Map charts
	MapType     string `json:"map,omitempty"`
	CoordSystem string `json:"coordinateSystem,omitempty"`

	// Line charts
	// 是否是阶梯线图。可以设置为 true 显示成阶梯线图。
	Step bool `json:"step,omitempty"`
	// 是否平滑曲线显示
	Smooth bool `json:"smooth,omitempty"`

	// TODO: example
	// pie charts
	// 是否展示成南丁格尔图，通过半径区分数据大小。可选择两种模式：
	// "radius": 扇区圆心角展现数据的百分比，半径展现数据的大小。
	// "area": 所有扇区圆心角相同，仅通过半径展现数据大小。
	RoseType interface{} `json:"roseType,omitempty"`
	// 饼图的中心（圆心）坐标，数组的第一项是横坐标，第二项是纵坐标。
	// 支持设置成百分比，设置成百分比时第一项是相对于容器宽度，第二项是相对于容器高度。
	// 示例
	// center: [400, 300] 设置成绝对的像素值
	// center: ["50%", "50%"] 设置成相对的百分比
	// 默认
	// ["50%", "50%"]
	Center interface{} `json:"center,omitempty"`
	// 饼图的半径。可以为如下类型：
	// number：直接指定外半径值。
	// string：例如，"20%"，表示外半径为可视区尺寸（容器高宽中较小一项）的 20% 长度。
	// Array.<number|string>：数组的第一项是内半径，第二项是外半径。每一项遵从上述 number string 的描述。
	// 默认
	// [0, "75%"]
	Radius interface{} `json:"radius,omitempty"`

	// wordCloud
	// 词云图形状
	// 可选 "circle", "cardioid", "diamond", "triangle-forward", "triangle", "pentagon", "star"
	// 默认 "circle"
	Shape string `json:"shape,omitempty"`
	// 字体大小范围
	SizeRange []float32 `json:"sizeRange,omitempty"`
	// 字体倾斜角度范围
	RotationRange []float32 `json:"rotationRange,omitempty"`

	// Liquid
	LiquidOutlineOpts `json:"outline,omitempty"`
	IsWaveAnimation   bool `json:"waveAnimation"`

	// 系列数据项
	Data interface{} `json:"data"`

	// 系列配置项
	ItemStyleOpts    `json:"itemStyle,omitempty"`
	LabelTextOpts    `json:"label,omitempty"`
	MarkLineOpts     `json:"markLine,omitempty"`
	MarkPointOpts    `json:"markPoint,omitempty"`
	RippleEffectOpts `json:"rippleEffect,omitempty"`
	LineStyleOpts    `json:"lineStyle,omitempty"`
	AreaStyleOpts    `json:"areaStyle,omitempty"`
	TextStyleOpts    `json:"textStyle,omitempty"`
}

// 设置 singleSeries 配置项
func (s *singleSeries) switchSeriesOpts(options ...interface{}) {
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
			s.MarkLineOpts.Data = append(s.MarkLineOpts.Data, option.(MLNameTypeItem))
		case MLNameXAxisItem:
			s.MarkLineOpts.Data = append(s.MarkLineOpts.Data, option.(MLNameXAxisItem))
		case MLNameYAxisItem:
			s.MarkLineOpts.Data = append(s.MarkLineOpts.Data, option.(MLNameYAxisItem))
		case MLNameCoordItem:
			m := option.(MLNameCoordItem)
			s.MarkLineOpts.Data = append(
				s.MarkLineOpts.Data, []MLNameCoord{{Name: m.Name, Coord: m.Coord0}, {Coord: m.Coord1}})
		case MLStyleOpts:
			s.MarkLineOpts.MLStyleOpts = option.(MLStyleOpts)

			// MarkPoint 配置项
		case MPNameTypeItem:
			s.MarkPointOpts.Data = append(s.MarkPointOpts.Data, option.(MPNameTypeItem))
		case MPNameCoordItem:
			s.MarkPointOpts.Data = append(s.MarkPointOpts.Data, option.(MPNameCoordItem))
		case MPStyleOpts:
			s.MarkPointOpts.MPStyleOpts = option.(MPStyleOpts)

		case BarOpts:
			opt := option.(BarOpts)
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
		case WordCLoudOpts:
			opt := option.(WordCLoudOpts)
			opt.setChartOpt(s)
		}
	}
}

func (s *singleSeries) setSingleSeriesOpts(options ...interface{}) {
	s.switchSeriesOpts(options...)
}

type Series []singleSeries

type serieser interface {
	exportSeries() Series
}

func (series *Series) exportSeries() Series {
	return *series
}

func (series *Series) SetSeriesConfig(options ...interface{}) {
	tsl := *series
	for i := 0; i < len(tsl); i++ {
		tsl[i].switchSeriesOpts(options...)
	}
}
