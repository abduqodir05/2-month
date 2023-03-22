package controller

import (
	"net/http"
	"strings"
)

func (c *Controller) Product(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.store.Product().CreateProduct(w, r)
	}
	if r.Method == "GET" {
		path := strings.Split(r.URL.Path, "/")

		if len(path) > 2 {
			c.store.Product().GetProductById(w, r)
		} else {
			c.store.Product().GetListProduct(w, r)
		}
	}
	if r.Method == "PUT" {
		c.store.Product().UpdateProduct(w, r)
	}
	if r.Method == "DELETE" {
		c.store.Product().DeleteProduct(w, r)
	}
}

// func (c *Controller) CreateProduct(req *models.CreateProduct) (string, error) {
// 	id, err := c.store.Product().Create(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	return id, nil
// }

// func (c *Controller) DeleteProduct(req *models.ProductPrimaryKey) error {
// 	err := c.store.Product().Delete(req)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *Controller) UpdateProduct(req *models.UpdateProduct, productId string) error {
// 	err := c.store.Product().Update(req, productId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (c *Controller) GetByIdProduct(req *models.ProductPrimaryKey) (models.Product, error) {
// 	product, err := c.store.Product().GetByID(req)
// 	if err != nil {
// 		return models.Product{}, err
// 	}

// 	category, err := c.store.Category().GetByID(&models.CategoryPrimaryKey{Id: product.CategoryID})
// 	if err != nil {
// 		return models.Product{}, err
// 	}

// 	return models.Product{
// 		Id:       product.Id,
// 		Name:     product.Name,
// 		Price:    product.Price,
// 		Category: category,
// 	}, nil
// }

// func (c *Controller) GetAllProduct(req *models.GetListProductRequest) (models.GetListProduct, error) {
// 	products, err := c.store.Product().GetAll(req)
// 	if err != nil {
// 		return models.GetListProduct{}, err
// 	}
// 	return products, nil
// }

// // func (c *Controller)TopSellingProducts(filter string) (models.GetProductCount, error) {
// // 	product, err := c.store.Product().GetAll()
// // }

// func (c *Controller) SelledProductThroughId(req *models.ProductPrimaryKey) (models.CountProduct, error) {
// 	shopCarts, err := c.store.ShopCart().GetByProductId(&models.ProductPrimaryKey{Id: req.Id})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	soldCount := 0

// 	for _, v := range shopCarts {
// 		soldCount += v.Count
// 	}
// 	productName, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: req.Id})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
	
// 	countProduct := models.CountProduct{Count: soldCount, Name: productName.Name}

// 	return countProduct, nil
// }
