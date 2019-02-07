package charts

type Chart3D struct {
	BaseOpts
	Series

	XAxis XAxis3DOpts
	YAxis YAxis3DOpts

	Grid3DOpacity float32
	Grid3DShading string
	XAxis3DType   string
	YAxis3DType   string
	ZAxis3DType   string
}

// 设置 Chart3D 全局配置项
func (c *Chart3D) setRectGlobalOptions(options ...globalOptser) {
	c.BaseOpts.setBaseGlobalOptions(options...)
}

func (c *Chart3D) initChart3D() {
	c.JSAssets.Add("echarts-gl.min.js")
	c.Has3DAxis = true
}

type XAxis3DOpts struct {
	Show        bool        `json:"show"`
	Name        bool        `json:"name"`
	Grid3DIndex int         `json:"grid3DIndex"`
	Type        string      `json:"type"`
	Min         interface{} `json:"min"`
	Max         interface{} `json:"max"`
	Data        interface{} `json:"data"`
}

type YAxis3DOpts struct {
	Show        bool        `json:"show"`
	Name        bool        `json:"name"`
	Grid3DIndex int         `json:"grid3DIndex"`
	Type        string      `json:"type"`
	Min         interface{} `json:"min"`
	Max         interface{} `json:"max"`
	Data        interface{} `json:"data"`
}
