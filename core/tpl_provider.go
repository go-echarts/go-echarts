package core

type PageTemplateProvider interface {
	Provide() *Page
}

type ContainerTemplateProvider interface {
	Provide() *Container
}
