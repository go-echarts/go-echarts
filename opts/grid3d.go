package opts

// Grid3D contains options for the 3D coordinate.
type Grid3D struct {
	// Whether to show the coordinate.
	Show bool `json:"show,omitempty"`

	// 3D Cartesian coordinates width
	// default 100
	BoxWidth float32 `json:"boxWidth,omitempty"`

	// 3D Cartesian coordinates height
	// default 100
	BoxHeight float32 `json:"boxHeight,omitempty"`

	// 3D Cartesian coordinates depth
	// default 100
	BoxDepth float32 `json:"boxDepth,omitempty"`

	// Rotate or scale fellows the mouse
	ViewControl *ViewControl `json:"viewControl,omitempty"`
}
