package templates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJSONFiles(t *testing.T) {
	assert.NotEmpty(t, BaseTpl)
	assert.NotEmpty(t, ChartTpl)
	assert.NotEmpty(t, HeaderTpl)
	assert.NotEmpty(t, PageTpl)
	assert.NotEmpty(t, RoutersTpl)
}
