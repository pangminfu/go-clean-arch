package testdata

import (
	"errors"

	"github.com/pangminfu/go-clean-arch/pkg/product"
)

var (
	ProductZeroValue = product.Product{}

	ProductValid = product.Product{
		Code: "SG001",
		Name: "Macbook",
		Desc: "This is a macbook",
	}

	ProductsValid = []product.Product{
		product.Product{
			Code: "SG001",
			Name: "Macbook",
			Desc: "This is a macbook",
		},
		product.Product{
			Code: "SG002",
			Name: "Asus",
			Desc: "This is a Asus",
		},
	}

	RepoErr = errors.New("repo return error")
)
