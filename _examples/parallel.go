package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/charts"
)

var parallelDataBJ = [][]interface{}{
	{1, 55, 9, 56, 0.46, 18, 6, "良"},
	{2, 25, 11, 21, 0.65, 34, 9, "优"},
	{3, 56, 7, 63, 0.3, 14, 5, "良"},
	{4, 33, 7, 29, 0.33, 16, 6, "优"},
	{5, 42, 24, 44, 0.76, 40, 16, "优"},
	{6, 82, 58, 90, 1.77, 68, 33, "良"},
	{7, 74, 49, 77, 1.46, 48, 27, "良"},
	{8, 78, 55, 80, 1.29, 59, 29, "良"},
	{9, 267, 216, 280, 4.8, 108, 64, "重度污染"},
	{10, 185, 127, 216, 2.52, 61, 27, "中度污染"},
	{11, 39, 19, 38, 0.57, 31, 15, "优"},
	{12, 41, 11, 40, 0.43, 21, 7, "优"},
	{13, 64, 38, 74, 1.04, 46, 22, "良"},
	{14, 108, 79, 120, 1.7, 75, 41, "轻度污染"},
	{15, 108, 63, 116, 1.48, 44, 26, "轻度污染"},
	{16, 33, 6, 29, 0.34, 13, 5, "优"},
	{17, 94, 66, 110, 1.54, 62, 31, "良"},
	{18, 186, 142, 192, 3.88, 93, 79, "中度污染"},
	{19, 57, 31, 54, 0.96, 32, 14, "良"},
	{20, 22, 8, 17, 0.48, 23, 10, "优"},
	{21, 39, 15, 36, 0.61, 29, 13, "优"},
}

var parallelDataGZ = [][]interface{}{
	{1, 26, 37, 27, 1.163, 27, 13, "优"},
	{2, 85, 62, 71, 1.195, 60, 8, "良"},
	{3, 78, 38, 74, 1.363, 37, 7, "良"},
	{4, 21, 21, 36, 0.634, 40, 9, "优"},
	{5, 41, 42, 46, 0.915, 81, 13, "优"},
	{6, 56, 52, 69, 1.067, 92, 16, "良"},
	{7, 64, 30, 28, 0.924, 51, 2, "良"},
	{8, 55, 48, 74, 1.236, 75, 26, "良"},
	{9, 76, 85, 113, 1.237, 114, 27, "良"},
	{10, 91, 81, 104, 1.041, 56, 40, "良"},
	{11, 84, 39, 60, 0.964, 25, 11, "良"},
	{12, 64, 51, 101, 0.862, 58, 23, "良"},
	{13, 70, 69, 120, 1.198, 65, 36, "良"},
	{14, 77, 105, 178, 2.549, 64, 16, "良"},
	{15, 109, 68, 87, 0.996, 74, 29, "轻度污染"},
	{16, 73, 68, 97, 0.905, 51, 34, "良"},
	{17, 54, 27, 47, 0.592, 53, 12, "良"},
	{18, 51, 61, 97, 0.811, 65, 19, "良"},
	{19, 91, 71, 121, 1.374, 43, 18, "良"},
	{20, 73, 102, 182, 2.787, 44, 19, "良"},
	{21, 73, 50, 76, 0.717, 31, 20, "良"},
}

var parallelDataSH = [][]interface{}{
	{1, 91, 45, 125, 0.82, 34, 23, "良"},
	{2, 65, 27, 78, 0.86, 45, 29, "良"},
	{3, 83, 60, 84, 1.09, 73, 27, "良"},
	{4, 109, 81, 121, 1.28, 68, 51, "轻度污染"},
	{5, 106, 77, 114, 1.07, 55, 51, "轻度污染"},
	{6, 109, 81, 121, 1.28, 68, 51, "轻度污染"},
	{7, 106, 77, 114, 1.07, 55, 51, "轻度污染"},
	{8, 89, 65, 78, 0.86, 51, 26, "良"},
	{9, 53, 33, 47, 0.64, 50, 17, "良"},
	{10, 80, 55, 80, 1.01, 75, 24, "良"},
	{11, 117, 81, 124, 1.03, 45, 24, "轻度污染"},
	{12, 99, 71, 142, 1.1, 62, 42, "良"},
	{13, 95, 69, 130, 1.28, 74, 50, "良"},
	{14, 116, 87, 131, 1.47, 84, 40, "轻度污染"},
	{15, 108, 80, 121, 1.3, 85, 37, "轻度污染"},
	{16, 134, 83, 167, 1.16, 57, 43, "轻度污染"},
	{17, 79, 43, 107, 1.05, 59, 37, "良"},
	{18, 71, 46, 89, 0.86, 64, 25, "良"},
	{19, 97, 71, 113, 1.17, 88, 31, "良"},
	{20, 84, 57, 91, 0.85, 55, 31, "良"},
	{21, 87, 63, 101, 0.9, 56, 41, "良"},
}

var parallelAxis = []charts.PAOpts{
	{Dim: 0, Name: "日期", Inverse: true, Max: 31, NameLocation: "start"},
	{Dim: 1, Name: "AQI"},
	{Dim: 2, Name: "PM2.5"},
	{Dim: 3, Name: "PM10"},
	{Dim: 4, Name: "CO"},
	{Dim: 5, Name: "NO2"},
	{Dim: 6, Name: "SO2"},
	{Dim: 7, Name: "等级", Type: "category",
		Data: []string{"优", "良", "轻度污染", "中度污染", "重度污染", "严重污染"}},
}

func parallelBase() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(
		charts.TitleOpts{Title: "Parallel-示例图"},
		charts.ParallelAxisOpts{
			{Dim: 0, Name: "日期", Inverse: true, Max: 31, NameLocation: "start"},
			{Dim: 1, Name: "AQI"},
			{Dim: 2, Name: "PM2.5"},
			{Dim: 3, Name: "PM10"},
			{Dim: 4, Name: "CO"},
			{Dim: 5, Name: "NO2"},
			{Dim: 6, Name: "SO2"},
			{Dim: 7, Name: "等级", Type: "category",
				Data: []string{"优", "良", "轻度污染", "中度污染", "重度污染", "严重污染"}},
		},
	)
	parallel.Add("北京", parallelDataBJ)
	return parallel
}

func parallelComponent() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(
		charts.TitleOpts{Title: "Parallel-ParallelComponent"},
		charts.ParallelComponentOpts{
			Left:   "15%",
			Right:  "13%",
			Bottom: "10%",
			Top:    "20%",
		},
	)
	parallel.ParallelAxisOpts = parallelAxis
	parallel.Add("北京", parallelDataBJ)
	return parallel
}

func parallelMulti() *charts.Parallel {
	parallel := charts.NewParallel()
	parallel.SetGlobalOptions(charts.TitleOpts{Title: "Parallel-多 Series"})
	parallel.ParallelAxisOpts = parallelAxis
	parallel.
		Add("北京", parallelDataBJ).
		Add("广州", parallelDataGZ).
		Add("上海", parallelDataSH)
	return parallel
}

func parallelHandler(w http.ResponseWriter, _ *http.Request) {
	page := charts.NewPage(orderRouters("parallel")...)
	page.Add(
		parallelBase(),
		parallelComponent(),
		parallelMulti(),
	)
	f, err := os.Create(getRenderPath("parallel.html"))
	if err != nil {
		log.Println(err)
	}
	page.Render(w, f)
}
