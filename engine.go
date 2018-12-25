package goecharts

import (
	"bytes"
	"html/template"
	"io"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/gobuffalo/packr"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	chartIdSize   = 12
)

// 渲染图表
func renderChart(chart interface{}, w io.Writer, name string) error {
	fileNames := []string{"header.html", "routers.html", "base.html", name + ".html"}
	contents, err := extractTplContents(fileNames...)
	if err != nil {
		return err
	}
	tpl := template.Must(template.New("").Parse(contents[0]))
	mustTpl(tpl, contents[1:]...)
	if err = tpl.ExecuteTemplate(w, name, chart); err != nil {
		return err
	}
	return nil
}

func mustTpl(tpl *template.Template, html ...string) *template.Template {
	for i := 0; i < len(html); i++ {
		tpl = template.Must(tpl.Parse(html[i]))
	}
	return tpl
}

func extractTplContents(fileNames ...string) ([]string, error) {
	box := packr.NewBox("./templates")

	contents := make([]string, 0)
	var (
		content string
		err     error
	)
	for i := 0; i < len(fileNames); i++ {
		content, err = box.FindString(fileNames[i])
		if err != nil {
			return []string{}, err
		}
		contents = append(contents, content)
	}
	return contents, nil
}

func renderToWriter(chart interface{}, renderName string, w ...io.Writer) error {
	var b bytes.Buffer
	if err := renderChart(chart, &b, renderName); err != nil {
		return err
	}
	res := replaceRender(b)
	for i := 0; i < len(w); i++ {
		w[i].Write(res)
	}
	return nil
}

// 随机种子
var seed = rand.NewSource(time.Now().UnixNano())

// 生成唯一且随机的图表 ID
func genChartID() string {
	b := make([]byte, chartIdSize)
	for i, cache, remain := chartIdSize-1, seed.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = seed.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// 过滤替换渲染结果
func replaceRender(b bytes.Buffer) []byte {
	// __x__ 与模板占位符相匹配
	idPat, _ := regexp.Compile(`(__x__")|("__x__)`)
	// 替换并转为 []byte 类型
	content := idPat.ReplaceAllString(b.String(), "_x_")
	unusedObj := []string{
		`,"inRange":{}`,
		`,"label":{}`,
		`,"markLine":{}`,
		`,"markPoint":{}`,
		`,"areaStyle":{}`,
		`,"lineStyle":{}`,
		`,"rippleEffect":{}`,
	}
	res := []byte(removeUnusedObj(content, unusedObj...))
	return res
}

// 移除无用的 JSON object
// 另一种解决方案是使用 *struct
func removeUnusedObj(content string, pat ...string) string {
	c := content
	for i := 0; i < len(pat); i++ {
		p, _ := regexp.Compile(pat[i])
		c = p.ReplaceAllString(c, "")
	}
	return c
}

// 为结构体设置默认值
// 部分代码参考 https://github.com/mcuadros/go-defaults
func setDefaultValue(ptr interface{}) error {
	var err error
	//需要参数为指针类型，指针才能改变值
	if reflect.TypeOf(ptr).Kind() != reflect.Ptr {
		return err
	}

	elem := reflect.ValueOf(ptr).Elem()
	t := elem.Type()

	//判断是否为结构体
	if t.Kind() != reflect.Struct {
		return err
	}

	for i := 0; i < t.NumField(); i++ {
		//如果没有 `default` tag 则不作处理
		if defaultVal := t.Field(i).Tag.Get("default"); defaultVal != "" {
			setField(elem.Field(i), defaultVal)
		}
	}
	return nil
}

// 为具体字段设置默认值
func setField(field reflect.Value, defaultVal string) {
	// 目前只判断 string, bool 两种变量类型
	switch field.Kind() {
	// string 类型
	case reflect.String:
		if field.String() == "" {
			field.Set(reflect.ValueOf(defaultVal).Convert(field.Type()))
		}
		// bool 类型
	case reflect.Bool:
		if val, err := strconv.ParseBool(defaultVal); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	}
}
