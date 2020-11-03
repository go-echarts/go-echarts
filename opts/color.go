package opts

import (
	"encoding/json"
	"fmt"
)

type ColorVar interface {
	json.Marshaler
}

type ColorString string

func (h ColorString) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, h)), nil
}

type RGB struct {
	R uint16
	G uint16
	B uint16
	A *float32
}

func NewRGB(R, G, B uint16) RGB {
	return RGB{
		R: R,
		G: G,
		B: B,
	}
}

func NewRGBA(R, G, B uint16, A float32) RGB {
	return RGB{
		R: R,
		G: G,
		B: B,
		A: &A,
	}
}

func (r RGB) MarshalJSON() ([]byte, error) {
	var text string
	if r.A == nil {
		text = fmt.Sprintf("rgb(%d, %d, %d)", r.R, r.G, r.B)
	} else {
		text = fmt.Sprintf("rgba(%d, %d, %d, %f)", r.R, r.G, r.B, *r.A)
	}
	return []byte(fmt.Sprintf(`"%s"`, text)), nil
}

type HSL struct {
	H float32
	S float32
	L float32
	A *float32
}

func NewHSL(H, S, L float32) HSL {
	return HSL{
		H: H,
		S: S,
		L: L,
	}
}

func NewHSLA(H, S, L, A float32) HSL {
	return HSL{
		H: H,
		S: S,
		L: L,
		A: &A,
	}
}

func (h HSL) MarshalJSON() ([]byte, error) {
	var text string
	if h.A == nil {
		text = fmt.Sprintf("hsl(%f, %f%%, %f%%)", h.H, h.S, h.L)
	} else {
		text = fmt.Sprintf("hsla(%f, %f%%, %f%%, %f)", h.H, h.S, h.L, *h.A)
	}
	return []byte(fmt.Sprintf(`"%s"`, text)), nil
}
