---
id: gauge
title: Gauge（仪表盘）
sidebar_label: Gauge（仪表盘）
---

## API
```go
// 实例化图表
func NewGauge(routers ...HTTPRouter) *Gauge {}
// 新增数据及配置项
func Add(name string, data map[string]interface{}, options ...seriesOptser) *Gauge {}
```

## Demo

### Gauge-示例图
```go
gauge := charts.NewGauge()
gauge.SetGlobalOptions(charts.TitleOpts{Title: "Gauge-示例图"})
m := make(map[string]interface{})
m["工作进度"] = rand.Intn(50)
gauge.Add("gauge", m)
```
![](https://user-images.githubusercontent.com/19553554/52332988-0b243180-2a37-11e9-9db8-eb6b8c86a0de.png)


### Gauge-定时器
```go
gauge := charts.NewGauge()

m := make(map[string]interface{})
m["工作进度"] = rand.Intn(50)
gauge.Add("gauge1", m)
gauge.SetGlobalOptions(charts.TitleOpts{Title: "Gauge-定时器"})
fn := fmt.Sprintf(`setInterval(function () {
        option_%s.series[0].data[0].value = (Math.random() * 100).toFixed(2) - 0;
        myChart_%s.setOption(option_%s, true);
    }, 2000);`, gauge.ChartID, gauge.ChartID, gauge.ChartID)
gauge.AddJSFuncs(fn)
```
![](https://user-images.githubusercontent.com/19553554/52333017-24c57900-2a37-11e9-8027-c7a623932dd4.gif)
