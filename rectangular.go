package goecharts

import (
	"bytes"
	"io"
)

// X 轴配置项
type XAxisOptions struct {
	// X 轴名称
	Name string `json:"name,omitempty"`
	// 是否显示 X 轴
	Show bool `json:"show,omitempty"`
	// X 轴数据项
	Data interface{} `json:"data,omitempty"`
}

// Y 轴配置项
type YAxisOptions struct {
	// Y 轴名称
	Name string `json:"name,omitempty"`
	// 是否显示 Y 轴
	Show bool `json:"show,omitempty"`
	// Y 轴数据项
	Data interface{} `json:"data,omitempty"`
}

// XY 轴配置项
type XYOptions struct {
	XAxisOptions
	YAxisOptions
}

// 设置 XYOptions 全局配置项
func (opt *XYOptions) setXYGlobalConfig(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case XAxisOptions:
			opt.XAxisOptions = option.(XAxisOptions)
		case YAxisOptions:
			opt.YAxisOptions = option.(YAxisOptions)
		}
	}
}

type RectOptions struct {
	BaseOptions
	XYOptions
}

// 设置 RectOptions 全局配置项
func (rect *RectOptions) setRectGlobalConfig(options ...interface{}) {
	rect.BaseOptions.setBaseGlobalConfig(options...)
	rect.XYOptions.setXYGlobalConfig(options...)
}

type RectChart struct {
	RectOptions
	SeriesList
	MarketLine

	HasXYAxis bool
	xAxisData interface{}
}

func (rc *RectChart) SetGlobalConfig(options ...interface{}) *RectChart {
	rc.RectOptions.setRectGlobalConfig(options...)
	return rc
}

func (rc *RectChart) SetSeriesConfig(options ...interface{}) *RectChart {
	rc.SeriesList.setSeriesOptions(options...)
	return rc
}

func (rc *RectChart) verifyOpts() {
	rc.XAxisOptions.Data = rc.xAxisData
	rc.verifyInitOpt()
}

func (rc *RectChart) Render(w ...io.Writer) {
	rc.XAxisOptions.Data = rc.xAxisData
	rc.verifyOpts()

	var b bytes.Buffer
	renderChart(rc, &b, "chart")
	res := replaceRender(b)
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}

//TODO:设计 markLine&markPoint 接口
type marketLineData struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type MarketLine struct {
	Data       []marketLineData `json:"data"`
	Symbol     string           `json:"symbol,omitempty"`
	SymbolSize string           `json:"symbolSize,omitempty"`
}
