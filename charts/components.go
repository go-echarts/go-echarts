package charts

// VMInRange is a visual map instance in a range.
type VMInRange struct {
	// 图元的颜色
	Color []string `json:"color,omitempty"`
	// 图元的图形类别
	// 可选 'circle', 'rect', 'roundRect', 'triangle', 'diamond', 'pin', 'arrow', 'none'
	Symbol string `json:"symbol,omitempty"`
	// 图元的大小
	SymbolSize float32 `json:"symbolSize,omitempty"`
}

// SplitAreaOpts is the option set for a split area.
type SplitAreaOpts struct {
	// 是否显示分隔区域
	Show bool `json:"show"`
	// 风格区域风格
	AreaStyle AreaStyleOpts `json:"areaStyle,omitempty"`
}

// SplitLineOpts is the option set for a split line.
type SplitLineOpts struct {
	// 是否显示分隔线
	Show bool `json:"show"`
	// 分割线风格
	LineStyle LineStyleOpts `json:"lineStyle,omitempty"`
}

// FuncOpts is the option set for handling function type.
func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
}
