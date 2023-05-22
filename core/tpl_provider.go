// Package core
/**
An Template Interface for user to build chart easily.
*/
package core

type PageTemplateProvider interface {
	Provide() *Page
}

type ContainerTemplateProvider interface {
	Provide() *Container
}
