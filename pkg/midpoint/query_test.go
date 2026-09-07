package midpoint

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
	}

	for _, test := range tests {
		actual := Query(test.Query)
		expected := test.Expected
		if actual != expected {
			t.Logf("expected: >%s<", expected)
			t.Logf("actual  : >%s<", actual)
			t.Fail()
		}
	}
}
