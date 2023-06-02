package canvas

import (
	"github.com/go-echarts/go-echarts/v2/component"
	page "github.com/go-echarts/go-echarts/v2/page"
	"github.com/go-echarts/go-echarts/v2/render"
)

type Builder struct {
	pageBuilder     *PageBuilder
	chartsBuilder   *ChartsBuilder
	rendererBuilder *RendererBuilder
}

type RendererBuilder struct {
	builder      *Builder
	renderer     render.Renderer
	renderConfig render.RenderProvider
	writerConfig render.WriterProvider
}

type PageBuilder struct {
	builder    *Builder
	page       *page.Page
	pageConfig page.Config
}

func (cb *Builder) Page() *PageBuilder {
	cb.pageBuilder = &PageBuilder{
		builder: cb,
		page:    page.NewPage(),
	}
	return cb.pageBuilder
}

func (pb *PageBuilder) UseTemplate(provider page.TemplateProvider) *PageBuilder {
	pb.page = provider.Provide()
	return pb
}

func (pb *PageBuilder) PageConfig(config page.Config) *PageBuilder {
	pb.pageConfig = config
	return pb
}

func (pb *PageBuilder) doBuildPage() *page.Page {
	pb.page.Config(pb.pageConfig)
	return pb.page
}

func (pb *PageBuilder) Charts() *ChartsBuilder {
	pb.builder.chartsBuilder = &ChartsBuilder{
		builder:    pb.builder,
		containers: []*component.Container{},
	}
	return pb.builder.chartsBuilder

}

type ChartsBuilder struct {
	builder    *Builder
	containers []*component.Container
}

func (cb *ChartsBuilder) AddCharts(charts ...component.Chart) *RendererBuilder {
	for _, chart := range charts {
		cb.containers = append(cb.containers, chart.GetContainer())
	}

	cb.builder.rendererBuilder = &RendererBuilder{
		builder: cb.builder,
	}
	cb.builder.pageBuilder.page.AddCharts(charts...)
	return cb.builder.rendererBuilder
}

// Render Use f... to make this function accept f or not since they may not need the dest with custom renderer
func (rb *RendererBuilder) Render(f ...string) {
	page := rb.builder.pageBuilder.doBuildPage()

	file := ""

	if f != nil {
		file = f[0]
	}

	render := rb.builder.rendererBuilder.renderer

	if render == nil {
		rb.builder.rendererBuilder.renderer = NewDefaultRenderer()
		rb.builder.rendererBuilder.doBuildRenderer()
		render = rb.builder.rendererBuilder.renderer
	}

	render.Resolve(page, file)
}

// CustomRenderer allow to replace whole renderer directly, such as PNG renderer
func (rb *RendererBuilder) CustomRenderer(re render.Renderer) *RendererBuilder {
	if re != nil {
		rb.renderer = re
	}
	return rb
}

// RendererConfig allow to replace either Render or IWriter based on default renderer
func (rb *RendererBuilder) RendererConfig(renderProvider render.RenderProvider, writerProvider render.WriterProvider) *RendererBuilder {
	if renderProvider != nil {
		rb.renderConfig = renderProvider
	}

	if writerProvider != nil {
		rb.writerConfig = writerProvider
	}
	return rb
}

func (rb *RendererBuilder) doBuildRenderer() render.Renderer {
	if rb.renderConfig != nil {
		re := rb.renderer.(*render.DefaultRenderer).GetRender()
		rb.renderer.SetRender(rb.renderConfig(re))
	}

	if rb.writerConfig != nil {
		we := rb.renderer.(*render.DefaultRenderer).GetWriter()
		rb.renderer.SetWriter(rb.writerConfig(we))
	}
	return rb.renderer
}

func NewDefaultRenderer() *render.DefaultRenderer {
	return &render.DefaultRenderer{}
}
