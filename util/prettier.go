package util

import "reflect"

// ConfigPrettier remove all the empty structs
func ConfigPrettier(ptr interface{}) {
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

			if f.IsZero() {
				field := val.Field(i)
				field.Set(reflect.Zero(field.Type()).Convert(field.Type()))
				continue
			} else {
				doPretty(f)
			}
		}
	}
}
