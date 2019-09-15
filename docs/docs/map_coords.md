---
id: map_coords
title: 地理坐标
sidebar_label: 地理坐标
---

## 数据格式

go-echarts 内置了一些常用的城市地理坐标数据，这些数据保存 [coordinates.json](https://github.com/go-echarts/go-echarts/blob/master/datasets/coordinates.json) 文件中。数据在编译程序时会被加载进内存中。

JSON 格式
```
{<name>: [<longitude>, <latitude>]}

// 如
{
    "阿城": [126.58, 45.32],
    "阿克苏": [80.19, 41.09],
    "阿勒泰": [88.12, 47.50],
}
```

## 定制坐标

数据加载进内存后可通过 `Coordinates` 引用和修改

```go
import "github.com/go-echarts/go-echarts/datasets"

// 打印
fmt.Println(datasets.Coordinates)
map[元氏县:[114.52 37.75] 万宁市:[110.4 18.8] 西夏区:[106.18 38.48] 吉水县:[115.13 27.22]...

// 修改
datasets.Coordinates["A城"] = [2]float32{11.11, 12.12}

// 判断一个地区的坐标是否存在
coord, ok := datasets.Coordinates["北京"]
if ok {
    fmt.Println(coord)
}
```

## 坐标来源

坐标精确数据来源可参考 [pyecharts/geo-region-coords](https://github.com/pyecharts/geo-region-coords)，从 [coords.json](https://github.com/pyecharts/geo-region-coords/blob/master/coords.json) 中检索出相应的地理坐标位置后有两种方式更新坐标数据。

1. 更新 `datasets.Coordinates` 变量。
2. 修改 [coordinates.json](https://github.com/go-echarts/go-echarts/blob/master/datasets/coordinates.json) 文件，不过这种方式在升级 go-echarts 的时候数据会被抹除！
