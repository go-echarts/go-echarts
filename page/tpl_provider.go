package page

/**
Template Provider Interface for user to implement templates and build charts easily.
*/

type TemplateProvider interface {
	Provide() *Page
}
