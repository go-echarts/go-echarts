package core

import "github.com/go-echarts/go-echarts/v2/primitive"

// Event mount on chart instance / container
// myChart.EventBinder("EventType', <params>, function{})
/**
myChart.on('click', 'series', function() {});

myChart.on('mouseover', { seriesIndex: 1, name: 'xx' }, function() {
});

myChart.dispatchAction({
    type: 'downplay',
    seriesIndex: 0,
    dataIndex: currentIndex
  });


*/
type Event struct {
	EventBinder primitive.String
	EventType   primitive.String
	Params      primitive.Mixed
	Action      primitive.Mixed
}
