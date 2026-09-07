package user

import (
	"context"
	"log/slog"
	"os"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
	"github.com/dihedron/midpoint-go-sdk/pkg/midpoint"
)

type Create struct {
	User *midpoint.User `short:"d" long:"data" description:"New user's data, either as an inline value or as a @file (in JSON or YAML format)."`
	base.Command
}

func (cmd *Create) Execute(args []string) error {
	slog.Debug("running user create command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password, "user", *cmd.User)
	mp := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password)
	id, err := mp.User.Create(context.Background(), cmd.User)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return err
	}
	slog.Debug("user created", "id", id)
	cmd.Write(os.Stdout, midpoint.ID{ID: id})
	return nil
}
