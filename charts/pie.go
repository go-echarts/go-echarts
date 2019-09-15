package charts

import (
	"github.com/go-echarts/go-echarts/datatypes"
	"io"
)

// Pie represents a pie chart.
type Pie struct {
	BaseOpts
	Series
}

func (Pie) chartType() string { return ChartType.Pie }

// PieOpts is the option set for a pie chart.
type PieOpts struct {
	// 是否展示成南丁格尔图，通过半径区分数据大小。可选择两种模式：
	// 1."radius": 扇区圆心角展现数据的百分比，半径展现数据的大小。
	// 2."area": 所有扇区圆心角相同，仅通过半径展现数据大小。
	RoseType string
	// 饼图的中心（圆心）坐标，数组的第一项是横坐标，第二项是纵坐标。
	// 支持设置成百分比，设置成百分比时第一项是相对于容器宽度，第二项是相对于容器高度
	// 使用示例
	// 设置成绝对的像素值: center: [400, 300]
	// 设置成相对的百分比: center: ['50%', '50%']
	// 默认 ["50%", "50%"]
	Center interface{}
	// 饼图的半径。可以为如下类型：
	// 1.number：直接指定外半径值。
	// 2.string：例如，'20%'，表示外半径为可视区尺寸（容器高宽中较小一项）的 20% 长度。
	// 3.Array.<number|string>：数组的第一项是内半径，第二项是外半径。
	// 每一项遵从上述 number string 的描述。
	// 默认 [0, "75%"]
	Radius interface{}
}

func (PieOpts) markSeries() {}

func (opt *PieOpts) setChartOpt(s *singleSeries) {
	s.RoseType = opt.RoseType
	s.Center = opt.Center
	s.Radius = opt.Radius
}

// NewPie creates a new gauge chart.
func NewPie(routers ...RouterOpts) *Pie {
	chart := new(Pie)
	chart.initBaseOpts(routers...)
	return chart
}

// Add adds new data sets.
func (c *Pie) Add(name string, data map[string]interface{}, options ...seriesOptser) *Pie {
	nvs := make([]datatypes.NameValueItem, 0)
	for k, v := range data {
		nvs = append(nvs, datatypes.NameValueItem{Name: k, Value: v})
	}
	series := singleSeries{Name: name, Type: ChartType.Pie, Data: nvs}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

// SetGlobalOptions sets options for the Pie instance.
func (c *Pie) SetGlobalOptions(options ...globalOptser) *Pie {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Pie) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Pie) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
