package templates

import _ "embed"

// BaseTpl Should check the BaseActions field before call JSONNotEscapedAction since BaseActions only exist in RectCharts
//
//go:embed base.tpl
var BaseTpl string

//go:embed chart.tpl
var ChartTpl string

//go:embed header.tpl
var HeaderTpl string

//go:embed page.tpl
var PageTpl string
