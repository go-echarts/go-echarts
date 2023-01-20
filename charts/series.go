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
	EdgeSymbol         interface{} `json:"edgeSymbol,omitempty"`
	EdgeSymbolSize     interface{} `json:"edgeSymbolSize,omitempty"`
	EdgeLabel          interface{} `json:"edgeLabel,omitempty"`
	Draggable          bool        `json:"draggable,omitempty"`
	FocusNodeAdjacency bool        `json:"focusNodeAdjacency,omitempty"`

	// Line
	Step         interface{} `json:"step,omitempty"`
	Smooth       bool        `json:"smooth"`
	ConnectNulls bool        `json:"connectNulls"`
	ShowSymbol   bool        `json:"showSymbol"`

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

	// TreeMap
	LeafDepth  int         `json:"leafDepth,omitempty"`
	Levels     interface{} `json:"levels,omitempty"`
	UpperLabel interface{} `json:"upperLabel,omitempty"`

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
	*opts.Encode        `json:"encode,omitempty"`
	*opts.ItemStyle     `json:"itemStyle,omitempty"`
	*opts.Label         `json:"label,omitempty"`
	*opts.LabelLine     `json:"labelLine,omitempty"`
	*opts.Emphasis      `json:"emphasis,omitempty"`
	*opts.MarkLines     `json:"markLine,omitempty"`
	*opts.MarkAreas     `json:"markArea,omitempty"`
	*opts.MarkPoints    `json:"markPoint,omitempty"`
	*opts.RippleEffect  `json:"rippleEffect,omitempty"`
	*opts.LineStyle     `json:"lineStyle,omitempty"`
	*opts.AreaStyle     `json:"areaStyle,omitempty"`
	*opts.TextStyle     `json:"textStyle,omitempty"`
	*opts.CircularStyle `json:"circular,omitempty"`
}

type SeriesOpts func(s *SingleSeries)

// WithLabelOpts sets the label.
func WithLabelOpts(opt opts.Label) SeriesOpts {
	return func(s *SingleSeries) {
		s.Label = &opt
	}
}

// WithEmphasisOpts sets the emphasis.
func WithEmphasisOpts(opt opts.Emphasis) SeriesOpts {
	return func(s *SingleSeries) {
		s.Emphasis = &opt
	}
}

// WithAreaStyleOpts sets the area style.
func WithAreaStyleOpts(opt opts.AreaStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.AreaStyle = &opt
	}
}

// WithItemStyleOpts sets the item style.
func WithItemStyleOpts(opt opts.ItemStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.ItemStyle = &opt
	}
}

// WithRippleEffectOpts sets the ripple effect.
func WithRippleEffectOpts(opt opts.RippleEffect) SeriesOpts {
	return func(s *SingleSeries) {
		s.RippleEffect = &opt
	}
}

// WithLineStyleOpts sets the line style.
func WithLineStyleOpts(opt opts.LineStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.LineStyle = &opt
	}
}

// With CircularStyle Opts
func WithCircularStyleOpts(opt opts.CircularStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.CircularStyle = &opt
	}
}

/* Chart Options */

// WithBarChartOpts sets the BarChart option.
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
	}
}

// WithSunburstOpts sets the SunburstChart option.
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

// WithGraphChartOpts sets the GraphChart option.
func WithGraphChartOpts(opt opts.GraphChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Layout = opt.Layout
		s.Force = opt.Force
		s.Roam = opt.Roam
		s.EdgeSymbol = opt.EdgeSymbol
		s.EdgeSymbolSize = opt.EdgeSymbolSize
		s.Draggable = opt.Draggable
		s.FocusNodeAdjacency = opt.FocusNodeAdjacency
		s.Categories = opt.Categories
		s.EdgeLabel = opt.EdgeLabel
	}
}

// WithHeatMapChartOpts sets the HeatMapChart option.
func WithHeatMapChartOpts(opt opts.HeatMapChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

// WithLineChartOpts sets the LineChart option.
func WithLineChartOpts(opt opts.LineChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.YAxisIndex = opt.YAxisIndex
		s.Stack = opt.Stack
		s.Smooth = opt.Smooth
		s.ShowSymbol = opt.ShowSymbol
		s.Step = opt.Step
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.ConnectNulls = opt.ConnectNulls
	}
}

// WithPieChartOpts sets the PieChart option.
func WithPieChartOpts(opt opts.PieChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.RoseType = opt.RoseType
		s.Center = opt.Center
		s.Radius = opt.Radius
	}
}

// WithScatterChartOpts sets the ScatterChart option.
func WithScatterChartOpts(opt opts.ScatterChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
	}
}

// WithLiquidChartOpts sets the LiquidChart option.
func WithLiquidChartOpts(opt opts.LiquidChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.IsLiquidOutline = opt.IsShowOutline
		s.IsWaveAnimation = opt.IsWaveAnimation
	}
}

// WithBar3DChartOpts sets the Bar3DChart option.
func WithBar3DChartOpts(opt opts.Bar3DChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shading = opt.Shading
	}
}

// WithTreeOpts sets the TreeChart option.
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

// WithTreeMapOpts sets the TreeMapChart options.
func WithTreeMapOpts(opt opts.TreeMapChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Animation = opt.Animation
		s.LeafDepth = opt.LeafDepth
		s.Roam = opt.Roam
		s.Levels = opt.Levels
		s.UpperLabel = opt.UpperLabel
		s.Right = opt.Right
		s.Left = opt.Left
		s.Top = opt.Top
		s.Bottom = opt.Bottom
	}
}

// WithWorldCloudChartOpts sets the WorldCloudChart option.
func WithWorldCloudChartOpts(opt opts.WordCloudChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.Shape = opt.Shape
		s.SizeRange = opt.SizeRange
		s.RotationRange = opt.RotationRange
	}
}

// WithMarkLineNameTypeItemOpts sets the type of the MarkLine.
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

// WithMarkLineStyleOpts sets the style of the MarkLine.
func WithMarkLineStyleOpts(opt opts.MarkLineStyle) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}

		s.MarkLines.MarkLineStyle = opt
	}
}

// WithMarkLineNameCoordItemOpts sets the coordinates of the MarkLine.
func WithMarkLineNameCoordItemOpts(opt ...opts.MarkLineNameCoordItem) SeriesOpts {
	type MLNameCoord struct {
		Name  string        `json:"name,omitempty"`
		Coord []interface{} `json:"coord"`
	}
	return func(s *SingleSeries) {
		if s.MarkLines == nil {
			s.MarkLines = &opts.MarkLines{}
		}
		for _, o := range opt {
			s.MarkLines.Data = append(s.MarkLines.Data, []MLNameCoord{{Name: o.Name, Coord: o.Coordinate0}, {Coord: o.Coordinate1}})
		}
	}
}

// WithMarkLineNameXAxisItemOpts sets the X axis of the MarkLine.
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

// WithMarkLineNameYAxisItemOpts sets the Y axis of the MarkLine.
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

// WithMarkAreaNameTypeItemOpts sets the type of the MarkArea.
func WithMarkAreaNameTypeItemOpts(opt ...opts.MarkAreaNameTypeItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		for _, o := range opt {
			s.MarkAreas.Data = append(s.MarkAreas.Data, o)
		}
	}
}

// WithMarkAreaStyleOpts sets the style of the MarkArea.
func WithMarkAreaStyleOpts(opt opts.MarkAreaStyle) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}

		s.MarkAreas.MarkAreaStyle = opt
	}
}

// WithMarkAreaNameCoordItemOpts sets the coordinates of the MarkLine.
func WithMarkAreaNameCoordItemOpts(opt ...opts.MarkAreaNameCoordItem) SeriesOpts {
	type MANameCoord struct {
		Name      string          `json:"name,omitempty"`
		ItemStyle *opts.ItemStyle `json:"itemStyle"`
		Coord     []interface{}   `json:"coord"`
	}
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		for _, o := range opt {
			s.MarkAreas.Data = append(s.MarkAreas.Data, []MANameCoord{{Name: o.Name, ItemStyle: o.ItemStyle, Coord: o.Coordinate0}, {Coord: o.Coordinate1}})
		}
	}
}

// WithMarkAreaNameXAxisItemOpts sets the X axis of the MarkLine.
func WithMarkAreaNameXAxisItemOpts(opt ...opts.MarkAreaNameXAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		for _, o := range opt {
			s.MarkAreas.Data = append(s.MarkAreas.Data, o)
		}
	}
}

// WithMarkAreaNameYAxisItemOpts sets the Y axis of the MarkLine.
func WithMarkAreaNameYAxisItemOpts(opt ...opts.MarkAreaNameYAxisItem) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		for _, o := range opt {
			s.MarkAreas.Data = append(s.MarkAreas.Data, o)
		}
	}
}

// WithMarkPointNameTypeItemOpts sets the type of the MarkPoint.
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

// WithMarkPointStyleOpts sets the style of the MarkPoint.
func WithMarkPointStyleOpts(opt opts.MarkPointStyle) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkPoints == nil {
			s.MarkPoints = &opts.MarkPoints{}
		}

		s.MarkPoints.MarkPointStyle = opt
	}
}

// WithMarkPointNameCoordItemOpts sets the coordinated of the MarkPoint.
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

func (s *SingleSeries) ConfigureSeriesOpts(options ...SeriesOpts) {
	for _, opt := range options {
		opt(s)
	}
}

// MultiSeries represents multiple series.
type MultiSeries []SingleSeries

// SetSeriesOptions sets options for all the series.
// Previous options will be overwrote every time hence setting them on the `AddSeries` if you want
// to customize each series individually
//
//	here -> ↓ <-
//
// func (c *Bar) AddSeries(name string, data []opts.BarData, options ...SeriesOpts)
func (ms *MultiSeries) SetSeriesOptions(opts ...SeriesOpts) {
	s := *ms
	for i := 0; i < len(s); i++ {
		s[i].ConfigureSeriesOpts(opts...)
	}
}

// WithEncodeOpts Set encodes for dataSets
func WithEncodeOpts(opt opts.Encode) SeriesOpts {
	return func(s *SingleSeries) {
		s.Encode = &opt
	}
}
