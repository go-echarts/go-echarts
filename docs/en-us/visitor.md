# ConfigurationVisitor

When you want to resolve some specific options which is out of control by
go-echarts, you could try to use
the [ConfigurationVisitor](https://github.com/go-echarts/go-echarts/blob/master/charts/visitor.go).

It is a variant Visitor pattern to help you modify chart's option on demands.

```go
type ConfigurationVisitor interface {
    // Visit called after all builtin options settled
    Visit(chart map[string]interface{})
    VisitTitleOpt(title opts.Title) interface{}
    VisitLegendOpt(legend opts.Legend) interface{}
    VisitTooltipOpt(tooltip opts.Tooltip) interface{}
    VisitSeriesOpt(series MultiSeries) interface{}
    VisitDatasets(datasets ...opts.Dataset) interface{}
    VisitAxisPointer(axisPointer *opts.AxisPointer) interface{}
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

```

i.e.

I want to add a non-exist option for title option, the key is `go-echarts`.

Define a Visitor: `MyVisitor`, and then, override the title visitor.

```go
type MyVisitor struct {
    charts.BaseConfigurationVisitor
}

func (mv MyVisitor) VisitTitleOpt(title opts.Title) interface{} {
    type MyTitle struct {
        opts.Title
        // addition field
        GoEcharts string `json:"go-echarts"`
    }
    return &MyTitle{
        Title:     title,
        GoEcharts: "here!",
        }
}

```

> A simple bar chart for verification.

```go
func barTitle() *charts.Bar {
    bar := charts.NewBar()
    bar.SetGlobalOptions(
        charts.WithTitleOpts(opts.Title{
            Title: "title visitor option",
        })
    )
    bar.SetXAxis(weeks). AddSeries("Category A", generateBarItems())
    bar.Accept(&MyVisitor{})
return bar
}

```

The output for the `title` option should be:

```json
{
  ...
  "title": {
    "text": "title visitor option",
    "go-echarts": "here!"
  },
  ...
}

```


