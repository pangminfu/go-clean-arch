package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pangminfu/go-clean-arch/pkg/product"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMysql(db *gorm.DB) product.Repository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) Add(p product.Product) (product.Product, error) {
	return product.Product{}, errors.New("no implementation")
}

func (r *MySQLRepository) Products() ([]product.Product, error) {
	return nil, errors.New("no implementation")
}
func (r *MySQLRepository) Search(code string) (product.Product, error) {
	return product.Product{}, errors.New("no implementation")
}
func (r *MySQLRepository) Update(p product.Product) (product.Product, error) {
	return product.Product{}, errors.New("no implementation")
}
func (r *MySQLRepository) Delete(id int) error {
	return errors.New("no implementation")
}
