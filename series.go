package goecharts

// 图形上的文本标签配置项
type LabelTextOpts struct {
	// 是否显示标签
	Show bool `json:"show,omitempty"`
	// 文字的颜色
	Color string `json:"color,omitempty"`
	// 标签的位置
	// 通过相对的百分比或者绝对像素值表示标签相对于图形包围盒左上角的位置。示例：
	// 绝对的像素值	position: [10, 10],
	// 相对的百分比	position: ['50%', '50%']
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
}

// MarkLine 配置项
type MarkPointOpts struct {
	Data []interface{} `json:"data,omitempty"`
	MPStyleOpts
}

// MarkPoint 风格配置项
type MPStyleOpts struct {
	// 图元的图形类别
	// 可选 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
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
}

// MarkPoint 数据 Name-Coordinates
type MPNameCoordItem struct {
	// 标记点名称
	Name string `json:"name"`
	// 标记点坐标
	Coord []interface{} `json:"coord"`
}

// MarkLine 配置项
type MarkLineOpts struct {
	Data []interface{} `json:"data,omitempty"`
	MLStyleOpts
}

// MarkLine 风格配置项
type MLStyleOpts struct {
	// 图元的图形类别
	// 可选 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
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
}

// MarkLine 数据 Name-YAxis
type MLNameYAxisItem struct {
	// 标记线名称
	Name string `json:"name"`
	// Y 轴数据
	YAxis interface{} `json:"yAxis"`
}

// MarkLine 数据 Name-XAxis
type MLNameXAxisItem struct {
	// 标记线名称
	Name string `json:"name"`
	// X 轴数据
	XAxis interface{} `json:"xAxis"`
}

// MarkLine 数据 Name-Coordinates
type MLNameCoordItem struct {
	// 标记线名称
	Name string
	// 标记线起始坐标
	Coord0 []interface{}
	// 标记线结束坐标
	Coord1 []interface{}
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

// Series 配置项
type singleSeries struct {
	// series 名称
	Name string `json:"name,omitempty"`
	// series 类型
	Type string `json:"type"`

	Stack      string `json:"stack,omitempty"`
	XAxisIndex int    `json:"xAxisIndex,omitempty"`
	YAxisIndex int    `json:"yAxisIndex,omitempty"`

	CoordSystem string `json:"coordinateSystem,omitempty"`

	// series 数据项
	Data interface{} `json:"data"`

	// series options
	ItemStyleOpts    `json:"itemStyle,omitempty"`
	LabelTextOpts    `json:"label,omitempty"`
	MarkLineOpts     `json:"markLine,omitempty"`
	MarkPointOpts    `json:"markPoint,omitempty"`
	RippleEffectOpts `json:"rippleEffect,omitempty"`
	LineStyleOpts    `json:"lineStyle,omitempty"`
	AreaStyleOpts    `json:"areaStyle,omitempty"`
}

// 设置 Series 配置项
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
		}
	}
}

// 设置 singleSeries 配置项
func (s *singleSeries) setSingleSeriesOpts(options ...interface{}) {
	s.switchSeriesOpts(options...)
}

// Series 列表
type Series []singleSeries

// 设置 Series 配置项
func (series *Series) SetSeriesConfig(options ...interface{}) {
	tsl := *series
	for i := 0; i < len(tsl); i++ {
		tsl[i].switchSeriesOpts(options...)
	}
}
