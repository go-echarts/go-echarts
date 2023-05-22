package core

import "github.com/go-echarts/go-echarts/v2/primitive"

// Event mount on chart instance / container
type Event struct {
	Event     primitive.String
	EventType primitive.String
	Action    primitive.Mixed
}
