package types

// thoughts on those boxed type for default value solution...
type (
	Bool    *bool
	Integer *int
)

func newBool(val bool) Bool {
	return &val
}

func newInteger(val int) Integer {
	return &val
}
