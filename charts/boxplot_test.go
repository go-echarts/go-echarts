package charts

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoxPlotAssetsBeforeRender(t *testing.T) {
	boxPlot := NewBoxPlot()
	assert.Equal(t, []string{"echarts.min.js"}, boxPlot.JSAssets.Values)
}

func TestBoxPlotAssetsAfterRender(t *testing.T) {
	boxPlot := NewBoxPlot()
	err := boxPlot.Render(ioutil.Discard)
	assert.NoError(t, err)
	assert.Equal(t, []string{host + "echarts.min.js"}, boxPlot.JSAssets.Values)
}
