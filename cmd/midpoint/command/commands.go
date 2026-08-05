package command

import (
	"github.com/dihedron/midpoint-go-sdk/cmd/midpoint/command/roles"
	"github.com/dihedron/midpoint-go-sdk/cmd/midpoint/command/users"
)

// Commands is the set of root command groups.
type Commands struct {
	Roles roles.Roles `command:"roles" alias:"role" alias:"r" description:"Manage midPoint roles."`
	Users users.Users `command:"uers" alias:"user" alias:"u" description:"Manage midPoint roles."`
}
