package controller

import (
	"net/http"
	"strings"
)

func (c *Controller) ShopCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c.AddShopCartController(w, r)
	}
	if r.Method == "GET" {
		path := strings.Split(r.URL.Path, "/")

		if len(path) > 2 {
			c.GetUserShopCarts(w, r)
		} else {
			c.store.User().GetList(w, r)
		}
	}
	if r.Method == "PUT" {
		c.store.User().UpdateUser(w, r)
	}
	if r.Method == "DELETE" {
		c.RemoveShopCartController(w, r)
	}
}

//!!Exam tasks
// func (c *Controller) AddShopCartController(req *models.Add) (string, error) {
// 	_, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: req.UserId})
// 	if err != nil {
// 		return "", err
// 	}

// 	_, err = c.store.Product().GetById(&models.ProductPrimaryKey{Id: req.ProductId})
// 	if err != nil {
// 		return "", err
// 	}

// 	id, err := c.store.ShopCart().AddShopCart(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	return id, nil
// }

// func (c *Controller) RemoveShopCartController(req *models.Remove) error {
// 	err := c.store.ShopCart().RemoveShopCart(req)
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

// func (c *Controller) CalculateTotal(req *models.UserPrimaryKey, status string, discount float64) (float64, error) {
// 	_, err := c.store.User().GetByID(req)
// 	if err != nil {
// 		return 0, err
// 	}

// 	users, err := c.store.ShopCart().GetUserShopCarts(req)
// 	if err != nil {
// 		return 0, err
// 	}

// 	var total float64
// 	for _, v := range users {
// 		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: v.ProductId})
// 		if err != nil {
// 			return 0, err
// 		}
// 		if status == "fixed" {
// 			total += float64(v.Count) * (product.Price - discount)
// 		} else if status == "percent" {
// 			if discount < 0 || discount > 100 {
// 				return 0, errors.New("Invalid discount range")
// 			}
// 			total += float64(v.Count) * (product.Price - (product.Price*discount)/100)
// 		} else {
// 			return 0, errors.New("Invalid status name")
// 		}
// 	}

// 	if total < 0 {
// 		return 0, nil
// 	}
// 	return total, nil
// }

// func (c *Controller) GetFilteredShopCarts(req *models.TimeFilter) (res *models.GetShopCartsList, err error) {

// 	data, err := c.store.ShopCart().Read()
// 	if err != nil {
// 		return nil, err
// 	}

// 	filtered_ShopCarts := []models.ShopCart{}
// 	for _, v := range data {
// 		if req.DateFrom >= v.Time && req.DateTo <= v.Time {
// 			filtered_ShopCarts = append(filtered_ShopCarts, v)
// 		}
// 	}
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.GetShopCartsList{ShopCarts: filtered_ShopCarts}, nil
// }
// func (c *Controller) StatistikaInShopCart() error {
// 	// =======================================================================================================================
// 	// statistika
// 	shopCarts, err := c.store.ShopCart().GetAllShopCarts()
// 	if err != nil {
// 		return err
// 	}

// 	hash := map[string]int{}

// 	fmt.Println("Statistika:")
// 	for _, v := range shopCarts {
// 		if v.Status {
// 			product, e := c.store.Product().GetByID(&models.ProductPrimaryKey{
// 				Id: v.ProductId,
// 			})
// 			if e != nil {
// 				return e
// 			}

// 			category, e := c.store.Category().GetByID(&models.CategoryPrimaryKey{
// 				Id: product.CategoryID,
// 			})
// 			if e != nil {
// 				return e
// 			}

// 			val, ok := hash[category.Name]
// 			if ok {
// 				hash[category.Name] = val + v.Count
// 			} else {
// 				hash[category.Name] = v.Count
// 			}

// 		}
// 	}
// 	// iterate
// 	for key, v := range hash {
// 		fmt.Printf("%s - %d\n", key, v)
// 	}

// 	// ============================================================================================================
// 	return nil
// }

// func (c *Controller) TopSellingProducts() (models.GetShopCartsList, error) {

// 	shopCarts, err := c.store.ShopCart().GetAllShopCarts()

// 	if err != nil {
// 		return models.GetShopCartsList{}, err
// 	}

// 	for i := 0; i < len(shopCarts)-1; i++ {
// 		for j := 0; j < len(shopCarts)-i-1; j++ {
// 			if shopCarts[j].Count < shopCarts[j+1].Count {
// 				shopCarts[j], shopCarts[j+1] = shopCarts[j+1], shopCarts[j]
// 			}
// 		}
// 	}

// 	ten := shopCarts[:10]

// 	for _, v := range ten {
// 		fmt.Println(v.Count)
// 	}
// 	return models.GetShopCartsList{ShopCarts: ten}, nil
// }

// func (c *Controller) GetTopSoldProducts() (models.GetShopCartsNameCount, error) {
// 	shopCarts, err := c.store.ShopCart().GetAllShopCarts()

// 	if err != nil {
// 		return models.GetShopCartsNameCount{}, err
// 	}

// 	hash := map[string]int{}

// 	for _, shopcart := range shopCarts {
// 		hash[shopcart.ProductId] += shopcart.Count
// 	}

// 	topProducts := []models.GetNameCount{}
// 	for key, val := range hash {
// 		product, err := c.store.Product().GetByID(&models.ProductPrimaryKey{Id: key})
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		filterProduct := models.GetNameCount{Name: product.Name, Count: val}
// 		topProducts = append(topProducts, filterProduct)
// 	}
// 	for i := 0; i < len(topProducts)-1; i++ {
// 		for j := 0; j < len(topProducts)-i-1; j++ {
// 			if topProducts[i].Count > topProducts[j].Count {
// 				fmt.Println("high", topProducts[:11])
// 			}
// 			//  if topProducts[i].Count < topProducts[j].Count {
// 			// 	fmt.Println("low", topProducts[:11])
// 			// }
// 		}
// 	}
// 	ten := topProducts[1:12]
// 	for key, v := range ten {
// 		fmt.Printf("%v - %v\n", key, v)
// 	}
// 	return models.GetShopCartsNameCount{GetnameCount: ten}, nil
// }

// func (c *Controller) MostActiveBuyer(req *models.ShopCartPrimaryKey) error {
// 	shopcarts, err := c.store.ShopCart().GetAllShopCarts()

// 	if err != nil {
// 		return err
// 	}

// 	myMap := map[string]int{}
// 	for _, v := range shopcarts {
// 		myMap[v.UserId] = myMap[v.UserId] + v.Count
// 	}
// 	counts := 0
// 	var userId string
// 	for key, val := range myMap {
// 		if counts < val {
// 			counts = val
// 			userId = key
// 		}
// 	}
// 	user, err := c.store.User().GetByID(&models.UserPrimaryKey{Id: userId})

// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("most Active user:\n\tName", user.Name, "\n\tOverall products count:", counts)
// 	return nil
// }
