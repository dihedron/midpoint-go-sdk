package midpoint

type UserService struct {
	Service
}

type User struct {
	Name       *string `json:"name,omitempty" yaml:"name,omitempty"`
	FullName   *string `json:"fullName,omitempty" yaml:"fullName,omitempty"`
	GivenName  *string `json:"givenName,omitempty" yaml:"givenName,omitempty"`
	FamilyName *string `json:"familyName,omitempty" yaml:"familyName,omitempty"`
}

type CreateUserOptions struct {
	Name    *string
	Surname *string
}

/*
func (s *UserService) Create(opts CreateUserOptions) error {
	body := struct {
		User User `json:"user"`
	}{
		User: {},
	}
	s.client.client.Post()
}
*/
