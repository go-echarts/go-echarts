package datasets

import (
	_ "embed"
	"encoding/json"
)

var Coordinates map[string][2]float32

//go:embed coordinates.json
var coordinatesJSON string

var MapFiles map[string]string

//go:embed mapfiles.json
var mapFilesJSON string

func init() {
	if err := json.Unmarshal([]byte(coordinatesJSON), &Coordinates); err != nil {
		panic("BUG: illegal coordinates json, err:" + err.Error())
	}
	if err := json.Unmarshal([]byte(mapFilesJSON), &MapFiles); err != nil {
		panic("BUG: illegal mapfiles json, err:" + err.Error())
	}
}
