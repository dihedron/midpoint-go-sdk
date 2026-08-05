package midpoint

type Query struct {
	Data Data `json:"query"`
}

type Data struct {
	Filter Filter `json:"filter"`
}

type Filter struct {
	Text string `json:"text"`
}

func NewQuery(text string) Query {
	return Query{
		Data: Data{
			Filter: Filter{
				Text: text,
			},
		},
	}
}
