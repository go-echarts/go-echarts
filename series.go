package geocharts

// 图形上的文本标签配置项
type LabelOptions struct {
	Show     bool   `json:"show" default:"false"`
	Position string `json:"position" default:"top"`
	Color    string `json:"color"`
}

// 为 LabelOptions 设置字段默认值
func (opt *LabelOptions) SetDefault() {
	err := SetDefaultValue(opt);
	checkError(err)
}

// Series 配置项
type Series struct {
	// series 名称
	Name string `json:"name"`
	// series 类型
	Type string `json:"type"`
	// series 数据项
	Data         interface{} `json:"data"`
	LabelOptions `json:"label"`
}

// 设置 Series 配置项
func (series *Series) setSingleSeriesOptions(options ...interface{}){
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case LabelOptions:
			series.LabelOptions = option.(LabelOptions)
		}
	}
}

// Series 列表
type SeriesList []Series

// 设置 SeriesList 配置项
func (sl *SeriesList) setSeriesOptions(options ...interface{}) {
	tsl := *sl
	for i := 0; i < len(tsl); i++ {
		for j := 0; j < len(options); j++ {
			option := options[j]
			switch option.(type) {
			case LabelOptions:
				tsl[i].LabelOptions = option.(LabelOptions)
			}
		}
	}
}
