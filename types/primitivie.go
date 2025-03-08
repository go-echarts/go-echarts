package types

type (
	// Bool a wrapper type of *bool, use opts.Bool to simply convert it.
	Bool *bool
	// Int a wrapper type of *int, use opts.Int to simply convert it.
	Int    *int
	Float  *float32
	String string
)

// FuncStr a pure JavaScrip function string or special formatted string
// use opts.FuncOpts or opts.FuncStripCommentsOpts to embed JavaScript.
type FuncStr string
