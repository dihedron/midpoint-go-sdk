package self

import (
	"context"
	"log/slog"
	"os"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
	"github.com/dihedron/midpoint-go-sdk/pkg/midpoint"
)

type Read struct {
	base.Command
}

func (cmd *Read) Execute(args []string) error {
	slog.Debug("running self command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password)
	options := []midpoint.Option{
		midpoint.WithDebug(cmd.Debug),
	}
	if cmd.Impersonate != nil {
		options = append(options, midpoint.WithImpersonation(*cmd.Impersonate))
	}
	mp := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password, options...)
	self, err := mp.Self.Read(context.Background())
	if err != nil {
		slog.Error("error reading self", "error", err)
		return err
	}
	return cmd.Write(os.Stdout, self)
}
