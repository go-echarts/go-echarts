package charts

import (
	"io"
)

// XY 轴配置项
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

// 扩展新增 X 轴
func (xy *XYAxis) ExtendXAxis(xAxis ...XAxisOpts) {
	for i := 0; i < len(xAxis); i++ {
		xy.XAxisOptsList = append(xy.XAxisOptsList, xAxis[i])
	}
}

// 扩展新增 Y 轴
func (xy *XYAxis) ExtendYAxis(yAxis ...YAxisOpts) {
	for i := 0; i < len(yAxis); i++ {
		xy.YAxisOptsList = append(xy.YAxisOptsList, yAxis[i])
	}
}

// 直角坐标系配置项
type RectOpts struct {
	BaseOpts
	XYAxis
}

// 设置 RectOptions 全局配置项
func (rect *RectOpts) setRectGlobalOptions(options ...globalOptser) {
	rect.BaseOpts.setBaseGlobalOptions(options...)
	rect.XYAxis.setXYGlobalOptions(options...)
}

// 直角坐标系图表
type RectChart struct {
	RectOpts
	Series

	xAxisData interface{}
}

// RectChart 设置全局配置项
func (rc *RectChart) SetGlobalOptions(options ...globalOptser) *RectChart {
	rc.RectOpts.setRectGlobalOptions(options...)
	return rc
}

// 结合不同类型图表叠加画在同张图上
// 只适用于 RectChart 图表，RectChart 图表包括 Bar/Line/Scatter/EffectScatter/Kline/HeatMap
// 将 RectChart 图表的 Series 追加到调用者的 Series 里面，Series 是完全独立的
// 而全局配置使用的是调用者的配置项
func (rc *RectChart) Overlap(a ...serieser) {
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

// RectChart 渲染图表
func (rc *RectChart) Render(w ...io.Writer) error {
	rc.validateOpts()
	return renderToWriter(rc, "chart", []string{}, w...)
}
