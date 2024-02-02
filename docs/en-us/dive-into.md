# Summary

?> go-echarts is a golang lib providing the power to create [echarts](https://echarts.apache.org/) styling charts in go.

---

# What go-echarts does

- Provide the generic HTML templates with echarts resources.
- Generate echarts required options and datasets based on go programs.
- Mount options to echarts' instance and render all in one place.

---

# What go-echarts changes and provide

## Option Types

echarts is a pure JS lib which provides lots of flexible abilities on types and functions. 
In golang, go-echarts is not feasible to achieve it in a simple way.
Besides, some of the options may have default values which may not consistent with go types.
Therefore, we change some types following this guideline for now:

- If the bool option with default(missing) value `(true)`, define it `*bool`(`types.Bool`) and provide the
  handy func `opts.Bool()` to do convert easily.
- If the option is `string|string array` or `number|number array`, define it in `type array` or `interface{}`.
- If the option is a complex struct, define it `interface{}` or providing an usable function, convert it then.

## JS Function and Formatter String

> Formatter string sounds like a DSL for echarts, such as `formatter: '{b0}: {c0}<br />{b1}: {c1}'` or a pure JS
> function.

Since go-echarts renders the options to charts pure in go,
there is no way to apply JS functions either on echarts options or echarts instance directly.  
Instead, we provide the ability to insert pure JS function string
and some useful functional wrapper api to help users do operations on echarts and echarts instance.  
The type of those kind of `Function String` is `types.FuncStr`.

*When you find the option type is `types.FuncStr`, you should notice it. :)*

## How to insert the pure JS function

- Use `opts.FuncOpts` or `opts.FuncStripCommentsOpts` to create your JS string.
- Use `myChart.AddJSFuncStrs()` to add the JS function into your charts.

For example:

```go
     actionWithEchartsInstance := `
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
pie.AddJSFuncStrs(opts.FuncOpts(actionWithEchartsInstance)
```

> BTW: use `const` to define a pure function string can do the builtin convert directly.

## How to get echarts instance

!> Especially, you may notice the special var `%MY_ECHARTS%` in the example above.

The `%MY_ECHARTS%` is the placeholder of the `echarts instance` in go-echarts.  
It will inject each of the real `echarts instance` of the chart into your functions.  
We believe it could bring more convenient to you when you want to do some advance operates on charts.  
Normally, you don't need it when you don't need do functions on chart, since go-echarts wants keep it always simple.

---

# Where to find options you need

go-echarts' options are almost align to [echarts' official options](https://echarts.apache.org/en/option.html#title).  
If you want to find the options of your specific charts, all in the echarts options list.  
And if you find some options go-echarts hasn't supported yet, kindly create a PR or raise an option request.
---

# Where to find custom echarts resources

Currently, we support echarts' version is `v5.4.3` and
the compatible version `v4.3` for those 3d charts and customized charts, such as word-cloud.  
All the resources are hosted in [go-echarts-assets](https://github.com/go-echarts/go-echarts-assets) by go-echarts
itself.

If you want to use your custom echarts resources.
You could find in the [jsdelivr of echarts](https://www.jsdelivr.com/package/npm/echarts) or other CDN providers.  
Besides, you can set your own assets host as well.

We support 2 ways to customize the echarts resources.

- Only change the `Asserts Host`, which needs the echarts resources naming `echarts.min.js`
- Clean all the preset assets and add custom resources all by yourself.



