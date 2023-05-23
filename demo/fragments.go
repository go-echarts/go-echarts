package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/core"
	"strings"
)

type MyMockPNGRenderer struct {
}

func (mcr *MyMockPNGRenderer) Resolve(p *core.Page, dest string) {
	title := p.Title.StringVal()
	suffix := "Option Way"
	if strings.Contains(title, "Flow") {
		suffix = "Flow Way"
	}
	// Process render PNG to dest
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	fmt.Print("\n")
	fmt.Println("\n Process render PNG to " + dest + " in " + suffix)
	fmt.Print("\n")
}

func (mcr *MyMockPNGRenderer) GetRenderer() *core.DefaultRenderer {
	return nil
}

func (mcr *MyMockPNGRenderer) SetRender(r core.Render) {
	panic("implement me")
}

func (mcr *MyMockPNGRenderer) SetWriter(w core.Writer) {
	panic("implement me")
}

type MyPageTplProvider struct {
}

func (mptp *MyPageTplProvider) Provide() *core.Page {
	p := core.NewPage()
	p.JSAssets.Reset()
	p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts/dist/echarts.min.js")
	p.JSAssets.Add("https://cdn.jsdelivr.net/npm/echarts-gl/dist/echarts-gl.min.js")
	p.JSAssets.Add("My.js")
	return p
}
