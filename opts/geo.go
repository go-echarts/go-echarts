package opts

// GeoComponent is the option set for geo component.
// https://echarts.apache.org/en/option.html#geo
type GeoComponent struct {
	// Map charts.
	Map string `json:"map,omitempty"`

	// Graphic style of Map Area Border, emphasis is the style when it is highlighted,
	// like being hovered by mouse, or highlighted via legend connect.
	ItemStyle *ItemStyle `json:"itemStyle,omitempty"`

	// Set this to true, to prevent interaction with the axis.
	Silent bool `json:"silent,omitempty"`
}
