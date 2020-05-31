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
    <a href="https://ci.appveyor.com/project/go-echarts/go-echarts">
        <img src="https://ci.appveyor.com/api/projects/status/kdxi0s1nc1t6dqn0?svg=true" alt="Build Status">
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

> 如果一门语言可以用来写爬虫，那么它就需要一个优雅的数据可视化库。 --- by chenjiandongx

在 Golang 这门语言中，目前数据可视化的第三方库还是特别少，[go-echarts](https://github.com/chenjiandongx/go-echarts) 的开发就是为了填补这部分的空隙。[Echarts](https://echarts.baidu.com) 是百度开源的非常优秀的可视化图表库，凭借着良好的交互性，精巧的图表设计，得到了众多开发者的认可。也有其他语言为其实现了相应语言版本的接口，如 Python 的 [pyecharts](https://github.com/pyecharts/pyecharts)，go-echarts 也是借鉴了 pyecharts 的一些设计思想。


### 🔰 安装

```shell
$ go get -u github.com/go-echarts/go-echarts/...
```

### ✨ 特性

* 简洁的 API 设计，使用如丝滑般流畅
* 囊括了 **25+** 种常见图表，应有尽有
* 高度灵活的配置项，可轻松搭配出精美的图表
* 详细的文档和示例，帮助开发者更快的上手项目
* 多达 **400+** 地图，为地理数据可视化提供强有力的支持

### 📝 使用

仅需要几行核心代码就可画出美观的图表

<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52524229-bf42e800-2cd5-11e9-9eb8-47d8e3f4052b.png" width="80%" height="80%" />
</p>

生成的 bar.html 是这样的。Cool！

<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52524101-34152280-2cd4-11e9-87c6-bbf5e388fe23.png" width="80%" height="80%" />
</p>

当然你也可以使用更加 `golang` 的方式，利用 `net/http`

<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52524272-2c567d80-2cd6-11e9-8a73-29ba059b8bb5.png"
 width="80%" height="80%" />
</p>

打开浏览器访问 http://localhost:8081 也可以看到同样的效果！

### 🔖 Demo

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

运行 example/main.go 可预览所有示例
```shell
$ cd your/gopath/src/github.com/go-echarts/go-echarts/example
$ go build .
$ ./example
```

了解更多文档的内容请访问 [go-echarts.github.io/go-echarts/](https://go-echarts.github.io/go-echarts/)

### 📃 LICENSE

MIT [©chenjiandongx](https://github.com/chenjiandongx)
