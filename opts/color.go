package opts

import (
	"bytes"
	"fmt"
	"strings"
)

type Color string

type ColorVar interface {
	Text() Color
}

type rgb struct {
	r uint16
	g uint16
	b uint16
	a *float32
}

func (c rgb) Text() Color {
	if c.a == nil {
		return Color(fmt.Sprintf("rgb(%d, %d, %d)", c.r, c.g, c.b))
	}
	return Color(fmt.Sprintf("rgba(%d, %d, %d, %f)", c.r, c.g, c.b, *c.a))
}

func HexColor(str string) Color {
	defaultVar := Color("#eeeeee")
	if len(str) == 0 {
		return defaultVar
	}

	if str[0] != '#' {
		return defaultVar
	}

	if len(str) == 7 {
		buf := bytes.NewBufferString(strings.ToLower(str))
		var r, g, b uint16
		_, err := fmt.Fscanf(buf, "#%02d%02d%02d", &r, &g, &b)
		if err != nil {
			return defaultVar
		}
		return RGB(r, g, b)
	}

	if len(str) == 4 {
		c1 := fmt.Sprintf("%c%c", str[1], str[1])
		c2 := fmt.Sprintf("%c%c", str[2], str[2])
		c3 := fmt.Sprintf("%c%c", str[3], str[3])

		text := "#" + c1 + c2 + c3
		buf := bytes.NewBufferString(strings.ToLower(text))
		var r, g, b uint16
		_, err := fmt.Fscanf(buf, "#%02d%02d%02d", &r, &g, &b)
		if err != nil {
			return defaultVar
		}
		return RGB(r, g, b)
	}

	return defaultVar
}

func RGB(r, g, b uint16) Color {
	color := rgb{
		r: r,
		g: g,
		b: b,
	}
	return color.Text()
}

func RGBA(r, g, b uint16, a float32) Color {
	color := rgb{
		r: r,
		g: g,
		b: b,
		a: &a,
	}
	return color.Text()
}

type hsl struct {
	h float32
	s float32
	l float32
	a *float32
}

func (c hsl) Text() Color {
	if c.a == nil {
		return Color(fmt.Sprintf("hsl(%f, %f%%, %f%%)", c.h, c.s, c.l))
	}
	return Color(fmt.Sprintf("hsla(%f, %f%%, %f%%, %f)", c.h, c.s, c.l, *c.a))
}

func HSL(h, s, l float32) Color {
	color := hsl{
		h: h,
		s: s,
		l: l,
	}
	return color.Text()
}

func HSLA(h, s, l, a float32) Color {
	color := hsl{
		h: h,
		s: s,
		l: l,
		a: &a,
	}
	return color.Text()
}
