package reusable

type InMemProductRepository struct {
	Products []*Product
}

func NewInMemProductRepository(r []*Product) ProductRepository {
	return &InMemProductRepository{
		Products: r,
	}
}

func (repo InMemProductRepository) Create(p *Product) (*Product, error) {
	return nil, nil
}
func (repo InMemProductRepository) List() ([]*Product, error) {
	return nil, nil
}
func (repo InMemProductRepository) GetByCode(code string) (*Product, error) {
	return nil, nil
}
func (repo InMemProductRepository) Update(p *Product) (*Product, error) {
	return nil, nil
}

func (repo InMemProductRepository) Delete(id string) error {
	return nil
}
