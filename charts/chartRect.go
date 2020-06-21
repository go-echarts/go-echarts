package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/opts"
)

type Overlaper interface {
	overlap() MultiSeries
}

// XYAxis represent the X and Y axis in the rectangular coordinate.
type XYAxis struct {
	XAxisList []opts.XAxis
	YAxisList []opts.YAxis
}

func (xy *XYAxis) initXYAxis() {
	xy.XAxisList = append(xy.XAxisList, opts.XAxis{})
	xy.YAxisList = append(xy.YAxisList, opts.YAxis{})
}

// ExtendXAxis adds new X axes.
func (xy *XYAxis) ExtendXAxis(xAxis ...opts.XAxis) {
	xy.XAxisList = append(xy.XAxisList, xAxis...)
}

// ExtendYAxis adds new Y axes.
func (xy *XYAxis) ExtendYAxis(yAxis ...opts.YAxis) {
	xy.YAxisList = append(xy.YAxisList, yAxis...)
}

func WithXAxisOpts(opt opts.XAxis, index ...int) GlobalOpts {
	return func(bc *BaseConfiguration) {
		if len(index) == 0 {
			index = []int{0}
		}
		for i := 0; i < len(index); i++ {
			bc.XYAxis.XAxisList[index[i]] = opt
		}
	}
}

func WithYAxisOpts(opt opts.YAxis, index ...int) GlobalOpts {
	return func(bc *BaseConfiguration) {
		if len(index) == 0 {
			index = []int{0}
		}
		for i := 0; i < len(index); i++ {
			bc.XYAxis.YAxisList[index[i]] = opt
		}
	}
}

// RectConfiguration contains options for the rectangular coordinate.
type RectConfiguration struct {
	BaseConfiguration
}

// 设置 RectOptions 全局配置项
func (rect *RectConfiguration) setRectGlobalOptions(opts ...GlobalOpts) {
	rect.BaseConfiguration.setBaseGlobalOptions(opts...)
}

// RectChart is a chart in RectChart coordinate.
type RectChart struct {
	RectConfiguration
	MultiSeries

	xAxisData interface{}
}

// 设置 RectOptions 全局配置项
func (rc *RectChart) overlap() MultiSeries {
	return rc.MultiSeries
}

// SetGlobalOptions sets options for the RectChart instance.
func (rc *RectChart) SetGlobalOptions(opts ...GlobalOpts) *RectChart {
	rc.RectConfiguration.setRectGlobalOptions(opts...)
	return rc
}

// Overlap composites multiple charts into one single canvas.
// 结合不同类型图表叠加画在同张图上
// 只适用于 RectChart 图表，其实现了 rectCharter 接口
// RectChart 图表包括 Bar/BoxPlot/Line/Scatter/EffectScatter/Kline/HeatMap
// 将 RectChart 图表的 Series 追加到调用者的 Series 里面，Series 是完全独立的
// 而全局配置使用的是调用者的配置项
func (rc *RectChart) Overlap(a ...Overlaper) {
	for i := 0; i < len(a); i++ {
		rc.MultiSeries = append(rc.MultiSeries, a[i].overlap()...)
	}
}

// RectChart 校验器
func (rc *RectChart) Validate() {
	// 确保 X 轴数据不会因为设置了 XAxisOpts 而被抹除
	rc.XAxisList[0].Data = rc.xAxisData
	// 确保 Y 轴数标签正确显示
	for i := 0; i < len(rc.YAxisList); i++ {
		rc.YAxisList[i].AxisLabel.Show = true
	}
	rc.Assets.Validate(rc.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (rc *RectChart) Render(w io.Writer) error {
	rc.Validate()
	return renderToWriter(rc, "chart", w)
}
