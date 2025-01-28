package charts

import (
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Overlaper interface {
	asRectChart() *RectChart
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
}

func (rect *RectConfiguration) setRectGlobalOptions(options ...GlobalOpts) {
	rect.BaseConfiguration.setBaseGlobalOptions(options...)
}

// RectChart is a chart in RectChart coordinate.
type RectChart struct {
	RectConfiguration

	xAxisData interface{}
}

func (rc *RectChart) asRectChart() *RectChart {
	return rc
}

// The default RectAggregator for the series data merge
func (rc *RectChart) overlap() RectAggregator {
	return func(this *RectChart, that *RectChart) *RectChart {
		this.MultiSeries = append(this.MultiSeries, that.MultiSeries...)
		return this
	}
}

// SetGlobalOptions sets options for the RectChart instance.
func (rc *RectChart) SetGlobalOptions(options ...GlobalOpts) *RectChart {
	rc.RectConfiguration.setRectGlobalOptions(options...)
	return rc
}

// Overlap composes multiple charts' Series data into one single canvas.
// It is only suited for some of the charts which are in rectangular coordinate.
// Supported charts: Bar/BoxPlot/Line/Scatter/EffectScatter/Kline/HeatMap/Custom
func (rc *RectChart) Overlap(a ...Overlaper) {
	rc.Aggregate(rc.overlap(), a...)
}

// RectAggregator is a binary operator for the RectChart merge function
type RectAggregator func(this *RectChart, that *RectChart) *RectChart

// Aggregate aggregation multiple rect charts into one single canvas.
// An aggregator will merge all charts in orders into the final one.
// It is only suited for some of the charts which are in rectangular coordinate.
func (rc *RectChart) Aggregate(aggregator RectAggregator, charts ...Overlaper) {
	identity := rc
	for _, o := range charts {
		c := o.asRectChart()
		identity = aggregator(identity, c)
	}
}

// Validate validates the given configuration.
func (rc *RectChart) Validate() {
	// Make sure that the data of X axis won't be cleaned for XAxisOpts
	rc.XAxisList[0].Data = rc.xAxisData
	// Make sure that the labels of Y axis show correctly
	for i := 0; i < len(rc.YAxisList); i++ {
		rc.YAxisList[i].AxisLabel.Show = opts.Bool(true)
	}
	rc.Assets.Validate(rc.AssetsHost)
}
