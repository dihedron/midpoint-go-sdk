package users

// Commands is the set of root command groups.
type Users struct {
	Read   Read   `command:"read" alias:"r" description:"Read a MidPoint user object."`
	Create Create `command:"create" alias:"c" description:"Create a new MidPoint user object."`
	// Users users.Users `command:"users" alias:"user" alias:"u" description:"Manage MidPoint users."`
	// Self  self.Read   `command:"self" alias:"s" description:"View information about oneself."`
}
