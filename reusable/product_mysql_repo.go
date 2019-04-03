package reusable

import (
	"database/sql"
)

type MySqlProductRepository struct {
	database *sql.DB
}

func NewMySqlProductRepository(db *sql.DB) ProductRepository {
	return &MySqlProductRepository{
		database: db,
	}
}

func (repo MySqlProductRepository) Create(p *Product) (*Product, error) {
	return nil, nil
}

func (repo MySqlProductRepository) List() ([]*Product, error) {
	return nil, nil
}

func (repo MySqlProductRepository) GetByCode(code string) (*Product, error) {
	return nil, nil
}

func (repo MySqlProductRepository) Update(p *Product) (*Product, error) {
	return nil, nil
}

func (repo MySqlProductRepository) Delete(id int) error {
	return nil
}
