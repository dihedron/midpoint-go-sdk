package users

import (
	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
)

type Users struct {
	base.Command
}

func (cmd *Users) Execute(args []string) error {
	/*
		slog.Debug("running self command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password)
		mp := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password)
		self, err := mp.User.Read(context.Background())
		if err != nil {
			slog.Error("error reading self", "error", err)
			return err
		}
		return cmd.Write(os.Stdout, self)
	*/
	return nil
}
