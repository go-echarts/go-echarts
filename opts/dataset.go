package opts

// Dataset brings convenience in data management separated with styles and enables data reuse by different series.
// More importantly, it enables data encoding from data to visual, which brings convenience in some scenarios.
// https://echarts.apache.org/en/option.html#dataset
type Dataset struct {
	// source
	Source interface{} `json:"source"`
}
