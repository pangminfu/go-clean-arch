package main

import (
	"github.com/pangminfu/go-clean-arch/reusable"
	"github.com/urfave/cli"
)

func init() {
	flags := []cli.Flag{
		cli.IntFlag{
			Name:  "id",
			Usage: "product id",
		},
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
				p := handleArgs(c)
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				_, err := svc.Create(p)
				return err
			},
		},
		{
			Name:  "list",
			Usage: "list all product",
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
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "code",
					Usage: "product code",
				},
			},
			Action: func(c *cli.Context) error {
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				_, err := svc.SearchByCode(c.String("code"))
				return err
			},
		},
		{
			Name:  "update",
			Usage: "update product",
			Flags: flags,
			Action: func(c *cli.Context) error {
				p := handleArgs(c)
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				_, err := svc.UpdateProduct(p)
				return err
			},
		},
		{
			Name:  "delete",
			Usage: "delete product",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "id",
					Usage: "product id",
				},
			},
			Action: func(c *cli.Context) error {
				repo := reusable.NewInMemProductRepository(nil)
				svc := reusable.NewService(repo)
				err := svc.DeleteProduct(c.Int("id"))
				return err
			},
		},
	}
}

func handleArgs(c *cli.Context) *reusable.Product {
	product := &reusable.Product{
		Id:   c.Int("id"),
		Code: c.String("code"),
		Name: c.String("name"),
		Desc: c.String("desc"),
	}

	return product
}
