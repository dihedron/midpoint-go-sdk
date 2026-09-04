package users

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
	"github.com/dihedron/midpoint-go-sdk/pkg/midpoint"
)

type Read struct {
	base.Command
}

func (cmd *Read) Execute(args []string) error {
	slog.Debug("running user read command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password, "ids", args)
	if len(args) == 0 {
		slog.Error("no ids provided")
		return fmt.Errorf("at least one ID must be provided")
	}
	options := []midpoint.Option{
		midpoint.WithDebug(cmd.Debug),
	}
	if cmd.Impersonate != nil {
		options = append(options, midpoint.WithImpersonation(*cmd.Impersonate))
	}
	mp := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password, options...)

	var result error
	for _, arg := range args {
		slog.Debug("reading user", "id", arg)
		self, err := mp.User.Read(context.Background(), arg)
		if err != nil {
			slog.Error("error reading user", "id", arg, "error", err)
			errors.Join(result, err)
			continue
		}
		if err = cmd.Write(os.Stdout, self); err != nil {
			errors.Join(result, err)
		}
	}
	return result
}
