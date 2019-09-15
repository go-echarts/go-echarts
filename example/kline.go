package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

type klineData struct {
	date string
	data [4]float32
}

var kd = [...]klineData{
	{date: "2018/1/24", data: [4]float32{2320.26, 2320.26, 2287.3, 2362.94}},
	{date: "2018/1/25", data: [4]float32{2300, 2291.3, 2288.26, 2308.38}},
	{date: "2018/1/28", data: [4]float32{2295.35, 2346.5, 2295.35, 2346.92}},
	{date: "2018/1/29", data: [4]float32{2347.22, 2358.98, 2337.35, 2363.8}},
	{date: "2018/1/30", data: [4]float32{2360.75, 2382.48, 2347.89, 2383.76}},
	{date: "2018/1/31", data: [4]float32{2383.43, 2385.42, 2371.23, 2391.82}},
	{date: "2018/2/1", data: [4]float32{2377.41, 2419.02, 2369.57, 2421.15}},
	{date: "2018/2/4", data: [4]float32{2425.92, 2428.15, 2417.58, 2440.38}},
	{date: "2018/2/5", data: [4]float32{2411, 2433.13, 2403.3, 2437.42}},
	{date: "2018/2/6", data: [4]float32{2432.68, 2434.48, 2427.7, 2441.73}},
	{date: "2018/2/7", data: [4]float32{2430.69, 2418.53, 2394.22, 2433.89}},
	{date: "2018/2/8", data: [4]float32{2416.62, 2432.4, 2414.4, 2443.03}},
	{date: "2018/2/18", data: [4]float32{2441.91, 2421.56, 2415.43, 2444.8}},
	{date: "2018/2/19", data: [4]float32{2420.26, 2382.91, 2373.53, 2427.07}},
	{date: "2018/2/20", data: [4]float32{2383.49, 2397.18, 2370.61, 2397.94}},
	{date: "2018/2/21", data: [4]float32{2378.82, 2325.95, 2309.17, 2378.82}},
	{date: "2018/2/22", data: [4]float32{2322.94, 2314.16, 2308.76, 2330.88}},
	{date: "2018/2/25", data: [4]float32{2320.62, 2325.82, 2315.01, 2338.78}},
	{date: "2018/2/26", data: [4]float32{2313.74, 2293.34, 2289.89, 2340.71}},
	{date: "2018/2/27", data: [4]float32{2297.77, 2313.22, 2292.03, 2324.63}},
	{date: "2018/2/28", data: [4]float32{2322.32, 2365.59, 2308.92, 2366.16}},
	{date: "2018/3/1", data: [4]float32{2364.54, 2359.51, 2330.86, 2369.65}},
	{date: "2018/3/4", data: [4]float32{2332.08, 2273.4, 2259.25, 2333.54}},
	{date: "2018/3/5", data: [4]float32{2274.81, 2326.31, 2270.1, 2328.14}},
	{date: "2018/3/6", data: [4]float32{2333.61, 2347.18, 2321.6, 2351.44}},
	{date: "2018/3/7", data: [4]float32{2340.44, 2324.29, 2304.27, 2352.02}},
	{date: "2018/3/8", data: [4]float32{2326.42, 2318.61, 2314.59, 2333.67}},
	{date: "2018/3/11", data: [4]float32{2314.68, 2310.59, 2296.58, 2320.96}},
	{date: "2018/3/12", data: [4]float32{2309.16, 2286.6, 2264.83, 2333.29}},
	{date: "2018/3/13", data: [4]float32{2282.17, 2263.97, 2253.25, 2286.33}},
	{date: "2018/3/14", data: [4]float32{2255.77, 2270.28, 2253.31, 2276.22}},
	{date: "2018/3/15", data: [4]float32{2269.31, 2278.4, 2250, 2312.08}},
	{date: "2018/3/18", data: [4]float32{2267.29, 2240.02, 2239.21, 2276.05}},
	{date: "2018/3/19", data: [4]float32{2244.26, 2257.43, 2232.02, 2261.31}},
	{date: "2018/3/20", data: [4]float32{2257.74, 2317.37, 2257.42, 2317.86}},
	{date: "2018/3/21", data: [4]float32{2318.21, 2324.24, 2311.6, 2330.81}},
	{date: "2018/3/22", data: [4]float32{2321.4, 2328.28, 2314.97, 2332}},
	{date: "2018/3/25", data: [4]float32{2334.74, 2326.72, 2319.91, 2344.89}},
	{date: "2018/3/26", data: [4]float32{2318.58, 2297.67, 2281.12, 2319.99}},
	{date: "2018/3/27", data: [4]float32{2299.38, 2301.26, 2289, 2323.48}},
	{date: "2018/3/28", data: [4]float32{2273.55, 2236.3, 2232.91, 2273.55}},
	{date: "2018/3/29", data: [4]float32{2238.49, 2236.62, 2228.81, 2246.87}},
	{date: "2018/4/1", data: [4]float32{2229.46, 2234.4, 2227.31, 2243.95}},
	{date: "2018/4/2", data: [4]float32{2234.9, 2227.74, 2220.44, 2253.42}},
	{date: "2018/4/3", data: [4]float32{2232.69, 2225.29, 2217.25, 2241.34}},
	{date: "2018/4/8", data: [4]float32{2196.24, 2211.59, 2180.67, 2212.59}},
	{date: "2018/4/9", data: [4]float32{2215.47, 2225.77, 2215.47, 2234.73}},
	{date: "2018/4/10", data: [4]float32{2224.93, 2226.13, 2212.56, 2233.04}},
	{date: "2018/4/11", data: [4]float32{2236.98, 2219.55, 2217.26, 2242.48}},
	{date: "2018/4/12", data: [4]float32{2218.09, 2206.78, 2204.44, 2226.26}},
	{date: "2018/4/15", data: [4]float32{2199.91, 2181.94, 2177.39, 2204.99}},
	{date: "2018/4/16", data: [4]float32{2169.63, 2194.85, 2165.78, 2196.43}},
	{date: "2018/4/17", data: [4]float32{2195.03, 2193.8, 2178.47, 2197.51}},
	{date: "2018/4/18", data: [4]float32{2181.82, 2197.6, 2175.44, 2206.03}},
	{date: "2018/4/19", data: [4]float32{2201.12, 2244.64, 2200.58, 2250.11}},
	{date: "2018/4/22", data: [4]float32{2236.4, 2242.17, 2232.26, 2245.12}},
	{date: "2018/4/23", data: [4]float32{2242.62, 2184.54, 2182.81, 2242.62}},
	{date: "2018/4/24", data: [4]float32{2187.35, 2218.32, 2184.11, 2226.12}},
	{date: "2018/4/25", data: [4]float32{2213.19, 2199.31, 2191.85, 2224.63}},
	{date: "2018/4/26", data: [4]float32{2203.89, 2177.91, 2173.86, 2210.58}},
	{date: "2018/5/2", data: [4]float32{2170.78, 2174.12, 2161.14, 2179.65}},
	{date: "2018/5/3", data: [4]float32{2179.05, 2205.5, 2179.05, 2222.81}},
	{date: "2018/5/6", data: [4]float32{2212.5, 2231.17, 2212.5, 2236.07}},
	{date: "2018/5/7", data: [4]float32{2227.86, 2235.57, 2219.44, 2240.26}},
	{date: "2018/5/8", data: [4]float32{2242.39, 2246.3, 2235.42, 2255.21}},
	{date: "2018/5/9", data: [4]float32{2246.96, 2232.97, 2221.38, 2247.86}},
	{date: "2018/5/10", data: [4]float32{2228.82, 2246.83, 2225.81, 2247.67}},
	{date: "2018/5/13", data: [4]float32{2247.68, 2241.92, 2231.36, 2250.85}},
	{date: "2018/5/14", data: [4]float32{2238.9, 2217.01, 2205.87, 2239.93}},
	{date: "2018/5/15", data: [4]float32{2217.09, 2224.8, 2213.58, 2225.19}},
	{date: "2018/5/16", data: [4]float32{2221.34, 2251.81, 2210.77, 2252.87}},
	{date: "2018/5/17", data: [4]float32{2249.81, 2282.87, 2248.41, 2288.09}},
	{date: "2018/5/20", data: [4]float32{2286.33, 2299.99, 2281.9, 2309.39}},
	{date: "2018/5/21", data: [4]float32{2297.11, 2305.11, 2290.12, 2305.3}},
	{date: "2018/5/22", data: [4]float32{2303.75, 2302.4, 2292.43, 2314.18}},
	{date: "2018/5/23", data: [4]float32{2293.81, 2275.67, 2274.1, 2304.95}},
	{date: "2018/5/24", data: [4]float32{2281.45, 2288.53, 2270.25, 2292.59}},
	{date: "2018/5/27", data: [4]float32{2286.66, 2293.08, 2283.94, 2301.7}},
	{date: "2018/5/28", data: [4]float32{2293.4, 2321.32, 2281.47, 2322.1}},
	{date: "2018/5/29", data: [4]float32{2323.54, 2324.02, 2321.17, 2334.33}},
	{date: "2018/5/30", data: [4]float32{2316.25, 2317.75, 2310.49, 2325.72}},
	{date: "2018/5/31", data: [4]float32{2320.74, 2300.59, 2299.37, 2325.53}},
	{date: "2018/6/3", data: [4]float32{2300.21, 2299.25, 2294.11, 2313.43}},
	{date: "2018/6/4", data: [4]float32{2297.1, 2272.42, 2264.76, 2297.1}},
	{date: "2018/6/5", data: [4]float32{2270.71, 2270.93, 2260.87, 2276.86}},
	{date: "2018/6/6", data: [4]float32{2264.43, 2242.11, 2240.07, 2266.69}},
	{date: "2018/6/7", data: [4]float32{2242.26, 2210.9, 2205.07, 2250.63}},
	{date: "2018/6/13", data: [4]float32{2190.1, 2148.35, 2126.22, 2190.1}},
}

func klineBase() *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([][4]float32, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, kd[i].data)
	}

	kline.AddXAxis(x).AddYAxis("kline", y)
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: "Kline-示例图"},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	return kline
}

func klineDataZoomInside() *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([][4]float32, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, kd[i].data)
	}

	kline.AddXAxis(x).AddYAxis("kline", y)
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: "Kline-DataZoom(inside)"},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	return kline
}

func klineDataZoomBoth() *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([][4]float32, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, kd[i].data)
	}

	kline.AddXAxis(x).AddYAxis("kline", y)
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: "Kline-DataZoom(inside+slider)"},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "inside", XAxisIndex: []int{0}, Start: 50, End: 100},
		charts.DataZoomOpts{Type: "slider", XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	return kline
}

func klineDataZoomYAxis() *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([][4]float32, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, kd[i].data)
	}

	kline.AddXAxis(x).AddYAxis("kline", y)
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: "Kline-DataZoom(yAxis)"},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{Type: "slider", YAxisIndex: []int{0}, Start: 50, End: 100},
	)
	return kline
}

func klineStyle() *charts.Kline {
	kline := charts.NewKLine()

	x := make([]string, 0)
	y := make([][4]float32, 0)
	for i := 0; i < len(kd); i++ {
		x = append(x, kd[i].date)
		y = append(y, kd[i].data)
	}
	kline.AddXAxis(x)
	kline.AddYAxis("kline", y)
	kline.SetGlobalOptions(
		charts.TitleOpts{Title: "Kline-不同风格"},
		charts.XAxisOpts{SplitNumber: 20},
		charts.YAxisOpts{Scale: true},
		charts.DataZoomOpts{XAxisIndex: []int{0}, Start: 50, End: 100},
	)
	kline.SetSeriesOptions(
		charts.MPNameTypeItem{Name: "highest value", Type: "max", ValueDim: "highest"},
		charts.MPNameTypeItem{Name: "lowest value", Type: "min", ValueDim: "lowest"},
		charts.MPStyleOpts{Label: charts.LabelTextOpts{Show: true}},
		charts.ItemStyleOpts{
			Color: "#ec0000", Color0: "#00da3c", BorderColor: "#8A0000", BorderColor0: "#008F28"},
	)
	return kline
}

func klineHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("kline")...)
	page.Add(
		klineBase(),
		klineDataZoomInside(),
		klineDataZoomBoth(),
		klineDataZoomYAxis(),
		klineStyle(),
	)
	f, err := os.Create(getRenderPath("kline.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
