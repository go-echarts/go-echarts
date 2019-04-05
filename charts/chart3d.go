package charts

import (
	"io"
)

// Chart3D is a chart in 3D coordinate.
type Chart3D struct {
	BaseOpts
	Series

	XAxis3D XAxis3DOpts
	YAxis3D YAxis3DOpts
	ZAxis3D ZAxis3DOpts
	Grid3D  Grid3DOpts

	xData, yData interface{}
}

func (c *Chart3D) initChart3D() {
	c.JSAssets.Add("echarts-gl.min.js")
	c.Has3DAxis = true
}

// SetGlobalOptions sets options for the Chart3D instance.
func (c *Chart3D) SetGlobalOptions(options ...globalOptser) *Chart3D {
	c.BaseOpts.setBaseGlobalOptions(options...)
	c.setChart3DGlobalOptions(options...)
	return c
}

func (c *Chart3D) setChart3DGlobalOptions(options ...globalOptser) {
	for i := 0; i < len(options); i++ {
		option := options[i]
		switch option.(type) {
		case XAxis3DOpts:
			c.XAxis3D = option.(XAxis3DOpts)
		case YAxis3DOpts:
			c.YAxis3D = option.(YAxis3DOpts)
		case ZAxis3DOpts:
			c.ZAxis3D = option.(ZAxis3DOpts)
		case Grid3DOpts:
			c.Grid3D = option.(Grid3DOpts)
		}
	}
}

func (c *Chart3D) validateOpts() {
	// 确保 XY 轴数据项不会被抹除
	if c.XAxis3D.Data == nil {
		c.XAxis3D.Data = c.xData
	}
	if c.YAxis3D.Data == nil {
		c.YAxis3D.Data = c.yData
	}
	c.validateAssets(c.AssetsHost)
}

func (c *Chart3D) addZAxis(chartType, name string, zAxis interface{}, options ...seriesOptser) {
	series := singleSeries{
		Name:        name,
		Type:        chartType,
		Data:        zAxis,
		CoordSystem: ChartType.Cartesian3D,
	}
	series.setSingleSeriesOpts(options...)
	c.Series = append(c.Series, series)
	c.setColor(options...)
}

// Render renders the chart and writes the output to given writers.
func (c *Chart3D) Render(w ...io.Writer) error {
	c.insertSeriesColors(c.appendColor)
	c.validateOpts()
	return renderToWriter(c, "chart", []string{}, w...)
}

// Grid3DOpts contains options for the 3D coordinate.
type Grid3DOpts struct {
	// 是否显示三维笛卡尔坐标系
	Show bool `json:"show,omitempty"`
	// 三维笛卡尔坐标系组件在三维场景中的宽度
	// 默认 100
	BoxWidth float32 `json:"boxWidth,omitempty"`
	// 三维笛卡尔坐标系组件在三维场景中的高度
	// 默认 100
	BoxHeight float32 `json:"boxHeight,omitempty"`
	// 三维笛卡尔坐标系组件在三维场景中的深度
	// 默认 100
	BoxDepth float32 `json:"boxDepth,omitempty"`
	// 用于鼠标的旋转，缩放等视角控制
	ViewControl ViewControlOpts `json:"viewControl,omitempty"`
}

func (Grid3DOpts) markGlobal() {}

// ViewControlOpts contains options for view controlling.
type ViewControlOpts struct {
	// 是否开启视角绕物体的自动旋转查看
	AutoRotate bool `json:"autoRotate,omitempty"`
	// 物体自转的速度。单位为角度 / 秒，默认为 10 ，也就是 36 秒转一圈
	AutoRotateSpeed float32 `json:"autoRotateSpeed,omitempty"`
}

// XAxis3DOpts contains options for X axis in the 3D coordinate.
type XAxis3DOpts struct {
	// 是否显示 3D X 轴
	Show bool `json:"show,omitempty"`
	// X 坐标轴名称
	Name bool `json:"name,omitempty"`
	// X 坐标轴使用的 grid3D 组件的索引。默认使用第一个 grid3D 组件
	Grid3DIndex int `json:"grid3DIndex,omitempty"`
	// X 坐标轴类型，可选：
	// "value" 数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// X 坐标轴刻度最小值。
	// 可以设置成特殊值 "dataMin"，此时取数据在该轴上的最小值作为最小刻度。
	// 不设置时会自动计算最小值保证坐标轴刻度的均匀分布
	Min interface{} `json:"min,omitempty"`
	// X 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度。
	// 不设置时会自动计算最大值保证坐标轴刻度的均匀分布
	Max interface{} `json:"max,omitempty"`
	// 类目数据，在类目轴（type: 'category'）中有效
	Data interface{} `json:"data,omitempty"`
}

func (XAxis3DOpts) markGlobal() {}

// YAxis3DOpts contains options for Y axis in the 3D coordinate.
type YAxis3DOpts struct {
	// 是否显示 3D Y 轴
	Show bool `json:"show,omitempty"`
	// Y 坐标轴名称
	Name bool `json:"name,omitempty"`
	// Y 坐标轴使用的 grid3D 组件的索引。默认使用第一个 grid3D 组件
	Grid3DIndex int `json:"grid3DIndex,omitempty"`
	// Y 坐标轴类型，可选：
	// "value" 数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// Y 坐标轴刻度最小值。
	// 可以设置成特殊值 "dataMin"，此时取数据在该轴上的最小值作为最小刻度。
	// 不设置时会自动计算最小值保证坐标轴刻度的均匀分布
	Min interface{} `json:"min,omitempty"`
	// Y 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度。
	// 不设置时会自动计算最大值保证坐标轴刻度的均匀分布
	Max interface{} `json:"max,omitempty"`
	// 类目数据，在类目轴（type: 'category'）中有效
	Data interface{} `json:"data,omitempty"`
}

func (YAxis3DOpts) markGlobal() {}

// ZAxis3DOpts contains options for Z axis in the 3D coordinate.
type ZAxis3DOpts struct {
	// 是否显示 3D Z 轴
	Show bool `json:"show,omitempty"`
	// Z 坐标轴名称
	Name bool `json:"name,omitempty"`
	// Z 坐标轴使用的 grid3D 组件的索引。默认使用第一个 grid3D 组件
	Grid3DIndex int `json:"grid3DIndex,omitempty"`
	// Z 坐标轴类型，可选：
	// "value" 数值轴，适用于连续数据。
	// "category" 类目轴，适用于离散的类目数据，为该类型时必须通过 data 设置类目数据。
	// "log" 对数轴。适用于对数数据。
	Type string `json:"type,omitempty"`
	// Z 坐标轴刻度最小值。
	// 可以设置成特殊值 "dataMin"，此时取数据在该轴上的最小值作为最小刻度。
	// 不设置时会自动计算最小值保证坐标轴刻度的均匀分布
	Min interface{} `json:"min,omitempty"`
	// Z 坐标轴刻度最大值。
	// 可以设置成特殊值 "dataMax"，此时取数据在该轴上的最大值作为最大刻度。
	// 不设置时会自动计算最大值保证坐标轴刻度的均匀分布
	Max interface{} `json:"max,omitempty"`
	// 类目数据，在类目轴（type: 'category'）中有效
	Data interface{} `json:"data,omitempty"`
}

func (ZAxis3DOpts) markGlobal() {}
