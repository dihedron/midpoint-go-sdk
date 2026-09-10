package query

import "testing"

func TestNewQuery(t *testing.T) {
	tests := []struct {
		Query    string
		Expected string
	}{
		{
			Query:    `*`,
			Expected: `{"query":""}`,
		},
		{
			Query:    ``,
			Expected: `{"query":""}`,
		},
		{
			Query:    `givenName startsWith "A"`,
			Expected: `{"query":{"filter":{"text":"givenName startsWith \"A\""}}}`,
		},
		{
			Query:    `requestable = true AND parentOrgRef/@/name = "Role Catalog"`,
			Expected: `{"query":{"filter":{"text":"requestable = true AND parentOrgRef/@/name = \"Role Catalog\""}}}`,
		},
	}

	for _, test := range tests {
		actual := New(test.Query).String()
		expected := test.Expected
		if actual != expected {
			t.Logf("expected: >%s<", expected)
			t.Logf("actual  : >%s<", actual)
			t.Fail()
		} else {
			t.Logf("Query >%s< passed: >%s< == >%s<", test.Query, actual, expected)
		}
	}
}
