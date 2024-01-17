package opts

import "github.com/go-echarts/go-echarts/v2/types"

func Bool(val bool) types.Bool {
	return &val
}

func Int(val int) types.Int {
	return &val
}

func Float(val float32) types.Float {
	return &val
}

func Str(val string) types.String {
	return types.String(val)
}
