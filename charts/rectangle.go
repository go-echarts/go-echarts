package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Overlaper interface {
	overlap() MultiSeries
}

// XYAxis represent the X and Y axis in the rectangular coordinates.
type XYAxis struct {
	XAxisList []opts.XAxis `json:"xaxis"`
	YAxisList []opts.YAxis `json:"yaxis"`
}

func (xy *XYAxis) initXYAxis() {
	xy.XAxisList = append(xy.XAxisList, opts.XAxis{})
	xy.YAxisList = append(xy.YAxisList, opts.YAxis{})
}

// ExtendXAxis adds new X axes.
func (xy *XYAxis) ExtendXAxis(xAxis ...opts.XAxis) {
	xy.XAxisList = append(xy.XAxisList, xAxis...)
}

// ExtendYAxis adds new Y axes.
func (xy *XYAxis) ExtendYAxis(yAxis ...opts.YAxis) {
	xy.YAxisList = append(xy.YAxisList, yAxis...)
}

// WithXAxisOpts sets the X axis.
func WithXAxisOpts(opt opts.XAxis, index ...int) GlobalOpts {
	return func(bc *BaseConfiguration) {
		if len(index) == 0 {
			index = []int{0}
		}
		for i := 0; i < len(index); i++ {
			bc.XYAxis.XAxisList[index[i]] = opt
		}
	}
}

// WithYAxisOpts sets the Y axis.
func WithYAxisOpts(opt opts.YAxis, index ...int) GlobalOpts {
	return func(bc *BaseConfiguration) {
		if len(index) == 0 {
			index = []int{0}
		}
		for i := 0; i < len(index); i++ {
			bc.XYAxis.YAxisList[index[i]] = opt
		}
	}
}

// RectConfiguration contains options for the rectangular coordinates.
type RectConfiguration struct {
	BaseConfiguration
	BaseActions
}

func (rect *RectConfiguration) setRectGlobalOptions(options ...GlobalOpts) {
	rect.BaseConfiguration.setBaseGlobalOptions(options...)
}

func (rect *RectConfiguration) setRectGlobalActions(options ...GlobalActions) {
	rect.BaseActions.setBaseGlobalActions(options...)
}

// RectChart is a chart in RectChart coordinate.
type RectChart struct {
	RectConfiguration

	xAxisData interface{}
}

func (rc *RectChart) overlap() MultiSeries {
	return rc.MultiSeries
}

// SetGlobalOptions sets options for the RectChart instance.
func (rc *RectChart) SetGlobalOptions(options ...GlobalOpts) *RectChart {
	rc.RectConfiguration.setRectGlobalOptions(options...)
	return rc
}

//SetDispatchActions sets actions for the RectChart instance.
func (rc *RectChart) SetDispatchActions(options ...GlobalActions) *RectChart {
	rc.RectConfiguration.setRectGlobalActions(options...)
	return rc
}

// Overlap composes multiple charts into one single canvas.
// It is only suited for some of the charts which are in rectangular coordinate.
// Supported charts: Bar/BoxPlot/Line/Scatter/EffectScatter/Kline/HeatMap
func (rc *RectChart) Overlap(a ...Overlaper) {
	for i := 0; i < len(a); i++ {
		rc.MultiSeries = append(rc.MultiSeries, a[i].overlap()...)
	}
}

// Validate validates the given configuration.
func (rc *RectChart) Validate() {
	// Make sure that the data of X axis won't be cleaned for XAxisOpts
	rc.XAxisList[0].Data = rc.xAxisData
	// Make sure that the labels of Y axis show correctly
	for i := 0; i < len(rc.YAxisList); i++ {
		rc.YAxisList[i].AxisLabel.Show = true
	}
	rc.Assets.Validate(rc.AssetsHost)
}
