package charts

import (
	"io"
)

type rectCharter interface {
	markRectChart()
	exportSeries() Series
}

// XYAxis represent the X and Y axis in the rectangular coordinate.
type XYAxis struct {
	XAxisOptsList []XAxisOpts
	YAxisOptsList []YAxisOpts
}

func (xy *XYAxis) initXYOpts() {
	xy.XAxisOptsList = append(xy.XAxisOptsList, XAxisOpts{})
	xy.YAxisOptsList = append(xy.YAxisOptsList, YAxisOpts{})
}

// 设置 XYOptions 全局配置项
func (xy *XYAxis) setXYGlobalOptions(options ...globalOptser) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case XAxisOpts:
			xy.XAxisOptsList[0] = option.(XAxisOpts)
		case YAxisOpts:
			xy.YAxisOptsList[0] = option.(YAxisOpts)
		}
	}
}

// ExtendXAxis adds new X axes.
func (xy *XYAxis) ExtendXAxis(xAxis ...XAxisOpts) {
	for i := 0; i < len(xAxis); i++ {
		xy.XAxisOptsList = append(xy.XAxisOptsList, xAxis[i])
	}
}

// ExtendYAxis adds new Y axes.
func (xy *XYAxis) ExtendYAxis(yAxis ...YAxisOpts) {
	for i := 0; i < len(yAxis); i++ {
		xy.YAxisOptsList = append(xy.YAxisOptsList, yAxis[i])
	}
}

// RectOpts contains options for the rectangular coordinate.
type RectOpts struct {
	BaseOpts
	XYAxis
}

// 设置 RectOptions 全局配置项
func (rect *RectOpts) setRectGlobalOptions(options ...globalOptser) {
	rect.BaseOpts.setBaseGlobalOptions(options...)
	rect.XYAxis.setXYGlobalOptions(options...)
}

// RectChart is a chart in RectChart coordinate.
type RectChart struct {
	RectOpts
	Series

	xAxisData interface{}
}

func (RectChart) markRectChart() {}

// SetGlobalOptions sets options for the RectChart instance.
func (rc *RectChart) SetGlobalOptions(options ...globalOptser) *RectChart {
	rc.RectOpts.setRectGlobalOptions(options...)
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
		rc.Series = append(rc.Series, a[i].exportSeries()...)
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
