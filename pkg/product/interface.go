package product

type Usecase interface {
	Add(p Product) (Product, error)
	Products() ([]Product, error)
	Search(code string) (Product, error)
	Update(p Product) (Product, error)
	Delete(id int) error
}

type Repository interface {
	Add(p Product) (Product, error)
	Products() ([]Product, error)
	Search(code string) (Product, error)
	Update(p Product) (Product, error)
	Delete(id int) error
}
