package product

import (
	"github.com/gin-gonic/gin"
	productusecase "github.com/pangminfu/go-clean-arch/pkg/product"
)

type Controller struct {
	uc productusecase.Usecase
}

func New(uc productusecase.Usecase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (ctrl *Controller) PostProduct(c *gin.Context) {

}

func (ctrl *Controller) GetProducts(c *gin.Context) {

}
