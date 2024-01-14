package opts

import (
	"fmt"
	"regexp"
)

const funcMarker = "__f__"

var newlineTabPat = regexp.MustCompile(`\n|\t`)
var commentPat = regexp.MustCompile(`(//.*)\n`)

type JSFunctions struct {
	Fns []string
}

// AddJSFuncs adds a new JS function.
func (f *JSFunctions) AddJSFuncs(fn ...string) {
	for i := 0; i < len(fn); i++ {
		f.Fns = append(f.Fns, newlineTabPat.ReplaceAllString(fn[i], ""))
	}
}

// FuncOpts returns a string suitable for options expecting pure JavaScript code insertion.
func FuncOpts(fn string) string {
	return replaceJsFuncs(fn)
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
