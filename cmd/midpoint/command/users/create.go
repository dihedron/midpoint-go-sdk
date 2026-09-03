package users

import (
	"context"
	"log/slog"
	"os"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
	"github.com/dihedron/midpoint-go-sdk/pkg/midpoint"
)

type Create struct {
	User *midpoint.User `short:"u" long:"user" description:"New user's data, either as an inline value or as a @file (in JSON or YAML format)."`
	base.Command
}

func (cmd *Create) Execute(args []string) error {
	slog.Debug("running user create command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password, "name", args)
	mp := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password)
	user := cmd.User
	cmd.Write(os.Stdout, user)
	return mp.User.Create(context.Background(), user)
}
