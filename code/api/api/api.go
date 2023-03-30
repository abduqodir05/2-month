package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {

	handler := handler.NewHandler(cfg, store, logger)

	//!! BOOK ----------------------------------
	r.POST("/book", handler.CreateBook)
	r.GET("/book/:id", handler.GetByIdBook)
	r.GET("/book", handler.GetListBook)
	r.PUT("/book/:id", handler.UpdateBook)
	r.PATCH("/book/:id", handler.UpdatePatchBook)
	r.DELETE("/book/:id", handler.DeleteBook)
	
	//!! USER ----------------------------------
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.PATCH("/user/:id", handler.UpdatePatchUser)
	r.DELETE("/user/:id", handler.DeleteUser)
	
	//!! Category ----------------------------------
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.PATCH("/category/:id", handler.UpdatePatchCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)
	
	//!! Courier ----------------------------------
	r.POST("/courier", handler.CreateCourier)
	r.GET("/courier/:id", handler.GetByIdCourier)
	r.GET("/courier", handler.GetListCourier)
	r.PUT("/courier/:id", handler.UpdateCourier)
	r.PATCH("/courier/:id", handler.UpdatePatchCourier)
	r.DELETE("/courier/:id", handler.DeleteCourier)

	//!! Customer ----------------------------------
	r.POST("/customer", handler.CreateCustomer)
	r.GET("/customer/:id", handler.GetByIdCustomer)
	r.GET("/customer", handler.GetListCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.PATCH("/customer/:id", handler.UpdatePatchCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
