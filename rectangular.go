package goecharts

import (
	"bytes"
	"io"
)

type XAxisOptions struct {
	Name string      `json:"name,omitempty"`
	Show bool        `json:"show,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (opt *XAxisOptions) SetDefault() {
	err := SetDefaultValue(opt);
	checkError(err)
}

type YAxisOptions struct {
	Name string      `json:"name,omitempty"`
	Show bool        `json:"show,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (opt *YAxisOptions) SetDefault() {
	err := SetDefaultValue(opt);
	checkError(err)
}

type XYOptions struct {
	XAxisOptions
	YAxisOptions
}

type RectOptions struct {
	BaseOptions
	XYOptions
}

func (rect *RectOptions) setRectGlobalConfig(options ...interface{}) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case InitOptions:
			rect.InitOptions = option.(InitOptions)
		case TitleOptions:
			rect.TitleOptions = option.(TitleOptions)
		case LegendOptions:
			rect.LegendOptions = option.(LegendOptions)
		case XAxisOptions:
			rect.XAxisOptions = option.(XAxisOptions)
		case YAxisOptions:
			rect.YAxisOptions = option.(YAxisOptions)
		}
	}
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

func (rc *RectChart) Render(w ...io.Writer) {
	rc.XAxisOptions.Data = rc.xAxisData
	rc.InitOptions.SetDefault()
	rc.InitOptions.ValidateID()

	var b bytes.Buffer
	renderChart(rc, &b)
	res := replaceRender(b)
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
}

type marketLineData struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type MarketLine struct {
	Data       []marketLineData `json:"data"`
	Symbol     string           `json:"symbol,omitempty"`
	SymbolSize string           `json:"symbolSize,omitempty"`
}
