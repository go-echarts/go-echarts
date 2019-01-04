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

	xAxisData interface{}
}

// RectChart 设置全局配置项
func (rc *RectChart) SetGlobalConfig(options ...interface{}) *RectChart {
	rc.RectOpts.setRectGlobalConfig(options...)
	return rc
}

// RectChart 校验器
func (rc *RectChart) validateOpts() {
	rc.XAxisOpts.Data = rc.xAxisData
	rc.validateInitOpt()
	rc.validateAssets(rc.AssetsHost)
}

// RectChart 渲染图表
func (rc *RectChart) Render(w ...io.Writer) error {
	rc.XAxisOpts.Data = rc.xAxisData
	rc.validateOpts()
	return renderToWriter(rc, "chart", w...)
}
