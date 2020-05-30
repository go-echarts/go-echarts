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
	BaseTpl, err = box.FindString("base.gohtml")
	if err != nil {
		log.Printf("packr load templates error: %v", err)
	}
	ChartTpl, err = box.FindString("chart.gohtml")
	if err != nil {
		log.Printf("packr load templates error: %v", err)
	}
	HeaderTpl, err = box.FindString("header.gohtml")
	if err != nil {
		log.Printf("packr load templates error: %v", err)
	}
	PageTpl, err = box.FindString("page.gohtml")
	if err != nil {
		log.Printf("packr load templates error: %v", err)
	}
	RoutersTpl, err = box.FindString("routers.gohtml")
	if err != nil {
		log.Printf("packr load templates error: %v", err)
	}
}

func LoadTemplates(loader http.FileSystem) {
	var err error
	BaseTpl, err = _bytesToString(loader.Open("base.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	ChartTpl, err = _bytesToString(loader.Open("chart.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	HeaderTpl, err = _bytesToString(loader.Open("header.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	PageTpl, err = _bytesToString(loader.Open("page.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
	RoutersTpl, err = _bytesToString(loader.Open("routers.gohtml"))
	if err != nil {
		log.Fatal(err)
	}
}

func _bytesToString(file http.File, err error) (string, error) {
	content, err := ioutil.ReadAll(file)
	return string(content), err
}
