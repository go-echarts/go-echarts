package components

import (
	"github.com/go-echarts/go-echarts/v2/core"
	"github.com/go-echarts/go-echarts/v2/primitive"
	"github.com/go-echarts/go-echarts/v2/util"
)

var DefaultColors = []string{
	"#5470c6", "#91cc75", "#fac858", "#ee6666", "#73c0de",
	"#3ba272", "#fc8452", "#9a60b4", "#ea7ccc"}

// BaseConfiguration represents basic options set needed by all chart types.
type BaseConfiguration struct {
	*core.Page      `json:"-"`
	*core.Container `json:"-"`
	*Title          `json:"title,omitempty"`
	*Legend         `json:"legend,omitempty,reserved"`
	*Tooltip        `json:"tooltip,omitempty,reserved"`
	*Toolbox        `json:"toolbox,omitempty"`

	Legends []primitive.String `json:"legends,omitempty"`
	// Colors is the color list of palette.
	// If no color is set in series, the colors would be adopted sequentially and circularly
	// from this list as the colors of series.
	Colors []primitive.String `json:"colors,omitempty"`
	//appendColor      []string // append customize color to the Colors(reverse order)

	// Array of dataset
	Datasets []*Dataset `json:"dataset,omitempty"`

	DataZooms  []*DataZoom `json:"datazoom,omitempty"`
	VisualMaps *VisualMap  `json:"visualmap,omitempty"`
}

func (bc BaseConfiguration) New() *BaseConfiguration {

	return &BaseConfiguration{
		Title:   &Title{},
		Legend:  &Legend{},
		Tooltip: &Tooltip{},
		Toolbox: &Toolbox{},
		Colors:  util.StringsConv(DefaultColors...),
	}
}
