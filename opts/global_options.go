package opts

import (
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/go-echarts/go-echarts/types"
)

// Initialization contains options for the canvas.
type Initialization struct {
	// HTML title
	PageTitle string `default:"Awesome go-echarts"`
	// Canvas width
	Width string `default:"900px"`
	// Canvas Height
	Height string `default:"500px"`
	// Canvas Background Color
	BackgroundColor string `json:"backgroundColor,omitempty"`
	// Chart unique ID
	ChartID string
	// Assets host
	AssetsHost string `default:"https://go-echarts.github.io/go-echarts-assets/assets/"`
	// Chart Theme
	Theme string `default:"white"`
}

// 验证初始化参数，确保图形能够得到正确渲染
func (opt *Initialization) Validate() {
	setDefaultValue(opt)
	if opt.ChartID == "" {
		opt.ChartID = genChartID()
	}
}

// 为结构体设置默认值
// 部分代码参考 https://github.com/mcuadros/go-defaults
func setDefaultValue(ptr interface{}) {
	elem := reflect.ValueOf(ptr).Elem()
	t := elem.Type()

	for i := 0; i < t.NumField(); i++ {
		// handle `default` tag only
		if defaultVal := t.Field(i).Tag.Get("default"); defaultVal != "" {
			setField(elem.Field(i), defaultVal)
		}
	}
}

func setField(field reflect.Value, defaultVal string) {
	switch field.Kind() {
	case reflect.String:
		if field.String() == "" {
			field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
		}
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	}
}

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	chartIDSize   = 12
)

var seed = rand.NewSource(time.Now().UnixNano())

// 生成唯一且随机的图表 ID
func genChartID() string {
	b := make([]byte, chartIDSize)
	for i, cache, remain := chartIDSize-1, seed.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = seed.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// Title is the option set for a title component.
type Title struct {
	// 主标题
	Title string `json:"text,omitempty"`
	// 主标题样式配置项
	TitleStyle TextStyle `json:"textStyle,omitempty"`
	// 副标题
	Subtitle string `json:"subtext,omitempty"`
	// 副标题样式配置项
	SubtitleStyle TextStyle `json:"subtextStyle,omitempty"`
	// 主标题文本超链接
	Link string `json:"link,omitempty"`
	// 指定窗口打开主标题超链接
	// 'self' 当前窗口打开
	// 'blank' 新窗口打开
	Target string `json:"target,omitempty"`
	// grid 组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐
	Top string `json:"top,omitempty"`
	// 图例组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
	// 图例组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// 图例组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
}

// Legend is the option set for a legend component.
type Legend struct {
	// 图例组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// 图例组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐。
	Top string `json:"top,omitempty"`
	// 图例组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
	// 图例组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Bottom string `json:"bottom,omitempty"`
	// Legend 数据项
	// 如果需要隐藏 Legend 则把 Data 设置为 []string{}
	Data interface{} `json:"data,omitempty"`
	// 除此之外也可以设成 "single" 或者 "multiple" 使用单选或者多选模式。默认 "multiple"
	SelectedMode string `json:"selectedMode,omitempty"`
	// 图例的公用文本样式
	TextStyle TextStyle `json:"textStyle,omitempty"`
}

// Tooltip is the option set for a tooltip component.
type Tooltip struct {
	// 是否显示提示框
	Show bool `json:"show,omitempty"`
	// 触发类型。
	// "item": 数据项图形触发，主要在散点图，饼图等无类目轴的图表中使用。
	// "axis": 坐标轴触发，主要在柱状图，折线图等会使用类目轴的图表中使用。
	// "none": 什么都不触发。
	Trigger string `json:"trigger,omitempty"`
	// 提示框触发的条件，可选：
	// "mousemove": 鼠标移动时触发。
	// "click": 鼠标点击时触发。
	// "mousemove|click": 同时鼠标移动和点击时触发。
	// "none": 不在 "mousemove" 或 "click" 时触发
	TriggerOn string `json:"triggerOn,omitempty"`
	// 1, 字符串模板
	// 模板变量有 {a}, {b}，{c}，{d}，{e}，分别表示系列名，数据名，数据值等。
	// 在 trigger 为 'axis' 的时候，会有多个系列的数据，此时可以通过 {a0}, {a1}, {a2}
	// 这种后面加索引的方式表示系列的索引。 不同图表类型下的 {a}，{b}，{c}，{d} 含义不一样。
	// 其中变量{a}, {b}, {c}, {d} 在不同图表类型下代表数据含义为：
	// 折线（区域）图、柱状（条形）图、K 线图 : {a}（系列名称），{b}（类目值），{c}（数值）, {d}（无）
	// 散点图（气泡）图 : {a}（系列名称），{b}（数据名称），{c}（数值数组）, {d}（无）
	// 地图 : {a}（系列名称），{b}（区域名称），{c}（合并数值）, {d}（无）
	// 饼图、仪表盘、漏斗图: {a}（系列名称），{b}（数据项名称），{c}（数值）, {d}（百分比）
	//
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

// Toolbox is the option set for a toolbox component.
type Toolbox struct {
	// 是否显示工具栏组件
	Show bool `json:"show,omitempty"`
	// 工具箱功能种类，不支持自定义
	ToolBoxFeature `json:"feature,omitempty"`
}

// ToolBoxFeature is a feature component under toolbox.
type ToolBoxFeature struct {
	// 保存为图片
	SaveAsImage *ToolBoxFeatureSaveAsImage `json:"saveAsImage"`
	// 数据区域缩放。目前只支持直角坐标系的缩放
	DataZoom *ToolBoxFeatureDataZoom `json:"dataZoom"`
	// 数据视图工具，可以展现当前图表所用的数据，编辑后可以动态更新
	DataView *ToolBoxFeatureDataView `json:"dataView"`
	// 配置项还原
	Restore *ToolBoxFeatureRestore `json:"restore"`
}

type ToolBoxFeatureSaveAsImage struct {
}

type ToolBoxFeatureDataZoom struct {
}

type ToolBoxFeatureDataView struct {
}

type ToolBoxFeatureRestore struct {
}

// XAxis is the option set for X axis.
type XAxis struct {
	// X 轴名称
	Name string `json:"name,omitempty"`
	// X 坐标轴类型，可选：
	// "value"：数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "time" 时间轴，适用于连续的时序数据，与数值轴相比时间轴带有时间的格式化，
	// 在刻度计算上也有所不同，例如会根据跨度的范围来决定使用月，星期，日还是小时范围的刻度。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// 是否显示 X 轴
	Show bool `json:"show,omitempty"`
	// X 轴数据项
	Data interface{} `json:"data,omitempty"`
	// X 坐标轴的分割段数，需要注意的是这个分割段数只是个预估值，
	// 最后实际显示的段数会在这个基础上根据分割后坐标轴刻度显示的易读程度作调整。
	// 在类目轴中无效
	SplitNumber int `json:"splitNumber,omitempty"`
	// 只在数值轴中（type: 'value'）有效。
	// 是否是脱离 0 值比例。设置成 true 后坐标刻度不会强制包含零刻度。在双数值轴的散点图中比较有用。
	// 在设置 min 和 max 之后该配置项无效
	// 默认为 false
	Scale bool `json:"scale,omitempty"`
	// X 坐标轴刻度最小值
	// 可以设置成特殊值 'dataMin'，此时取数据在该轴上的最小值作为最小刻度，数值轴有效
	Min interface{} `json:"min,omitempty"`
	// X 坐标轴刻度最大值
	// 可以设置成特殊值 'dataMax'，此时取数据在该轴上的最小值作为最小刻度，数值轴有效
	Max interface{} `json:"max,omitempty"`
	// X 轴所在的 grid 的索引
	// 默认 0
	GridIndex int `json:"gridIndex,omitempty"`
	// X 轴在 grid 区域中的分隔区域配置项
	*SplitArea `json:"splitArea,omitempty"`
	// X 轴在 grid 区域中的分隔线配置项
	*SplitLine `json:"splitLine,,omitempty"`
}

// YAxis is the option set for Y axis.
type YAxis struct {
	// Y 轴名称
	Name string `json:"name,omitempty"`
	// Y 坐标轴类型，可选：
	// "value"：数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "time" 时间轴，适用于连续的时序数据，与数值轴相比时间轴带有时间的格式化，
	// 在刻度计算上也有所不同，例如会根据跨度的范围来决定使用月，星期，日还是小时范围的刻度。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// 是否显示 Y 轴
	Show bool `json:"show,omitempty"`
	// 刻度标签的内容格式器，支持字符串模板和回调函数两种形式
	// 1.使用字符串模板，模板变量为刻度默认标签 {value}
	// formatter: '{value} kg'
	// 2.使用函数模板，函数参数分别为刻度数值（类目），刻度的索引
	// formatter: function (value, index) {
	//    // 格式化成月/日，只在第一个刻度显示年份
	//    var date = new Date(value);
	//    var texts = [(date.getMonth() + 1), date.getDate()];
	//    if (index === 0) {
	//        texts.unshift(date.getYear());
	//    }
	//    return texts.join('/');
	// }
	AxisLabel Label `json:"axisLabel,omitempty"`
	// Y 轴数据项
	Data interface{} `json:"data,omitempty"`
	// Y 坐标轴的分割段数，需要注意的是这个分割段数只是个预估值，
	// 最后实际显示的段数会在这个基础上根据分割后坐标轴刻度显示的易读程度作调整。
	// 在类目轴中无效
	SplitNumber int `json:"splitNumber,omitempty"`
	// 只在数值轴中（type: 'value'）有效。
	// 是否是脱离 0 值比例。设置成 true 后坐标刻度不会强制包含零刻度。在双数值轴的散点图中比较有用。
	// 在设置 min 和 max 之后该配置项无效
	// 默认为 false
	Scale bool `json:"scale,omitempty"`
	// Y 坐标轴刻度最小值
	// 可以设置成特殊值 'dataMin'，此时取数据在该轴上的最小值作为最小刻度，数值轴有效
	Min interface{} `json:"min,omitempty"`
	// Y 坐标轴刻度最大值
	// 可以设置成特殊值 'dataMax'，此时取数据在该轴上的最小值作为最小刻度，数值轴有效
	Max interface{} `json:"max,omitempty"`
	// Y 轴所在的 grid 的索引
	// 默认 0
	GridIndex int `json:"gridIndex,omitempty"`
	// Y 轴在 grid 区域中的分隔区域配置项
	*SplitArea `json:"splitArea,omitempty"`
	// Y 轴在 grid 区域中的分隔线配置项
	*SplitLine `json:"splitLine,,omitempty"`
}

// TextStyle is the option set for a text style component.
type TextStyle struct {
	// 文字字体颜色
	Color string `json:"color,omitempty"`
	// 文字字体的风格
	// 可选  'normal', 'italic', 'oblique'
	FontStyle string `json:"fontStyle,omitempty"`
	// 字体大小
	FontSize int `json:"fontSize,omitempty"`
	// 递归结构，为了兼容 wordCloud
	Normal *TextStyle `json:"normal,omitempty"`
}

// SplitArea is the option set for a split area.
type SplitArea struct {
	// 是否显示分隔区域
	Show bool `json:"show"`
	// 风格区域风格
	*AreaStyle `json:"areaStyle,omitempty"`
}

// SplitLine is the option set for a split line.
type SplitLine struct {
	// 是否显示分隔线
	Show bool `json:"show"`
	// 分割线风格
	*LineStyle `json:"lineStyle,omitempty"`
}

// VisualMap is the option set for a visual map component.
// 用于进行『视觉编码』，也就是将数据映射到视觉元素（视觉通道）
type VisualMap struct {
	// 映射类型，可选 "continuous", "piecewise"
	Type string `json:"type,omitempty" default:"continuous"`
	// 是否显示拖拽用的手柄（手柄能拖拽调整选中范围）
	Calculable bool `json:"calculable"`
	// VisualMap 组件的允许的最小值
	Min float32 `json:"min,omitempty"`
	// VisualMap 组件的允许的最大值
	Max float32 `json:"max,omitempty"`
	// 指定手柄对应数值的位置。range 应在 min max 范围内
	Range []float32 `json:"range,omitempty"`
	// 两端的文本，如 ['High', 'Low']
	Text []string `json:"text,omitempty"`
	// 定义在选中范围中的视觉元素
	InRange VisualMapInRange `json:"inRange,omitempty"`
}

// VisualMapInRange is a visual map instance in a range.
type VisualMapInRange struct {
	// 图元的颜色
	Color []string `json:"color,omitempty"`
	// 图元的图形类别
	// 可选 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	Symbol string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
}

// DataZoom is the option set for a zoom component.
type DataZoom struct {
	// 缩放类型，可选 "inside", "slider"
	Type string `json:"type" default:"inside"`
	// 数据窗口范围的起始百分比。范围是：0 ~ 100。表示 0% ~ 100%。
	// 默认: 0
	Start float32 `json:"start,omitempty"`
	// 数据窗口范围的结束百分比。范围是：0 ~ 100。
	// 默认: 100
	End float32 `json:"end,omitempty"`
	// 触发视图刷新的频率。单位为毫秒（ms）。
	// 默认: 100
	Throttle float32 `json:"throttle,omitempty"`
	// 设置 dataZoom 组件控制的 X 轴
	// 不指定时，当 dataZoom-inside.orient 为 'horizontal'时，默认控制和 dataZoom 平行的第一个 xAxis。
	// 但是不建议使用默认值，建议显式指定。
	// 如果是 number 表示控制一个轴，如果是 Array 表示控制多个轴。
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`
	// 设置 dataZoom 组件控制的 Y 轴
	// 不指定时，当 dataZoom-slider.orient 为 'vertical'时，默认控制和 dataZoom 平行的第一个 yAxis。
	// 但是不建议使用默认值，建议显式指定。
	// 如果是 number 表示控制一个轴，如果是 Array 表示控制多个轴。
	YAxisIndex interface{} `json:"yAxisIndex,omitempty"`
}

// SingleAxis is the option set for single axis.
type SingleAxis struct {
	// 坐标轴刻度最小值。
	// 可以设置成特殊值 "dataMin"，此时取数据在该轴上的最小值作为最小刻度
	Min interface{} `json:"min,omitempty"`
	// 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度
	Max interface{} `json:"max,omitempty"`
	// 坐标轴类型
	// "value" 数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "time" 时间轴，适用于连续的时序数据，与数值轴相比时间轴带有时间的格式化，
	// 在刻度计算上也有所不同，例如会根据跨度的范围来决定使用月，星期，日还是小时范围的刻度。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// single 组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 "20%" 这样相对于容器高宽的百分比，
	// 也可以是 "left", "center", "right"。
	// 如果 left 的值为 "left", "center", "right"，组件会根据相应的位置自动对齐
	Left string `json:"left,omitempty"`
	// single 组件离容器右侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 "20%" 这样相对于容器高宽的百分比，
	// 也可以是 "left", "center", "right"。
	// 如果 left 的值为 "left", "center", "right"，组件会根据相应的位置自动对齐
	Right string `json:"right,omitempty"`
	// single 组件离容器顶侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 "20%" 这样相对于容器高宽的百分比，
	// 也可以是 "left", "center", "right"。
	// 如果 left 的值为 "left", "center", "right"，组件会根据相应的位置自动对齐
	Top string `json:"top,omitempty"`
	// single 组件离容器底侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 "20%" 这样相对于容器高宽的百分比，
	// 也可以是 "left", "center", "right"。
	// 如果 left 的值为 "left", "center", "right"，组件会根据相应的位置自动对齐
	Bottom string `json:"bottom,omitempty"`
}

// Indicator is the option set for a radar chart.
type Indicator struct {
	// 指示器名称
	Name string `json:"name,omitempty"`
	// 指示器的最大值，可选，建议设置
	Max float32 `json:"max,omitempty"`
	// 指示器的最小值，可选，默认为 0
	Min float32 `json:"min,omitempty"`
	// 标签特定的颜色
	Color string `json:"color,omitempty"`
}

// RadarComponent is the option set for a radar component.
type RadarComponent struct {
	// 雷达图的指示器，用来指定雷达图中的多个变量（维度）
	Indicator []Indicator `json:"indicator,omitempty"`
	// 雷达图绘制类型，支持 "polygon" 和 "circle"
	Shape string `json:"shape,omitempty"`
	// 指示器轴的分割段数。默认 5
	SplitNumber int `json:"splitNumber,omitempty"`
	// 坐标轴在 grid 区域中的分隔区域
	SplitArea SplitArea `json:"splitArea,omitempty"`
	// 坐标轴在 grid 区域中的分隔线
	SplitLine SplitLine `json:"splitLine,omitempty"`
}

// ParallelComponent is the option set for parallel component.
type ParallelComponent struct {
	// parallel 组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// parallel 组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐。
	Top string `json:"top,omitempty"`
	// parallel 组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
	// parallel 组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应
	Bottom string `json:"bottom,omitempty"`
}

// PAOpts is the option set for a parallel axis.
type ParallelAxis struct {
	// 坐标轴的维度序号
	Dim int `json:"dim,omitempty"`
	// 坐标轴名称
	Name string `json:"name,omitempty"`
	// 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度。
	// 不设置时会自动计算最大值保证坐标轴刻度的均匀分布
	Max interface{} `json:"max,omitempty"`
	// 是否是反向坐标轴
	Inverse bool `json:"inverse,omitempty"`
	// 坐标轴名称显示位置，可选 "start", "middle", "center", "end"
	// 默认 "end"
	NameLocation string `json:"nameLocation,omitempty"`
	// 坐标轴类型，可选：
	// "value"：数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// 类目数据，在类目轴（type: "category"）中有效
	Data interface{} `json:"data,omitempty"`
}

type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, replaceJsFuncs(fn[i]))
	}
}

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	pat, _ := regexp.Compile(`\n|\t`)
	fn = pat.ReplaceAllString(fn, "")
	return "__x__" + fn + "__x__"
}

type Colors []string

// AssetsOpts contains options for static assets.
type Assets struct {
	JSAssets  types.OrderedSet
	CSSAssets types.OrderedSet
}

// 初始化静态资源配置项
func (opt *Assets) InitAssets() {
	opt.JSAssets.Init("echarts.min.js")
	opt.CSSAssets.Init("bulma.min.css")
}

// 校验静态资源配置项，追加 host
func (opt *Assets) Validate(host string) {
	for i := 0; i < len(opt.JSAssets.Values); i++ {
		opt.JSAssets.Values[i] = host + opt.JSAssets.Values[i]
	}
	for j := 0; j < len(opt.CSSAssets.Values); j++ {
		opt.CSSAssets.Values[j] = host + opt.CSSAssets.Values[j]
	}
}

// Router contains information for routing.
type Router struct {
	URL  string // 路由 URL
	Text string // 路由显示文字
}

// XAxis3DOpts contains options for X axis in the 3D coordinate.
type XAxis3D struct {
	// 是否显示 3D X 轴
	Show bool `json:"show,omitempty"`
	// X 坐标轴名称
	Name bool `json:"name,omitempty"`
	// X 坐标轴使用的 grid3D 组件的索引。默认使用第一个 grid3D 组件
	Grid3DIndex int `json:"grid3DIndex,omitempty"`
	// X 坐标轴类型，可选：
	// "value" 数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// X 坐标轴刻度最小值。
	// 可以设置成特殊值 "dataMin"，此时取数据在该轴上的最小值作为最小刻度。
	// 不设置时会自动计算最小值保证坐标轴刻度的均匀分布
	Min interface{} `json:"min,omitempty"`
	// X 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度。
	// 不设置时会自动计算最大值保证坐标轴刻度的均匀分布
	Max interface{} `json:"max,omitempty"`
	// 类目数据，在类目轴（type: 'category'）中有效
	Data interface{} `json:"data,omitempty"`
}

// YAxis3D contains options for Y axis in the 3D coordinate.
type YAxis3D struct {
	// 是否显示 3D Y 轴
	Show bool `json:"show,omitempty"`
	// Y 坐标轴名称
	Name bool `json:"name,omitempty"`
	// Y 坐标轴使用的 grid3D 组件的索引。默认使用第一个 grid3D 组件
	Grid3DIndex int `json:"grid3DIndex,omitempty"`
	// Y 坐标轴类型，可选：
	// "value" 数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// Y 坐标轴刻度最小值。
	// 可以设置成特殊值 "dataMin"，此时取数据在该轴上的最小值作为最小刻度。
	// 不设置时会自动计算最小值保证坐标轴刻度的均匀分布
	Min interface{} `json:"min,omitempty"`
	// Y 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度。
	// 不设置时会自动计算最大值保证坐标轴刻度的均匀分布
	Max interface{} `json:"max,omitempty"`
	// 类目数据，在类目轴（type: 'category'）中有效
	Data interface{} `json:"data,omitempty"`
}

// ZAxis3D contains options for Z axis in the 3D coordinate.
type ZAxis3D struct {
	// 是否显示 3D Z 轴
	Show bool `json:"show,omitempty"`
	// Z 坐标轴名称
	Name bool `json:"name,omitempty"`
	// Z 坐标轴使用的 grid3D 组件的索引。默认使用第一个 grid3D 组件
	Grid3DIndex int `json:"grid3DIndex,omitempty"`
	// Z 坐标轴类型，可选：
	// "value" 数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// Z 坐标轴刻度最小值。
	// 可以设置成特殊值 "dataMin"，此时取数据在该轴上的最小值作为最小刻度。
	// 不设置时会自动计算最小值保证坐标轴刻度的均匀分布
	Min interface{} `json:"min,omitempty"`
	// Z 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度。
	// 不设置时会自动计算最大值保证坐标轴刻度的均匀分布
	Max interface{} `json:"max,omitempty"`
	// 类目数据，在类目轴（type: 'category'）中有效
	Data interface{} `json:"data,omitempty"`
}

// Grid3D contains options for the 3D coordinate.
type Grid3D struct {
	// 是否显示三维笛卡尔坐标系
	Show bool `json:"show,omitempty"`
	// 三维笛卡尔坐标系组件在三维场景中的宽度
	// 默认 100
	BoxWidth float32 `json:"boxWidth,omitempty"`
	// 三维笛卡尔坐标系组件在三维场景中的高度
	// 默认 100
	BoxHeight float32 `json:"boxHeight,omitempty"`
	// 三维笛卡尔坐标系组件在三维场景中的深度
	// 默认 100
	BoxDepth float32 `json:"boxDepth,omitempty"`
	// 用于鼠标的旋转，缩放等视角控制
	ViewControl ViewControl `json:"viewControl,omitempty"`
}

// ViewControl contains options for view controlling.
type ViewControl struct {
	// 是否开启视角绕物体的自动旋转查看
	AutoRotate bool `json:"autoRotate,omitempty"`
	// 物体自转的速度。单位为角度 / 秒，默认为 10 ，也就是 36 秒转一圈
	AutoRotateSpeed float32 `json:"autoRotateSpeed,omitempty"`
}

// FuncOpts is the option set for handling function type.
func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
}
