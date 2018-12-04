package geocharts

import "log"

type Series struct {
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	YValue interface{} `json:"data"`
}

type InitOptions struct {
	PageTitle   string `default:"Awesome go-echarts"`
	Width       string `default:"800px"`
	Height      string `default:"500px"`
	ContainerID string `default:"main"`
	Title       string
	SubTitle    string
}

func (opt *InitOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type LegendOptions struct {
	Show   bool   `json:"show" default:"true"`
	Left   string `json:"left,omitempty"`
	Top    string `json:"top,omitempty"`
	Right  string `json:"right,omitempty"`
	Bottom string `json:"bottom,omitempty"`
}

func (opt *LegendOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type LabelOptions struct {
	Show bool `default:"false"`
}

func (opt *LabelOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type XAxisOptions struct {
	XAxisName string
	Show      bool `default:"true"`
}

func (opt *XAxisOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}

type YAxisOptions struct {
	YAxisName string
	Show      bool `default:"true"`
}

func (opt *YAxisOptions) SetDefault() {
	if err := SetDefaultValue(opt); err != nil {
		log.Println(err)
	}
}
