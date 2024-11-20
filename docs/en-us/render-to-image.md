# Render to Image

## snapshot-chromedp

Currently, we support render to image
via [go-echarts/snapshot-chromedp](https://github.com/go-echarts/snapshot-chromedp) approach.

How does the trick?  
It uses the headless browser to generate the HTML and capture the chart base64 data via echarts api.
Which means it actually generated the HTML, and then, snapshot the chart base64 image.
By default, we will delete the mid-product, the HTML.
It literally likely gets the image directly.

The seamless function is

```go
    render.MakeChartSnapshot(myChart.RenderContent(), "my-chart.png")
```

Furthermore, if you want to get it all by yourself, a fully `SnapshotConfig` config here.

```go
type SnapshotConfig struct {
    // RenderContent the content bytes of charts after rendered
    RenderContent []byte
    // Path the path to save image
    Path string
    // FileName image name
    FileName string
    // Suffix image format, png, jpeg
    Suffix string
    // Quality the generated image quality, aka pixelRatio
    Quality int
    // KeepHtml whether keep the generated html also, default false
    KeepHtml bool
    // HtmlPath where to keep the generated html, default same to image path
    HtmlPath string
    // Timeout  the timeout config
    Timeout time.Duration
    // MultiCharts Only enable it when you have multi charts in the single page, better to set larger quality
    MultiCharts bool
}
```

More details see [go-echarts/snapshot-chromedp](https://github.com/go-echarts/snapshot-chromedp).



