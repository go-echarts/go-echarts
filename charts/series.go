package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"github.com/go-echarts/go-echarts/v2/util"
)

type SingleSeries struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Id   string `json:"id,omitempty"`

	// Rectangular charts
	// Line | Bar
	Stack string `json:"stack,omitempty"`
	// Line | Bar | Scatter
	XAxisIndex int `json:"xAxisIndex,omitempty"`
	// Line | Bar | Scatter
	YAxisIndex int `json:"yAxisIndex,omitempty"`

	// Bar
	BarGap string `json:"barGap,omitempty"`
	// Bar
	BarCategoryGap string `json:"barCategoryGap,omitempty"`
	// Bar
	ShowBackground types.Bool `json:"showBackground,omitempty"`
	// Bar
	RoundCap types.Bool `json:"roundCap,omitempty"`

	// Bar3D
	Shading string `json:"shading,omitempty"`

	// Graph
	Links              interface{} `json:"links,omitempty"`
	Layout             string      `json:"layout,omitempty"`
	Force              interface{} `json:"force,omitempty"`
	Categories         interface{} `json:"categories,omitempty"`
	Roam               types.Bool  `json:"roam,omitempty"`
	EdgeSymbol         interface{} `json:"edgeSymbol,omitempty"`
	EdgeSymbolSize     interface{} `json:"edgeSymbolSize,omitempty"`
	EdgeLabel          interface{} `json:"edgeLabel,omitempty"`
	Draggable          types.Bool  `json:"draggable,omitempty"`
	FocusNodeAdjacency types.Bool  `json:"focusNodeAdjacency,omitempty"`
	// Line | Radar
	SymbolKeepAspect types.Bool `json:"symbolKeepAspect,omitempty"`

	// BarWidth The width options of the bar. Adaptive when not specified.
	// Can be an absolute value like 40 or a percent value like '60%'.
	// Configurable charts: bar | kline
	BarWidth    string `json:"barWidth,omitempty"`
	BarMinWidth string `json:"barMinWidth,omitempty"`
	BarMaxWidth string `json:"barMaxWidth,omitempty"`

	// Line | Bar | Pie | Scatter | Radar
	ColorBy string `json:"colorBy,omitempty"`
	// Line | Bar
	PolarIndex int `json:"polarIndex,omitempty"`
	// Line
	Step interface{} `json:"step,omitempty"`
	// Line
	Smooth types.Bool `json:"smooth,omitempty"`
	// Line
	ConnectNulls types.Bool `json:"connectNulls,omitempty"`
	// Line
	ShowSymbol types.Bool `json:"showSymbol,omitempty"`
	// Line | Scatter | Radar
	Symbol string `json:"symbol,omitempty"`
	Color  string `json:"color,omitempty"`

	// Liquid
	IsLiquidOutline types.Bool `json:"outline,omitempty"`
	IsWaveAnimation types.Bool `json:"waveAnimation,omitempty"`

	// Map
	MapType string `json:"map,omitempty"`
	// Map | Line | Bar | Pie | Scatter
	CoordSystem string `json:"coordinateSystem,omitempty"`

	// Pie
	RoseType string `json:"roseType,omitempty"`
	// Pie
	Center interface{} `json:"center,omitempty"`
	// Pie
	Radius interface{} `json:"radius,omitempty"`

	// Line | Scatter | Radar
	SymbolSize interface{} `json:"symbolSize,omitempty"`

	// Tree
	Orient            string      `json:"orient,omitempty"`
	ExpandAndCollapse types.Bool  `json:"expandAndCollapse,omitempty"`
	InitialTreeDepth  int         `json:"initialTreeDepth,omitempty"`
	Leaves            interface{} `json:"leaves,omitempty"`
	Left              string      `json:"left,omitempty"`
	Right             string      `json:"right,omitempty"`
	Top               string      `json:"top,omitempty"`
	Bottom            string      `json:"bottom,omitempty"`

	// Sankey
	NodeWidth types.Int `json:"nodeWidth,omitempty"`
	NodeGap   types.Int `json:"nodeGap,omitempty"`
	NodeAlign string    `json:"nodeAlign,omitempty"`

	// Radar
	RadarIndex int `json:"radarIndex,omitempty"`

	// TreeMap
	LeafDepth  int         `json:"leafDepth,omitempty"`
	Levels     interface{} `json:"levels,omitempty"`
	UpperLabel interface{} `json:"upperLabel,omitempty"`

	// WordCloud
	Shape         string    `json:"shape,omitempty"`
	SizeRange     []float32 `json:"sizeRange,omitempty"`
	RotationRange []float32 `json:"rotationRange,omitempty"`

	// Sunburst
	NodeClick string `json:"nodeClick,omitempty"`
	Sort      string `json:"sort,omitempty"`

	// Custom
	RenderItem types.FuncStr `json:"renderItem,omitempty"`

	// series data
	Data         interface{} `json:"data,omitempty"`
	DatasetIndex int         `json:"datasetIndex,omitempty"`

	// Gauge
	Progress   *opts.Progress  `json:"progress,omitempty"`
	AxisTick   *opts.AxisTick  `json:"axisTick,omitempty"`
	AxisLabel  *opts.AxisLabel `json:"axisLabel,omitempty"`
	AxisLine   *opts.AxisLine  `json:"axisLine,omitempty"`
	Pointer    *opts.Pointer   `json:"pointer,omitempty"`
	SplitLine  *opts.SplitLine `json:"splitLine,omitempty"`
	Detail     *opts.Detail    `json:"detail,omitempty"`
	Title      *opts.Title     `json:"title,omitempty"`
	Min        int             `json:"min,omitempty"`
	Max        int             `json:"max,omitempty"`
	StartAngle float64         `json:"startAngle,omitempty"`
	EndAngle   float64         `json:"endAngle,omitempty"`

	Large               types.Bool `json:"large,omitempty"`
	LargeThreshold      int        `json:"largeThreshold,omitempty"`
	HoverLayerThreshold int        `json:"hoverLayerThreshold,omitempty"`
	UseUTC              types.Bool `json:"useUTC,omitempty"`

	// Animation related configs
	Animation               types.Bool `json:"animation,omitempty"`
	AnimationThreshold      int        `json:"animationThreshold,omitempty"`
	AnimationDuration       int        `json:"animationDuration,omitempty"`
	AnimationEasing         string     `json:"animationEasing,omitempty"`
	AnimationDelay          int        `json:"animationDelay,omitempty"`
	AnimationDurationUpdate int        `json:"animationDurationUpdate,omitempty"`
	AnimationEasingUpdate   string     `json:"animationEasingUpdate,omitempty"`
	AnimationDelayUpdate    int        `json:"animationDelayUpdate,omitempty"`
	RenderLabelForZeroData  types.Bool `json:"renderLabelForZeroData,omitempty"`
	SelectedMode            types.Bool `json:"selectedMode,omitempty"`

	// series options
	*opts.Encode        `json:"encode,omitempty"`
	*opts.ItemStyle     `json:"itemStyle,omitempty"`
	*opts.Label         `json:"label,omitempty"`
	*opts.LabelLayout   `json:"labelLayout,omitempty"`
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
	*opts.SeriesTooltip `json:"tooltip,omitempty"`

	// Calendar
	CalendarIndex int `json:"calendarIndex,omitempty"`
}

type SeriesOpts func(s *SingleSeries)

type SingleSeriesOptFunc func(s *SingleSeries)

// WithSeriesOpts If the WithXXX helper method is not good enough, use this directly!
func WithSeriesOpts(opf SingleSeriesOptFunc) SeriesOpts {
	return func(s *SingleSeries) {
		opf(s)
	}
}

func WithSeriesId(id string) SeriesOpts {
	return func(s *SingleSeries) {
		s.Id = id
	}
}

func WithCoordinateSystem(cs string) SeriesOpts {
	return func(s *SingleSeries) {
		s.CoordSystem = cs
	}
}

func WithCalendarIndex(index int) SeriesOpts {
	return func(s *SingleSeries) {
		s.CalendarIndex = index
	}
}

func WithSeriesAnimation(enable bool) SeriesOpts {
	return func(s *SingleSeries) {
		s.Animation = opts.Bool(enable)
	}
}

func WithSeriesSymbolKeepAspect(enable bool) SeriesOpts {
	return func(s *SingleSeries) {
		s.SymbolKeepAspect = opts.Bool(enable)
	}
}

func WithAnimationOpts(opt opts.Animation) SeriesOpts {
	return func(s *SingleSeries) {
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

// WithLabelOpts sets the label.
func WithLabelOpts(opt opts.Label) SeriesOpts {
	return func(s *SingleSeries) {
		s.Label = &opt
	}
}

// WithLabelLayoutOpts sets the label.
func WithLabelLayoutOpts(opt opts.LabelLayout) SeriesOpts {
	return func(s *SingleSeries) {
		s.LabelLayout = &opt
	}
}

// WithLabelLineOpts sets the label.
func WithLabelLineOpts(opt opts.LabelLine) SeriesOpts {
	return func(s *SingleSeries) {
		s.LabelLine = &opt
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

// WithCircularStyleOpts With CircularStyle Opts
func WithCircularStyleOpts(opt opts.CircularStyle) SeriesOpts {
	return func(s *SingleSeries) {
		s.CircularStyle = &opt
	}
}

// WithSeriesTooltipOpts With Tooltip Opts
func WithSeriesTooltipOpts(opt opts.SeriesTooltip) SeriesOpts {
	return func(s *SingleSeries) {
		s.SeriesTooltip = &opt
	}
}

/* Chart Options */

// WithBarChartOpts sets the BarChart option.
func WithBarChartOpts(opt opts.BarChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.ColorBy = opt.ColorBy
		s.CoordSystem = opt.CoordSystem
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.PolarIndex = opt.PolarIndex
		s.RoundCap = opt.RoundCap
		s.ShowBackground = opt.ShowBackground
		s.Stack = opt.Stack
		s.BarGap = opt.BarGap
		s.BarCategoryGap = opt.BarCategoryGap
		s.BarWidth = opt.BarWidth
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
		s.SymbolKeepAspect = opt.SymbolKeepAspect
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
		s.ColorBy = opt.ColorBy
		s.CoordSystem = opt.CoordSystem
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.PolarIndex = opt.PolarIndex
		s.Symbol = opt.Symbol
		s.SymbolSize = opt.SymbolSize
		s.SymbolKeepAspect = opt.SymbolKeepAspect
		s.ShowSymbol = opt.ShowSymbol
		s.Stack = opt.Stack
		s.Smooth = opt.Smooth
		s.ConnectNulls = opt.ConnectNulls
		s.Step = opt.Step
	}
}

// WithLineChartOpts sets the LineChart option.
func WithKlineChartOpts(opt opts.KlineChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.BarWidth = opt.BarWidth
		s.BarMinWidth = opt.BarMinWidth
		s.BarMaxWidth = opt.BarMaxWidth
	}
}

// WithPieChartOpts sets the PieChart option.
func WithPieChartOpts(opt opts.PieChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.ColorBy = opt.ColorBy
		s.CoordSystem = opt.CoordSystem
		s.RoseType = opt.RoseType
		s.Center = opt.Center
		s.Radius = opt.Radius
	}
}

// WithScatterChartOpts sets the ScatterChart option.
func WithScatterChartOpts(opt opts.ScatterChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.ColorBy = opt.ColorBy
		s.CoordSystem = opt.CoordSystem
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.Symbol = opt.Symbol
		s.SymbolSize = opt.SymbolSize
		s.SymbolKeepAspect = opt.SymbolKeepAspect
	}
}

// WithEffectScatterChartOpts sets the ScatterChart option.
func WithEffectScatterChartOpts(opt opts.EffectScatterChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.ColorBy = opt.ColorBy
		s.CoordSystem = opt.CoordSystem
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.Symbol = opt.Symbol
		s.SymbolSize = opt.SymbolSize
		s.SymbolKeepAspect = opt.SymbolKeepAspect
	}
}

// WithRadarChartOpts sets the RadarChart option.
func WithRadarChartOpts(opt opts.RadarChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.ColorBy = opt.ColorBy
		s.RadarIndex = opt.RadarIndex
		s.Symbol = opt.Symbol
		s.SymbolSize = opt.SymbolSize
		s.SymbolKeepAspect = opt.SymbolKeepAspect
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
		s.SymbolKeepAspect = opt.SymbolKeepAspect
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
			s.MarkLines.Data = append(
				s.MarkLines.Data,
				[]MLNameCoord{{Name: o.Name, Coord: o.Coordinate0}, {Coord: o.Coordinate1}},
			)
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
			s.MarkAreas.Data = append(
				s.MarkAreas.Data,
				[]MANameCoord{
					{Name: o.Name, ItemStyle: o.ItemStyle, Coord: o.Coordinate0},
					{Coord: o.Coordinate1},
				},
			)
		}
	}
}

// WithMarkAreaData0 sets the markArea.data.0
func WithMarkAreaData0(data0 opts.MarkAreaData0) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		s.MarkAreas.Data = append(s.MarkAreas.Data, data0)
	}
}

// WithMarkAreaData1 sets the markArea.data.1
func WithMarkAreaData1(data1 opts.MarkAreaData1) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		s.MarkAreas.Data = append(s.MarkAreas.Data, data1)
	}
}

// WithMarkAreaData sets the markArea.data each item as array
// See https://echarts.apache.org/en/option.html#series-candlestick.markArea.data
func WithMarkAreaData(datas ...[]opts.MarkAreaData) SeriesOpts {
	return func(s *SingleSeries) {
		if s.MarkAreas == nil {
			s.MarkAreas = &opts.MarkAreas{}
		}
		for _, d := range datas {
			s.MarkAreas.Data = append(s.MarkAreas.Data, d)
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

func (s *SingleSeries) InitSeriesDefaultOpts(c BaseConfiguration) {
	util.SetDefaultValue(s)
	// some special inherited options from BaseConfiguration
	s.Animation = c.Animation
}

func (s *SingleSeries) ConfigureSeriesOpts(options ...SeriesOpts) {
	for _, opt := range options {
		opt(s)
	}
}

// MultiSeries represents multiple series.
type MultiSeries []SingleSeries

// SetSeriesOptions sets options for all the series.
// NOTE:
// It should be called after AddSeries, otherwise, the Options is no place to add on.
// Previous options will be overwritten every time hence setting them on the `AddSeries` if you want
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

// WithDatasetIndex sets the datasetIndex option.
func WithDatasetIndex(index int) SeriesOpts {
	return func(s *SingleSeries) {
		s.DatasetIndex = index
	}
}

// WithCustomChartOpts sets the CustomChart option.
func WithCustomChartOpts(opt opts.CustomChart) SeriesOpts {
	return func(s *SingleSeries) {
		s.XAxisIndex = opt.XAxisIndex
		s.YAxisIndex = opt.YAxisIndex
		s.RenderItem = opt.RenderItem
	}
}
