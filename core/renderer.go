package core

type RenderProvider func(old Render) Render
type WriterProvider func(old Writer) Writer

type RenderExposer interface {
	Render(source *Page, dest string)
	GetRenderer() *Renderer
}

type Renderer struct {
	render Render
	writer Writer
}

func (r *Renderer) GetRenderer() *Renderer {
	return r
}

type Render interface {
	Render(page *Page) []byte
}

type Writer interface {
	Write(data []byte)
}

type DefaultRenderer struct {
	Renderer
	File string
}

func (dr *DefaultRenderer) Render(page *Page, dest string) {
	dr.File = dest
	data := dr.render.Render(page)
	dr.writer.Write(data)
}

func NewDefaultRenderer(file string) *DefaultRenderer {

	return &DefaultRenderer{
		Renderer: Renderer{
			render: &DefaultRender{},
			writer: &DefaultFileWriter{File: file},
		},
		File: file,
	}
}
