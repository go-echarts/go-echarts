package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
)

// ConfigurationVisitor visitor, also the transformer.
// Allow to visit and modify/enhance chart options before rendering, even though the option structure isn't built-in.
type ConfigurationVisitor interface {
	// Visit called after all builtin options settled
	Visit(chart map[string]interface{})
	VisitTitleOpt(title opts.Title) interface{}
	VisitLegendOpt(legend opts.Legend) interface{}
	VisitTooltipOpt(tooltip opts.Tooltip) interface{}
	VisitSeriesOpt(series MultiSeries) interface{}
	VisitDatasets(datasets ...opts.Dataset) interface{}
	VisitAxisPointer(axisPointer *opts.AxisPointer) interface{}
	VisitAria(axisPointer *opts.Aria) interface{}
	VisitPolar(polar opts.Polar) interface{}
	VisitAngleAxis(angleAxis opts.AngleAxis) interface{}
	VisitRadiusAxis(radiusAxis opts.RadiusAxis) interface{}
	VisitGeo(geo opts.GeoComponent) interface{}
	VisitRadar(radar opts.RadarComponent) interface{}
	VisitParallel(parallel opts.ParallelComponent) interface{}
	VisitParallelAxis(parallelAxes []opts.ParallelAxis) interface{}
	VisitSingleAxis(singleAxis opts.SingleAxis) interface{}
	VisitToolbox(toolbox opts.Toolbox) interface{}
	VisitDataZooms(dataZooms []opts.DataZoom) interface{}
	VisitVisualMaps(visualMaps []opts.VisualMap) interface{}
	VisitXAxis(xAxis []opts.XAxis) interface{}
	VisitYAxis(yAxis []opts.YAxis) interface{}
	VisitGrid(grid []opts.Grid) interface{}
	VisitXAxis3D(xAxis3D opts.XAxis3D) interface{}
	VisitYAxis3D(yAxis3D opts.YAxis3D) interface{}
	VisitZAxis3D(zAxis3D opts.ZAxis3D) interface{}
	VisitGrid3D(grid3D opts.Grid3D) interface{}
	VisitBrush(brush opts.Brush) interface{}
	VisitCalendar(calendar []*opts.Calendar) interface{}
}

// BaseConfigurationVisitor a default visitor, noop
type BaseConfigurationVisitor struct {
}

func (b BaseConfigurationVisitor) Visit(chart map[string]interface{}) {
	// noop
}

func (b BaseConfigurationVisitor) VisitTitleOpt(title opts.Title) interface{} {
	return title
}

func (b BaseConfigurationVisitor) VisitLegendOpt(legend opts.Legend) interface{} {
	return legend
}

func (b BaseConfigurationVisitor) VisitTooltipOpt(tooltip opts.Tooltip) interface{} {
	return tooltip
}

func (b BaseConfigurationVisitor) VisitSeriesOpt(series MultiSeries) interface{} {
	return series
}

func (b BaseConfigurationVisitor) VisitDatasets(datasets ...opts.Dataset) interface{} {
	return datasets
}

func (b BaseConfigurationVisitor) VisitAxisPointer(axisPointer *opts.AxisPointer) interface{} {
	return axisPointer
}

func (b BaseConfigurationVisitor) VisitAria(aria *opts.Aria) interface{} {
	return aria
}

func (b BaseConfigurationVisitor) VisitPolar(polar opts.Polar) interface{} {
	return polar
}

func (b BaseConfigurationVisitor) VisitAngleAxis(angleAxis opts.AngleAxis) interface{} {
	return angleAxis
}

func (b BaseConfigurationVisitor) VisitRadiusAxis(radiusAxis opts.RadiusAxis) interface{} {
	return radiusAxis
}

func (b BaseConfigurationVisitor) VisitGeo(geo opts.GeoComponent) interface{} {
	return geo
}

func (b BaseConfigurationVisitor) VisitRadar(radar opts.RadarComponent) interface{} {
	return radar
}

func (b BaseConfigurationVisitor) VisitParallel(parallel opts.ParallelComponent) interface{} {
	return parallel
}

func (b BaseConfigurationVisitor) VisitParallelAxis(parallelAxes []opts.ParallelAxis) interface{} {
	return parallelAxes
}

func (b BaseConfigurationVisitor) VisitSingleAxis(singleAxis opts.SingleAxis) interface{} {
	return singleAxis
}

func (b BaseConfigurationVisitor) VisitToolbox(toolbox opts.Toolbox) interface{} {
	return toolbox
}

func (b BaseConfigurationVisitor) VisitDataZooms(dataZooms []opts.DataZoom) interface{} {
	return dataZooms
}

func (b BaseConfigurationVisitor) VisitVisualMaps(visualMaps []opts.VisualMap) interface{} {
	return visualMaps
}

func (b BaseConfigurationVisitor) VisitXAxis(xAxis []opts.XAxis) interface{} {
	return xAxis
}

func (b BaseConfigurationVisitor) VisitYAxis(yAxis []opts.YAxis) interface{} {
	return yAxis
}

func (b BaseConfigurationVisitor) VisitGrid(grid []opts.Grid) interface{} {
	return grid
}

func (b BaseConfigurationVisitor) VisitXAxis3D(xAxis3D opts.XAxis3D) interface{} {
	return xAxis3D
}

func (b BaseConfigurationVisitor) VisitYAxis3D(yAxis3D opts.YAxis3D) interface{} {
	return yAxis3D
}

func (b BaseConfigurationVisitor) VisitZAxis3D(zAxis3D opts.ZAxis3D) interface{} {
	return zAxis3D
}

func (b BaseConfigurationVisitor) VisitGrid3D(grid3D opts.Grid3D) interface{} {
	return grid3D
}

func (b BaseConfigurationVisitor) VisitBrush(brush opts.Brush) interface{} {
	return brush
}

func (b BaseConfigurationVisitor) VisitCalendar(calendar []*opts.Calendar) interface{} {
	return calendar
}
