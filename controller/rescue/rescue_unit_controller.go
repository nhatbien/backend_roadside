package rescue_controller

import (
	"backend/biedeptrai"
	"backend/log"
	"backend/model"
	"backend/model/request"
	RescueRepo "backend/repository/rescue"
	"backend/security"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RescueUnitController struct {
	RescueUnitRepo RescueRepo.RescueUnitRepo
}

func NewRescueUnitController(rescueUnitRepo RescueRepo.RescueUnitRepo) *RescueUnitController {
	return &RescueUnitController{RescueUnitRepo: rescueUnitRepo}
}

func (n *RescueUnitController) Login(c echo.Context) error {
	request := request.LoginRescueUnitRequest{}
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
	rescueUnit, err := n.RescueUnitRepo.LoginRescueUnit(c.Request().Context(), request.Phone, request.Password)
	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	isTheSame := security.ComparePasswords(rescueUnit.Password, []byte(request.Password))

	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  false,
			Message: "Đăng nhậP thất bại",
			Data:    nil,
		})
	}

	//gentoken is require
	token, err := security.GenTokenResuceUnit(rescueUnit)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	rescueUnit.Token = token

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Đăng nhập thành công",
		Data:    rescueUnit,
	})
}

func (n *RescueUnitController) UpdateLocationRescueUnit(c echo.Context) error {

	request := request.UpdateLocationResuceUnitRequest{}
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

	rescueUnit, err := n.RescueUnitRepo.UpdateLocationRescueUnit(c.Request().Context(), claims.Id, request.Lat, request.Lng)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Cập nhật thành công",
		Data:    rescueUnit,
	})
}

func (n *RescueUnitController) SaveRescueUnit(c echo.Context) error {
	request := request.SaveRescueUnitRequest{}

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
	hash := security.HashAndSalt([]byte(request.Password))
	userId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	resuceUnit := model.RescueUnit{
		Id:       userId.String(),
		Name:     request.Name,
		Phone:    request.Phone,
		Email:    request.Email,
		Address:  request.Address,
		Avatar:   request.Avatar,
		Password: hash,
		Lat:      request.Lat,
		Lng:      request.Lng,
		Type:     request.Type,
		Status:   request.Status,
	}
	rescueUnit, err := n.RescueUnitRepo.SaveRescueUnit(c.Request().Context(), resuceUnit)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lưu thành công",
		Data:    rescueUnit,
	})

}

func (n *RescueUnitController) GetRescueUnit(c echo.Context) error {
	rescueId := c.Param("id")
	rescueUnit, err := n.RescueUnitRepo.GetRescueUnit(c.Request().Context(), rescueId)
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
		Data:    rescueUnit,
	})
}

func (n *RescueUnitController) GetRescueUnits(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if claims.Role.RoleName != "admin" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: biedeptrai.ErrorRoleUser.Error(),
			Data:    nil,
		})
	}

	rescueUnits, err := n.RescueUnitRepo.GetRescueUnits(c.Request().Context())
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
		Data:    rescueUnits,
	})
}

func (n *RescueUnitController) GetRescueUnitsByLocation(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if claims.Role.RoleName != "admin" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: biedeptrai.ErrorRoleUser.Error(),
			Data:    nil,
		})
	}

	queryLat := c.QueryParam("lat")
	queryLng := c.QueryParam("lng")

	if queryLat == "" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "lat is null",
			Data:    nil,
		})
	}
	if queryLng == "" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: "lng is null",
			Data:    nil,
		})
	}

	lat, err := strconv.ParseFloat(queryLat, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	lng, err := strconv.ParseFloat(queryLng, 64)
	if err != nil {
		fmt.Println("Error parsing float:", err)
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	if err != nil {
		fmt.Println("Error parsing float:", err)
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	rescueUnits, err := n.RescueUnitRepo.GetRescueUnitsByLocation(c.Request().Context(), lat, lng, 1)
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
		Data:    rescueUnits,
	})
}
