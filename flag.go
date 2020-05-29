package main

import "github.com/urfave/cli/v2"

func NewFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Required: true,
		},
		&cli.BoolFlag{
			Name: "docker",
			Aliases:  []string{"d"},
		},
		&cli.BoolFlag{
			Name: "docker-compose",
			Aliases:  []string{"c"},
		},
	}
}
