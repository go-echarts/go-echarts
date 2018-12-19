package goecharts

// 图形上的文本标签配置项
type LabelTextOptions struct {
	Show     bool   `json:"show,omitempty"`
	Color    string `json:"color,omitempty"`
	Position string `json:"position,omitempty"`
}

// MarkLine 配置项
type MarkPointOptions struct {
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
type MarkLineOptions struct {
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
type Series struct {
	// series 名称
	Name string `json:"name,omitempty"`
	// series 类型
	Type string `json:"type"`
	// series 数据项
	Data                 interface{} `json:"data"`
	LabelTextOptions     `json:"label,omitempty"`
	MarkLineOptions      `json:"markLine,omitempty"`
	MarkPointOptions     `json:"markPoint,omitempty"`
	*RippleEffectOptions `json:"rippleEffect,omitempty"`
}

// 设置 Series 配置项
func (series *Series) setSingleSeriesOptions(options ...interface{}) {

	// 方法内部大写 struct 类型不会对外暴露
	type MLNameCoord struct {
		Name  string        `json:"name,omitempty"`
		Coord []interface{} `json:"coord"`
	}

	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case LabelTextOptions:
			series.LabelTextOptions = option.(LabelTextOptions)
		case RippleEffectOptions:
			tmp := new(RippleEffectOptions)
			*tmp = option.(RippleEffectOptions)
			series.RippleEffectOptions = tmp

			// MarkLine 配置项
		case MLNameType:
			series.MarkLineOptions.Data = append(series.MarkLineOptions.Data, option.(MLNameType))
		case MLNameXAxis:
			series.MarkLineOptions.Data = append(series.MarkLineOptions.Data, option.(MLNameXAxis))
		case MLNameYAxis:
			series.MarkLineOptions.Data = append(series.MarkLineOptions.Data, option.(MLNameYAxis))
		case MLNameCoords:
			m := option.(MLNameCoords)
			series.MarkLineOptions.Data = append(
				series.MarkLineOptions.Data, []MLNameCoord{{Name: m.Name, Coord: m.Coord0}, {Coord: m.Coord1}})
		case MarkLineStyle:
			series.MarkLineOptions.MarkLineStyle = option.(MarkLineStyle)

			// MarkPoint 配置项
		case MPNameType:
			series.MarkPointOptions.Data = append(series.MarkPointOptions.Data, option.(MPNameType))
		case MPNameCoord:
			series.MarkPointOptions.Data = append(series.MarkPointOptions.Data, option.(MPNameCoord))
		case MarkPointStyle:
			series.MarkPointOptions.MarkPointStyle = option.(MarkPointStyle)
		}
	}
}

// Series 列表
type SeriesList []Series

// TODO: 两个 setSeriesOptions 整合
// TODO: MarkLine&MarkPoint StyleOptions

// 设置 SeriesList 配置项
func (sl *SeriesList) setSeriesOptions(options ...interface{}) {
	tsl := *sl
	for i := 0; i < len(tsl); i++ {
		for j := 0; j < len(options); j++ {
			option := options[j]
			switch option.(type) {
			case LabelTextOptions:
				tsl[i].LabelTextOptions = option.(LabelTextOptions)
			case RippleEffectOptions:
				tmp := new(RippleEffectOptions)
				*tmp = option.(RippleEffectOptions)
				tsl[i].RippleEffectOptions = tmp
			case MarkLineStyle:
				tsl[i].MarkLineOptions.MarkLineStyle = option.(MarkLineStyle)
			case MarkPointStyle:
				tsl[i].MarkPointOptions.MarkPointStyle = option.(MarkPointStyle)
			}
		}
	}
}
