package app

import "github.com/mamtaharris/shopping-portal/controllers"

func mappings() {
	router.GET("/", controllers.HomeHandler)

	router.POST("/user/create", controllers.CreateUserHandler)
	router.POST("/user/login", controllers.LoginHandler)
	router.GET("/user/list", controllers.ListUsersHandler)

	router.POST("/item/create", controllers.CreateItemHandler)
	router.GET("/item/list", controllers.ListItemsHandler)

	router.POST("/cart/add", controllers.AddToCartHandler)
	router.POST("/cart/:cartId/complete", controllers.OrderFromCartHandler)
	router.GET("/cart/list", controllers.ListCartsHandler)

	router.GET("/order/list", controllers.ListOrdersHandler)
}
