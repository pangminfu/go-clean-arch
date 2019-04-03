package main

import "github.com/urfave/cli"

func init() {
	cmds = []cli.Command{
		{
			Name:  "create",
			Usage: "create new produt",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "pub",
					Usage: "publication code",
				},
				cli.StringFlag{
					Name:  "host",
					Usage: "FTP host name",
				},
				cli.StringFlag{
					Name:  "user",
					Usage: "FTP username",
				},
				cli.StringFlag{
					Name:  "pass",
					Usage: "FTP password",
				},
				cli.StringFlag{
					Name:  "port",
					Value: "21",
					Usage: "FTP port number default is 21",
				},
				cli.StringFlag{
					Name:  "directory, dir",
					Usage: "FTP directory to sync",
				},
			},
			// Action: epapersyncCLI,
		},
	}
}
