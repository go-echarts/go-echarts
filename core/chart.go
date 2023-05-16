package core

type Chart interface {
	GetChart() interface{}
	GetChartName() string
	GetContainer() *Container
	GetPage() *Page
	// GetChartToJson()
	// GetChartToMap()
}
