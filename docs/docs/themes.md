---
id: themes
title: 主题定制
sidebar_label: 主题定制
---

## 主题种类

> go-echarts 内置了多种主题，可通过 `common.ThemeType` 查看主题种类

下面是示例代码
```go
package main

import (
    "log"
    "math/rand"
    "os"
    "time"

    "github.com/go-echarts/go-echarts/charts"
)

var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var seed = rand.NewSource(time.Now().UnixNano())

func randInt() []int {
    cnt := len(nameItems)
    r := make([]int, 0)
    for i := 0; i < cnt; i++ {
        r = append(r, int(seed.Int63())%50)
    }
    return r
}

func main() {
    bar := charts.NewBar()
    bar.SetGlobalOptions(
        charts.TitleOpts{Title: "Theme-Walden"},
        charts.InitOpts{Theme: charts.ThemeType.Walden},
        charts.ToolboxOpts{Show: true},
    )
    bar.AddXAxis(nameItems).
        AddYAxis("商家A", randInt()).
        AddYAxis("商家B", randInt()).
        AddYAxis("商家C", randInt()).
        AddYAxis("商家D", randInt())
    f, err := os.Create("bar.html")
    if err != nil {
        log.Println(err)
    }
    bar.Render(f)
}
```

### Default

![default](https://user-images.githubusercontent.com/19553554/52496989-dddeab80-2c0f-11e9-93a6-fe0371fde99d.png)

### Chalk

![chalk](https://user-images.githubusercontent.com/19553554/52496994-dfa86f00-2c0f-11e9-88c2-8bd7f05027cd.png)

### Essos

![essos](https://user-images.githubusercontent.com/19553554/52496997-e0410580-2c0f-11e9-8a47-7d3025184217.png)

### Infographic

![infographic](https://user-images.githubusercontent.com/19553554/52497000-e0d99c00-2c0f-11e9-8f5d-19dccf8c4a7c.png)

### Macarons

![macarons](https://user-images.githubusercontent.com/19553554/52496991-de774200-2c0f-11e9-8c2f-1d25f2d4df6e.png)

### PurplePassion

![purplepassion](https://user-images.githubusercontent.com/19553554/52497002-e1723280-2c0f-11e9-8022-d74923278bcc.png)

### Roma

![roms](https://user-images.githubusercontent.com/19553554/52496990-de774200-2c0f-11e9-8b12-3bc2a87523cb.png)

### Romantic

![romantic](https://user-images.githubusercontent.com/19553554/52496993-df0fd880-2c0f-11e9-9f63-7b374826ae6c.png)

### Shine

![shine](https://user-images.githubusercontent.com/19553554/52496996-e0410580-2c0f-11e9-9626-de5451661758.png)

### Vintage

![vintage](https://user-images.githubusercontent.com/19553554/52496999-e0410580-2c0f-11e9-8999-d4988c15adba.png)

### Walden

![walden](https://user-images.githubusercontent.com/19553554/52497001-e1723280-2c0f-11e9-924e-3fb0505d42c0.png)

### Westeros

![westeros](https://user-images.githubusercontent.com/19553554/52497003-e20ac900-2c0f-11e9-9b6b-b94ee89a05e6.png)

### Wonderland

![wonderland](https://user-images.githubusercontent.com/19553554/52497004-e20ac900-2c0f-11e9-8ec7-75525ab9ef37.png)

## 定制主题

Echarts 提供了 [主题构建工具](http://echarts.baidu.com/theme-builder/)，你可以从中构建喜欢的主题，如 myTheme.js。然后将 myTheme.js 放入到 `go-echarts-assets/assets/themes` 中，然后启动自己的本地服务器或远程服务器，这部分内容可参考 [docs/assets_file](docs/assets_file)。