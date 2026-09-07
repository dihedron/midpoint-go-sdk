package midpoint

import "testing"

func TestIDString(t *testing.T) {
	value := ID{ID: "1234"}.String()
	if value != "id: 1234" {
		t.Fail()
	}
}
