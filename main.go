package main

import (
	rescue_controller "backend/controller/rescue"
	user_controller "backend/controller/user"
	"backend/db"
	"backend/helper"
	rescue_impl "backend/repository/rescue/repo_impl"
	userImpl "backend/repository/user/repo_impl"
	"backend/router"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	sql := &db.Sql{
		Host:     "localhost",
		Port:     "3306",
		User:     "roadside_assistance",
		Password: "KhongPhaiLaMatKhau123456@",
		Dbname:   "roadside_assistance",
	}
	sql.Connect()
	e := echo.New()
	e.Use(middleware.CORS())

	structValidator := helper.NewStructValidaten()
	structValidator.RegisterValidate()
	e.Validator = structValidator

	userController := user_controller.UserController{
		UserRepo: userImpl.NewUserRepo(sql)}

	orderController := user_controller.OrderController{
		OrderUserRepo: userImpl.NewOrderRepo(sql),
	}
	rescueUnitController := rescue_controller.RescueUnitController{
		RescueUnitRepo: rescue_impl.NewRescueUnitRepo(sql),
	}
	rescueUnitOrderController := rescue_controller.OrderRescueUnitController{
		OrderRescueUnitRepo: rescue_impl.NewOrderRepo(sql),
	}
	statsController := user_controller.StatsController{
		StatsRepo: userImpl.NewStatsRepo(sql),
	}

	api := router.API{
		Echo:                 e,
		UserController:       userController,
		RescueUnitController: rescueUnitController,
		OrderController:      orderController,
		OrderUnitController:  rescueUnitOrderController,
		StatsController:      statsController,
	}

	api.SetupRouter()

	e.Logger.Fatal(e.Start(":3001"))
	//	defer sql.Close()

}
