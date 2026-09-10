package query

import (
	"encoding/json"
)

/*
type Node interface {
	MarshalJSON() ([]byte, error)
}

type queryWrapper struct {
	Query *Query `json:"query"`
}

type Query struct {
	Filter struct {
		Text *string `json:"text"`
		And  []And   `json:"and"`
	} `json:"filter"`
}

type And struct {
	Equal struct {
		Path  string `json:"path"`
		Value string `json:"value"`
	} `json:"equal,omitempty"`
	Ref struct {
		Path  string `json:"path"`
		Value []struct {
			Oid string `json:"oid"`
		} `json:"value"`
	} `json:"ref,omitempty"`
}

*/

type QueryWrapper struct {
	Query *Query `json:"query,omitempty"`
}

type Query struct {
	Filter *Filter `json:"filter,omitempty"`
}

func (q QueryWrapper) MarshalJSON() ([]byte, error) {
	if q.Query == nil {
		return []byte(`{"query":""}`), nil
	}
	return json.Marshal(struct {
		Query *Query `json:"query"`
	}{Query: q.Query})
}

type Filter struct {
	Text *string `json:"text,omitempty"`
}

func (q QueryWrapper) String() string {
	value, _ := json.Marshal(q)
	return string(value)
}

func New(text string) *QueryWrapper {
	if text == "" || text == "*" {
		return &QueryWrapper{Query: nil}
	}
	return &QueryWrapper{
		Query: &Query{
			Filter: &Filter{
				Text: new(text),
			},
		},
	}
}
