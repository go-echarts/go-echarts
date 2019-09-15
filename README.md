<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52535979-c0d0e680-2d8f-11e9-85c8-2e9f659e7c6f.png" width=300 height=300 />
</p>

<h1 align="center">go-echarts</h1>
<p align="center">
    <em>🎨 The adorable charts library for Golang.</em>
</p>
<p align="center">
    <a href="https://travis-ci.org/chenjiandongx/go-echarts">
        <img src="https://travis-ci.org/chenjiandongx/go-echarts.svg?branch=master" alt="Build Status">
    </a>
    <a href="https://ci.appveyor.com/project/chenjiandongx/go-echarts">
        <img src="https://ci.appveyor.com/api/projects/status/kdxi0s1nc1t6dqn0?svg=true" alt="Build Status">
    </a>
    <a href="https://goreportcard.com/report/github.com/chenjiandongx/go-echarts">
        <img src="https://goreportcard.com/badge/github.com/chenjiandongx/go-echarts" alt="Go Report Card">
    </a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-brightgreen.svg" alt="MIT License">
    </a>
        <a href="https://godoc.org/github.com/chenjiandongx/go-echarts">
        <img src="https://godoc.org/github.com/chenjiandongx/go-echarts?status.svg" alt="GoDoc">
    </a>
</p>

> If a language can be used to build web scrapers, it definitely needs to provide a graceful data visualizing library. --- by chenjiandongx

In the Golang ecosystem, there are not many choices for data visualizing libraries. The development of [go-echarts](https://github.com/chenjiandongx/go-echarts) aims to provide a simple yet powerful data visualizing library for Golang. [Echarts](https://echarts.baidu.com) is an outstanding charting and visualizing library opensourced by Baidu, it supports adorable chart types and various interactive features. There are many language bindings for Echarts, for example, [pyecharts](https://github.com/pyecharts/pyecharts). go-echarts learns from pyecharts and has evolved a lot.

[中文 README](README_CN.md)

### 🔰 Installation

```shell
$ go get -u github.com/go-echarts/go-echarts/...
```

### ✨ Features

* Clean and comprehensive API.
* Visualize your data in 25+ different ways.
* Highly configurable chart options.
* Detailed documentation and a rich collection of examples.
* Visualize your geographical data with 400+ maps.

### 📝 Usage

It's easy to get started with go-echarts. In this example, we create a simple bar chart with only a few lines of code.

<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52524229-bf42e800-2cd5-11e9-9eb8-47d8e3f4052b.png" width="80%" height="80%" />
</p>

And the generated bar.html is rendered as below. Isn't that cool！

<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52524101-34152280-2cd4-11e9-87c6-bbf5e388fe23.png" width="80%" height="80%" />
</p>

Of course we can also start a listening web server with `net/http`.

<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52524272-2c567d80-2cd6-11e9-8a73-29ba059b8bb5.png"
 width="80%" height="80%" />
</p>

Now visit http://localhost:8081 in your browser and you'll see the same bar chart.

### 🔖 Gallery

<div align="center">
<img src="https://user-images.githubusercontent.com/19553554/52197440-843a5200-289a-11e9-8601-3ce8d945b04a.gif" width="33%" height="33%" alt="bar"/>
<img src="https://user-images.githubusercontent.com/19553554/52360729-ad640980-2a77-11e9-84e2-feff7e11aea5.gif" width="33%" height="33%" alt="boxplot"/>
<img src="https://user-images.githubusercontent.com/19553554/52535290-4b611800-2d87-11e9-8bf2-b43a54a3bda8.png" width="33%" height="33%" alt="effectScatter"/>
<img src="https://user-images.githubusercontent.com/19553554/52332816-ac5eb800-2a36-11e9-8227-3538976f447d.gif" width="33%" height="33%" alt="funnel"/>
<img src="https://user-images.githubusercontent.com/19553554/52332988-0b243180-2a37-11e9-9db8-eb6b8c86a0de.png" width="33%" height="33%" alt="gague"/>
<img src="https://user-images.githubusercontent.com/19553554/52344575-133f9980-2a56-11e9-93e0-568e484936ce.gif" width="33%" height="33%" alt="geo"/>
<img src="https://user-images.githubusercontent.com/19553554/52727805-f7f20280-2ff0-11e9-91ab-cd99848e3127.gif" width="33%" height="33%" alt="graph"/>
<img src="https://user-images.githubusercontent.com/19553554/52345115-6534ef00-2a57-11e9-80cd-9cbfed252139.gif" width="33%" height="33%" alt="heatmap"/>
<img src="https://user-images.githubusercontent.com/19553554/52345490-4a16af00-2a58-11e9-9b43-7bbc86aa05b6.gif" width="33%" height="33%" alt="kline"/>
<img src="https://user-images.githubusercontent.com/19553554/52346064-b7770f80-2a59-11e9-9e03-6dae3a8c637d.gif" width="33%" height="33%" alt="line"/>
<img src="https://user-images.githubusercontent.com/19553554/52347117-248ba480-2a5c-11e9-8402-5a94054dca50.gif" width="33%" height="33%" alt="liquid"/>
<img src="https://user-images.githubusercontent.com/19553554/52347915-0a52c600-2a5e-11e9-8039-41268238576c.gif" width="33%" height="33%" alt="map"/>
<img src="https://user-images.githubusercontent.com/19553554/52535013-e48e2f80-2d83-11e9-8886-ac0d2122d6af.png" width="33%" height="33%" alt="parallel"/>
<img src="https://user-images.githubusercontent.com/19553554/52348202-bb596080-2a5e-11e9-84a7-60732be0743a.gif" width="33%" height="33%" alt="pie"/>
<img src="https://user-images.githubusercontent.com/19553554/52533994-932b7380-2d76-11e9-93b4-0de3132eb941.gif" width="33%" height="33%" alt="radar"/>
<img src="https://user-images.githubusercontent.com/19553554/52348431-420e3d80-2a5f-11e9-8cab-7b415592dc77.gif" width="33%" height="33%" alt="scatter"/>
<img src="https://user-images.githubusercontent.com/19553554/52348737-01fb8a80-2a60-11e9-94ac-dacbd7b58811.png" width="33%" height="33%" alt="wordCloud"/>
<img src="https://user-images.githubusercontent.com/19553554/52433989-4f075b80-2b49-11e9-9979-ef32c2d17c96.gif" width="33%" height="33%" alt="bar3D"/>
<img src="https://user-images.githubusercontent.com/19553554/52464826-4baab900-2bb7-11e9-8299-776f5ee43670.gif" width="33%" height="33%" alt="line3D"/>
<img src="https://user-images.githubusercontent.com/19553554/52802261-8d0cfe00-30ba-11e9-8ae7-ae0773770a59.gif" width="33%" height="33%" alt="sankey"/>
<img src="https://user-images.githubusercontent.com/19553554/52464647-aee81b80-2bb6-11e9-864e-c544392e523a.gif" width="33%" height="33%" alt="scatter3D"/>
<img src="https://user-images.githubusercontent.com/19553554/52465183-a55fb300-2bb8-11e9-8c10-4519c4e3f758.gif" width="33%" height="33%" alt="surface3D"/>
<img src="https://user-images.githubusercontent.com/19553554/52798246-7ebae400-30b2-11e9-8489-6c10339c3429.gif" width="33%" height="33%" alt="themeRiver"/>
<img src="https://user-images.githubusercontent.com/19553554/52349544-c2ce3900-2a61-11e9-82af-28aaaaae0d67.gif" width="33%" height="33%" alt="overlap"/>
</div>

Run the demo program under `example/` to view all the above chart examples.
```shell
$ cd your/gopath/src/github.com/go-echarts/go-echarts/example
$ go build .
$ ./example
```

For more information, please refer to [go-echarts.chenjiandongx.com](http://go-echarts.chenjiandongx.com).

### 📃 LICENSE

MIT [©chenjiandongx](https://github.com/chenjiandongx)
