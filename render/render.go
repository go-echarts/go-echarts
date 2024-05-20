package render

import (
	"io"
)

// The DefaultRender implementation of Renderer, make it easier to extend Renderer
type DefaultRender struct {
}

// Render renders the chart(s) into the given io.Writer.
func (r *DefaultRender) Render(w io.Writer) error {
	content := r.RenderContent()
	_, err := w.Write(content)
	return err
}

func (r *DefaultRender) RenderContent() []byte {
	panic("unsupported render content in current Render!")
}

func (r *DefaultRender) RenderSnippet() ChartSnippet {
	panic("unsupported render snippets in current Render!")
}
