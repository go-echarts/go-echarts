# Container of echarts

A `container` is used to define a root dom box of echarts instance and how echarts does.
It contains the height, width...etc, more
in [echarts Chart Container](https://echarts.apache.org/handbook/en/concepts/chart-size)

It focuses on the single echarts box instead of the whole pages.
And go-echarts follows it on each chart as well.
Obviously, go-echarts needs the HTML page management thing before do echarts stuff.

# Container of go-echarts

In go-echarts, there have two approaches to manage charts, for single chart and for multi charts.  
It has a little different points more than a simple aggregation either.  
The `Container` is more close to a `HTML` page result in go-echarts.

- On single chart, it only has the single Chart Container with Chart Render.
- On multi charts, it only has the single Page Container with Page Render.

!> **There is no `1:n` relationship between Chart Container and Page Container.**

We hold the `Charter` instances in program level for `Chart Container` and `Page Container`, it is more straightforward.

## Chart Container

## Page Container