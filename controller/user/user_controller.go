package user_controller

import (
	"backend/biedeptrai"
	"backend/log"
	"backend/model"
	"backend/model/request"
	repository "backend/repository/user"
	"backend/security"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserRepo repository.UserRepo
}

func (u *UserController) Signup(c echo.Context) error {

	request := request.UserSignupRequest{}

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
	//role := model.MEMBER.String()
	//role := model.ADMIN.String()
	userId, err := uuid.NewUUID()
	if err != nil {
		return c.JSON(http.StatusForbidden, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	userModel := model.User{
		Id:        userId.String(),
		FullName:  request.FullName,
		Phone:     request.Phone,
		Avatar:    request.Avatar,
		Email:     request.Email,
		Username:  request.Username,
		Age:       request.Age,
		Address:   request.Address,
		Password:  hash,
		RoleId:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user, err := u.UserRepo.SaveUser(c.Request().Context(), userModel)
	if err != nil {
		return c.JSON(http.StatusConflict, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Đăng ký thành công",
		Data:    user,
	})
}

func (u *UserController) Login(c echo.Context) error {
	request := request.UserLoginRequest{}
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

	user, err := u.UserRepo.Login(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	isTheSame := security.ComparePasswords(user.Password, []byte(request.Password))

	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			Status:  false,
			Message: "Đăng nhậP thất bại",
			Data:    nil,
		})
	}

	//gentoken is require
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Đăng nhập thành công",
		Data:    user,
	})
}

func (u *UserController) Update(c echo.Context) error {
	request := request.UserUpdateRequest{}

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

	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	user := model.User{
		Id:        claims.Id,
		Username:  request.Username,
		Email:     request.Email,
		Phone:     request.Phone,
		Avatar:    request.Avatar,
		FullName:  request.FullName,
		Status:    request.Status,
		Age:       request.Age,
		Address:   request.Address,
		UpdatedAt: time.Now(),
	}

	_, err := u.UserRepo.UpdateUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  false,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lưu thành công",
	})
}

func (u *UserController) UpdateRole(c echo.Context) error {
	request := request.UserUpdateRoleRequest{}
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if claims.Role.RoleName != "admin" {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: biedeptrai.ErrorRoleUser.Error(),
			Data:    nil,
		})
	}

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

	/* err := u.UserRepo.UpdateRole(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
	} */
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lưu thành công",
	})
}

func (u *UserController) GetProfile(c echo.Context) error {

	tokenData := c.Get("user").(*jwt.Token)

	claims := tokenData.Claims.(*model.JwtCustomClaims)

	user, err := u.UserRepo.GetProfile(c.Request().Context(), claims.Id)

	if err != nil {
		return c.JSON(http.StatusNotFound, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Lấy thông tin thành công",
		Data:    user,
	})
}

func (n *UserController) UpdateLocationUser(c echo.Context) error {

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

	rescueUnit, err := n.UserRepo.UpdateLocationUser(c.Request().Context(), claims.Id, request.Lat, request.Lng)
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
