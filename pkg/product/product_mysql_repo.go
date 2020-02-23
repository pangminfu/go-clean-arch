package reusable

import (
	"database/sql"
	"fmt"
)

type MySqlProductRepository struct {
	database *sql.DB
}

func NewMySqlProductRepository(db *sql.DB) ProductRepository {
	return &MySqlProductRepository{
		database: db,
	}
}

func ConnectMysql(u, p, h, dbname string) (*sql.DB, error) {
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?timeout=10s",
		u, p, h, dbname,
	)

	// Use db to perform SQL operations on database
	db, err := sql.Open("mysql", dnsStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (repo MySqlProductRepository) Create(p *Product) (*Product, error) {
	preStatement, err := repo.database.Prepare("INSERT INTO product (code, name, desc) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer preStatement.Close()

	result, err := preStatement.Exec(p.Code, p.Name, p.Desc)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	p.Id = int(id)

	return p, nil
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
	preStatement, err := repo.database.Prepare("SELECT id, code, name, desc, mtime FROM product WHERE code=?")
	if err != nil {
		return nil, err
	}
	defer preStatement.Close()

	resultList, err := preStatement.Query(code)
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

	return products[0], nil
}

func (repo MySqlProductRepository) Update(p *Product) (*Product, error) {
	preStatement, err := repo.database.Prepare("UPDATE product SET code = ?, name = ?, desc = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer preStatement.Close()

	result, err := preStatement.Exec(p.Code, p.Name, p.Desc, p.Id)
	if err != nil {
		return nil, err
	}
	id, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	p.Id = int(id)

	return p, nil
}

func (repo MySqlProductRepository) Delete(id int) error {
	preStatement, err := repo.database.Prepare("DELETE FROM product WHERE id=?")
	if err != nil {
		return err
	}
	defer preStatement.Close()

	_, err = preStatement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
