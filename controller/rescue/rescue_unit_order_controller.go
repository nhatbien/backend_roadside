package rescue_controller

import (
	"backend/model"
	"backend/model/request"
	RescueRepo "backend/repository/rescue"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type OrderRescueUnitController struct {
	OrderRescueUnitRepo RescueRepo.OrderRepo
}

func NewOrderRescueUnitController(orderRescueUnitRepo RescueRepo.OrderRepo) *OrderRescueUnitController {
	return &OrderRescueUnitController{OrderRescueUnitRepo: orderRescueUnitRepo}
}

func (n *OrderRescueUnitController) GetOrderByNear(c echo.Context) error {

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	orders, err := n.OrderRescueUnitRepo.GetOrdersNearbyUnit(c.Request().Context(), claims.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	if len(orders) == 0 {
		return c.JSON(http.StatusOK, model.Response{
			Status:  true,
			Message: "Không có đơn hàng nào gần bạn",
			Data:    []model.Order{},
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lấy danh sách thành công",
		Data:    orders,
	})
}

func (n *OrderRescueUnitController) SelectOrder(c echo.Context) error {

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	idOrder, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	order, err := n.OrderRescueUnitRepo.SelectOrder(c.Request().Context(), claims.Id, idOrder)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lấy đơn thành công",
		Data:    order,
	})
}

func (n *OrderRescueUnitController) GetOrdersPending(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	/* if claims.Role.RoleName != "rescue_unit" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: biedeptrai.ErrorRoleUser.Error(),
			Data:    nil,
		})
	} */

	orders, err := n.OrderRescueUnitRepo.GetOrdersPending(c.Request().Context(), claims.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lấy dữ liệu thành công",
		Data:    orders,
	})
}

func (n *OrderRescueUnitController) PutOrder(c echo.Context) error {

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

	idOrder, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	orderModel := model.Order{
		Id:     idOrder,
		Status: request.Status,
	}
	order, err := n.OrderRescueUnitRepo.PutOrder(c.Request().Context(), orderModel)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lấy đơn thành công",
		Data:    order,
	})
}
