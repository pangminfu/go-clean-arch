package product

import "testing"

func TestCreate_ErrNil(t *testing.T) {
	repo := NewInMemProductRepository(nil)
	svc := NewService(repo)

	samples := []*Product{
		&Product{
			Code: "A8965",
			Name: "test",
			Desc: "test desc positive",
		},
		&Product{
			Code: "A0077",
			Name: "test name",
			Desc: "test desc not too long desc",
		},
		&Product{
			Code: "A112",
			Name: "test",
			Desc: "test desc i dont see error",
		},
	}

	for i, sample := range samples {
		_, err := svc.Create(sample)
		if err != nil {
			t.Errorf("Sample %d should return error nil instead of %s", i+1, err.Error())
		}
	}
}

func TestCreate_ErrNotNil(t *testing.T) {
	repo := NewInMemProductRepository(nil)
	svc := NewService(repo)

	samples := []*Product{
		nil,
		&Product{},
		&Product{
			Code: "8965",
			Name: "test",
			Desc: "test desc",
		},
		&Product{
			Code: "A0077",
			Name: "",
			Desc: "test desc not too long desc",
		},
		&Product{
			Code: "A112",
			Name: "test",
			Desc: "",
		},
		&Product{
			Code: "112A",
			Name: "test",
			Desc: "",
		},
		&Product{
			Code: "112",
			Name: "test",
			Desc: "test desc",
		},
		&Product{
			Code: "A112",
			Name: "te",
			Desc: "test desc",
		},
		&Product{
			Code: "A112",
			Name: "test",
			Desc: "te",
		},
		&Product{
			Code: "a112",
			Name: "test",
			Desc: "test desc i dont see error",
		},
		&Product{
			Code: "",
			Name: "",
			Desc: "",
		},
	}

	for i, sample := range samples {
		_, err := svc.Create(sample)
		if err == nil {
			t.Errorf("Sample %d should return error instead of error nil", i+1)
		}
	}
}

func TestUpdateProduct_ErrNil(t *testing.T) {
	repo := NewInMemProductRepository(InMemProducts)
	svc := NewService(repo)

	samples := []*Product{
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

	for i, sample := range samples {
		_, err := svc.UpdateProduct(sample)
		if err != nil {
			t.Errorf("Sample %d should return error nil instead of %s", i+1, err.Error())
		}
	}
}

func TestUpdateProduct_ErrNotNil(t *testing.T) {
	repo := NewInMemProductRepository(InMemProducts)
	svc := NewService(repo)

	samples := []*Product{
		nil,
		&Product{},
		&Product{
			Code: "",
			Name: "",
			Desc: "",
		},
		&Product{
			Code: "8965",
			Name: "test",
			Desc: "test desc",
		},
		&Product{
			Code: "A0077",
			Name: "",
			Desc: "test desc not too long desc",
		},
		&Product{
			Code: "A112",
			Name: "test",
			Desc: "",
		},
		&Product{
			Code: "112A",
			Name: "test",
			Desc: "",
		},
		&Product{
			Code: "112",
			Name: "test",
			Desc: "test desc",
		},
		&Product{
			Code: "A112",
			Name: "te",
			Desc: "test desc",
		},
		&Product{
			Code: "A112",
			Name: "test",
			Desc: "te",
		},
		&Product{
			Code: "a112",
			Name: "test",
			Desc: "test desc i dont see error",
		},
		&Product{
			Code: "A8965",
			Name: "test",
			Desc: "test desc positive",
		},
	}

	for i, sample := range samples {
		_, err := svc.UpdateProduct(sample)
		if err == nil {
			t.Errorf("Sample %d should return error instead of error nil", i+1)
		}
	}
}

func TestIsProductInfoValid_ErrNil(t *testing.T) {
	samples := []*Product{
		&Product{
			Code: "A8965",
			Name: "test",
			Desc: "test desc positive",
		},
		&Product{
			Code: "A0077",
			Name: "test name",
			Desc: "test desc not too long desc",
		},
		&Product{
			Code: "A112",
			Name: "test",
			Desc: "test desc i dont see error",
		},
	}
	for i, sample := range samples {
		err := isProductInfoValid(sample)
		if err != nil {
			t.Errorf("Sample %d should not return error %s", i+1, err.Error())
		}
	}
}

func TestIsProductInfoValid_ErrNotNil(t *testing.T) {
	samples := []*Product{
		nil,
		&Product{},
	}
	for i, sample := range samples {
		err := isProductInfoValid(sample)
		if err == nil {
			t.Errorf("Sample %d should return error instead of nil", i+1)
		}
	}
}

func TestIsProductNameValid_ErrNil(t *testing.T) {
	samples := []string{
		"test",
		"test row",
		"how are you?",
		"F&N nestle",
	}
	for i, sample := range samples {
		err := isProductNameValid(sample)
		if err != nil {
			t.Errorf("Sample %d should not return error %s", i+1, err.Error())
		}
	}
}

func TestIsProductNameValid_ErrNotNil(t *testing.T) {
	samples := []string{
		"",
		"tes",
		"h",
	}
	for i, sample := range samples {
		err := isProductNameValid(sample)
		if err == nil {
			t.Errorf("Sample %d should return error instead of nil", i+1)
		}
	}
}

func TestIsProductDescValid_ErrNil(t *testing.T) {
	samples := []string{
		"10 character",
		"12345678910",
		"how are you?",
		"this must more than 10 character",
	}
	for i, sample := range samples {
		err := isProductDescValid(sample)
		if err != nil {
			t.Errorf("Sample %d should not return error %s", i+1, err.Error())
		}
	}
}

func TestIsProductDescValid_ErrNotNil(t *testing.T) {
	samples := []string{
		"1234567890",
		"te",
		"h",
		"",
	}
	for i, sample := range samples {
		err := isProductDescValid(sample)
		if err == nil {
			t.Errorf("Sample %d should return error instead of nil", i+1)
		}
	}
}

func TestIsProductCodeValid_ErrNil(t *testing.T) {
	samples := []string{
		"A1",
		"B123",
		"C123",
	}
	for i, sample := range samples {
		err := isProductCodeValid(sample)
		if err != nil {
			t.Errorf("Sample %d should not return error %s", i+1, err.Error())
		}
	}
}

func TestIsProductCodeValid_ErrNotNil(t *testing.T) {
	samples := []string{
		"",
		"123A",
		"a234",
		"Ax123",
		"A123L",
	}
	for i, sample := range samples {
		err := isProductCodeValid(sample)
		if err == nil {
			t.Errorf("Sample %d should return error instead of nil", i+1)
		}
	}
}

func TestSearchByCode_ErrNil(t *testing.T) {
	repo := NewInMemProductRepository(InMemProducts)
	svc := NewService(repo)

	samples := []string{
		"A1",
		"B123",
		"C123",
	}

	for i, sample := range samples {
		_, err := svc.SearchByCode(sample)
		if err != nil {
			t.Errorf("Sample %d should return error nil instead of %s", i+1, err.Error())
		}
	}
}

func TestSearchByCode_ErrNotNil(t *testing.T) {
	repo := NewInMemProductRepository(InMemProducts)
	svc := NewService(repo)

	samples := []string{
		"",
		"123",
		"c123",
	}

	for i, sample := range samples {
		_, err := svc.SearchByCode(sample)
		if err == nil {
			t.Errorf("Sample %d should return error instead of nil", i+1)
		}
	}
}

func TestDeleteProduct_ErrNil(t *testing.T) {
	repo := NewInMemProductRepository(InMemProducts)
	svc := NewService(repo)

	samples := []int{
		1,
		2,
		3,
	}

	for i, sample := range samples {
		err := svc.DeleteProduct(sample)
		if err != nil {
			t.Errorf("Sample %d should return error nil instead of %s", i+1, err.Error())
		}
	}
}

func TestDeleteProduct_ErrNotNil(t *testing.T) {
	repo := NewInMemProductRepository(InMemProducts)
	svc := NewService(repo)

	samples := []int{
		0,
		5,
		100,
	}

	for i, sample := range samples {
		err := svc.DeleteProduct(sample)
		if err == nil {
			t.Errorf("Sample %d should return error instead of nil", i+1)
		}
	}
}
