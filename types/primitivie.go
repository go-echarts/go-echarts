package types

type Bool *bool
type Int *int
type Float *float32
type String string

func IsTrue(b Bool) bool {
	return *b
}
