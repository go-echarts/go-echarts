package util

import (
	"reflect"
	"strings"
)

const ReversedTag = "reserved"

// Prettier remove all the empty structs
func Prettier(ptr interface{}) {
	elem := reflect.ValueOf(ptr).Elem()
	doPretty(elem)
}

func doPretty(val reflect.Value) {
	t := val.Type()

	for i := 0; i < t.NumField(); i++ {
		f := val.Field(i)
		if f.Kind() == reflect.Ptr {
			f = f.Elem()
		}

		if f.Kind() == reflect.Struct {

			t.Field(i)
			if f.IsZero() && !reserved(t.Field(i)) {
				field := val.Field(i)
				field.Set(reflect.Zero(field.Type()).Convert(field.Type()))
				continue
			} else {
				doPretty(f)
			}
		}
	}
}

func reserved(field reflect.StructField) bool {
	var f string
	if f = field.Tag.Get("json"); f == "" {
		return false
	}

	return strings.Contains(f, ReversedTag)

}
