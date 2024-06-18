## Grid

Drawing grid in rectangular coordinate. In a single grid, at most two X and Y axes each is allowed. Line chart, bar chart, and scatter chart (bubble chart) can be drawn in grid.

## Available Option

```go
package opts

// Grid Drawing grid in rectangular coordinate.
// https://echarts.apache.org/en/option.html#grid
type Grid struct {
	// Whether to show the grid in rectangular coordinate.
	Show types.Bool `json:"show,omitempty"`

	// Distance between grid component and the left side of the container.
	Left string `json:"left,omitempty"`

	// Distance between grid component and the top side of the container.
	Top string `json:"top,omitempty"`

	// Distance between grid component and the right side of the container.
	Right string `json:"right,omitempty"`

	// Distance between grid component and the bottom side of the container.
	Bottom string `json:"bottom,omitempty"`

	// Width of grid component.
	Width string `json:"width,omitempty"`

	// Height of grid component. Adaptive by default.
	Height string `json:"height,omitempty"`

	ContainLabel types.Bool `json:"containLabel,omitempty"`

	Tooltip *Tooltip `json:"tooltip,omitempty"`
}

```

## Expect Format
```js
option = {
  grid: {
    width: "26%",
    height: "33%",
    containLabel: true
  },
    ...
}
```