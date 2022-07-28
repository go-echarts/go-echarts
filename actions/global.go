package actions

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Type string

type Areas struct {
	BrushType  string      `json:"BrushType,omitempty"`
	CoordRange []string    `json:"coordRange,omitempty"`
	XAxisIndex interface{} `json:"xAxisIndex,omitempty"`
}
