package usecase

import (
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
	var product product.Product
	product, err := uc.repo.Search(code)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (uc *useCaseImpl) Update(p product.Product) (product.Product, error) {
	var product product.Product
	product, err := uc.repo.Update(p)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (uc *useCaseImpl) Delete(id int) error {
	return uc.repo.Delete(id)
}
