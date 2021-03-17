package charts

import "github.com/go-echarts/go-echarts/v2/opts"

type SingleSeries struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	// Rectangular charts
	Stack      string `json:"stack,omitempty"`
	XAxisIndex int    `json:"xAxisIndex,omitempty"`
	YAxisIndex int    `json:"yAxisIndex,omitempty"`

	// Bar
	BarGap         string `json:"barGap,omitempty"`
	BarCategoryGap string `json:"barCategoryGap,omitempty"`
	ShowBackground bool   `json:"showBackground,omitempty"`
	RoundCap       bool   `json:"roundCap,omitempty"`

	// Bar3D
	Shading string `json:"shading,omitempty"`

	// Graph
	Links              interface{} `json:"links,omitempty"`
	Layout             string      `json:"layout,omitempty"`
	Force              interface{} `json:"force,omitempty"`
	Categories         interface{} `json:"categories,omitempty"`
	Roam               bool        `json:"roam,omitempty"`
	FocusNodeAdjacency bool        `json:"focusNodeAdjacency,omitempty"`

	// Line
	Step         bool `json:"step,omitempty"`
	Smooth       bool `json:"smooth,omitempty"`
	ConnectNulls bool `json:"connectNulls,omitempty"`

	// Liquid
	IsLiquidOutline bool `json:"outline,omitempty"`
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

	// Tree
	Orient            string      `json:"orient,omitempty"`
	ExpandAndCollapse bool        `json:"expandAndCollapse,omitempty"`
	InitialTreeDepth  int         `json:"initialTreeDepth,omitempty"`
	Leaves            interface{} `json:"leaves,omitempty"`
	Left              string      `json:"left,omitempty"`
	Right             string      `json:"right,omitempty"`
	Top               string      `json:"top,omitempty"`
	Bottom            string      `json:"bottom,omitempty"`

	// WordCloud
	Shape         string    `json:"shape,omitempty"`
	SizeRange     []float32 `json:"sizeRange,omitempty"`
	RotationRange []float32 `json:"rotationRange,omitempty"`

	// Sunburst
	NodeClick               string `json:"nodeClick,omitempty"`
	Sort                    string `json:"sort,omitempty"`
	RenderLabelForZeroData  bool   `json:"renderLabelForZeroData"`
	SelectedMode            bool   `json:"selectedMode"`
	Animation               bool   `json:"animation"`
	AnimationThreshold      int    `json:"animationThreshold,omitempty"`
	AnimationDuration       int    `json:"animationDuration,omitempty"`
	AnimationEasing         string `json:"animationEasing,omitempty"`
	AnimationDelay          int    `json:"animationDelay,omitempty"`
	AnimationDurationUpdate int    `json:"animationDurationUpdate,omitempty"`
	AnimationEasingUpdate   string `json:"animationEasingUpdate,omitempty"`
	AnimationDelayUpdate    int    `json:"animationDelayUpdate,omitempty"`

	// series data
	Data interface{} `json:"data"`

	// series options
	*opts.ItemStyle    `json:"itemStyle,omitempty"`
	*opts.Label        `json:"label,omitempty"`
	*opts.LabelLine    `json:"labelLine,omitempty"`
	*opts.Emphasis     `json:"emphasis,omitempty"`
	*opts.MarkLines    `json:"markLine,omitempty"`
	*opts.MarkPoints   `json:"markPoint,omitempty"`
	*opts.RippleEffect `json:"rippleEffect,omitempty"`
	*opts.LineStyle    `json:"lineStyle,omitempty"`
	*opts.AreaStyle    `json:"areaStyle,omitempty"`
	*opts.TextStyle    `json:"textStyle,omitempty"`
}

type SeriesOpts func(s *SingleSeries)

// WithLabelOpts
func WithLabelOpts(opt opts.Label) SeriesOpts {
	return func(s *SingleSeries) {
		s.Label = &opt
	}
}

// WithEmphasisOpts
func WithEmphasisOpts(opt opts.Emphasis) SeriesOpts {
	return func(s *SingleSeries) {
		s.Emphasis = &opt
	}
}

// WithAreaStyleOpts
func WithAreaStyleOpts(opt opts.AreaStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.AreaStyle = &opt
	}
}

// WithItemStyleOpts
func WithItemStyleOpts(opt opts.ItemStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.ItemStyle = &opt
	}
}

// WithRippleEffectOpts
func WithRippleEffectOpts(opt opts.RippleEffect) SeriesOpts {
	return func(s *SingleSeries) {
		s.RippleEffect = &opt
	}
}

// WithLineStyleOpts
func WithLineStyleOpts(opt opts.LineStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.LineStyle = &opt
	}
}

/* Chart Options */

// WithBarChartOpts
func WithBarChartOpts(opt opts.BarChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Stack = opt.Stack
		s.BarGap = opt.BarGap
		s.BarCategoryGap = opt.BarCategoryGap
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.ShowBackground = opt.ShowBackground
		s.RoundCap = opt.RoundCap
		s.CoordSystem = opt.CoordSystem
		s.Type = opt.Type
	}
}

func WithSunburstOpts(opt opts.SunburstChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.NodeClick = opt.NodeClick
		s.Sort = opt.Sort
		s.RenderLabelForZeroData = opt.RenderLabelForZeroData
		s.SelectedMode = opt.SelectedMode
		s.Animation = opt.Animation
		s.AnimationThreshold = opt.AnimationThreshold
		s.AnimationDuration = opt.AnimationDuration
		s.AnimationEasing = opt.AnimationEasing
		s.AnimationDelay = opt.AnimationDelay
		s.AnimationDurationUpdate = opt.AnimationDurationUpdate
		s.AnimationEasingUpdate = opt.AnimationEasingUpdate
		s.AnimationDelayUpdate = opt.AnimationDelayUpdate
	}
}

// WithGraphChartOpts
func WithGraphChartOpts(opt opts.GraphChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Layout = opt.Layout
		s.Force = opt.Force
		s.Roam = opt.Roam
		s.FocusNodeAdjacency = opt.FocusNodeAdjacency
		s.Categories = opt.Categories
	}
}

// WithHeatMapChartOpts
func WithHeatMapChartOpts(opt opts.HeatMapChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

// WithLineChartOpts
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

// WithPieChartOpts
func WithPieChartOpts(opt opts.PieChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.RoseType = opt.RoseType
		s.Center = opt.Center
		s.Radius = opt.Radius
	}
}

// WithScatterChartOpts
func WithScatterChartOpts(opt opts.ScatterChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

// WithLiquidChartOpts
func WithLiquidChartOpts(opt opts.LiquidChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.IsLiquidOutline = opt.IsShowOutline
		s.IsWaveAnimation = opt.IsWaveAnimation
	}
}

// WithBar3DChartOpts
func WithBar3DChartOpts(opt opts.Bar3DChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shading = opt.Shading
	}
}

// WithTreeOpts
func WithTreeOpts(opt opts.TreeChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Layout = opt.Layout
		s.Orient = opt.Orient
		s.ExpandAndCollapse = opt.ExpandAndCollapse
		s.InitialTreeDepth = opt.InitialTreeDepth
		s.Roam = opt.Roam
		s.Label = opt.Label
		s.Leaves = opt.Leaves
		s.Right = opt.Right
		s.Left = opt.Left
		s.Top = opt.Top
		s.Bottom = opt.Bottom
	}
}

// WithWorldCloudChartOpts
func WithWorldCloudChartOpts(opt opts.WordCloudChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.SizeRange = opt.SizeRange
		s.RotationRange = opt.RotationRange
	}
}

// WithMarkLineNameTypeItemOpts
func WithMarkLineNameTypeItemOpts(opt ...opts.MarkLineNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkLineNameXAxisItemOpts
func WithMarkLineNameXAxisItemOpts(opt ...opts.MarkLineNameXAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkLineNameYAxisItemOpts
func WithMarkLineNameYAxisItemOpts(opt ...opts.MarkLineNameYAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, o)
		}
	}
}

// WithMarkPointNameTypeItemOpts
func WithMarkPointNameTypeItemOpts(opt ...opts.MarkPointNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}
		for _, o := range opt {
			s.MarkPoints.Data = append(s.MarkPoints.Data, o)
		}
	}
}

// WithMarkPointStyleOpts
func WithMarkPointStyleOpts(opt opts.MarkPointStyle) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}

		s.MarkPoints.MarkPointStyle = opt
	}
}

// WithMarkPointNameCoordItemOpts
func WithMarkPointNameCoordItemOpts(opt ...opts.MarkPointNameCoordItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}
		for _, o := range opt {
			s.MarkPoints.Data = append(s.MarkPoints.Data, o)
		}
	}
}

func (s *SingleSeries) configureSeriesOpts(options ...SeriesOpts) {
	for _, opt := range options {
		opt(s)
	}
}

// MultiSeries represents multiple series.
type MultiSeries []SingleSeries

// SetSeriesOptions sets options for all the series.
// Previous options will be overwrote every time hence setting them on the `AddSeries` if you want
// to customize each series individually
// 															 here -> ↓ <-
// func (c *Bar) AddSeries(name string, data []opts.BarData, options ...SeriesOpts)
func (ms *MultiSeries) SetSeriesOptions(opts ...SeriesOpts) {
	s := *ms
	for i := 0; i < len(s); i++ {
		s[i].configureSeriesOpts(opts...)
	}
}
