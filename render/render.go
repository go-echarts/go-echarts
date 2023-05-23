package render

import (
	"github.com/go-echarts/go-echarts/v2/core"
)

func Render(page *core.Page, dest string) {
	core.NewDefaultRenderer().Resolve(page, dest)
}

func RenderWithConfig(config *Config, page *core.Page, dest string) {
	var renderer core.RendererExposer
	if config.CustomerRenderer != nil {
		renderer = config.CustomerRenderer
	}

	if config.CustomerRender != nil {
		renderer.SetRender(config.CustomerRender)
	}

	if config.CustomerWriter != nil {
		renderer.SetWriter(config.CustomerWriter)
	}
	renderer.Resolve(page, dest)
}

type Config struct {
	CustomerRenderer core.RendererExposer
	CustomerRender   core.Render
	CustomerWriter   core.Writer
}
