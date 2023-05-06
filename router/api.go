package router

import (
	rescue_controller "backend/controller/rescue"
	user_controller "backend/controller/user"
	"backend/middleware"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo                 *echo.Echo
	UserController       user_controller.UserController
	RescueUnitController rescue_controller.RescueUnitController
	OrderUnitController  rescue_controller.OrderRescueUnitController
	OrderController      user_controller.OrderController
	StatsController      user_controller.StatsController
}

func (api *API) SetupRouter() {

	v1 := api.Echo.Group("/api/v1")

	user := v1.Group("/user")

	//user.POST("/user/login", api.UserController.Login)
	user.POST("/signup", api.UserController.Signup)
	user.POST("/signin", api.UserController.Login)
	user.POST("/update", api.UserController.Update, middleware.JWTMiddleware())
	user.POST("/update-role", api.UserController.UpdateRole, middleware.JWTMiddleware())
	user.POST("/update/location", api.UserController.UpdateLocationUser, middleware.JWTMiddleware())

	user.GET("/profile", api.UserController.GetProfile, middleware.JWTMiddleware())

	rescueUnit := v1.Group("/rescue-unit")
	rescueUnit.POST("/login", api.RescueUnitController.Login)
	rescueUnit.POST("/signup", api.RescueUnitController.SaveRescueUnit)
	rescueUnit.POST("/update/location", api.RescueUnitController.UpdateLocationRescueUnit, middleware.JWTMiddleware())
	rescueUnit.GET("/:id", api.RescueUnitController.GetRescueUnit, middleware.JWTMiddleware())
	//rescueUnit.GET("/all", api.RescueUnitController.GetRescueUnits, middleware.JWTMiddleware())
	rescueUnit.GET("/all", api.RescueUnitController.GetRescueUnitsByLocation, middleware.JWTMiddleware())
	rescueUnit.GET("/order/all", api.OrderUnitController.GetOrderByNear, middleware.JWTMiddleware())
	rescueUnit.POST("/order/:id/select", api.OrderUnitController.SelectOrder, middleware.JWTMiddleware())
	rescueUnit.PUT("/order/:id", api.OrderUnitController.PutOrder, middleware.JWTMiddleware())
	rescueUnit.GET("/order/pending", api.OrderUnitController.GetOrdersPending, middleware.JWTMiddleware())
	rescueUnit.GET("/order/history", api.OrderUnitController.GetOrdersByUserId, middleware.JWTMiddleware())

	order := v1.Group("/order")
	order.POST("/save", api.OrderController.SaveOrder, middleware.JWTMiddleware())
	order.PUT("/:id", api.OrderController.PutStatsOrder, middleware.JWTMiddleware())
	order.GET("/:id", api.OrderController.GetOrder, middleware.JWTMiddleware())
	order.GET("/history", api.OrderController.GetOrdersByUserId, middleware.JWTMiddleware())

	stats := v1.Group("/stats")
	stats.GET("/vehicle", api.StatsController.StatsVehicle)
	stats.GET("/rescue", api.StatsController.StatsRescueUnit)
	stats.GET("/order", api.StatsController.StatsOrder)
	stats.GET("/order-by-date", api.StatsController.StatsOrderByDate)
	stats.GET("/rating", api.StatsController.StatsOrderRating)
}
