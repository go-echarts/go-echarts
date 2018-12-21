package goecharts

// 图形上的文本标签配置项
type LabelTextOpts struct {
	Show     bool   `json:"show,omitempty"`
	Color    string `json:"color,omitempty"`
	Position string `json:"position,omitempty"`
}

// MarkLine 配置项
type MarkPointOpts struct {
	Data []interface{} `json:"data,omitempty"`
	MarkPointStyle
}

// MarkPoint 风格配置项
type MarkPointStyle struct {
	Symbol     string  `json:"symbol,omitempty"`
	SymbolSize float32 `json:"symbolSize,omitempty"`
}

// MarkPoint 数据 Name-Type
type MPNameType struct {
	// 标记点名称
	Name string `json:"name"`
	// 内置类型，可选"average", "min", "max"
	Type string `json:"type"`
}

// MarkPoint 数据 Name-Coordinates
type MPNameCoord struct {
	// 标记点名称
	Name string `json:"name"`
	// 标记点坐标
	Coord []interface{} `json:"coord"`
}

// MarkLine 配置项
type MarkLineOpts struct {
	Data []interface{} `json:"data,omitempty"`
	MarkLineStyle
}

// MarkLine 风格配置项
type MarkLineStyle struct {
	Symbol     []string `json:"symbol,omitempty"`
	SymbolSize float32  `json:"symbolSize,omitempty"`
}

// MarkLine 数据 Name-Type
type MLNameType struct {
	// 标记线名称
	Name string `json:"name"`
	// 内置类型，可选"average", "min", "max"
	Type string `json:"type"`
}

// MarkLine 数据 Name-YAxis
type MLNameYAxis struct {
	// 标记线名称
	Name string `json:"name"`
	// Y 轴数据
	YAxis interface{} `json:"yAxis"`
}

// MarkLine 数据 Name-XAxis
type MLNameXAxis struct {
	// 标记线名称
	Name string `json:"name"`
	// X 轴数据
	XAxis interface{} `json:"xAxis"`
}

// MarkLine 数据 Name-Coordinates
type MLNameCoords struct {
	// 标记线名称
	Name string
	// 标记线起始坐标
	Coord0 []interface{}
	// 标记线结束坐标
	Coord1 []interface{}
}

// Series 配置项
type singleSeries struct {
	// series 名称
	Name string `json:"name,omitempty"`
	// series 类型
	Type string `json:"type"`
	// series 数据项
	Data                 interface{} `json:"data"`
	LabelTextOpts     `json:"label,omitempty"`
	MarkLineOpts      `json:"markLine,omitempty"`
	MarkPointOpts     `json:"markPoint,omitempty"`
	*RippleEffectOpts `json:"rippleEffect,omitempty"`
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
			tmp := new(RippleEffectOpts)
			*tmp = option.(RippleEffectOpts)
			s.RippleEffectOpts = tmp

			// MarkLine 配置项
		case MLNameType:
			s.MarkLineOpts.Data = append(s.MarkLineOpts.Data, option.(MLNameType))
		case MLNameXAxis:
			s.MarkLineOpts.Data = append(s.MarkLineOpts.Data, option.(MLNameXAxis))
		case MLNameYAxis:
			s.MarkLineOpts.Data = append(s.MarkLineOpts.Data, option.(MLNameYAxis))
		case MLNameCoords:
			m := option.(MLNameCoords)
			s.MarkLineOpts.Data = append(
				s.MarkLineOpts.Data, []MLNameCoord{{Name: m.Name, Coord: m.Coord0}, {Coord: m.Coord1}})
		case MarkLineStyle:
			s.MarkLineOpts.MarkLineStyle = option.(MarkLineStyle)

			// MarkPoint 配置项
		case MPNameType:
			s.MarkPointOpts.Data = append(s.MarkPointOpts.Data, option.(MPNameType))
		case MPNameCoord:
			s.MarkPointOpts.Data = append(s.MarkPointOpts.Data, option.(MPNameCoord))
		case MarkPointStyle:
			s.MarkPointOpts.MarkPointStyle = option.(MarkPointStyle)
		}
	}
}

// 设置 Series 配置项
func (s *singleSeries) setSingleSeriesOpts(options ...interface{}) {
	s.switchSeriesOpts(options...)
}

// Series 列表
type Series []singleSeries

// 设置 SeriesList 配置项
func (series *Series) setAllSeriesOpts(options ...interface{}) {
	tsl := *series
	for i := 0; i < len(tsl); i++ {
		tsl[i].switchSeriesOpts(options...)
	}
}
