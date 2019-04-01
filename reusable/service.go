package reusable

type ProductService struct {
	ProductRepo ProductRepository
}

func New(productRepo ProductRepository) Service {
	return &ProductService{
		ProductRepo: productRepo,
	}
}

func (svc ProductService) Create(product *Product) (*Product, error) {
	return nil, nil
}

func (svc ProductService) ListProduct() ([]*Product, error) {
	return nil, nil
}

func (svc ProductService) SearchByCode(code string) (*Product, error) {
	return nil, nil
}

func (svc ProductService) UpdateProduct(p *Product) (*Product, error) {
	return nil, nil
}

func (svc ProductService) DeleteProduct(id string) error {
	return nil
}
