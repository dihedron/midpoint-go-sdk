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
	mp := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password)
	self, err := mp.Self.Read(context.Background())
	if err != nil {
		slog.Error("error reading self", "error", err)
		return err
	}
	return cmd.Write(os.Stdout, self)
}
