---
id: map_file
title: 地图定制
sidebar_label: 地图定制
---

## 地图资源
> go-echarts 提供了全国 370+ 城市以及 28+ 省份的地图

具体数据映射关系保存在 [datasets/map_filename.json](https://github.com/chenjiandongx/go-echarts/blob/master/datasets/map_filename.json) 中，JSON 数据格式
```
{
  "china": "china",
  "world": "world",
  "广东": "guangdong",
  "安徽": "anhui",
  "福建": "fujian",
  "甘肃": "gansu",
  "广西": "guangxi",
  "贵州": "guizhou",
  "海南": "hainan",
  ....
}
```

同样的，JSON 格式的数据会在编译时被加载进内容中，可通过 `MapFileNames` 引用和修改
```go
import "github.com/chenjiandongx/go-echarts/datasets"

// 打印
fmt.Println(datasets.MapFileNames)
map[山东:shandong 延边朝鲜族自治州:ji2_lin2_yan2_bian1_zhao1_xian1_zu2_zi4_zhi4_zhou1 

// 判断一个地图是否存在
region, ok := datasets.MapFileNames["北京"]
if ok {
    fmt.Println(region)
}
```

## 定制地图

由于 Echarts 已经关闭了地图定制功能

![](https://user-images.githubusercontent.com/19553554/52518306-b295a400-2c83-11e9-986b-f2e76fc2621d.png)

所以现在已经不支持自己定制地图了，但是 go-echarts 本身提供的 400+ 地图应该能够满足开发者的日常开发需要了！
