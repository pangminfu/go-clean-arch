package reusable

type Service interface {
	Create(p *Product) (*Product, error)
	ListProduct() ([]*Product, error)
	SearchByCode(code string) (*Product, error)
	UpdateProduct(p *Product) (*Product, error)
	DeleteProduct(id string) error
}

type ProductRepository interface {
	Create(p *Product) (*Product, error)
	List() ([]*Product, error)
	GetByCode(code string) (*Product, error)
	Update(p *Product) (*Product, error)
	Delete(id string) error
}