package render

import (
	"github.com/go-echarts/go-echarts/v2/page"
	"os"
)

type RenderProvider func(defaultRender IRender) (newProvider IRender)
type WriterProvider func(defaultWriter IWriter) (newWriter IWriter)

type IRender interface {
	Render(page *page.Page) []byte
}

type IWriter interface {
	Write(data []byte, dest string)
}

// Renderer A renderer can render *Page and write to any dest format such as `.html, .png` ...etc.
type Renderer interface {
	Resolve(source *page.Page, dest string)
	SetRender(r IRender)
	SetWriter(w IWriter)
	GetRender() IRender
	GetWriter() IWriter
}

type DefaultRenderer struct {
	render IRender
	writer IWriter
}

func (r *DefaultRenderer) Resolve(source *page.Page, dest string) {
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

func (r *DefaultRenderer) SetRender(render IRender) {
	r.render = render
}

func (r *DefaultRenderer) SetWriter(writer IWriter) {
	r.writer = writer
}

func (r *DefaultRenderer) GetRender() IRender {
	return r.render
}

func (r *DefaultRenderer) GetWriter() IWriter {
	return r.writer
}
