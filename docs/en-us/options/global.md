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

## Create Chart in go-echarts

Now, let me explain how to create a chart instance in go-echarts.

```go
func PieWithDispatchAction() *charts.Pie {
	// define a pureJS function which you want embed it in the HTML page
	const actionWithEchartsInstance = `
		let currentIndex = -1;
		setInterval(function() {
		  const myChart = %MY_ECHARTS%;
		  var dataLen = myChart.getOption().series[0].data.length;
		  myChart.dispatchAction({
			type: 'downplay',
			seriesIndex: 0,
			dataIndex: currentIndex
		  });
		  currentIndex = (currentIndex + 1) % dataLen;
		  myChart.dispatchAction({
			type: 'highlight',
			seriesIndex: 0,
			dataIndex: currentIndex
		  });
		  myChart.dispatchAction({
			type: 'showTip',
			seriesIndex: 0,
			dataIndex: currentIndex
		  });
		}, 1000);
`

	// new a chart instance
	pie := charts.NewPie()
	// Add custom JSFunc
	pie.AddJSFuncStrs(actionWithEchartsInstance)
	// Add GlobalOptions including assets and `Options` (how the chart show)
	// There is lots of options you can put it in, ensure it works for you chart type tho.
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "dispatchAction pie",
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Trigger:   "item",
			Formatter: "{a} <br/>{b} : {c} ({d}%)",
		}),
		charts.WithLegendOpts(opts.Legend{
			Left:   "left",
			Orient: "vertical",
		}),
	)
    
	// Add Series, the datasets and the `Series` options (how the series show)
	pie.AddSeries("pie action", generatePieItems()).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:      opts.Bool(true),
				Formatter: "{b}: {c}",
			}),
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"55%"},
				Center: []string{"50%", "60%"},
			}),

			charts.WithEmphasisOpts(opts.Emphasis{
				ItemStyle: &opts.ItemStyle{
					ShadowBlur:    10,
					ShadowOffsetX: 0,
					ShadowColor:   "rgba(0, 0, 0, 0.5)",
				},
			}),
		)

	return pie
}
```

Additionally, when you have different types of series, you need the `Overlap` functions.
```js
series: [
    {
      name: 'Precipitation',
      type: 'bar',
      data: [
        2.6, 5.9, 9.0, 26.4, 28.7, 70.7, 175.6, 182.2, 48.7, 18.8, 6.0, 2.3
      ]
    },
    {
      name: 'Temperature',
      type: 'line',
      data: [2.0, 2.2, 3.3, 4.5, 6.3, 10.2, 20.3, 23.4, 23.0, 16.5, 12.0, 6.2]
    }
  ]
```

In go-echarts:

```go
    line := charts.NewLine()
    bar := charts.NewBar()
    ...
    bar.Overlap(line)
```

It will add the line's all series into bar's and the `Options` is only obtained from bar's. 

As you can see, there is two parts of the charts function.
- SetGlobalOptions
- SetSeriesOptions

They map to what the configs we mentioned early with lots of `WithXXX` helper functions.
Normally, the `WithXXX` helper functions should meet your needs.
And you can set those options all from the Charts' instance side either.
