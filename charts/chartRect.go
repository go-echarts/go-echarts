package charts

import (
	"io"

	"github.com/go-echarts/go-echarts/opts"
)

//type rectCharter interface {
//	markRectChart()
//	exportSeries() Series
//}

// XYAxis represent the X and Y axis in the rectangular coordinate.
type XYAxis struct {
	XAxisList []opts.XAxis
	YAxisList []opts.YAxis
}

func WithXAxisOpts(opt opts.XAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
	}
}

func WithYAxisOpts(opt opts.YAxis) GlobalOpts {
	return func(bc *BaseConfiguration) {
	}
}

func (xy *XYAxis) initXYAxis() {
	xy.XAxisList = append(xy.XAxisList, opts.XAxis{})
	xy.YAxisList = append(xy.YAxisList, opts.YAxis{})
}

// 设置 XYOptions 全局配置项
func (xy *XYAxis) setXYGlobalOptions(options ...GlobalOptser) {
	for i := 0; i < len(options); i++ {
		switch option := options[i].(type) {
		case XAxisOpts:
			xy.XAxisOptsList[0] = option
		case YAxisOpts:
			xy.YAxisOptsList[0] = option
		}
	}
}

// ExtendXAxis adds new X axes.
func (xy *XYAxis) ExtendXAxis(xAxis ...opts.XAxis) {
	xy.XAxisList = append(xy.XAxisList, xAxis...)
}

// ExtendYAxis adds new Y axes.
func (xy *XYAxis) ExtendYAxis(yAxis ...opts.YAxis) {
	xy.YAxisList = append(xy.YAxisList, yAxis...)
}

// RectConfiguration contains options for the rectangular coordinate.
type RectConfiguration struct {
	BaseConfiguration
	XYAxis
}

// 设置 RectOptions 全局配置项
func (rect *RectConfiguration) setRectGlobalOptions(opts ...GlobalOpts) {
	rect.BaseConfiguration.setBaseGlobalOptions(opts...)
	rect.XYAxis.setXYGlobalOptions(opts...)
}

// RectChart is a chart in RectChart coordinate.
type RectChart struct {
	RectConfiguration
	MultiSeries

	xAxisData interface{}
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
func (rc *RectChart) Overlap(a ...rectCharter) {
	for i := 0; i < len(a); i++ {
		rc.MultiSeries = append(rc.MultiSeries, a[i].exportSeries()...)
	}
}

// RectChart 校验器
func (rc *RectChart) validateOpts() {
	// 确保 X 轴数据不会因为设置了 XAxisOpts 而被抹除
	rc.XAxisOptsList[0].Data = rc.xAxisData
	// 确保 Y 轴数标签正确显示
	for i := 0; i < len(rc.YAxisOptsList); i++ {
		rc.YAxisOptsList[i].AxisLabel.Show = true
	}
	rc.validateAssets(rc.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (rc *RectChart) Render(w ...io.Writer) error {
	rc.validateOpts()
	return renderToWriter(rc, "chart", []string{}, w...)
}
