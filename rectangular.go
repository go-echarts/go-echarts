package goecharts

import (
	"bytes"
	"io"
)

// X 轴配置项
type XAxisOpts struct {
	// X 轴名称
	Name string `json:"name,omitempty"`
	// 是否显示 X 轴
	Show bool `json:"show,omitempty"`
	// X 轴数据项
	Data interface{} `json:"data,omitempty"`
}

// Y 轴配置项
type YAxisOpts struct {
	// Y 轴名称
	Name string `json:"name,omitempty"`
	// 是否显示 Y 轴
	Show bool `json:"show,omitempty"`
	// Y 轴数据项
	Data interface{} `json:"data,omitempty"`
}

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

type RectOptions struct {
	BaseOpts
	XYOpts
}

// 设置 RectOptions 全局配置项
func (rect *RectOptions) setRectGlobalConfig(options ...interface{}) {
	rect.BaseOpts.setBaseGlobalConfig(options...)
	rect.XYOpts.setXYGlobalConfig(options...)
}

type RectChart struct {
	RectOptions
	Series

	HasXYAxis bool
	xAxisData interface{}
}

func (rc *RectChart) SetGlobalConfig(options ...interface{}) *RectChart {
	rc.RectOptions.setRectGlobalConfig(options...)
	return rc
}

func (rc *RectChart) SetSeriesConfig(options ...interface{}) *RectChart {
	rc.Series.setAllSeriesOpts(options...)
	return rc
}

func (rc *RectChart) verifyOpts() {
	rc.XAxisOpts.Data = rc.xAxisData
	rc.verifyInitOpt()
}

func (rc *RectChart) Render(w ...io.Writer) {
	rc.XAxisOpts.Data = rc.xAxisData
	rc.verifyOpts()

	var b bytes.Buffer
	renderChart(rc, &b, "chart")
	res := replaceRender(b)
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}
