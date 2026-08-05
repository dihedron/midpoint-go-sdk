package users

import (
	"context"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
	"github.com/dihedron/midpoint-go-sdk/pkg/midpoint"
)

type Users struct {
	base.Command
}

func (cmd *Users) Execute(args []string) error {
	//http://localhost:8080/midpoint/ws/rest/self?options=raw
	client := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password)
	client.Do(context.Background(), "GET", "self", nil, nil)

	return nil
}
