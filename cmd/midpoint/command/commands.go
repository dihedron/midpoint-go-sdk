package command

import (
	"github.com/dihedron/midpoint-go-sdk/cmd/midpoint/command/roles"
	"github.com/dihedron/midpoint-go-sdk/cmd/midpoint/command/self"
	"github.com/dihedron/midpoint-go-sdk/cmd/midpoint/command/users"
	"github.com/dihedron/midpoint-go-sdk/internal/command/version"
)

// Commands is the set of root command groups.
type Commands struct {
	// Roles manages MidPoint roles.
	//lint:ignore SA5008 duplicate alias tags are legitimate
	Roles roles.Roles `command:"roles" alias:"role" alias:"r" description:"Manage MidPoint roles."`
	// Users manages MidPoint users.
	//lint:ignore SA5008 duplicate alias tags are legitimate
	Users users.Users `command:"users" alias:"user" alias:"u" description:"Manage MidPoint users."`
	// Self prints information about the current user.
	Self self.Read `command:"self" alias:"s" description:"View information about oneself."`
	// Version prints the program version information.
	//lint:ignore SA5008 duplicate alias tags are legitimate
	Version version.Version `command:"version" alias:"ver" alias:"v" description:"Print program version information."`
}
