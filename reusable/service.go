package reusable

import (
	"errors"
	"regexp"
)

type ProductService struct {
	ProductRepo ProductRepository
}

func NewService(productRepo ProductRepository) Service {
	return &ProductService{
		ProductRepo: productRepo,
	}
}

func (svc ProductService) Create(product *Product) (*Product, error) {
	if product == new(Product) || product == nil {
		return nil, errors.New("Product is empty")
	}

	if err := isProductNameValid(product.Name); err != nil {
		return nil, err
	}

	if err := isProductDescValid(product.Desc); err != nil {
		return nil, err
	}

	if err := isProductCodeValid(product.Code); err != nil {
		return nil, err
	}

	return svc.ProductRepo.Create(product)
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

func isProductNameValid(name string) error {
	if len(name) <= 3 {
		return errors.New("Product Name must be more than 3 charater")
	}

	return nil
}

func isProductDescValid(desc string) error {
	if len(desc) <= 10 {
		return errors.New("Product Description must be more than 10 character")
	}

	return nil
}

func isProductCodeValid(code string) error {
	re := regexp.MustCompile(`^[A-Z]([0-9]+)`)
	if re.MatchString(code) {
		return nil
	}

	return errors.New("Invalid Product code format")
}
