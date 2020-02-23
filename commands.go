package main

import (
	"log"

	"github.com/pangminfu/go-clean-arch/pkg/product"
	"github.com/urfave/cli"
)

var DATA = []*product.Product{
	&product.Product{
		Id:   1,
		Name: "testname1",
		Code: "A123",
		Desc: "i am first product",
	},
	&product.Product{
		Id:   2,
		Name: "testname2",
		Code: "A222",
		Desc: "i am second product",
	},
}

func init() {
	//initialize db connection
	// db, err := product.ConnectMysql(username, password, host, dbname)
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
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
				//can replace with product.NewMySqlProductRepository(db) after db connection is establish
				repo := product.NewInMemProductRepository(DATA) //repo := product.NewMySqlProductRepository(db)
				svc := product.NewService(repo)
				created, err := svc.Create(p)
				if err != nil {
					return err
				}
				log.Printf("created : %+v", *created)
				return nil
			},
		},
		{
			Name:  "list",
			Usage: "list all product",
			Action: func(c *cli.Context) error {
				//can replace with product.NewMySqlProductRepository(db) after db connection is establish
				repo := product.NewInMemProductRepository(DATA) //repo := product.NewMySqlProductRepository(db)
				svc := product.NewService(repo)
				list, err := svc.ListProduct()
				if err != nil {
					return err
				}
				for i, p := range list {
					log.Printf("%d : %+v", i+1, p)
				}
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
				//can replace with product.NewMySqlProductRepository(db) after db connection is establish
				repo := product.NewInMemProductRepository(DATA) //repo := product.NewMySqlProductRepository(db)
				svc := product.NewService(repo)
				result, err := svc.SearchByCode(c.String("code"))
				if err != nil {
					return err
				}
				log.Printf("result : %+v", result)
				return nil
			},
		},
		{
			Name:  "update",
			Usage: "update product",
			Flags: flags,
			Action: func(c *cli.Context) error {
				p := handleArgs(c)
				//can replace with product.NewMySqlProductRepository(db) after db connection is establish
				repo := product.NewInMemProductRepository(DATA) //repo := product.NewMySqlProductRepository(db)
				svc := product.NewService(repo)
				updated, err := svc.UpdateProduct(p)
				if err != nil {
					return err
				}
				log.Printf("updated : %+v", updated)
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
				//can replace with product.NewMySqlProductRepository(db) after db connection is establish
				repo := product.NewInMemProductRepository(DATA) //repo := product.NewMySqlProductRepository(db)
				svc := product.NewService(repo)
				err := svc.DeleteProduct(c.Int("id"))
				if err != nil {
					return err
				}
				log.Printf("deleted")
				log.Printf("product remain :")
				for i, p := range DATA {
					log.Printf("%d : %+v", i+1, p)
				}
				return nil
			},
		},
	}
}

func handleArgs(c *cli.Context) *product.Product {
	product := &product.Product{
		Id:   c.Int("id"),
		Code: c.String("code"),
		Name: c.String("name"),
		Desc: c.String("desc"),
	}

	return product
}
