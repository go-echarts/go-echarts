package datasets

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
)

var MapFileNames map[string]string
var Coordinates map[string][2]float32

func init() {
	box := packr.NewBox(".")
	maps, err := box.Find("map_filename.json")
	if err == nil {
		json.Unmarshal(maps, &MapFileNames)
	}

	coordinates, err := box.Find("coordinates.json")
	if err == nil {
		json.Unmarshal(coordinates, &Coordinates)
	}
}

func LoadAssets(loader http.FileSystem) {
	fMaps, err := loader.Open("map_filename.json")
	maps, err := ioutil.ReadAll(fMaps)
	json.Unmarshal(maps, &MapFileNames)

	fCoordinates, err := loader.Open("coordinates.json")
	coordinates, err := ioutil.ReadAll(fCoordinates)
	json.Unmarshal(coordinates, &Coordinates)

	if err != nil {
		log.Fatal(err)
	}
}
