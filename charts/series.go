package charts

import "github.com/go-echarts/go-echarts/opts"

type SingleSeries struct {
	// series name
	Name string `json:"name,omitempty"`
	// series type
	Type string `json:"type"`

	// Rectangular charts
	Stack      string `json:"stack,omitempty"`
	XAxisIndex int    `json:"xAxisIndex,omitempty"`
	YAxisIndex int    `json:"yAxisIndex,omitempty"`

	// Bar
	BarGap         string `json:"barGap,omitempty"`
	BarCategoryGap string `json:"barCategoryGap,omitempty"`

	// Bar3D
	Shading string `json:"shading,omitempty"`

	// Graph
	Links              interface{}     `json:"links,omitempty"`
	Layout             string          `json:"layout,omitempty"`
	Force              GraphForce      `json:"force,omitempty"`
	Categories         []GraphCategory `json:"categories,omitempty"`
	Roam               bool            `json:"roam,omitempty"`
	FocusNodeAdjacency bool            `json:"focusNodeAdjacency,omitempty"`

	// Line
	Step         bool `json:"step,omitempty"`
	Smooth       bool `json:"smooth,omitempty"`
	ConnectNulls bool `json:"connectNulls,omitempty"`

	// Liquid
	LiquidOutline   `json:"outline,omitempty"`
	IsWaveAnimation bool `json:"waveAnimation"`

	// Map
	MapType     string `json:"map,omitempty"`
	CoordSystem string `json:"coordinateSystem,omitempty"`

	// Pie
	RoseType interface{} `json:"roseType,omitempty"`
	Center   interface{} `json:"center,omitempty"`
	Radius   interface{} `json:"radius,omitempty"`

	// Scatter
	SymbolSize float32 `json:"symbolSize,omitempty"`

	// WordCloud
	Shape         string    `json:"shape,omitempty"`
	SizeRange     []float32 `json:"sizeRange,omitempty"`
	RotationRange []float32 `json:"rotationRange,omitempty"`

	// 系列数据项
	Data interface{} `json:"data"`

	// 系列配置项
	opts.ItemStyle    `json:"itemStyle,omitempty"`
	opts.LabelText    `json:"label,omitempty"`
	opts.Emphasis     `json:"emphasis,omitempty"`
	opts.MarkLines    `json:"markLine,omitempty"`
	opts.MarkPoints   `json:"markPoint,omitempty"`
	opts.RippleEffect `json:"rippleEffect,omitempty"`
	opts.LineStyle    `json:"lineStyle,omitempty"`
	opts.AreaStyle    `json:"areaStyle,omitempty"`
	TextStyleOpts     `json:"textStyle,omitempty"`
}

type SeriesOpts func(s *SingleSeries)

func WithLabelTextOpts(opt opts.LabelText) SeriesOpts {
	return func(s *SingleSeries) {
		s.LabelText = opt
	}
}

func WithEmphasisOpts(opt opts.Emphasis) SeriesOpts {
	return func(s *SingleSeries) {
		s.Emphasis = opt
	}
}

func WithRippleEffectOpts(opt opts.RippleEffect) SeriesOpts {
	return func(s *SingleSeries) {
		s.RippleEffect = opt
	}
}

func WithLineStyleOpts(opt opts.LineStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.LineStyle = opt
	}
}

func WithBarChartOpts(opt opts.BarChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Stack = opt.Stack
		s.BarGap = opt.BarGap
		s.BarCategoryGap = opt.BarCategoryGap
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

func WithGraphChartOpts(opt opts.GraphChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Layout = opt.Layout
		s.Force = opt.Force
		s.Roam = opt.Roam
		s.FocusNodeAdjacency = opt.FocusNodeAdjacency
		s.Categories = opt.Categories
	}
}

func WithHeatMapChartOpts(opt opts.HeatMapChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

func WithLineChartOpts(opt opts.LineChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.YAxisIndex = opt.YAxisIndex
		s.Stack = opt.Stack
		s.Smooth = opt.Smooth
		s.Step = opt.Step
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.ConnectNulls = opt.ConnectNulls
	}
}

func WithPieChartOpts(opt opts.PieChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.RoseType = opt.RoseType
		s.Center = opt.Center
		s.Radius = opt.Radius
	}
}

func WithScatterChartOpts(opt opts.ScatterChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

func WithLiquidChartOpts(opt opts.LiquidChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.LiquidOutlineOpts.Show = opt.IsShowOutline
		s.IsWaveAnimation = opt.IsWaveAnimation
	}
}

func WithBar3DChartOpts(opt opts.Bar3DChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shading = opt.Shading
	}
}

func WithWorldCloudChartOpts(opt opts.WordCloudChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.SizeRange = opt.SizeRange
		s.RotationRange = opt.RotationRange
	}
}

func WithMarkLineNameTypeItemOpts(opt opts.MarkLineNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		s.MarkLines.Data = append(s.MarkLines.Data, opt)
	}
}

func WithMarkLineNameXAxisItemOpts(opt opts.MarkLineNameXAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		s.MarkLines.Data = append(s.MarkLines.Data, opt)
	}
}

func WithMarkLineNameYAxisItemOpts(opt opts.MarkLineNameYAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		s.MarkLines.Data = append(s.MarkLines.Data, opt)
	}
}

func (s *SingleSeries) configureSeriesFns(fns ...SeriesOpts) {
	for _, fn := range fns {
		fn(s)
	}
}

// 设置 singleSeries 配置项
func (s *singleSeries) switchSeriesOpts(options ...SeriesOptser) {
	// 实际 MarkLevel Name Coordinates 结构
	type MLNameCoord struct {
		Name  string        `json:"name,omitempty"`
		Coord []interface{} `json:"coord"`
	}

	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option := option.(type) {
		//case LabelTextOpts:
		//	s.LabelTextOpts = option
		//case EmphasisOpts:
		//	s.EmphasisOpts = option
		//case RippleEffectOpts:
		//	s.RippleEffectOpts = option
		//case LineStyleOpts:
		//	s.LineStyleOpts = option
		//case AreaStyleOpts:
		//	s.AreaStyleOpts = option
		//case ItemStyleOpts:
		//	s.ItemStyleOpts = option
		case TextStyleOpts:
			s.TextStyleOpts = option
			s.TextStyleOpts.Normal = &TextStyleOpts{
				Color:     s.TextStyleOpts.Color,
				FontSize:  s.TextStyleOpts.FontSize,
				FontStyle: s.TextStyleOpts.FontStyle,
			}

			// MarkLine 配置项
		//case MLNameTypeItem:
		//	s.MarkLine.Data = append(s.MarkLine.Data, option)
		//case MLNameXAxisItem:
		//	s.MarkLine.Data = append(s.MarkLine.Data, option)
		//case MLNameYAxisItem:
		//	s.MarkLine.Data = append(s.MarkLine.Data, option)
		case MLNameCoordItem:
			m := option
			s.MarkLine.Data = append(
				s.MarkLine.Data, []MLNameCoord{{Name: m.Name, Coord: m.Coord0}, {Coord: m.Coord1}})
		//case MLStyleOpts:
		//	s.MarkLine.MLStyleOpts = option

		// MarkPoint 配置项
		case MPNameTypeItem:
			s.MarkPoint.Data = append(s.MarkPoint.Data, option)
		case MPNameCoordItem:
			s.MarkPoint.Data = append(s.MarkPoint.Data, option)
		case MPStyleOpts:
			s.MarkPoint.MPStyleOpts = option

			//case BarOpts:
			//	opt := option
			//	opt.setChartOpt(s)
			//case Bar3DOpts:
			//	opt := option
			//	opt.setChartOpt(s)
			//case GraphOpts:
			//	opt := option
			//	opt.setChartOpt(s)
			//case HeatMapOpts:
			//	opt := option
			//	opt.setChartOpt(s)
			//case LineOpts:
			//opt := option
			//option.setChartOpt(s)
			//case LiquidOpts:
			//	opt := option
			//	opt.setChartOpt(s)
			//case PieOpts:
			//	opt := option
			//	opt.setChartOpt(s)
			//case ScatterOpts:
			//	opt := option
			//	opt.setChartOpt(s)
			//case WordCloudOpts:
			//	opt := option
			//	opt.setChartOpt(s)
			//}
		}
	}
}

//func (s *singleSeries) setSingleSeriesOpts(options ...SeriesOptser) {
//	s.switchSeriesOpts(options...)
//}

// Series represents multiple series.
type MultiSeries []SingleSeries

//func (ms *MultiSeries) exportSeries() SingleSeries {
//
//}

// SetSeriesOptions sets options for the series.
func (ms *MultiSeries) SetSeriesOptions(fns ...SeriesOpts) {
	s := *ms
	for i := 0; i < len(s); i++ {
		s[i].configureSeriesFns(fns...)
	}
}
