package user_controller

import (
	"backend/model"
	"backend/model/request"
	userRepo "backend/repository/user"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

/*
status 1 : chờ xác nhận
status 2 : đang giao
status 3 : đã giao
status 4 : đã hủy


*/
type OrderController struct {
	OrderUserRepo userRepo.OrderUserRepo
}

func NewOrderController(orderRepo userRepo.OrderUserRepo) OrderController {
	return OrderController{OrderUserRepo: orderRepo}
}

func (o *OrderController) SaveOrder(c echo.Context) error {
	request := request.SaveOrderRequest{}

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	order := model.Order{
		UserId:    claims.Id,
		Note:      request.Note,
		Address:   request.Address,
		Status:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	order, err := o.OrderUserRepo.SaveOrder(c.Request().Context(), order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Tạo đơn hàng thành công",
		Data:    order,
	})

}
func (o *OrderController) PutOrder(c echo.Context) error {
	request := request.PutOrderRequest{}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	order := model.Order{
		Status:    request.Status,
		UpdatedAt: time.Now(),
	}
	order, err := o.OrderUserRepo.PutOrder(c.Request().Context(), order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})

	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Cập nhật thành công",
		Data:    order,
	})

}

func (o *OrderController) GetOrder(c echo.Context) error {
	id := c.Param("id")

	num, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "Lỗi Order ID",
			Data:    nil,
		})
	}
	fmt.Print(1)

	order, err := o.OrderUserRepo.GetOrder(c.Request().Context(), num)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lấy dữ liệu thành công",
		Data:    order,
	})
}

/*
func (o *OrderController) GetOrder(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	orders, err := o.OrderRepo.GetOrder(c.Request().Context(), claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, orders)
}

func (o *OrderController) GetOrderById(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)
	id := c.Param("id")
	order, err := o.OrderRepo.GetOrderById(c.Request().Context(), id, claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}
*/
