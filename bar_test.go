package goecharts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBarAssetsBeforeRender(t *testing.T) {
	bar := NewBar()
	assert.Equal(t, bar.JSAssets.Values, []string{"echarts.min.js"}, "")
	assert.Equal(t, bar.CSSAssets.Values, []string{"bulma.min.css"}, "")
}

func TestBarAssetsAfterRender(t *testing.T) {
	bar := NewBar()
	bar.Render()

	var host = "http://chenjiandongx.com/go-echarts-assets/assets/"
	assert.Equal(t, bar.JSAssets.Values, []string{host + "echarts.min.js"}, "")
	assert.Equal(t, bar.CSSAssets.Values, []string{host + "bulma.min.css"}, "")
}
