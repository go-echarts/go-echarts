package core

// Chart all the chart should implement the basic chart interface
type Chart interface {
	GetChart() interface{}
	GetChartName() string
	GetContainer() *Container
	GetPage() *Page
	// GetChartToJson()
	// GetChartToMap()
}
