package templates

import _ "embed"

//go:embed base.tpl
var BaseTpl string

//go:embed chart.tpl
var ChartTpl string

//go:embed header.tpl
var HeaderTpl string

//go:embed page.tpl
var PageTpl string
