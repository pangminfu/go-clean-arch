package usecase

import (
	"errors"

	"github.com/pangminfu/go-clean-arch/pkg/product"
)

type useCaseImpl struct {
	repo product.Repository
}

func New(repo product.Repository) product.Usecase {
	return &useCaseImpl{
		repo: repo,
	}
}

func (uc *useCaseImpl) Add(p product.Product) (product.Product, error) {
	var result product.Product

	result, err := uc.repo.Add(p)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (uc *useCaseImpl) Products() ([]product.Product, error) {
	return uc.repo.Products()
}

func (uc *useCaseImpl) Search(code string) (product.Product, error) {
	return product.Product{}, errors.New("no implementation")
}

func (uc *useCaseImpl) Update(p product.Product) (product.Product, error) {
	return product.Product{}, errors.New("no implementation")
}

func (uc *useCaseImpl) Delete(id int) error {
	return errors.New("no implementation")
}
