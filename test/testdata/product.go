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

	RepoErr = errors.New("repo return error")
)
