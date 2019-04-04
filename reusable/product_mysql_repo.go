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
	preStatement, err := repo.database.Prepare("SELECT id, code, name, desc, mtime FROM product")
	if err != nil {
		return nil, err
	}
	defer preStatement.Close()

	resultList, err := preStatement.Query()
	if err != nil {
		return nil, err
	}
	defer resultList.Close()

	products := []*Product{}
	for resultList.Next() {
		var p Product
		err := resultList.Scan(&p.Id, &p.Code, &p.Name, &p.Desc)
		if err != nil {
			return nil, err
		}

		products = append(products, &p)
	}

	return products, nil
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
