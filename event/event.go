package event

import "github.com/go-echarts/go-echarts/v2/types"

type EventInterceptor struct {
	EventName string
	Query     types.FuncStr
	Handler   types.FuncStr
}
