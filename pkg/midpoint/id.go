package midpoint

import "fmt"

type ID struct {
	ID string `json:"id" yaml:"id"`
}

func (id ID) String() string {
	return fmt.Sprintf("id: %s", id.ID)
}
