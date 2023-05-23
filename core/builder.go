package core

type Builder struct {
	pageBuilder     *PageBuilder
	chartsBuilder   *ChartsBuilder
	rendererBuilder *RendererBuilder
}

type RendererBuilder struct {
	builder      *Builder
	renderer     RendererExposer
	renderConfig RenderProvider
	writerConfig WriterProvider
}

type PageBuilder struct {
	builder    *Builder
	page       *Page
	pageConfig PageConfig
}

func (cb *Builder) Page() *PageBuilder {
	cb.pageBuilder = &PageBuilder{
		builder: cb,
		page:    NewPage(),
	}
	return cb.pageBuilder
}

func (pb *PageBuilder) UseTemplate(provider PageTemplateProvider) *PageBuilder {
	pb.page = provider.Provide()
	return pb
}

func (pb *PageBuilder) PageConfig(config PageConfig) *PageBuilder {
	pb.pageConfig = config
	return pb
}

func (pb *PageBuilder) doBuildPage() *Page {
	pb.page.Config(pb.pageConfig)
	return pb.page
}

func (pb *PageBuilder) Charts() *ChartsBuilder {
	pb.builder.chartsBuilder = &ChartsBuilder{
		builder:    pb.builder,
		containers: []*Container{},
	}
	return pb.builder.chartsBuilder

}

type ChartsBuilder struct {
	builder    *Builder
	containers []*Container
}

func (cb *ChartsBuilder) AddCharts(charts ...Chart) *RendererBuilder {
	for _, chart := range charts {
		cb.containers = append(cb.containers, chart.GetContainer())
	}

	cb.builder.rendererBuilder = &RendererBuilder{
		builder: cb.builder,
	}
	cb.builder.pageBuilder.page.AddCharts(charts...)
	return cb.builder.rendererBuilder
}

// Render Use f... to make this function accept f or not
func (rb *RendererBuilder) Render(f ...string) {
	page := rb.builder.pageBuilder.doBuildPage()

	file := ""

	if f != nil {
		file = f[0]
	}

	if rb.builder.rendererBuilder.renderer == nil {
		rb.builder.rendererBuilder.renderer = NewDefaultRenderer()
	}

	render := rb.builder.rendererBuilder.doBuildRenderer()

	render.Resolve(page, file)
}

// CustomRenderer allow to replace whole renderer directly, such as PNG renderer
func (rb *RendererBuilder) CustomRenderer(re RendererExposer) *RendererBuilder {
	if re != nil {
		rb.renderer = re
	}
	return rb
}

// RendererConfig allow to replace either Render or Writer
func (rb *RendererBuilder) RendererConfig(renderProvider RenderProvider, writerProvider WriterProvider) *RendererBuilder {
	if renderProvider != nil {
		rb.renderConfig = renderProvider
	}

	if writerProvider != nil {
		rb.writerConfig = writerProvider
	}
	return rb
}

func (rb *RendererBuilder) doBuildRenderer() RendererExposer {
	if rb.renderConfig != nil {
		re := rb.renderer.GetRenderer().render
		rb.renderer.SetRender(rb.renderConfig(re))
	}

	if rb.writerConfig != nil {
		we := rb.renderer.GetRenderer().writer
		rb.renderer.SetWriter(rb.writerConfig(we))
	}
	return rb.renderer
}
