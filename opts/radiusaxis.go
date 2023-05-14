package opts

type RadiusAxis struct {
	PolarAxisBase
	Name          string    `json:"name,omitempty"`
	NameLocation  string    `json:"nameLocation,omitempty"`
	NameTextStyle TextStyle `json:"nameTextStyle,omitempty"`
	NameGap       float64   `json:"nameGap,omitempty"`
	NameRadius    float64   `json:"nameRotate,omitempty"`
	Inverse       bool      `json:"inverse,omitempty"`
}
