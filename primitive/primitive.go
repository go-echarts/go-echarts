package primitive

type Bool *bool

func BoolOf(val bool) Bool {
	return &val
}

type Int *int

func IntOf(val int) Int {
	return &val
}

type Float *float32

func FloatOf(val float32) Float {
	return &val
}

// String makes the primitive consistent
type String string

func StringOf(val string) String {
	return String(val)
}

type Mixed interface{}
