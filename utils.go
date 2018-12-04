package geocharts

import (
	"fmt"
	"reflect"
	"strconv"
)

func SetDefaultValue(ptr interface{}) error {
	var err error
	//需要参数为指针类型，指针才能改变值
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return err
	}

	elem := reflect.ValueOf(ptr).Elem()
	t := elem.Type()

	//判断是否为结构体
	if t.Kind() != reflect.Struct {
		return err
	}

	for i := 0; i < t.NumField(); i++ {
		//如果没有 `default` tag 则不作处理
		fmt.Println(t.Field(i).Tag.Get("default"))
		if defaultVal := t.Field(i).Tag.Get("default"); defaultVal != "" {
			setField(elem.Field(i), defaultVal)
		}
	}
	return nil
}

func setField(field reflect.Value, defaultVal string) {
	switch field.Kind() {
	case reflect.String:
		field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	}
}



