package actions

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Type kind of dispatch action
type Type string

// Areas means select-boxes. Multi-boxes can be specified.
// If Areas is empty, all of the select-boxes will be deleted.
// The first area.
type Areas struct {

	//BrushType Optional: 'polygon', 'rect', 'lineX', 'lineY'
	BrushType string `json:"brushType,omitempty"`

	// CoordRange Only for "coordinate system area", define the area with the
	// coordinates.
	CoordRange []string `json:"coordRange,omitempty"`

	// XAxisIndex Assigns which of the xAxisIndex can use Area selecting.
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`
}
