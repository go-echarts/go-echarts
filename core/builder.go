package core

type ChartsBuilder struct {
	pageBuilder     *PageBuilder
	chartBuilder    *ChartBuilder
	rendererBuilder *RendererBuilder
}

func NewCanvas() *ChartsBuilder {
	return &ChartsBuilder{}
}

type RendererExchangeBuilder struct {
	builder      *ChartsBuilder
	renderer     RenderExposer
	renderConfig RenderProvider
	writerConfig WriterProvider
}

type RendererBuilder struct {
	builder      *ChartsBuilder
	renderer     RenderExposer
	renderConfig RenderProvider
	writerConfig WriterProvider
}

type PageBuilder struct {
	builder    *ChartsBuilder
	page       *Page
	pageConfig PageConfig
}

func (cb *ChartsBuilder) Page() *PageBuilder {
	cb.pageBuilder = &PageBuilder{
		builder: cb,
		page:    NewPage(),
	}
	return cb.pageBuilder
}

//func (pb *PageBuilder) UseTemplate(pageTemplate *Page) *PageBuilder {
//	pb.page = pageTemplate
//	return pb
//}

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

func (pb *PageBuilder) Charts() *ChartBuilder {
	pb.builder.chartBuilder = &ChartBuilder{
		builder:    pb.builder,
		containers: []*Container{},
	}
	return pb.builder.chartBuilder

}

type ChartBuilder struct {
	builder    *ChartsBuilder
	containers []*Container
}

// AddCharts TODO template support
func (cb *ChartBuilder) AddCharts(charts ...Chart) *RendererBuilder {
	for _, chart := range charts {
		cb.containers = append(cb.containers, chart.GetContainer())
	}

	cb.builder.rendererBuilder = &RendererBuilder{
		builder: cb.builder,
	}
	cb.builder.pageBuilder.page.AddCharts(charts...)
	return cb.builder.rendererBuilder
}

func (rb *RendererBuilder) Render(f ...string) {
	if f == nil {
		return
	}

	file := f[0]
	rb.builder.rendererBuilder.renderer = NewDefaultRenderer(file)

	page := rb.builder.pageBuilder.doBuildPage()

	render := rb.builder.rendererBuilder.doBuildRenderer()
	render.Render(page)
}

func (rb *RendererBuilder) CustomRenderer(re RenderExposer) *RendererBuilder {
	if re != nil {
		rb.renderer = re
	}
	return rb
}

func (rb *RendererBuilder) RendererConfig(renderProvider RenderProvider, writerProvider WriterProvider) *RendererBuilder {
	if renderProvider != nil {
		rb.renderConfig = renderProvider
	}

	if writerProvider != nil {
		rb.writerConfig = writerProvider
	}
	return rb
}

func (rb *RendererBuilder) doBuildRenderer() RenderExposer {
	if rb.renderConfig != nil {
		re := rb.renderer.GetRenderer().render
		rb.renderer.GetRenderer().render = rb.renderConfig(re)
	}

	if rb.writerConfig != nil {
		we := rb.renderer.GetRenderer().writer
		rb.renderer.GetRenderer().writer = rb.writerConfig(we)
	}
	return rb.renderer
}
