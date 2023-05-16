package types

// OrderedSet represents an ordered set.
type OrderedSet struct {
	filter map[string]bool
	Values []string
}

// Add adds a new item into the ordered set.
func (o *OrderedSet) Add(item string) *OrderedSet {
	if o.filter == nil {
		o.filter = make(map[string]bool)
	}
	if !o.filter[item] {
		o.filter[item] = true
		o.Values = append(o.Values, item)
	}
	return o
}
