# Series Options

`go-echarts` puts all the series options simply in one place.
And provide the `WithXXXChartsOpts` (e.g. `WithBarChartOpts`) to config the specific charts' options
easily.

Additionally, If the `WithXXXChartOpts` not meets your needs.
You can use the `WithSeriesOpts` to fully control the options as you want.

The Series Options list:
```go
type SingleSeries struct {
    Name string `json:"name,omitempty"`
    Type string `json:"type,omitempty"`

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
    Progress *opts.Progress `json:"progress,omitempty"`
    AxisTick *opts.AxisTick `json:"axisTick,omitempty"`
    Detail   *opts.Detail   `json:"detail,omitempty"`
    Title    *opts.Title    `json:"title,omitempty"`

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

    // Calendar
    CalendarIndex int `json:"calendarIndex,omitempty"`

}


// WithSeriesOpts If the WithXXX helper method is not good enough, use this directly!
func WithSeriesOpts(opf SingleSeriesOptFunc) SeriesOpts {
    return func (s *SingleSeries) {
    opf(s)
 }

```