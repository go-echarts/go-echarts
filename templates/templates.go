package templates

import (
	"log"

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
		log.Fatal(err)
	}
}
