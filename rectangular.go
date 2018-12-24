package goecharts

import (
	"io"
)

// XY 轴配置项
type XYOpts struct {
	XAxisOpts
	YAxisOpts
}

// 设置 XYOptions 全局配置项
func (opt *XYOpts) setXYGlobalConfig(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case XAxisOpts:
			opt.XAxisOpts = option.(XAxisOpts)
		case YAxisOpts:
			opt.YAxisOpts = option.(YAxisOpts)
		}
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

	HasXYAxis bool
	xAxisData interface{}
}

// RectChart 设置全局配置项
func (rc *RectChart) SetGlobalConfig(options ...interface{}) *RectChart {
	rc.RectOpts.setRectGlobalConfig(options...)
	return rc
}

// RectChart 设置 series 配置项
func (rc *RectChart) SetSeriesConfig(options ...interface{}) *RectChart {
	rc.Series.setAllSeriesOpts(options...)
	return rc
}

// RectChart 校验器
func (rc *RectChart) verifyOpts() {
	rc.XAxisOpts.Data = rc.xAxisData
	rc.verifyInitOpt()
	rc.verifyAssets(rc.AssetsHost)
}

// RectChart 渲染图表
func (rc *RectChart) Render(w ...io.Writer) {
	rc.XAxisOpts.Data = rc.xAxisData
	rc.verifyOpts()
	renderToWriter(rc, "chart", w...)
}
