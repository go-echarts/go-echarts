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

!> **There is no `1:n` strict relationship between Chart Container and Page Container.**

We hold the `Charter` instances in program level for `Chart Container` and `Page Container`, it is more straightforward.

## Chart Container

A Chart Container is made for a single chart.
The templates it required:

- header.tpl
- base.tpl
- chart.tpl

It allows to customize the asserts on JS and CSS.  
Because it is only one chart which means it is associated with single echarts container also.  
Any change only applies to the single chart, directly and plain.

## Page Container

A Page Container is made for multi chart.  
The templates it required:

- header.tpl
- base.tpl
- page.tpl

It allows to customize the assets on JS and CSS on page level and charts level.  
Because it contains multi charts in one page, it is associated with multi echarts containers also.  
There support layout types:

- 'none'
- 'center'
- 'flex'

Any changes on page level should align to chart level, otherwise, it may has side effect.

!> Known issue: Only clear preset assets on page level, the charts level's preset assets will replenish it instead.  
