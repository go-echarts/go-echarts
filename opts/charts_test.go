package opts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointerShowAboveMarshal(t *testing.T) {
	for _, tt := range []struct {
		name      string
		showAbove *bool
		want      string
	}{
		{"false", Bool(false), `{"showAbove":false}`},
		{"true", Bool(true), `{"showAbove":true}`},
		{"unset", nil, `{}`},
	} {
		t.Run(tt.name, func(t *testing.T) {
			b, err := json.Marshal(Pointer{ShowAbove: tt.showAbove})
			assert.NoError(t, err)
			assert.JSONEq(t, tt.want, string(b))
		})
	}
}
