package router

import (
	"github.com/gin-gonic/gin"
	productctrl "github.com/pangminfu/go-clean-arch/internal/product"
	"github.com/pangminfu/go-clean-arch/pkg/product/repository"
	"github.com/pangminfu/go-clean-arch/pkg/product/usecase"
	"github.com/spf13/viper"
)

func Init(cfg *viper.Viper) (*gin.Engine, error) {
	r := gin.Default()

	db, err := repository.ConnectMySQL(
		cfg.GetString("Database.Username"),
		cfg.GetString("Database.Password"),
		cfg.GetString("Database.Host"),
		cfg.GetString("Database.Port"),
		cfg.GetString("Database.DBName"),
	)
	if err != nil {
		return nil, err
	}

	repo := repository.NewMysql(db)
	uc := usecase.New(repo)
	ctrl := productctrl.New(uc)

	r.POST("/product", ctrl.PostProduct)
	r.GET("/products", ctrl.GetProducts)
	r.GET("/product/:code", ctrl.SearchProduct)
	r.PUT("/product/:id", ctrl.PutProduct)
	r.DELETE("/product/:id", ctrl.DeleteProduct)

	return r, nil
}
