package charts

// Name-Value 数据项
type nameValueItem struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

// 有序去重集合
type orderSet struct {
	filter map[string]bool
	Values []string
}

func (o *orderSet) init(item string) {
	o.filter = make(map[string]bool)
	o.Add(item)
}

func (o *orderSet) initWithoutArg() {
	o.filter = make(map[string]bool)
}

func (o *orderSet) Add(item string) {
	if !o.filter[item] {
		o.filter[item] = true
		o.Values = append(o.Values, item)
	}
}
