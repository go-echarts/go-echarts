package component

type TemplateProvider interface {
	Provide() *Container
}
