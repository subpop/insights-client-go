package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "3.99"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:      "config,c",
			Usage:     "load configuration from `FILE`",
			Value:     "/etc/insights-client/insights-client.conf",
			TakesFile: true,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "update",
			Usage: "update to latest core",
			Action: func(c *cli.Context) error {
				cfg := defaultConfig()
				return update(cfg)
			},
		},
		{
			Name:  "collect",
			Usage: "run data collection",
			Action: func(c *cli.Context) error {
				return collect()
			},
		},
		{
			Name:  "upload",
			Usage: "upload an archive",
			Action: func(c *cli.Context) error {
				archivePath := c.Args().First()
				if archivePath == "" {
					return fmt.Errorf("error: path to archive is required")
				}
				cfg := defaultConfig()
				return upload(cfg, archivePath)
			},
		},
		{
			Name:  "show",
			Usage: "show insights about this machine",
			Action: func(c *cli.Context) error {
				cfg := defaultConfig()
				return show(cfg)
			},
		},
		{
			Name: "status",
		},
		{
			Name: "unregister",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
