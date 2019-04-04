package datatypes

// NameValueItem represents a data type with a name and a value.
type NameValueItem struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
