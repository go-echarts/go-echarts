package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

type Sankey struct {
	BaseOpts
	Series
}

// 桑基图节点间的关系数据
type SankeyLink struct {
	// 边的源节点名称的字符串，也支持使用数字表示源节点的索引
	Source interface{} `json:"source,omitempty"`
	// 边的目标节点名称的字符串，也支持使用数字表示源节点的索引
	Target interface{} `json:"target,omitempty"`
	// 边的数值，可以在力引导布局中用于映射到边的长度
	Value float32 `json:"value,omitempty"`
}

// 桑基图节点数据
type SankeyNode struct {
	// 数据项名称
	Name string `json:"name,omitempty"`
	// 数据项值
	Value string `json:"value,omitempty"`
}

func (Sankey) chartType() string { return common.ChartType.Sankey }

func NewSankey(routers ...RouterOpts) *Sankey {
	chart := new(Sankey)
	chart.initBaseOpts(routers...)
	return chart
}

func (c *Sankey) Add(name string, nodes []SankeyNode, links []SankeyLink, options ...seriesOptser) *Sankey {
	series := singleSeries{Name: name, Type: common.ChartType.Sankey, Data: nodes, Links: links}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	return c
}

func (c *Sankey) SetGlobalOptions(options ...globalOptser) *Sankey {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Sankey) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Sankey) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
