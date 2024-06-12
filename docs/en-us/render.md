# Render

Render is an interface for render charts, alias the `Renderer`.

```go
// Renderer
// Any kinds of charts have their render implementation, and
// you can define your own render logic easily.
type Renderer interface {
    Render(w io.Writer) error
    RenderContent() []byte
    RenderSnippet() ChartSnippet
}
```

## `Render(w io.Writer) error`

The `Render(w io.Writer) error` function aims to generate chart to target `io.Writer`.  
Underlying, it depends on the `RenderContent() []byte` to generate the chart bytes streaming.

## `RenderContent() []byte`

The `RenderContent() []byte` uses templates to generate the chart bytes streaming.


## `RenderSnippet() ChartSnippet`

The `RenderSnippet() ChartSnippet` function can extract the element(chart container), the script
and the options from a chart.

!> It only works for Chart, Page is unsupported.

e.g.
> Full Code: [examples/renderer.go](https://github.com/go-echarts/examples/blob/master/examples/renderer.go)
```go
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Bar chart for Snippets",
	}))

	bar.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())

	// pure extracted snippets
	chartSnippet := bar.RenderSnippet()

	tmpl := "{{.Element}} {{.Script}} {{.Option}}"
	t := template.New("snippet")
	t, err := t.Parse(tmpl)
	if err != nil {
		panic(err)
	}

	data := struct {
		Element template.HTML
		Script  template.HTML
		Option  template.HTML
	}{
		Element: template.HTML(chartSnippet.Element),
		Script:  template.HTML(chartSnippet.Script),
		Option:  template.HTML(chartSnippet.Option),
	}

	// rerender the snippets out
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

```

## Custom Renderer
You can implement the `Renderer` all by yourself, such as
[MyRender Example](https://github.com/go-echarts/examples/blob/master/examples/renderer.go).  
If you just want to use one of the `Renderer` functions, you can use th `BaseRender`. e.g.

> The `BaseRender` is a base Render with default implementations of Renderer.
> It helps you can only implement your own necessary functions of `Renderer`.

```go
type chartRender struct {
    BaseRender
	// chart instance
    c interface{}
    // before the pre-process functions for chart render, it should only call once to support multi renders
    before []func ()
}

```

!> You need run the `before` functions to make sure all the charts options set in place.
