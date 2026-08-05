package roles

type Self struct{}

func (cmd *Self) Execute(args []string) error {
	//http://localhost:8080/midpoint/ws/rest/self?options=raw

	return nil
}
