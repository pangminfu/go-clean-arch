package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pangminfu/go-clean-arch/pkg/product"
)

type MySQLRepository struct {
	db *gorm.DB
}

const (
	Dialect      = "mysql"
	ProductTable = "product_tbl"
)

func ConnectMySQL(user, password, host, port, database string) (*gorm.DB, error) {
	return gorm.Open(Dialect, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?timeout=15s&parseTime=true", user, password, host, port, database))
}

func NewMysql(db *gorm.DB) product.Repository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) Add(p product.Product) (product.Product, error) {
	err := r.db.Table(ProductTable).Create(&p).Error
	if err != nil {
		return p, err
	}

	return p, nil
}

func (r *MySQLRepository) Products() ([]product.Product, error) {
	products := make([]product.Product, 0)
	err := r.db.Table(ProductTable).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
func (r *MySQLRepository) Search(code string) (product.Product, error) {
	var product product.Product
	err := r.db.Where("code = ?", code).First(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}
func (r *MySQLRepository) Update(p product.Product) (product.Product, error) {
	err := r.db.Save(&p).Error
	if err != nil {
		return p, err
	}

	return p, nil
}
func (r *MySQLRepository) Delete(id int) error {
	return r.db.Where("id = ?", id).Delete(&product.Product{}).Error
}
