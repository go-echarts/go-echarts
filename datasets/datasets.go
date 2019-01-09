package datasets

import (
	"encoding/json"
	"log"

	"github.com/gobuffalo/packr"
)

var MapFileNames map[string]string
var Coordinates map[string][2]float32

func init() {
	box := packr.NewBox(".")
	maps, err := box.Find("map_filename.json")
	json.Unmarshal(maps, &MapFileNames)

	coordinates, err := box.Find("coordinates.json")
	json.Unmarshal(coordinates, &Coordinates)

	if err != nil {
		log.Fatal(err)
	}
}
