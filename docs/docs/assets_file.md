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

## 编译须知

在将目标程序编译成可执行文件的时候，需要先执行以下操作（一次即可）。

原因是调试期间 [packr](https://github.com/gobuffalo/packr) 可以在开发环境中找到 `go-echarts` 的 `datasets` 和 `templates` 的路径，但是编译后的可执行文件，在其他无 `go-echarts` 的环境中，将会报错

```shell
FindFirstFile src\$GOPATH\src\github.com\chenjiandongx\go-echarts\templates\routers.html: The system cannot find the path specified.
```

因此使用 [packr](https://github.com/gobuffalo/packr#building-a-binary-the-hard-way) 提供的打包方式，将 `datasets` 和 `templates` 中的静态文件打包进可执行文件中，确保 `go-echarts` 的正常使用。

Linux/MacOS
```bash
$ cd $GOPATH/src/github.com/chenjiandongx/go-echarts
$ ./doPackr.sh
```

Windows
```dos
$ cd %GOPATH%\src\github.com\chenjiandongx\go-echarts
$ doPackr.bat
```

以上操作将会在 `%GOPATH%\src\github.com\chenjiandongx\go-echarts\datasets` 及 `%GOPATH%\src\github.com\chenjiandongx\go-echarts\templates` 目录下生成 `*-packr.go` 文件，如 `a_datasets-packr`，`a_templates-packr.go`，用于替换原有的静态资源文件，相当于将静态文件的内容存储到 golang 中。
