package reusable

import "testing"

func TestCreate_ErrNil(t *testing.T) {
	repo := NewInMemProductRepository(nil)
	svc := NewService(repo)

	samples := []*Product{
		&Product{
			Code: "A8965",
			Name: "test",
			Desc: "test desc",
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
