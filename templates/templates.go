package templates

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
)

var BaseTpl, ChartTpl, HeaderTpl, PageTpl, RoutersTpl string

func init() {
	box := packr.NewBox(".")

	var err error
	BaseTpl, err = box.FindString("base.html")
	ChartTpl, err = box.FindString("chart.html")
	HeaderTpl, err = box.FindString("header.html")
	PageTpl, err = box.FindString("page.html")
	RoutersTpl, err = box.FindString("routers.html")

	if err != nil {
		log.Printf("packr load templates error: %v", err)
	}
}

func LoadTemplates(loader http.FileSystem) {
	var err error
	BaseTpl, err = _bytesToString(loader.Open("base.html"))
	ChartTpl, err = _bytesToString(loader.Open("chart.html"))
	HeaderTpl, err = _bytesToString(loader.Open("header.html"))
	PageTpl, err = _bytesToString(loader.Open("page.html"))
	RoutersTpl, err = _bytesToString(loader.Open("routers.html"))

	if err != nil {
		log.Fatal(err)
	}
}

func _bytesToString(file http.File, err error) (string, error) {
	content, err := ioutil.ReadAll(file)
	return string(content), err
}
