package charts

import (
	"bytes"
	"encoding/json"
	"html/template"

	"github.com/go-echarts/go-echarts/v3/opts"
)

type BaseConfiguration struct {
	Title       *opts.Title
	Legend      *opts.Legend
	Tooltip     *opts.Tooltip
	AxisPointer *opts.AxisPointer
	Radar       *opts.Radar
	Series      interface{}
}

func (bc *BaseConfiguration) ToMap() map[string]interface{} {
	obj := map[string]interface{}{
		"title":   bc.Title,
		"legend":  bc.Legend,
		"tooltip": bc.Tooltip,
		"series":  bc.Series,
	}

	if bc.AxisPointer != nil {
		obj["axisPointer"] = bc.AxisPointer
	}
	return obj
}

func (bc *BaseConfiguration) JSONNotEscaped() template.HTML {
	obj := bc.ToMap()
	buf := bytes.NewBufferString("")
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	enc.Encode(obj)

	return template.HTML(buf.String())
}
