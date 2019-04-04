package main

import (
	"log"

	"github.com/pangminfu/go-clean-arch/reusable"
	"github.com/urfave/cli"
)

var DATA = []*reusable.Product{
	&reusable.Product{
		Id:   1,
		Name: "testname1",
		Code: "A123",
		Desc: "i am first product",
	},
	&reusable.Product{
		Id:   2,
		Name: "testname2",
		Code: "A222",
		Desc: "i am second product",
	},
}

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
				//can replace with reusable.NewMySqlProductRepository(db) after db connection is establish
				repo := reusable.NewInMemProductRepository(DATA) //repo := reusable.NewMySqlProductRepository(db)
				svc := reusable.NewService(repo)
				created, err := svc.Create(p)
				if err != nil {
					return err
				}
				log.Printf("created : %s", *created)
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "list all product",
			Action: func(c *cli.Context) error {
				//can replace with reusable.NewMySqlProductRepository(db) after db connection is establish
				repo := reusable.NewInMemProductRepository(DATA) //repo := reusable.NewMySqlProductRepository(db)
				svc := reusable.NewService(repo)
				list, err := svc.ListProduct()
				if err != nil {
					return err
				}
				log.Printf("list : %s", list)
				return nil
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
				//can replace with reusable.NewMySqlProductRepository(db) after db connection is establish
				repo := reusable.NewInMemProductRepository(DATA) //repo := reusable.NewMySqlProductRepository(db)
				svc := reusable.NewService(repo)
				result, err := svc.SearchByCode(c.String("code"))
				if err != nil {
					return err
				}
				log.Printf("result : %s", result)
				return nil
			},
		},
		{
			Name:  "update",
			Usage: "update product",
			Flags: flags,
			Action: func(c *cli.Context) error {
				p := handleArgs(c)
				//can replace with reusable.NewMySqlProductRepository(db) after db connection is establish
				repo := reusable.NewInMemProductRepository(DATA) //repo := reusable.NewMySqlProductRepository(db)
				svc := reusable.NewService(repo)
				updated, err := svc.UpdateProduct(p)
				if err != nil {
					return err
				}
				log.Printf("updated : %s", updated)
				return nil
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
				//can replace with reusable.NewMySqlProductRepository(db) after db connection is establish
				repo := reusable.NewInMemProductRepository(DATA) //repo := reusable.NewMySqlProductRepository(db)
				svc := reusable.NewService(repo)
				err := svc.DeleteProduct(c.Int("id"))
				if err != nil {
					return err
				}
				log.Printf("deleted")
				log.Printf("product remain : %s", DATA)
				return nil
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
