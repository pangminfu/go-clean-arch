package main

import (
	"github.com/pangminfu/go-clean-arch/reusable"
	"github.com/urfave/cli"
)

func init() {
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "code",
			Usage: "product code",
		},
		cli.StringFlag{
			Name:  "name",
			Usage: "product name",
		},
		cli.StringFlag{
			Name:  "desc",
			Usage: "product description",
		},
	}

	cmds = []cli.Command{
		{
			Name:  "create",
			Usage: "create new produt",
			Flags: flags,
			Action: func(c *cli.Context) error {
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				_, err := svc.Create(nil)
				return err
			},
		},
	}
}
