package repository

import (
	"errors"

	"github.com/pangminfu/go-clean-arch/pkg/product"
)

type inMemRepository struct {
	db []product.Product
}

func New(db []product.Product) product.Repository {
	return &inMemRepository{
		db: db,
	}
}

func (r *inMemRepository) Add(p product.Product) (product.Product, error) {
	return product.Product{}, errors.New("no implementation")
}

func (r *inMemRepository) Products() ([]product.Product, error) {
	return nil, errors.New("no implementation")
}
func (r *inMemRepository) Search(code string) (product.Product, error) {
	return product.Product{}, errors.New("no implementation")
}
func (r *inMemRepository) Update(p product.Product) (product.Product, error) {
	return product.Product{}, errors.New("no implementation")
}
func (r *inMemRepository) Delete(id int) error {
	return errors.New("no implementation")
}
