package core

import "os"

type RenderProvider func(defaultRender Render) (newProvider Render)
type WriterProvider func(defaultWriter Writer) (newWriter Writer)

type Render interface {
	Render(page *Page) []byte
}

type Writer interface {
	Write(data []byte, dest string)
}

// Renderer A renderer can render *Page and write to any dest format such as `.html, .png` ...etc.
type Renderer interface {
	Resolve(source *Page, dest string)
	SetRender(r Render)
	SetWriter(w Writer)
}

type DefaultRenderer struct {
	render Render
	writer Writer
}

func (r *DefaultRenderer) Resolve(source *Page, dest string) {
	data := r.Render(source)
	r.Write(data, dest)
}

func (r *DefaultRenderer) Write(data []byte, dest string) {
	writer, err := os.Create(dest)
	if err != nil {
		panic(err)
	}

	_, err = writer.Write(data)
	if err != nil {
		panic(err)
	}
}

func (r *DefaultRenderer) GetRenderer() *DefaultRenderer {
	return r
}

func (r *DefaultRenderer) SetRender(render Render) {
	r.render = render
}

func (r *DefaultRenderer) SetWriter(writer Writer) {
	r.writer = writer
}

func NewDefaultRenderer() *DefaultRenderer {
	return &DefaultRenderer{}
}
