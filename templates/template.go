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

//go:embed base_element.tpl
var BaseElementTpl string

//go:embed base_script.tpl
var BaseScriptTpl string

//go:embed base_option.tpl
var BaseOptionTpl string
