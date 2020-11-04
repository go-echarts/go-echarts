package opts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssets(t *testing.T) {
	assetsEntity := &Assets{}
	assetsEntity.InitAssets()

	assert.Equal(t, assetsEntity.JSAssets.Values, []string{"echarts.min.js"})
	assert.Equal(t, len(assetsEntity.CSSAssets.Values), 0)

	assetsEntity.JSAssets.Add("jquery.min.js")
	assetsEntity.AddCustomizedJSAssets("http://myhost/my.assets.js")

	const host = "https://go-echarts.github.io/go-echarts-assets/assets/"

	assetsEntity.Validate(host)
	assert.Equal(t, assetsEntity.JSAssets.Values, []string{
		host + "echarts.min.js",
		host + "jquery.min.js",
		"http://myhost/my.assets.js",
	})
}
