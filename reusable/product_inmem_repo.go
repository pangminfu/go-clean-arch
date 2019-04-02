package reusable

import "errors"

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

	return nil, nil
}

func (repo InMemProductRepository) Delete(id int) error {
	for i, product := range repo.Products {
		if product.Id == id {
			repo.Products[i] = repo.Products[len(repo.Products)-1]
			repo.Products = repo.Products[:len(repo.Products)-1]
			return nil
		}
	}
	return errors.New("Id not found")
}
