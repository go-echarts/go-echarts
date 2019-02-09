---
id: javascript
title: 原生 Javascript
sidebar_label: 原生 Javascript
---

## go-echarts 局限性

毕竟 Echarts 是原生 JS 项目，使用 JS 来开发肯定是最舒服的，可定制性最高的。go-echarts 只是做了 **语言转换**，将其配置项转换为 JSON 格式并将图渲染出来。但我也尽量的扩展 go-echarts 的灵活性，现在它已支持传入原生的 JS 函数，通过字符串的形式。

JS 函数用得最多的应该是 Formatter 配置项，因为原生的字符串模板并不足以满足开发需求。

## JS 函数示例

### 在 Map/Geo 图中

```go
var geoFormatter = `function (params) {
		return params.name + ' : ' + params.value[2];
}`

// 然后在 charts.LabelTextOpts, charts.TooltipOpts 等有用 Formatter 配置项中使用
// 必须使用 charts.FuncOpts(fn) 进行一层包装，会清除掉 \n \t 等特殊字符
charts.LabelTextOpts{Show: true, Formatter: charts.FuncOpts(fn)}
```

### 在 Gauge 图中
```go
gauge := charts.NewGauge()

m := make(map[string]interface{})
m["工作进度"] = rand.Intn(50)
gauge.Add("gauge1", m)
gauge.SetGlobalOptions(charts.TitleOpts{Title: "Gauge-定时器"})
// 使用了 JS 函数
fn := fmt.Sprintf(`setInterval(function () {
        option_%s.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
        myChart_%s.setOption(option_%s, true);
    }, 2000);`, gauge.ChartID, gauge.ChartID, gauge.ChartID)
gauge.AddJSFuncs(fn)
```

这里有一点需要特别提醒，在 go-echarts 中，每一个图表示例都有一个对应的的 ID，是图表的唯一标识，这是为了避免了在 Page 多图中渲染出现问题。所以当需要对指定图表使用 JS 函数时，需要在代码中指定 ID。如上面代码中的 `option_%s.series[0].data[0].value` 的 %s，后面用 `gauge.ChartID` 格式化了。