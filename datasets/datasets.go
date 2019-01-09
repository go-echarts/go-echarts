package datasets

import (
	"encoding/json"
	"fmt"
	
	"github.com/gobuffalo/packr"
)

var MapFileNames map[string]string
var Coordinates map[string][2]float32

func init() {
	box := packr.NewBox(".")
	maps, err := box.Find("map_filename.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(maps, &MapFileNames)

	coordinates, err := box.Find("coordinates.json")
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(coordinates, &Coordinates)
}
