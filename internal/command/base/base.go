package base

import (
	"encoding/json/jsontext"
	"encoding/json/v2"
	"fmt"
	"io"
	"log/slog"

	"github.com/joho/godotenv"
	"go.yaml.in/yaml/v3"
)

// Command is the base command.
type Command struct {
	Endpoint string  `short:"E" long:"endpoint" description:"The network endpoint for midPoint API access." required:"true" env:"MIDPOINT_ENDPOINT"`
	Username string  `short:"U" long:"username" description:"The username for authentication." required:"true" env:"MIDPOINT_USERNAME"`
	Password string  `short:"P" long:"password" description:"The password for authentication." required:"true" env:"MIDPOINT_PASSWORD"`
	Format   string  `short:"F" long:"format" description:"The format of the output." optional:"true" default:"yaml" choice:"text" choice:"json" choice:"yaml" choice:"none" env:"MIDPOINT_FORMAT"`
	DotEnv   *string `short:"D" long:"dotenv" description:"The path to the .env file." optional:"true" env:"MIDPOINT_DOTENV"`
}

func (cmd *Command) LoadEnv() error {
	if cmd.DotEnv == nil {
		return nil
	}
	if err := godotenv.Load(*cmd.DotEnv); err != nil {
		slog.Error("error loading .env file", "error", err)
	}
	slog.Info("successfully loaded .env file", "path", *cmd.DotEnv)
	return nil
}

func (cmd *Command) Write(stream io.Writer, object any) error {
	switch cmd.Format {
	case "yaml":
		data, err := yaml.Marshal(object)
		if err != nil {
			return err
		}
		fmt.Fprintf(stream, "%s", string(data))
	case "json":
		data, err := json.Marshal(
			object,
			json.OmitZeroStructFields(true),
			jsontext.WithIndent("  "),
		)
		if err != nil {
			return err
		}
		fmt.Fprintf(stream, "%s\n", string(data))
	case "text":
		fmt.Fprintf(stream, "%v\n", object)
	default:
		return nil
	}
	return nil
}
