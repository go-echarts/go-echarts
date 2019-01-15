package datasets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJSONFiles(t *testing.T) {
	assert.Equal(t, MapFileNames["china"], "china")
	assert.Equal(t, MapFileNames["world"], "world")
	assert.Equal(t, Coordinates["阿城"], [2]float32{126.58, 45.32})
	assert.Equal(t, Coordinates["阿克苏"], [2]float32{80.19, 41.09})
}
