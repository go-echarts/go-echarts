# go-eharts

## BaseOpts
> 所有图表都拥有的基本配置项
```go
type BaseOpts struct {
    InitOpts                   // 图形初始化配置项
    LegendOpts                 // 图例组件配置项
    TooltipOpts                // 提示框组件配置项
    ToolboxOpts                // 工具箱组件配置项
    TitleOpts                  // 标题组件配置项
    AssetsOpts                 // 静态资源配置项
    Colors            []string // 全局颜色列表
    appendColor       []string // 追加全局颜色列表
    HTTPRouters                // 路由列表
    DataZoomOptsList           // 区域缩放组件配置项列表
    VisualMapOptsList          // 视觉映射组件配置项列表
    GeoOpts                    // 地理坐标系组件配置项

    JSFunctions      // JS 函数列表
    HasXYAxis   bool // 图形是否拥有 XY 轴
}
```

## SeriesOpts

## 基本图表

### Bar

### EffectScatter

### Funnel

### Gauge

### Geo

### HeatMap

### Kline

### Line

### Liquid

### Map

### Pie

### Scatter

### WorldCloud


## 组合图表

### Grid

### Page

### Timeline
