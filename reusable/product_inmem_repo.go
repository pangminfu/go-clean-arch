package reusable

import "errors"

var InMemProducts = []*Product{
	&Product{
		Id:   1,
		Code: "A8965",
		Name: "test",
		Desc: "test desc positive",
	},
	&Product{
		Id:   2,
		Code: "A0077",
		Name: "test name",
		Desc: "test desc not too long desc",
	},
	&Product{
		Id:   3,
		Code: "A112",
		Name: "test",
		Desc: "test desc i dont see error",
	},
}

type InMemProductRepository struct {
	Products []*Product
}

func NewInMemProductRepository(r []*Product) ProductRepository {
	return &InMemProductRepository{
		Products: r,
	}
}

func (repo InMemProductRepository) Create(p *Product) (*Product, error) {
	repo.Products = append(repo.Products, p)
	return p, nil
}
func (repo InMemProductRepository) List() ([]*Product, error) {
	return repo.Products, nil
}
func (repo InMemProductRepository) GetByCode(code string) (*Product, error) {
	for _, p := range repo.Products {
		if p.Code == code {
			return p, nil
		}
	}

	return nil, nil
}
func (repo InMemProductRepository) Update(p *Product) (*Product, error) {
	for i, product := range repo.Products {
		if product.Id == p.Id {
			repo.Products[i] = p
			return repo.Products[i], nil
		}
	}

	return nil, errors.New("Id not found")
}

func (repo InMemProductRepository) Delete(id int) error {
	for i, product := range repo.Products {
		if product.Id == id {
			repo.Products = remove(repo.Products, i)
			return nil
		}
	}
	return errors.New("Id not found")
}

func remove(s []*Product, i int) []*Product {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil
	s = s[:len(s)-1]

	return s
}
