package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsonDb"
	"fmt"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()
	jsonDb, err := jsonDb.NewFileJson(&cfg)
	if err != nil {
		panic("error while connect to json file: " + err.Error())
	}
	defer jsonDb.CloseDB()
	c := controller.NewController(&cfg, jsonDb)

	// user crud api
	http.HandleFunc("/user", c.User)
	http.HandleFunc("/user/", c.User)

	// category crud api
	http.HandleFunc("/category", c.Category)
	http.HandleFunc("/category/", c.Category)

	// product api crud
	http.HandleFunc("/product", c.Product)
	http.HandleFunc("/product/", c.Product)

	// shopcart api crud
	http.HandleFunc("/shopcart", c.ShopCart)
	http.HandleFunc("/shopcart/", c.ShopCart)

	// branch api crud
	//?? http.HandleFunc("/branch", c.Branch)
	//?? http.HandleFunc("/branch/", c.branch)

	// Running server
	fmt.Println("Server running on port 4000")
	err = http.ListenAndServe("localhost:4000", nil)
	if err != nil {
		log.Println(err)
		return
	}
}


















// !! old Codes

// package main

// import (
// 	"app/config"
// 	"app/controller"
// 	"app/models"
// 	"app/storage/jsonDb"
// 	"fmt"
// 	"log"
// )

// func main() {
// 	cfg := config.Load()

// 	jsonDb, err := jsonDb.NewFileJson(&cfg)
// 	if err != nil {
// 		log.Fatal("error while connecting to database")
// 	}
// 	defer jsonDb.CloseDb()

// 	c := controller.NewController(&cfg, jsonDb)

// 	// Filter(c)
// 	// GetHistory(c)
// 	// GetTotalPrice(c)
// 	// GetSoldProducts(c)
// 	// TopSellingProducts(c)
// 	// GetSelledProducts(c)
// 	// TopTenProducts(c)
// 	// GetTopSoldProducts(c)
// 	ActiveClient(c)

// }

// func Filter(c *controller.Controller) {
// 	data, err := c.GetFilteredShopCarts(&models.TimeFilter{
// 		DateTo:   "2022-05-11",
// 		DateFrom: "2022-05-15",
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println(data)
// }
// func GetHistory(c *controller.Controller) {
// 	history, err := c.GetClientHistory(&models.UserPrimaryKey{
// 		Id: "807f6f90-5fc0-4887-9c3d-413fd7361c56",
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println(history)
// }
// func GetTotalPrice(c *controller.Controller) {
// 	total, err := c.GetUserTotalPriceController(&models.UserPrimaryKey{
// 		Id: "807f6f90-5fc0-4887-9c3d-413fd7361c56",
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println(total)
// }
// func GetSoldProducts(c *controller.Controller) {
// 	err := c.StatistikaInShopCart()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }
// func TopSellingProducts(c *controller.Controller) {
// 	top, err := c.TopSellingProducts()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	fmt.Println(top)
// }
// func TopTenProducts(c *controller.Controller) {
// 	top, err := c.GetTopSoldProducts()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	fmt.Println(top)
// }
// func GetSelledProducts(c *controller.Controller) {
// 	selled, err := c.SelledProductThroughId(&models.ProductPrimaryKey{
// 		Id: "ace982a1-d556-40b1-9089-45c7157624eb",
// 	})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	fmt.Println(selled)
// }
// func GetTopSoldProducts(c *controller.Controller)  {
// 	top, err := c.GetTopSoldProducts()
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	fmt.Println(top)
// }
// func ActiveClient(c *controller.Controller)  {
// 	 err := c.MostActiveBuyer(&models.ShopCartPrimaryKey{
// 		Id: "5ac1bb12-4f7b-46f7-96d8-de1ffef1306c",
// 	})
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }
// func Branch(c *controller.Controller) {
// 	// Create branch
// 	id, err := c.CreateBranch(models.BranchReq{
// 		Name: "Mirobod",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(id)

// 	// Get List
// 	branches, err := c.GetBranchList(&models.GetBranchListRequest{
// 		Offset: 1,
// 		Limit:  100,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(branches.Branches)

// 	// Get by id
// 	branch, err := c.GetBranchByIdController(&models.BranchPrimaryKey{
// 		Id: "4101931f-c1f3-45a2-9e54-e2699b14bd97",
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(branch)

// 	// Update branch
// 	b, e := c.UpdateBranchController(&models.Branch{
// 		Id:   "9ec528e7-4756-4005-9c4a-fa480ff63f9f",
// 		Name: "Beruniy",
// 	})
// 	if e != nil {
// 		log.Fatal(e)
// 	}
// 	fmt.Println(b)

// 	// Remove branch
// 	b, e = c.DeleteBranchController(&models.BranchPrimaryKey{
// 		Id: "b4bc93a4-e501-4ba1-9d20-232967649276",
// 	})
// 	if e != nil {
// 		log.Fatal(e)
// 	}
// 	fmt.Println(b)

// }
// // func Product(c *controller.Controller) {

// // c.CreateProduct(&models.CreateProduct{
// // 	Name:       "Smartfon vivo V25 8/256 GB",
// // 	Price:      4_860_000,
// // 	CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
// // })

// // product, err := c.GetByIdProduct(&models.ProductPrimaryKey{Id: "38292285-4c27-497b-bc5f-dfe418a9f959"})

// // if err != nil {
// // 	log.Println(err)
// // 	return
// // }

// // 	products, err := c.GetAllProduct(
// // 		&models.GetListProductRequest{
// // 			Offset:     0,
// // 			Limit:      1,
// // 			CategoryID: "6325b81f-9a2b-48ef-8d38-5cef642fed6b",
// // 		},
// // 	)

// // 	if err != nil {
// // 		log.Println(err)
// // 		return
// // 	}

// // 	for in, product := range products.Products {
// // 		fmt.Println(in+1, product)
// // 	}
// // }

// // func Category(c *controller.Controller) {
// // c.CreateCategory(&models.CreateCategory{
// // 	Name:     "Smartfonlar va telefonlar",
// // 	ParentID: "eed2e676-1f17-429f-b75c-899eda296e65",
// // })

// // 	category, err := c.GetByIdCategory(&models.CategoryPrimaryKey{Id: "eed2e676-1f17-429f-b75c-899eda296e65"})
// // 	if err != nil {
// // 		log.Println(err)
// // 		return
// // 	}

// // 	fmt.Println(category)

// // }

// // func User(c *controller.Controller) {

// // 	sender := "bbda487b-1c0f-4c93-b17f-47b8570adfa6"
// // 	receiver := "657a41b6-1bdc-47cc-bdad-1f85eb8fb98c"
// // 	err := c.MoneyTransfer(sender, receiver, 500_000)
// // 	if err != nil {
// // 		log.Println(err)
// // 	}
// // }
