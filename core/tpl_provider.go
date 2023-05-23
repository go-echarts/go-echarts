package core

/**
Template Provider Interface for user to implement templates and build charts easily.
*/

type PageTemplateProvider interface {
	Provide() *Page
}

type ContainerTemplateProvider interface {
	Provide() *Container
}
