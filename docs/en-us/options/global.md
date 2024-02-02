# General Topic

There has tons of configs/options in `echars`, as well as in `go-echarts`.
TBH, if we want set them correctly is not that quite easy.
Luckily, the [`echarts` document](https://echarts.apache.org/en/option.html#title) provides so much details on each
configuration.
So, We don't want duplicate the docs of options here.
I recommend that you could find the options you want
on that docs, and build it on `go-echarts`, it is a fast and practicable way.
If you already have some `JS` knowledgebase, it could be a big plus to help you convert all the functions
into `go-echarts`.

## Page/Charts level

On the chart level, or I should say on the `page` level.
What options we can image instantaneously about a chart or a HTML page of echarts?

- Assets
- Layout
- Size

Yes, that's it.

And this is what you need all of them. :)

```go

// WithInitializationOpts sets the initialization.
// Change Page title, page size, asset host and background color
func WithInitializationOpts(opt opts.Initialization) GlobalOpts { }

// ClearPresetAssets clear both the preset JS and CSS static assets.
func (opt *Assets) ClearPresetAssets() { }

// ClearPresetJSAssets only clear all the preset JS static assets.
func (opt *Assets) ClearPresetJSAssets() { }

// ClearPresetCSSAssets only clear all the preset CSS static assets.
func (opt *Assets) ClearPresetCSSAssets() { }

// AddCustomizedJSAssets adds the customized javascript assets which will not be added the `host` prefix.
func (opt *Assets) AddCustomizedJSAssets(assets ...string) { }

// AddCustomizedCSSAssets adds the customized css assets which will not be added the `host` prefix.
func (opt *Assets) AddCustomizedCSSAssets(assets ...string) { }

```

## Option level

As we all know, a typical echarts instance looks like this:
> example
> from [echarts Bar Chart with Negative Value](https://echarts.apache.org/examples/en/editor.html?c=bar-negative)

```js
option = {
    tooltip: {
        trigger: 'axis',
        axisPointer: {
            type: 'shadow'
        }
    },
    legend: {
        data: ['Profit', 'Expenses', 'Income']
    },
    grid: {
        left: '3%',
        right: '4%',
        bottom: '3%',
        containLabel: true
    },
    xAxis: [
        {
            type: 'value'
        }
    ],
    yAxis: [
        {
            type: 'category',
            axisTick: {
                show: false
            },
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
        }
    ],
    series: [
        {
            name: 'Profit',
            type: 'bar',
            label: {
                show: true,
                position: 'inside'
            },
            emphasis: {
                focus: 'series'
            },
            data: [200, 170, 240, 244, 200, 220, 210]
        },
        {
            name: 'Income',
            type: 'bar',
            stack: 'Total',
            label: {
                show: true
            },
            emphasis: {
                focus: 'series'
            },
            data: [320, 302, 341, 374, 390, 450, 420]
        },
        {
            name: 'Expenses',
            type: 'bar',
            stack: 'Total',
            label: {
                show: true,
                position: 'left'
            },
            emphasis: {
                focus: 'series'
            },
            data: [-120, -132, -101, -134, -190, -230, -210]
        }
    ]
};
```

All of those configs called `option` but we can see there has two parts.

- For Chart
    - `tooltip`
    - `legend`
    - `grid`
      ...
- For Dataset
    - `series[0]`
    - `series[1]`
    - `series[2]`
      ...

As per these characteristics, we separate the `option` into two things, the `Options` and the `Series`.

- `Options`
  How does this whole chart general look.
- `Series`
  How does each dataset styles and showing.

