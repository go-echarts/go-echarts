package core

import "os"

type RenderProvider func(defaultRender Render) Render
type WriterProvider func(defaultWriter Writer) Writer

type Render interface {
	Render(page *Page) []byte
}

type Writer interface {
	Write(data []byte, dest string)
}

// RendererExposer Renderer wrapper interface which exposing Render and GetRenderer return instance
type RendererExposer interface {
	Resolve(source *Page, dest string)
	GetRenderer() *Renderer
	SetRender(r Render)
	SetWriter(w Writer)
}

type Renderer struct {
	render Render
	writer Writer
}

func (r *Renderer) Resolve(source *Page, dest string) {
	data := r.Render(source)
	r.Write(data, dest)
}

func (r *Renderer) Write(data []byte, dest string) {
	writer, err := os.Create(dest)
	if err != nil {
		panic(err)
	}

	_, err = writer.Write(data)
	if err != nil {
		panic(err)
	}
}

func (r *Renderer) GetRenderer() *Renderer {
	return r
}

func (r *Renderer) SetRender(render Render) {
	r.render = render
}

func (r *Renderer) SetWriter(writer Writer) {
	r.writer = writer
}

func NewDefaultRenderer() *Renderer {

	return &Renderer{}
}
