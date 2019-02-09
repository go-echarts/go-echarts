---
id: assets_file
title: 静态文件
sidebar_label: 静态文件
---

## 远程引用

go-echarts 没有自带自身所需要引用的 JS/CSS 文件，均是通过远程引用，引用的地址为 http://chenjiandongx.com/go-echarts-assets/assets/，这是通过我的一个项目 [go-echarts-assets](https://github.com/chenjiandongx/go-echarts-assets) 实现的。

任何想使用自己远程 Host 的开发者都可以 Fork [go-echarts-assets](https://github.com/chenjiandongx/go-echarts-assets) 这个项目，使用 Github 的话可以直接使用 Github Page 部署，其他情况使用自己的服务器部署也行。

## 本地引用

下面以本地开启服务器为例

```shell
$ git clone https://github.com/chenjiandongx/go-echarts-assets.git
$ cd go-echarts-assets
$ go run main.go
```

mian.go
```go
package main

import "net/http"

func main() {
    fs := http.FileServer(http.Dir("."))
    http.Handle("/", fs)

    http.ListenAndServe(":8083", nil)
}
```

启动 main.go 将会启动在本地开启一个 8083 端口的服务器，此时你需要修改 `charts.InitOpts` 配置项

```go
// `快速开始` 中的示例
func main() {
    bar := charts.NewBar()
    bar.SetGlobalOptions(
        charts.TitleOpts{Title: "Bar-示例图"}, 
        charts.ToolboxOpts{Show: true},
        // AssetsHost 设置远程 Host
        charts.InitOpts{AssetsHost: "http://127.0.0.1::8083/assets/"},
    )
    bar.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
        AddYAxis("商家B", randInt())
    f, err := os.Create("bar.html")
    if err != nil {
        log.Println(err)
    }
    bar.Render(f)
}
```

在一次 `Render()` 方法中只需要配置一次 `charts.InitOpts` 即可，如在 Page 的 Add() 方法中，只要其中一个图表能正确引用到的话那所有图表也就都能正常显示，因为他们都在同一个 HTML 文件中。
