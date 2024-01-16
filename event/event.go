package event

import "github.com/go-echarts/go-echarts/v2/types"

// EventInterceptor is the abstract of events in echarts
// see https://echarts.apache.org/handbook/en/concepts/event
// which allow to listen special mouse and components event
// example:
//
//	bar := charts.NewBar()
//	JFunc := ` (params) => alert(params.name) `
//	bar.SetGlobalOptions(
//	charts.WithEventInterceptors(
//	       event.EventInterceptor{
//	         EventName: "click",
//	         Handler:   opts.FuncOpts(JFunc),
//	       },
//	       event.EventInterceptor{
//	         EventName: "mouseup",
//	         Query:     "'series'",
//	         Handler:   opts.FuncOpts(JFunc),
//	       },
//	       event.EventInterceptor{
//	         EventName: "mouseover",
//	         Query:     "{ seriesName: 'go-echarts' }",
//	         Handler:   opts.FuncOpts(JFunc),
//	       },
//
// ),
// )
type EventInterceptor struct {
	EventName string
	Query     types.FuncStr
	Handler   types.FuncStr
}
