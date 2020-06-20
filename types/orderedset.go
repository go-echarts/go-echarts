package types

// OrderedSet represents an ordered set.
type OrderedSet struct {
	filter map[string]bool
	Values []string
}

// Init creates a new OrderedSet instance, and adds any given items into this set.
func (o *OrderedSet) Init(items ...string) {
	o.filter = make(map[string]bool)
	for _, item := range items {
		o.Add(item)
	}
}

// Add adds a new item into the ordered set.
func (o *OrderedSet) Add(item string) {
	if !o.filter[item] {
		o.filter[item] = true
		o.Values = append(o.Values, item)
	}
}
