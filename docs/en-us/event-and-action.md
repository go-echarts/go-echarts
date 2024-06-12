# Event And Action

This selection talks about how go-echarts implement the events
and actions according to the [echarts Event and Action](https://echarts.apache.org/handbook/en/concepts/event).

## Event

> Users can trigger corresponding events by their operation.
> The developer can handle the callback function by listening
> to these events... ( from echarts docs )

A typical event handler in echarts:

```js
myChart.on('click', function (params) {
    // Print name in console
    console.log(params.name);
});
```

or

```js
myChart.on('mouseover', {seriesIndex: 1, name: 'foo'}, function () {
    // when data named 'foo' in series index 1 triggered 'mouseover'.
});
```

It is a pure JS function event handler registering to echarts instance.

In go-echarts, we provide the `event.Listener` as the wrapper of echarts event listener.

```go
type Listener struct {
EventName string
Query     types.FuncStr
Handler   types.FuncStr
}
```

The `EventName` is obvious the event, such as `click`, `mouseover`...etc,
see [Event types](https://echarts.apache.org/en/api.html#events).  
The `Query` is the filter for events, such as `'series.line'` or `{ seriesName: 'foo' }`.  
The `Handler` is the handler JS function string with the params, i.e. `(params)=> {...}`.

A full example on bar chart:

```go
    bar := charts.NewBar()
    JFunc := ` (params) => alert(params.name) `
    bar.SetGlobalOptions(
            charts.WithEventListeners(
            event.Listener{
            EventName: "click",
            Handler:   opts.FuncOpts(JFunc),
        },
            event.Listener{
            EventName: "mouseup",
            Query:     "'series'",
            Handler:   opts.FuncOpts(JFunc),
        },
            event.Listener{
            EventName: "mouseover",
            Query:     "{ seriesName: 'go-echarts' }",
            Handler:   opts.FuncOpts(JFunc),
        }
    )

```

## Action

The `Action` is considered to implement `myChart.dispatchAction({ type: '' })`, which is used to trigger the behavior.
It allows to manually trigger events on charts to make the chart dynamic.

To be honest, it is hard to implement the full functions since we can not run all things
like a pure JS.
Considering for a chart lib, the target is not implement an echarts engine in go.  
On the one hand, the `action`/`animation` things are moreover the `charts` scope, on the other hand, there is
hard to decide where to put the `dispatchAction` part, inside other JS functions? or a static one?
Hence, we haven't provided an action api yet, for the *static* `dispatchAction`.

So, is it no way to make it?  
#### *Actually, we do have one more thing...*

With the power of `%MY_ECHARTS` (see `dive-into` chapter),
absolutely, when you hold the echarts instance, you get the whole world.

!> Talk is cheap, show you the code :)  

The implementation of the example in
echarts [Writing Code to Trigger Component Action Manually](https://echarts.apache.org/handbook/en/concepts/event). 
![image](https://github.com/go-echarts/go-echarts/assets/33706142/dbfc22af-b628-4f3c-951d-2021e0f0e52b)


```go
var (
	itemCntPie = 5
	data       = []string{"Direct Access", "Email Marketing", "Affiliate Ads ", "Video Ads", "Search Engines"}
)

func generatePieItems() []opts.PieData {
    items := make([]opts.PieData, 0)
    for i := 0; i < itemCntPie; i++ {
        items = append(items, opts.PieData{Name: data[i], Value: rand.Intn(500)})
    }
    return items
}

func PieWithDispatchAction() *charts.Pie {
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

	pie := charts.NewPie()
	pie.AddJSFuncStrs(actionWithEchartsInstance)
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

So, you can make the `dispatchAction` purely in JS insertion instead of building some go types.

```js
const myChart = %MY_ECHARTS%;
myChart.dispatchAction({
    ...
});

```

Although it seems a little tricky, enough as a workaround. If you really need it.
