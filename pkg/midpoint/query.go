package midpoint

import "encoding/json"

type query struct {
	Data data `json:"query"`
}

type data struct {
	Filter filter `json:"filter"`
}

type filter struct {
	Text string `json:"text"`
}

func (q query) String() string {
	value, _ := json.Marshal(q)
	return string(value)
}

func Query(text string) string {
	if text == "" || text == "*" {
		return `{"query":""}`
	}
	return query{
		Data: data{
			Filter: filter{
				Text: text,
			},
		},
	}.String()
}
