package event

import "github.com/go-echarts/go-echarts/v2/types"

// Listener is the abstract of events listening in echarts
// see https://echarts.apache.org/handbook/en/concepts/event
// which allow to listen special mouse and components event
// example:
//
//	bar := charts.NewBar()
//	JFunc := ` (params) => alert(params.name) `
//	bar.SetGlobalOptions(
//	charts.WithEventListeners(
//	       event.Listener{
//	         EventName: "click",
//	         Handler:   opts.FuncOpts(JFunc),
//	       },
//	       event.Listener{
//	         EventName: "mouseup",
//	         Query:     "'series'",
//	         Handler:   opts.FuncOpts(JFunc),
//	       },
//	       event.Listener{
//	         EventName: "mouseover",
//	         Query:     "{ seriesName: 'go-echarts' }",
//	         Handler:   opts.FuncOpts(JFunc),
//	       },
//
// ),
// )
type Listener struct {
	EventName string
	Query     types.FuncStr
	Handler   types.FuncStr
}
