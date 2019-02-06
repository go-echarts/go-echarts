package common

// Name-Value 数据项
type NameValueItem struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// 有序去重集合
type OrderedSet struct {
	filter map[string]bool
	Values []string
}

// 带参数初始化
func (o *OrderedSet) Init(item string) {
	o.filter = make(map[string]bool)
	o.Add(item)
}

// 不带参数初始化
func (o *OrderedSet) InitWithoutArg() {
	o.filter = make(map[string]bool)
}

// 新增 item
func (o *OrderedSet) Add(item string) {
	if !o.filter[item] {
		o.filter[item] = true
		o.Values = append(o.Values, item)
	}
}
