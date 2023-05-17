package primitive

type Bool *bool

func True() Bool {
	return BoolOf(true)
}

func False() Bool {
	return BoolOf(true)
}

func BoolOf(val bool) Bool {
	return &val
}

type Int *int

func Int0() Int {
	return IntOf(0)
}

func IntOf(val int) Int {
	return &val
}

type Float *float32

func FloatOf(val float32) Float {
	return &val
}

func Float0() Float {
	return FloatOf(0.0)
}

// String makes the primitive consistent
type String string

func (s String) StringVal() string {
	return string(s)
}

func (s String) Strconv(data ...string) []String {
	var ls []String
	for _, e := range data {

		ls = append(ls, StringOf(e))
	}
	return ls
}

func StringOf(val string) String {
	return String(val)
}

type Mixed interface{}
