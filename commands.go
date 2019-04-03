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
			Usage: "create new product",
			Flags: flags,
			Action: func(c *cli.Context) error {
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				_, err := svc.Create(nil)
				return err
			},
		},
		{
			Name:  "list",
			Usage: "list all product",
			Flags: flags,
			Action: func(c *cli.Context) error {
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				_, err := svc.ListProduct()
				return err
			},
		},
		{
			Name:  "search",
			Usage: "search product by code",
			Flags: flags,
			Action: func(c *cli.Context) error {
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				_, err := svc.SearchByCode("")
				return err
			},
		},
		{
			Name:  "update",
			Usage: "update product",
			Flags: flags,
			Action: func(c *cli.Context) error {
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				_, err := svc.UpdateProduct(nil)
				return err
			},
		},
		{
			Name:  "delete",
			Usage: "delete product",
			Flags: flags,
			Action: func(c *cli.Context) error {
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				err := svc.DeleteProduct(0)
				return err
			},
		},
	}
}
