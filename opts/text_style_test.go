package opts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTextStyleRichOmittedByDefault(t *testing.T) {
	b, err := json.Marshal(TextStyle{Color: "red"})
	assert.NoError(t, err)
	assert.JSONEq(t, `{"color":"red"}`, string(b))
}

func TestTextStyleRichMarshal(t *testing.T) {
	ts := TextStyle{
		Rich: map[string]*TextStyle{
			"a": {Color: "red", FontSize: 18},
			"b": {BackgroundColor: "#eee", Padding: []int{4, 8}},
		},
	}
	b, err := json.Marshal(ts)
	assert.NoError(t, err)
	assert.JSONEq(t,
		`{"rich":{"a":{"color":"red","fontSize":18},"b":{"backgroundColor":"#eee","padding":[4,8]}}}`,
		string(b))
}
