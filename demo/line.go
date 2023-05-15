package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/primitive"
	"os"
	"reflect"
)

func main() {

	line := charts.NewLine()
	line.TitleV3.Text = "Title"
	line.TitleV3.SubText = "Subtitle"

	line.XAxis.Type = "category"
	line.XAxis.Data = []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	line.YAxis.Type = "value"
	line.Series.Data = []int{150, 230, 224, 218, 135, 147, 260}
	// show Toolbox but no other empty structs
	line.Toolbox.Show = primitive.BoolOf(true)
	data, _ := json.Marshal(line)
	fmt.Println(string(data))

	f, _ := os.Create("line.html")
	// TODO: 1.  allow process chart before render such as ConfigCleaner 2.page configs
	ConfigCleaner(line)
	_ = line.Render(f)
}

// ConfigCleaner remove all the empty structs
func ConfigCleaner(ptr interface{}) {
	elem := reflect.ValueOf(ptr).Elem()
	walkField(elem)
}

func walkField(val reflect.Value) {
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
				walkField(f)
			}
		}
	}
}
