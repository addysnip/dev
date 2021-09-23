/*
Simple Go Get Alias
Copyright (C) 2021 Daniel A. Hawton (daniel@hawton.org)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

package main

import (
	"os"

	"github.com/dhawton/hawton.dev/internal/server"
	"github.com/dhawton/log4g"
	"github.com/urfave/cli/v2"
)

var log = log4g.Category("main")

func main() {
	app := &cli.App{
		Name:  "hawton",
		Usage: "Hawton Go Get Alias",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "Set log level to debug",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "Start Server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "Set port to listen on",
						Value:   3000,
					},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("debug") {
						log4g.SetLogLevel(log4g.DEBUG)
					}
					log.Info("Starting server")
					server.Run(c.Int("port"))
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}
