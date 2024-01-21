# Container of echarts

A `Container` is used to define a root dom box of echarts instance and how to render charts.
It contains the height, width...etc, more
in [echarts Chart Container](https://echarts.apache.org/handbook/en/concepts/chart-size)

It focuses on the single echarts box instead of the whole pages,
go-echarts follows its rules on each chart level as well.  
Obviously, go-echarts needs the HTML page management thing before do echarts stuff also.

# Container of go-echarts

In go-echarts, there have two approaches to manage charts, for single chart and for multi charts.  
It has a little different points more than a simple aggregation operation either.  
The `Container` is more close to a `HTML` page result in go-echarts.

- On single chart, it only has the single Chart Container with Chart Render.
- On multi charts, it only has the single Page Container with Page Render.

!> **There is no strict `1:n` relationship between `Chart Container` and `Page Container`.**

We hold the `Charter` (interface for charts) instances in program level for `Chart Container` and `Page Container`
with individual templates.

## Chart Container

A Chart Container is made for a single chart.
The templates it required:

- header.tpl
- base.tpl
- chart.tpl

It allows to customize the asserts on JS and CSS, and provides the shortcuts for page settings.  
Because it is only one chart which means it is associated with single echarts container also.  
Any change only applies to the single chart page, directly and plain.

## Page Container

A Page Container is made for multi charts.  
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

Any changes on page level should compatible to chart level, otherwise, it may has side effect.

!> Attention :

- In `Page Container`, all the config for page related options on chart level will be ignored, such as `PageTitile`.
- Only clear preset assets on page level, the charts level's preset assets will replenish it, maybe you need clear both.  
