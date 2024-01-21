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

The `Action` is the target to implement `myChart.dispatchAction({ type: '' })`, which is used to trigger the behavior.
It allows to manually trigger events on charts to make the chart dynamic.

## More