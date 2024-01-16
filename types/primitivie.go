package types

// Bool a wrapper type of *bool, use opts.Bool to simply convert it.
type Bool *bool
type Int *int
type Float *float32
type String string

// FuncStr a pure JavaScrip function string or special formatted string
// use opts.FuncOpts or opts.FuncStripCommentsOpts to embed JavaScript.
type FuncStr string
