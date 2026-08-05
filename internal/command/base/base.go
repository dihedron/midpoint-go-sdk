package base

// Command is the base command.
type Command struct {
	Endpoint string `short:"E" long:"endpoint" description:"The network endpoint for midPoint API access." required:"true" env:"MIDPOINT_ENDPOINT"`
	Username string `short:"U" long:"username" description:"The username for authentication." required:"true" env:"MIDPOINT_USERNAME"`
	Password string `short:"P" long:"password" description:"The password for authentication." required:"true" env:"MIDPOINT_PASSWORD"`
	//Format   string `short:"F" long:"format" description:"The format of the output." optional:"true" default:"none" choice:"text" choice:"json" choice:"yaml" choice:"csv" choice:"excel" choice:"none" env:"EXCEL_FORMAT"`
}
