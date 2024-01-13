package opts

// Polar coordinate can be used in scatter and line chart. Every polar coordinate has an angleAxis and a radiusAxis.
type Polar struct {
	ID      string    `json:"id,omitempty"`
	Zlevel  int       `json:"zlevel,omitempty"`
	Z       int       `json:"z,omitempty"`
	Center  [2]string `json:"center,omitempty"`
	Radius  [2]string `json:"radius,omitempty"`
	Tooltip Tooltip   `json:"tooltip,omitempty"`
}
