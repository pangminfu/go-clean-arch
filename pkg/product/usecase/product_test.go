package usecase_test

import (
	"testing"

	"github.com/pangminfu/go-clean-arch/pkg/product"
	"github.com/pangminfu/go-clean-arch/pkg/product/mocks"
	"github.com/pangminfu/go-clean-arch/pkg/product/usecase"
	"github.com/pangminfu/go-clean-arch/test/testdata"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name        string
		product     product.Product
		repoRes     product.Product
		repoErr     error
		expectedRes product.Product
		expectedErr error
	}{
		{
			name:        "Given valid product product detail",
			product:     testdata.ProductValid,
			repoRes:     testdata.ProductValid,
			repoErr:     nil,
			expectedRes: testdata.ProductValid,
			expectedErr: nil,
		},
		{
			name:        "Given repo return error",
			product:     testdata.ProductValid,
			repoRes:     testdata.ProductZeroValue,
			repoErr:     testdata.RepoErr,
			expectedRes: testdata.ProductZeroValue,
			expectedErr: testdata.RepoErr,
		},
	}

	for _, testCase := range testCases {
		repo := new(mocks.Repository)

		repo.On("Add", testCase.product).
			Return(testCase.repoRes, testCase.repoErr)

		uc := usecase.New(repo)

		result, err := uc.Add(testCase.product)

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedRes, result)
	}
}

func TestProducts(t *testing.T) {
	testCases := []struct {
		name        string
		repoRes     []product.Product
		repoErr     error
		expectedRes []product.Product
		expectedErr error
	}{
		{
			name:        "Given products retrieve successfully",
			repoRes:     testdata.ProductsValid,
			repoErr:     nil,
			expectedRes: testdata.ProductsValid,
			expectedErr: nil,
		},
		{
			name:        "Given repo return error",
			repoRes:     nil,
			repoErr:     testdata.RepoErr,
			expectedRes: nil,
			expectedErr: testdata.RepoErr,
		},
	}

	for _, testCase := range testCases {
		repo := new(mocks.Repository)

		repo.On("Products").
			Return(testCase.repoRes, testCase.repoErr)

		uc := usecase.New(repo)

		result, err := uc.Products()

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedRes, result)
	}
}

func TestSearch(t *testing.T) {
	testCases := []struct {
		name        string
		code        string
		repoRes     product.Product
		repoErr     error
		expectedRes product.Product
		expectedErr error
	}{
		{
			name:        "Given code retrieve product successfully",
			code:        "SG001",
			repoRes:     testdata.ProductValid,
			repoErr:     nil,
			expectedRes: testdata.ProductValid,
			expectedErr: nil,
		},
		{
			name:        "Given repo return error",
			repoRes:     testdata.ProductZeroValue,
			repoErr:     testdata.RepoErr,
			expectedRes: testdata.ProductZeroValue,
			expectedErr: testdata.RepoErr,
		},
	}

	for _, testCase := range testCases {
		repo := new(mocks.Repository)

		repo.On("Search", testCase.code).
			Return(testCase.repoRes, testCase.repoErr)

		uc := usecase.New(repo)

		result, err := uc.Search(testCase.code)

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedRes, result)
	}
}

func TestUpdate(t *testing.T) {
	testCases := []struct {
		name        string
		product     product.Product
		repoRes     product.Product
		repoErr     error
		expectedRes product.Product
		expectedErr error
	}{
		{
			name:        "Given valid product product detail",
			product:     testdata.ProductValid,
			repoRes:     testdata.ProductValid,
			repoErr:     nil,
			expectedRes: testdata.ProductValid,
			expectedErr: nil,
		},
		{
			name:        "Given repo return error",
			product:     testdata.ProductValid,
			repoRes:     testdata.ProductZeroValue,
			repoErr:     testdata.RepoErr,
			expectedRes: testdata.ProductZeroValue,
			expectedErr: testdata.RepoErr,
		},
	}

	for _, testCase := range testCases {
		repo := new(mocks.Repository)

		repo.On("Update", testCase.product).
			Return(testCase.repoRes, testCase.repoErr)

		uc := usecase.New(repo)

		result, err := uc.Update(testCase.product)

		assert.Equal(t, testCase.expectedErr, err)
		assert.Equal(t, testCase.expectedRes, result)
	}
}
