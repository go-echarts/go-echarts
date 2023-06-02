package component

// Chart all the chart should implement the basic chart interface
type Chart interface {
	GetChart() interface{}
	GetChartType() string
	GetContainer() *Container
	// GetChartToJson()
	// GetChartToMap()
}
