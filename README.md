<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52535979-c0d0e680-2d8f-11e9-85c8-2e9f659e7c6f.png" width=300 height=300 />
</p>

<h1 align="center">go-echarts</h1>
<p align="center">
    <em>🎨 The adorable charts library for Golang.</em>
</p>
<p align="center">
    <a href="https://travis-ci.org/go-echarts/go-echarts">
        <img src="https://travis-ci.org/go-echarts/go-echarts.svg?branch=master" alt="Build Status">
    </a>
    <a href="https://goreportcard.com/report/github.com/go-echarts/go-echarts">
        <img src="https://goreportcard.com/badge/github.com/go-echarts/go-echarts" alt="Go Report Card">
    </a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-brightgreen.svg" alt="MIT License">
    </a>
        <a href="https://godoc.org/github.com/go-echarts/go-echarts">
        <img src="https://godoc.org/github.com/go-echarts/go-echarts?status.svg" alt="GoDoc">
    </a>
</p>

> If a language can be used to build web scrapers, it definitely needs to provide a graceful data visualizing library. --- by dongdong.

In the Golang ecosystem, there are not many choices for data visualizing libraries. The development of [go-echarts](https://github.com/go-echarts/go-echarts) aims to provide a simple yet powerful data visualizing library for Golang. [Echarts](https://echarts.apache.org/) is an outstanding charting and visualizing library opensourced by Baidu, it supports adorable chart types and various interactive features. There are many language bindings for Echarts, for example, [pyecharts](https://github.com/pyecharts/pyecharts). go-echarts learns from pyecharts and has evolved a lot.

[中文 README](README_CN.md)

### 🔰 Installation

Classic way to get go-echarts

```shell
$ go get -u github.com/go-echarts/go-echarts/...
```

Use gomod style

```shell
require "github.com/go-echarts/go-echarts/v2"
```

### ⏳ Version

The go-echarts project is being developed under v2 version and the codebase is on the master branch now.

v1 and v2 is incompatible which is mean that you cannot upgrade go-echarts from v1 to v2 smoothly. But I think is worth trying the new version.

### ✨ Features

* Clean and comprehensive API.
* Visualize your data in 25+ different ways.
* Highly configurable chart options.
* Detailed documentation and a rich collection of examples.
* Visualize your geographical data with 400+ maps.

### 📝 Usage

It's easy to get started with go-echarts. In this example, we create a simple bar chart with only a few lines of code.

```golang
package examples

import (
	"math/rand"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(300)})
	}
	return items
}

func main() {
    // create a new bar instance
    bar := charts.NewBar()
    
    // set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Bar-basic-example",
			Subtitle: "This is the subtitle.",
		}),
	)

    // Put some data in instance
	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
        AddSeries("Category B", generateBarItems()).
        Render()
}
```

And the generated bar.html is rendered as below. Isn't that cool！

![](https://user-images.githubusercontent.com/19553554/97107442-85f91880-1702-11eb-8b73-d0d8daedf549.png)

### 🔖 Gallery

<div align="center">
<img src="https://user-images.githubusercontent.com/19553554/52197440-843a5200-289a-11e9-8601-3ce8d945b04a.gif" width="33%" alt="bar"/>
<img src="https://user-images.githubusercontent.com/19553554/52360729-ad640980-2a77-11e9-84e2-feff7e11aea5.gif" width="33%" alt="boxplot"/>
<img src="https://user-images.githubusercontent.com/19553554/52535290-4b611800-2d87-11e9-8bf2-b43a54a3bda8.png" width="33%" alt="effectScatter"/>
<img src="https://user-images.githubusercontent.com/19553554/52332816-ac5eb800-2a36-11e9-8227-3538976f447d.gif" width="33%" alt="funnel"/>
<img src="https://user-images.githubusercontent.com/19553554/52332988-0b243180-2a37-11e9-9db8-eb6b8c86a0de.png" width="33%" alt="gague"/>
<img src="https://user-images.githubusercontent.com/19553554/52344575-133f9980-2a56-11e9-93e0-568e484936ce.gif" width="33%" alt="geo"/>
<img src="https://user-images.githubusercontent.com/19553554/52727805-f7f20280-2ff0-11e9-91ab-cd99848e3127.gif" width="33%" alt="graph"/>
<img src="https://user-images.githubusercontent.com/19553554/52345115-6534ef00-2a57-11e9-80cd-9cbfed252139.gif" width="33%" alt="heatmap"/>
<img src="https://user-images.githubusercontent.com/19553554/52345490-4a16af00-2a58-11e9-9b43-7bbc86aa05b6.gif" width="33%" alt="kline"/>
<img src="https://user-images.githubusercontent.com/19553554/52346064-b7770f80-2a59-11e9-9e03-6dae3a8c637d.gif" width="33%" alt="line"/>
<img src="https://user-images.githubusercontent.com/19553554/52347117-248ba480-2a5c-11e9-8402-5a94054dca50.gif" width="33%" alt="liquid"/>
<img src="https://user-images.githubusercontent.com/19553554/52347915-0a52c600-2a5e-11e9-8039-41268238576c.gif" width="33%" alt="map"/>
<img src="https://user-images.githubusercontent.com/19553554/52535013-e48e2f80-2d83-11e9-8886-ac0d2122d6af.png" width="33%" alt="parallel"/>
<img src="https://user-images.githubusercontent.com/19553554/52348202-bb596080-2a5e-11e9-84a7-60732be0743a.gif" width="33%" alt="pie"/>
<img src="https://user-images.githubusercontent.com/19553554/52533994-932b7380-2d76-11e9-93b4-0de3132eb941.gif" width="33%" alt="radar"/>
<img src="https://user-images.githubusercontent.com/19553554/52348431-420e3d80-2a5f-11e9-8cab-7b415592dc77.gif" width="33%" alt="scatter"/>
<img src="https://user-images.githubusercontent.com/19553554/52348737-01fb8a80-2a60-11e9-94ac-dacbd7b58811.png" width="33%" alt="wordCloud"/>
<img src="https://user-images.githubusercontent.com/19553554/52433989-4f075b80-2b49-11e9-9979-ef32c2d17c96.gif" width="33%" alt="bar3D"/>
<img src="https://user-images.githubusercontent.com/19553554/52464826-4baab900-2bb7-11e9-8299-776f5ee43670.gif" width="33%" alt="line3D"/>
<img src="https://user-images.githubusercontent.com/19553554/52802261-8d0cfe00-30ba-11e9-8ae7-ae0773770a59.gif" width="33%" alt="sankey"/>
<img src="https://user-images.githubusercontent.com/19553554/52464647-aee81b80-2bb6-11e9-864e-c544392e523a.gif" width="33%" alt="scatter3D"/>
<img src="https://user-images.githubusercontent.com/19553554/52465183-a55fb300-2bb8-11e9-8c10-4519c4e3f758.gif" width="33%" alt="surface3D"/>
<img src="https://user-images.githubusercontent.com/19553554/52798246-7ebae400-30b2-11e9-8489-6c10339c3429.gif" width="33%" alt="themeRiver"/>
<img src="https://user-images.githubusercontent.com/19553554/52349544-c2ce3900-2a61-11e9-82af-28aaaaae0d67.gif" width="33%" alt="overlap"/>
</div>

For more information, please refer to [go-echarts/examples](https://github.com/go-echarts/examples).

### 📃 License

MIT [©chenjiandongx](https://github.com/chenjiandongx)
