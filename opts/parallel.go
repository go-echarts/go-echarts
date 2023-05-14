package opts

// ParallelComponent is the option set for parallel component.
type ParallelComponent struct {
	// Distance between parallel component and the left side of the container.
	// Left value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%';
	// and it can also be 'left', 'center', or 'right'.
	// If the left value is set to be 'left', 'center', or 'right',
	// then the component will be aligned automatically based on position.
	Left string `json:"left,omitempty"`

	// Distance between parallel component and the top side of the container.
	// Top value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	// and it can also be 'top', 'middle', or 'bottom'.
	// If the left value is set to be 'top', 'middle', or 'bottom',
	// then the component will be aligned automatically based on position.
	Top string `json:"top,omitempty"`

	// Distance between parallel component and the right side of the container.
	// Right value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Right string `json:"right,omitempty"`

	// Distance between parallel component and the bottom side of the container.
	// Bottom value can be instant pixel value like 20.
	// It can also be a percentage value relative to container width like '20%'.
	Bottom string `json:"bottom,omitempty"`
}
