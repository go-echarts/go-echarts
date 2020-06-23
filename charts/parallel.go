package charts

import (
	"github.com/go-echarts/go-echarts/types"
	"io"
)

// ParallelAxisOpts is a list of ParallelAxisOpts.
// 平行坐标系中的坐标轴组件配置项
type ParallelAxisOpts []PAOpts

func (ParallelAxisOpts) MarkGlobal() {}

// PAOpts is the option set for a parallel axis.
type PAOpts struct {
	// 坐标轴的维度序号
	Dim int `json:"dim,omitempty"`
	// 坐标轴名称
	Name string `json:"name,omitempty"`
	// 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度。
	// 不设置时会自动计算最大值保证坐标轴刻度的均匀分布
	Max interface{} `json:"max,omitempty"`
	// 是否是反向坐标轴
	Inverse bool `json:"inverse,omitempty"`
	// 坐标轴名称显示位置，可选 "start", "middle", "center", "end"
	// 默认 "end"
	NameLocation string `json:"nameLocation,omitempty"`
	// 坐标轴类型，可选：
	// "value"：数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// 类目数据，在类目轴（type: "category"）中有效
	Data interface{} `json:"data,omitempty"`
}

//func (ParallelComponentOpts) MarkGlobal() {}

// Parallel represents a parallel axis.
type Parallel struct {
	BaseConfiguration
	MultiSeries
}

func (Parallel) Type() string { return types.ChartParallel }

// NewParallel creates a new parallel instance.
func NewParallel() *Parallel {
	chart := &Parallel{}
	chart.initBaseConfiguration()
	chart.HasParallel = true
	return chart
}

// Add adds new data sets.
func (c *Parallel) AddSeries(name string, data interface{}, opts ...SeriesOpts) *Parallel {
	series := SingleSeries{Name: name, Type: types.ChartParallel, Data: data}
	series.configureSeriesOpts(opts...)
	c.MultiSeries = append(c.MultiSeries, series)
	return c
}

// SetGlobalOptions sets options for the Parallel instance.
func (c *Parallel) SetGlobalOptions(opts ...GlobalOpts) *Parallel {
	c.BaseConfiguration.setBaseGlobalOptions(opts...)
	return c
}

func (c *Parallel) Validate() {
	c.Assets.Validate(c.AssetsHost)
}

// Render renders the chart and writes the output to given writers.
func (c *Parallel) Render(w io.Writer) error {
	c.Validate()
	return renderToWriter(c, ModChart, w)
}
