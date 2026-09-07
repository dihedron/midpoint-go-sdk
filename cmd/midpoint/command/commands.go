package command

import (
	"github.com/dihedron/midpoint-go-sdk/cmd/midpoint/command/roles"
	"github.com/dihedron/midpoint-go-sdk/cmd/midpoint/command/self"
	"github.com/dihedron/midpoint-go-sdk/cmd/midpoint/command/user"
	"github.com/dihedron/midpoint-go-sdk/internal/command/version"
)

// Commands is the set of root command groups.
type Commands struct {
	// Role manages MidPoint roles.
	Role roles.Roles `command:"role" alias:"r" description:"Manage MidPoint roles."`
	// User manages MidPoint users.
	User user.User `command:"user" alias:"u" description:"Manage MidPoint users."`
	// Self prints information about the current user.
	Self self.Read `command:"self" alias:"s" description:"View information about oneself."`
	// Version prints the program version information.
	Version version.Version `command:"version" alias:"v" description:"Print program version information."`
}
