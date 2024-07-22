package opts

import (
	"fmt"
	"regexp"

	"github.com/go-echarts/go-echarts/v2/types"
)

const funcMarker = "__f__"

var (
	newlineTabPat = regexp.MustCompile(`\n|\t`)
	commentPat    = regexp.MustCompile(`(//.*)\n`)
)

type JSFunctions struct {
	Fns []types.FuncStr
}

// AddJSFuncs adds a new JS function, for back compatibility before v2.4
// Plz use the AddJSFuncStrs instead
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, types.FuncStr(newlineTabPat.ReplaceAllString(fn[i], "")))
	}
}

// AddJSFuncStrs adds a new JS function string.
func (f *JSFunctions) AddJSFuncStrs(fn ...types.FuncStr) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, types.FuncStr(newlineTabPat.ReplaceAllString(string(fn[i]), "")))
	}
}

// FuncOpts returns a string suitable for options expecting pure JavaScript code insertion.
func FuncOpts(fn string) types.FuncStr {
	return types.FuncStr(replaceJsFuncs(fn))
}

// FuncStripCommentsOpts returns a string suitable for options expecting JavaScript code,
// stripping '//' comments.
func FuncStripCommentsOpts(fn string) string {
	fn = commentPat.ReplaceAllString(fn, "")
	return replaceJsFuncs(fn)
}

// replace and clear up js functions string
func replaceJsFuncs(fn string) string {
	fn = newlineTabPat.ReplaceAllString(fn, "")
	return fmt.Sprintf("%s%s%s", funcMarker, fn, funcMarker)
}
