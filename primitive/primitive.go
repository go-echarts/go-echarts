package primitive

type Bool *bool

func NewBool(b bool) Bool {
	//tmp := b
	return &b
}

type Int *int

type Float64 *float64
