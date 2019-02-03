package main

import (
	"log"
	"math/rand"
	"net/http"
	"path"

	"github.com/chenjiandongx/go-echarts/charts"
)

const (
	host   = "http://127.0.0.1:8080"
	maxNum = 50
)

type Routers struct {
	name string
	charts.HTTPRouter
}

var (
	nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}

	hours = [...]string{
		"12a", "1a", "2a", "3a", "4a", "5a", "6a", "7a", "8a", "9a", "10a", "11a",
		"12p", "1p", "2p", "3p", "4p", "5p", "6p", "7p", "8p", "9p", "10p", "11p",
	}

	days = [...]string{"Saturday", "Friday", "Thursday", "Wednesday", "Tuesday", "Monday", "Sunday"}

	mapData = map[string]float32{
		"北京":   float32(rand.Intn(150)),
		"上海":   float32(rand.Intn(150)),
		"深圳":   float32(rand.Intn(150)),
		"辽宁":   float32(rand.Intn(150)),
		"青岛":   float32(rand.Intn(150)),
		"山西":   float32(rand.Intn(150)),
		"陕西":   float32(rand.Intn(150)),
		"乌鲁木齐": float32(rand.Intn(150)),
		"齐齐哈尔": float32(rand.Intn(150)),
	}

	guangdongData = map[string]float32{
		"深圳市": float32(rand.Intn(150)),
		"广州市": float32(rand.Intn(150)),
		"湛江市": float32(rand.Intn(150)),
		"汕头市": float32(rand.Intn(150)),
		"东莞市": float32(rand.Intn(150)),
		"佛山市": float32(rand.Intn(150)),
		"云浮市": float32(rand.Intn(150)),
		"肇庆市": float32(rand.Intn(150)),
		"梅州市": float32(rand.Intn(150)),
	}

	shantouData = map[string]float32{
		"澄海区": float32(rand.Intn(150)),
		"潮阳区": float32(rand.Intn(150)),
		"潮南区": float32(rand.Intn(150)),
		"南澳县": float32(rand.Intn(150)),
	}

	routers = []Routers{
		{"bar", charts.HTTPRouter{URL: host + "/bar", Text: "Bar-(柱状图)"}},
		{"effectScatter", charts.HTTPRouter{URL: host + "/effectScatter", Text: "EffectScatter-(动态散点图)"}},
		{"funnel", charts.HTTPRouter{URL: host + "/funnel", Text: "Funnel-(漏斗图)"}},
		{"gauge", charts.HTTPRouter{URL: host + "/gauge", Text: "Gauge-仪表盘"}},
		{"geo", charts.HTTPRouter{URL: host + "/geo", Text: "Geo-地理坐标系"}},
		{"heatMap", charts.HTTPRouter{URL: host + "/heatMap", Text: "HeatMap-热力图"}},
		{"kline", charts.HTTPRouter{URL: host + "/kline", Text: "Kline-K 线图"}},
		{"line", charts.HTTPRouter{URL: host + "/line", Text: "Line-(折线图)"}},
		{"liquid", charts.HTTPRouter{URL: host + "/liquid", Text: "Liquid-(水球图)"}},
		{"map", charts.HTTPRouter{URL: host + "/map", Text: "Map-(地图)"}},
		{"pie", charts.HTTPRouter{URL: host + "/pie", Text: "Pie-(饼图)"}},
		{"scatter", charts.HTTPRouter{URL: host + "/scatter", Text: "Scatter-(散点图)"}},
		{"wordCloud", charts.HTTPRouter{URL: host + "/wordCloud", Text: "WordCloud-(词云图)"}},
		{"page", charts.HTTPRouter{URL: host + "/page", Text: "Page-(顺序多图)"}},
	}
)

func orderRouters(chartType string) []charts.HTTPRouter {
	for i := 0; i < len(routers); i++ {
		if routers[i].name == chartType {
			routers[i], routers[0] = routers[0], routers[i]
			break
		}
	}

	rs := make([]charts.HTTPRouter, 0)
	for i := 0; i < len(routers); i++ {
		rs = append(rs, routers[i].HTTPRouter)
	}
	return rs
}

func getRenderPath(f string) string {
	return path.Join("html", f)
}

func logTracing(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, rand.Intn(maxNum))
	}
	return r
}

func genKvData() map[string]interface{} {
	m := make(map[string]interface{})
	for i := 0; i < len(nameItems); i++ {
		m[nameItems[i]] = rand.Intn(50)
	}
	return m
}

func main() {
	http.HandleFunc("/bar", logTracing(BarHandler))
	http.HandleFunc("/effectScatter", logTracing(esHandler))
	http.HandleFunc("/funnel", logTracing(funnelHandler))
	http.HandleFunc("/gauge", logTracing(gaugeHandler))
	http.HandleFunc("/geo", logTracing(geoHandler))
	http.HandleFunc("/heatMap", logTracing(heatMapHandler))
	http.HandleFunc("/kline", logTracing(klineHandler))
	http.HandleFunc("/line", logTracing(lineHandler))
	http.HandleFunc("/liquid", logTracing(liquidHandler))
	http.HandleFunc("/map", logTracing(mapHandler))
	http.HandleFunc("/pie", logTracing(pieHandler))
	http.HandleFunc("/scatter", logTracing(scatterHandler))
	http.HandleFunc("/wordCloud", logTracing(wcHandler))
	http.HandleFunc("/page", logTracing(pageHandler))

	//box := packr.NewBox("./go-echarts-assets")
	//http.Handle("/", http.FileServer(box))

	log.Println("Run server at " + host)
	http.ListenAndServe(":8080", nil)
}
