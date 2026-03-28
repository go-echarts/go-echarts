package opts

import "github.com/go-echarts/go-echarts/v2/types"

// GeoComponent is the option set for geo component.
// https://echarts.apache.org/en/option.html#geo
type GeoComponent struct {
	// Map charts.
	Map string `json:"map,omitempty"`

	// Specifies which point on the map should be placed at center of viewport.
	// When roaming, values in center and zoom will be modified correspondingly.
	// Center is in longitude and latitude by default. Use the projected coordinates.
	// Example: [115.97, 29.71].
	// A percentage string can also be used, like '30%', based on the bounding rect.
	// You can use '0%' to place the top or left of bounding rect to the center,
	// or use '50%' to place the entire source map at the the center of the viewport.
	// Example: [115, '30%'].
	Center interface{} `json:"center,omitempty"`

	// Zoom rate of current viewport. The value less than 1 indicates zooming out,
	// while the value greater than 1 indicates zooming in.
	// When roaming, values in center and zoom will be modified correspondingly.
	Zoom types.Float `json:"zoom,omitempty"`

	// Whether to enable mouse or touch roam (move and zoom):
	// * false: roam is disabled.
	// * true: both zoom and move (translation) are available.
	// When roaming, values in center and zoom will be modified correspondingly.
	Roam types.Bool `json:"roam,omitempty"`

	// Graphic style of Map Area Border, emphasis is the style when it is highlighted,
	// like being hovered by mouse, or highlighted via legend connect.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Set this to true, to prevent interaction with the axis.
	Silent types.Bool `json:"silent,omitempty"`
}
