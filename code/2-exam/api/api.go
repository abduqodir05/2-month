package api

import (
	_ "app/api/docs"

	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)
	// category api
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	// brand api
	r.POST("/brand", handler.CreateBrand)
	r.GET("/brand/:id", handler.GetByIdBrand)
	r.GET("/brand", handler.GetListBrand)
	r.PUT("/brand/:id", handler.UpdateBrand)
	r.DELETE("/brand/:id", handler.DeleteBrand)

	// product api
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	// stock api  -- not ready for using
	r.POST("/stock", handler.CreateStock)
	r.GET("/stock/:id", handler.GetByIdStock)
	r.GET("/stock", handler.GetListStock)
	r.PUT("/stock/:id", handler.UpdateStock)
	r.DELETE("/stock/:id", handler.DeleteStock)
	r.PUT("/sendstock/:id", handler.SendProduct)

	// store api
	r.POST("/store", handler.CreateStore)
	r.GET("/store/:id", handler.GetByIdStore)
	r.GET("/store", handler.GetListStore)
	r.PUT("/store/:id", handler.UpdateStore)
	r.PATCH("/store/:id", handler.UpdatePatchStore)
	r.DELETE("/store/:id", handler.DeleteStore)

	// customer api
	r.POST("/customer", handler.CreateCustomer)
	r.GET("/customer/:id", handler.GetByIdCustomer)
	r.GET("/customer", handler.GetListCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.PATCH("/customer/:id", handler.UpdatePatchCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)

	// staff api
	r.POST("/staff", handler.CreateStaff)
	r.GET("/staff/:id", handler.GetByIdStaff)
	r.GET("/staff", handler.GetListStaff)
	r.PUT("/staff/:id", handler.UpdateStaff)
	r.PATCH("/staff/:id", handler.UpdatePatchStaff)
	r.DELETE("/staff/:id", handler.DeleteStaff)

	// order api
	// r.Use(cusomMIddleware())

	v1 := r.Group("/v1")

	v1.Use(handler.AuthMiddleware())

	r.POST("/order", handler.AuthMiddleware(), handler.CreateOrder)
	r.GET("/order/:id", handler.AuthMiddleware(), handler.GetByIdOrder)
	r.GET("/order",handler.AuthMiddleware(), handler.GetListOrder)
	r.PUT("/order/:id",handler.AuthMiddleware(), handler.UpdateOrder)
	r.PATCH("/order/:id",handler.AuthMiddleware(), handler.UpdatePatchOrder)
	r.DELETE("/order/:id",handler.AuthMiddleware(), handler.DeleteOrder)
	r.POST("/order_item/",handler.AuthMiddleware(), handler.CreateOrderItem)
	r.DELETE("/order_item/:id",handler.AuthMiddleware(), handler.DeleteOrderItem)
	r.GET("/ordersold",handler.AuthMiddleware(), handler.InfoOfSoldProductsByStaffer)

	// total price
	r.GET("/totalorder/total_price/:id", handler.TotalPriceOrder)
	
	// promo_code api
	r.POST("/promo_code", handler.CreatePromoCode)
	r.GET("/promo_code/:id", handler.GetByIdPromoCode)
	r.GET("/promo_code", handler.GetListPromoCode)
	r.DELETE("/promo_code/:id", handler.DeletePromoCode)
	
	
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
