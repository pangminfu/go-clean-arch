package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
	productservice "github.com/pangminfu/go-clean-arch/pkg/product"
)

type Controller struct {
	uc productservice.Usecase
}

func New(uc productservice.Usecase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (ctrl *Controller) PostProduct(c *gin.Context) {
	product := new(productservice.Product)
	if err := c.ShouldBindJSON(&product); err != nil {
		errResponse(c, err, http.StatusBadRequest)

		return
	}

	result, err := ctrl.uc.Add(*product)
	if err != nil {
		errResponse(c, err, http.StatusInternalServerError)

		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": result,
		},
	)
}

func (ctrl *Controller) GetProducts(c *gin.Context) {
	result, err := ctrl.uc.Products()
	if err != nil {
		errResponse(c, err, http.StatusInternalServerError)

		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": result,
		},
	)
}

func errResponse(c *gin.Context, err error, status int) {
	c.JSON(
		status,
		gin.H{
			"error": err.Error(),
		},
	)
}
