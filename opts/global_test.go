package opts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssets(t *testing.T) {
	assetsEntity := &Assets{}
	assetsEntity.InitAssets()

	assert.Equal(t, []string{"echarts.min.js"}, assetsEntity.JSAssets.Values)
	assert.Equal(t, 0, len(assetsEntity.CSSAssets.Values))

	assetsEntity.JSAssets.Add("jquery.min.js")
	assetsEntity.AddCustomizedJSAssets("http://myhost/my.assets.js")

	const host = "https://go-echarts.github.io/go-echarts-assets/assets/"

	assetsEntity.Validate(host)
	assert.Equal(t, []string{host + "echarts.min.js", host + "jquery.min.js"}, assetsEntity.JSAssets.Values)
	assert.Equal(t, []string{"http://myhost/my.assets.js"}, assetsEntity.CustomizedJSAssets.Values)
}

func TestGenerateUniqueID(t *testing.T) {
	var old string
	for i := 0; i < 10; i++ {
		new := generateUniqueID()
		assert.NotSame(t, old, new)
		old = new
	}
}
