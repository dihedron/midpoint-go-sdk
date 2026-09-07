package user

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
)

type Policy struct {
	base.Command
}

func (cmd *Policy) Execute(args []string) error {
	slog.Debug("running user password policy command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password, "ids", args)
	if len(args) == 0 {
		slog.Error("no ids provided")
		return fmt.Errorf("at least one ID must be provided")
	}

	mp := cmd.GetAPI()
	defer mp.Close()

	var result error
	for _, arg := range args {
		slog.Debug("reading user", "id", arg)
		self, err := mp.User.ReadPasswordPolicy(context.Background(), arg)
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
