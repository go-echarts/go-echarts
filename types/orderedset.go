package types

type Index struct {
	index int
}

// OrderedSet represents an ordered set.
type OrderedSet struct {
	filter map[string]*Index
	cur    int
	Values []string
}

// Init creates a new OrderedSet instance, and adds any given items into this set.
func (o *OrderedSet) Init(items ...string) {
	o.filter = make(map[string]*Index)
	o.cur = 0
	for _, item := range items {
		o.Add(item)
	}
}

// Add adds a new item into the ordered set
func (o *OrderedSet) Add(item string) {
	if o.filter[item] == nil {
		o.filter[item] = &Index{
			index: o.cur,
		}
		o.cur++
		o.Values = append(o.Values, item)
	}
}

func (o *OrderedSet) Remove(item string) {
	if o.filter[item] != nil {
		idx := o.filter[item].index
		o.Values = append(o.Values[:idx], o.Values[idx+1:]...)

		renew := &OrderedSet{}
		renew.Init(o.Values...)

		o.cur = renew.cur
		o.filter = renew.filter
		o.Values = renew.Values
	}
}

func (o *OrderedSet) Size() int {
	return o.cur
}

func (o *OrderedSet) Contains(item string) bool {
	return o.filter[item] != nil
}
