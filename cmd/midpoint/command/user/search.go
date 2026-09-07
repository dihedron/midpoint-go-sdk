package user

import (
	"context"
	"log/slog"
	"os"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
	"github.com/dihedron/midpoint-go-sdk/pkg/midpoint"
)

type Search struct {
	base.Command
	Query string `short:"q" long:"query" description:"The query to filter results." required:"true" default:"*"`
}

func (cmd *Search) Execute(args []string) error {
	slog.Debug("running user search command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password)
	options := []midpoint.Option{
		midpoint.WithDebug(cmd.Debug),
	}
	if cmd.Impersonate != nil {
		options = append(options, midpoint.WithImpersonation(*cmd.Impersonate))
	}
	mp := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password, options...)

	users, err := mp.User.Search(context.Background(), cmd.Query)
	if err != nil {
		slog.Error("error searching for users", "query", cmd.Query, "error", err)
		return err
	}
	if err = cmd.Write(os.Stdout, users); err != nil {
		return err
	}
	return nil
}
