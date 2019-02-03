package charts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPageAssetsBeforeRender(t *testing.T) {
	page := NewPage()
	assert.Equal(t, len(page.JSAssets.Values), 0)
	assert.Equal(t, len(page.CSSAssets.Values), 0)
}

func TestPageAssetsAfterRender(t *testing.T) {
	page := NewPage()

	page.Add(NewBar())
	page.Add(NewMap("china").SetGlobalOptions(InitOpts{Theme: "macarons"}))

	err := page.Render()
	assert.NoError(t, err)
	var host = "http://chenjiandongx.com/go-echarts-assets/assets/"
	assert.Equal(t, page.JSAssets.Values, []string{
		host + "echarts.min.js", host + "maps/china.js", host + "themes/macarons.js"},
	)
	assert.Equal(t, page.CSSAssets.Values, []string{host + "bulma.min.css"})
}

func TestUniqueChartID(t *testing.T) {
	bar := NewBar()
	line := NewLine()

	assert.NotEqual(t, bar.ChartID, line.ChartID)
}
