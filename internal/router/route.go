package router

import (
	"github.com/gin-gonic/gin"
	productctrl "github.com/pangminfu/go-clean-arch/internal/product"
	"github.com/pangminfu/go-clean-arch/pkg/product"
	"github.com/pangminfu/go-clean-arch/pkg/product/repository"
	"github.com/pangminfu/go-clean-arch/pkg/product/usecase"
)

var db []product.Product

func Init() (*gin.Engine, error) {
	r := gin.Default()

	repo := repository.New(db)
	uc := usecase.New(repo)
	ctrl := productctrl.New(uc)

	r.POST("/product", ctrl.PostProduct)
	r.GET("/products", ctrl.GetProducts)

	return r, nil
}
