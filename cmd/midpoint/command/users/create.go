package users

import (
	"context"
	"log/slog"
	"os"

	"github.com/dihedron/midpoint-go-sdk/internal/command/base"
	"github.com/dihedron/midpoint-go-sdk/pkg/midpoint"
)

type Create struct {
	// Name       *string        `short:"n" long:"name" description:"The username." `
	// GivenName  *string        `short:"g" long:"name" description:"The given name of the user."`
	// FamilyName *string        `short:"f" long:"name" description:"The family name of the user."`
	// FullName   *string        `short:"u" long:"name" description:"The full (given + family) name of the user."`
	User *midpoint.User `short:"u" long:"user" description:"New user's data, either as an inline value or as a @file (in JSON or YAML format)."`
	base.Command
}

func (cmd *Create) Execute(args []string) error {
	slog.Debug("running user create command", "endpoint", cmd.Endpoint, "username", cmd.Username, "password", cmd.Password, "name", args)
	mp := midpoint.New(cmd.Endpoint, cmd.Username, cmd.Password)

	user := cmd.User
	cmd.Write(os.Stdout, user)
	// if cmd.User == nil {
	// 	user = &midpoint.User{
	// 		Name:      cmd.Name,
	// 		GivenName: cmd.GivenName,
	// 	}
	// }
	return mp.User.Create(context.Background(), user)
	// self, err := mp.User.Read(context.Background(), arg)
	// var result error
	// for _, arg := range args {
	// 	slog.Debug("reading user", "id", arg)
	// 	self, err := mp.User.Read(context.Background(), arg)
	// 	if err != nil {
	// 		slog.Error("error reading user", "id", arg, "error", err)
	// 		errors.Join(result, err)
	// 		continue
	// 	}
	// 	if err = cmd.Write(os.Stdout, self); err != nil {
	// 		errors.Join(result, err)
	// 	}
	// }
	//return nil
}
