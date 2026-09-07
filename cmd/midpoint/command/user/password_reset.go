package user

import (
	"context"
	"log/slog"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
)

type Reset struct {
	//User *midpoint.User `short:"d" long:"data" description:"New user's data, either as an inline value or as a @file (in JSON or YAML format)."`
	base.Command
	Positional struct {
		Id       string `positional-arg-name:"USERID"`
		Password string `positional-arg-name:"PASSWORD"`
	} `positional-args:"true" required:"true"`
}

func (cmd *Reset) Execute(args []string) error {
	slog.Debug("running user password reset command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password, "id", cmd.Positional.Id, "password", cmd.Positional.Password)

	mp := cmd.GetAPI()
	defer mp.Close()

	err := mp.User.ResetPassword(context.Background(), cmd.Positional.Id, cmd.Positional.Password)
	if err != nil {
		slog.Error("failed to create user", "error", err)
		return err
	}
	slog.Debug("user's password reset", "id", cmd.Positional.Id, "password", cmd.Positional.Password)
	//cmd.Write(os.Stdout, midpoint.ID{ID: id})
	return nil
}
