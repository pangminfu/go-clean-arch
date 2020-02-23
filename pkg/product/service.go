package product

import (
	"errors"
	"regexp"
	"strings"
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
	if err := isProductInfoValid(product); err != nil {
		return nil, err
	}

	return svc.ProductRepo.Create(product)
}

func (svc ProductService) ListProduct() ([]*Product, error) {

	return svc.ProductRepo.List()
}

func (svc ProductService) SearchByCode(code string) (*Product, error) {
	code = strings.TrimSpace(code)
	if err := isProductCodeValid(code); err != nil {
		return nil, err
	}

	return svc.ProductRepo.GetByCode(code)
}

func (svc ProductService) UpdateProduct(product *Product) (*Product, error) {
	if err := isProductInfoValid(product); err != nil {
		return nil, err
	}

	return svc.ProductRepo.Update(product)
}

func (svc ProductService) DeleteProduct(id int) error {
	return svc.ProductRepo.Delete(id)
}

func isProductInfoValid(product *Product) error {
	if product == new(Product) || product == nil {
		return errors.New("Product is empty")
	}

	product.Name = strings.TrimSpace(product.Name)
	if err := isProductNameValid(product.Name); err != nil {
		return err
	}

	product.Desc = strings.TrimSpace(product.Desc)
	if err := isProductDescValid(product.Desc); err != nil {
		return err
	}

	product.Code = strings.TrimSpace(product.Code)
	if err := isProductCodeValid(product.Code); err != nil {
		return err
	}

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
	re := regexp.MustCompile(`^[A-Z]([0-9]+)$`)
	if re.MatchString(code) {
		return nil
	}

	return errors.New("Invalid Product code format")
}
