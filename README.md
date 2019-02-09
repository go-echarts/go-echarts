<p align="center">
<img src="https://user-images.githubusercontent.com/19553554/52387794-8680f400-2ac6-11e9-8f5e-cf7821f09a03.png" width=300 height=300 />
</p>

<h1 align="center">go-echarts</h1>
<p align="center">
    <em>🎨 The adorable charts library for Golang</em>
</p>
<p align="center">
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-brightgreen.svg" alt="MIT License">
    </a>
</p>

在 Golang 这门语言中，目前数据可视化的第三方库还是特别少，[go-echarts](https://github.com/chenjiandongx/go-echarts) 的开发就是为了填补这部分的空隙。[Echarts](https://echarts.baidu.com) 是百度开源的非常优秀的可视化图表库，凭借着良好的交互性，精巧的图表设计，得到了众多开发者的认可。也有其他语言为其实现了相应语言版本的接口，如 Python 的 [pyecharts](https://github.com/pyecharts/pyecharts)，go-echarts 也是借鉴了 pyecharts 的一些设计思想。


### 🔰 安装

```shell
$ go get -u github.com/chenjiandongx/go-echarts/...
```

### ✨ 特性

* 简洁的 API 设计，使用如丝滑般流畅
* 囊括了 20+ 种常见图表，应有尽有
* 高度灵活的配置项，可轻松搭配出精美的图表
* 详细的文档和示例，帮助开发者更快的上手项目
* 多达 400+ 地图，为地理数据可视化提供强有力的支持

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

想了解更多文档的内容请移步至 [...]()

### 📃 LICENSE

MIT [©chenjiandongx](https://github.com/chenjiandongx)
