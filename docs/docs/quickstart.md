---
id: quickstart
title: 快速开始
sidebar_label: 快速开始
---

### 安装

```shell
$ go get github.com/chenjiandongx/go-echarts/...
```

### 启航

```go
// example.go
package main

import (
    "log"
    "math/rand"
    "os"
    "time"

    "github.com/chenjiandongx/go-echarts/charts"
)

var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var seed = rand.NewSource(time.Now().UnixNano())

func randInt() []int {
    cnt := len(nameItems)
    r := make([]int, 0)
    for i := 0; i < cnt; i++ {
        r = append(r, int(seed.Int63()) % 50)
    }
    return r
}

func main() {
    bar := charts.NewBar()
    bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
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

运行代码

```shell
$ go run example.go
```

程序会生成 `bar.html` 文件在根目录下，使用浏览器打开。Surprise!

![](https://user-images.githubusercontent.com/19553554/52197440-843a5200-289a-11e9-8601-3ce8d945b04a.gif)

当然，你也可以使用 `net/http`

```go
func handler(w http.ResponseWriter, _ *http.Request) {
    bar := charts.NewBar()
    bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
    bar.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
        AddYAxis("商家B", randInt())
    f, err := os.Create("bar.html")
    if err != nil {
        log.Println(err)
    }
    bar.Render(w, f) // Render 可接收多个 io.Writer 接口
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8081", nil)
}
```

启动服务器

![](https://user-images.githubusercontent.com/19553554/52198001-6a9a0a00-289c-11e9-9daa-e6637b4ae527.png)
