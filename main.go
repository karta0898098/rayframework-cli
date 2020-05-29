package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

type (
	Config struct {
		Name             string
		UseDocker        bool
		UseDockerCompose bool
	}
)

var config Config

func main() {
	app := cli.NewApp()
	app.Name = "rayframework-cli"
	app.Usage = ""
	app.Flags = NewFlag()
	app.Action = Run
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Sorry can't create project reason:", err)
	}
}

func Run(c *cli.Context) error {
	config = Config{
		Name:             c.String("name"),
		UseDocker:        c.Bool("docker"),
		UseDockerCompose: c.Bool("docker-compose"),
	}
	return Exec()
}

func Exec() error {

	builder := &AppBuilder{}

	builder = builder.Name(config.Name)
	builder = builder.WorkingDir()
	builder = builder.Folder()
	builder = builder.Database()
	builder = builder.Router()
	builder = builder.Util()
	builder = builder.Templates()

	if config.UseDocker {
		builder = builder.Docker()
	}

	builder = builder.Config()
	builder = builder.Main()
	builder.Build()
	return nil
}
