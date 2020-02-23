package usecase_test

import (
	"testing"

	"github.com/pangminfu/go-clean-arch/pkg/product"
	"github.com/pangminfu/go-clean-arch/pkg/product/mocks"
	"github.com/pangminfu/go-clean-arch/pkg/product/usecase"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name        string
		product     product.Product
		expectedRes product.Product
		expectedErr error
	}{
		{
			name: "",
		},
	}

	for _, testCase := range testCases {
		repo := new(mocks.Repository)

		uc := usecase.New(repo)

		result, err := uc.Add(testCase.product)

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedRes, res)
	}
}
