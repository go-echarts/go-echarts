package config

import (
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/primitive"
)

// BaseConfiguration represents basic options set needed by all chart types.
type BaseConfiguration struct {
	*components.Page      `json:"-"`
	*components.Container `json:"-"`
	*opts.Title           `json:"title,omitempty"`
	*opts.Legend          `json:"legend,omitempty"`
	*opts.Tooltip         `json:"tooltip,omitempty"`
	*opts.Toolbox         `json:"toolbox,omitempty"`

	Legends []primitive.String `json:"legends,omitempty"`
	// Colors is the color list of palette.
	// If no color is set in series, the colors would be adopted sequentially and circularly
	// from this list as the colors of series.
	Colors []primitive.String `json:"colors,omitempty"`
	//appendColor      []string // append customize color to the Colors(reverse order)

	// Array of dataset
	Datasets []*opts.Dataset `json:"dataset,omitempty"`

	DataZooms  []*opts.DataZoom  `json:"datazoom,omitempty"`
	VisualMaps []*opts.VisualMap `json:"visualmap,omitempty"`

	*opts.Initialization
	*opts.Assets `json:"-"`
}

func (bc BaseConfiguration) New() *BaseConfiguration {
	return &BaseConfiguration{
		Title:      &opts.Title{},
		Legend:     &opts.Legend{},
		Tooltip:    &opts.Tooltip{},
		Toolbox:    &opts.Toolbox{},
		Legends:    nil,
		Colors:     nil,
		Datasets:   nil,
		DataZooms:  nil,
		VisualMaps: nil,
	}
}
