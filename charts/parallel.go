package charts

import (
	"io"

	"github.com/chenjiandongx/go-echarts/common"
)

// 平行坐标系组件配置项
type ParallelComponentOpts struct {
	// parallel 组件离容器左侧的距离。
	// left 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'left', 'center', 'right'。
	// 如果 left 的值为'left', 'center', 'right'，组件会根据相应的位置自动对齐。
	Left string `json:"left,omitempty"`
	// parallel 组件离容器上侧的距离。
	// top 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比
	// 也可以是 'top', 'middle', 'bottom'。
	// 如果 top 的值为'top', 'middle', 'bottom'，组件会根据相应的位置自动对齐。
	Top string `json:"top,omitempty"`
	// parallel 组件离容器右侧的距离。
	// right 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应。
	Right string `json:"right,omitempty"`
	// parallel 组件离容器下侧的距离。
	// bottom 的值可以是像 20 这样的具体像素值，可以是像 '20%' 这样相对于容器高宽的百分比。
	// 默认自适应
	Bottom string `json:"bottom,omitempty"`
}

// 平行坐标系中的坐标轴组件配置项
type ParallelAxisOpts []PAOpts

func (ParallelAxisOpts) markGlobal() {}

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

func (ParallelComponentOpts) markGlobal() {}

type Parallel struct {
	BaseOpts
	Series
}

func (Parallel) chartType() string { return common.ChartType.Parallel }

func NewParallel(routers ...RouterOpts) *Parallel {
	chart := new(Parallel)
	chart.initBaseOpts(false, routers...)
	chart.HasParallel = true
	return chart
}

func (c *Parallel) Add(name string, data interface{}, options ...seriesOptser) *Parallel {
	series := singleSeries{Name: name, Type: common.ChartType.Parallel, Data: data}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
	return c
}

func (c *Parallel) SetGlobalOptions(options ...globalOptser) *Parallel {
	c.BaseOpts.setBaseGlobalOptions(options...)
	return c
}

func (c *Parallel) validateOpts() {
	c.validateAssets(c.AssetsHost)
}

func (c *Parallel) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}
