package render

// BaseRender the default implementation of Renderer, make it easier to extend Renderer for individual render function
type BaseRender struct{}

func (r *BaseRender) RenderContent() []byte {
	panic("unsupported render content in current Render!")
}

func (r *BaseRender) RenderSnippet() ChartSnippet {
	panic("unsupported render snippets in current Render!")
}
