package goecharts

import (
	"io"
)

// XY 轴配置项
type XYOpts struct {
	XAxisOptsList []XAxisOpts
	YAxisOptsList []YAxisOpts
}

func (opt *XYOpts) initXYOpts() {
	opt.XAxisOptsList = append(opt.XAxisOptsList, XAxisOpts{})
	opt.YAxisOptsList = append(opt.YAxisOptsList, YAxisOpts{})
}

// 设置 XYOptions 全局配置项
func (opt *XYOpts) setXYGlobalConfig(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case XAxisOpts:
			opt.XAxisOptsList[0] = option.(XAxisOpts)
		case YAxisOpts:
			opt.YAxisOptsList[0] = option.(YAxisOpts)
		}
	}
}

// 扩展新增 X 轴
func (opt *XYOpts) ExtendXAxis(xAxis ...XAxisOpts) {
	for i := 0; i < len(xAxis); i++ {
		opt.XAxisOptsList = append(opt.XAxisOptsList, xAxis[i])
	}
}

// 扩展新增 Y 轴
func (opt *XYOpts) ExtendYAxis(yAxis ...YAxisOpts) {
	for i := 0; i < len(yAxis); i++ {
		opt.YAxisOptsList = append(opt.YAxisOptsList, yAxis[i])
	}
}

// 直角坐标系配置项
type RectOpts struct {
	BaseOpts
	XYOpts
}

// 设置 RectOptions 全局配置项
func (rect *RectOpts) setRectGlobalConfig(options ...interface{}) {
	rect.BaseOpts.setBaseGlobalConfig(options...)
	rect.XYOpts.setXYGlobalConfig(options...)
}

// 直角坐标系图表
type RectChart struct {
	RectOpts
	Series

	xAxisData interface{}
}

// RectChart 设置全局配置项
func (rc *RectChart) SetGlobalConfig(options ...interface{}) *RectChart {
	rc.RectOpts.setRectGlobalConfig(options...)
	return rc
}

// 结合不同类型图表叠加画在同张图上
// 将 RectChart 图表的 Series 追加到调用者的 Series 里面，Series 是完全独立的
// 而全局配置使用的是调用者的配置项
func (rc *RectChart) Overlap(a ...seriesI) {
	for i := 0; i < len(a); i++ {
		rc.Series = append(rc.Series, a[i].exportSeries()...)
	}
}

// RectChart 校验器
func (rc *RectChart) validateOpts() {
	// 确保 X 轴数据不会因为设置了 XAxisOpts 而被抹除
	rc.XAxisOptsList[0].Data = rc.xAxisData
	rc.validateAssets(rc.AssetsHost)
}

// RectChart 渲染图表
func (rc *RectChart) Render(w ...io.Writer) error {
	rc.validateOpts()
	return renderToWriter(rc, "chart", w...)
}
