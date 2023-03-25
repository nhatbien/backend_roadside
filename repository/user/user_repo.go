package user_repo

import (
	"backend/model"
	"backend/model/request"
	"context"
)

type UserRepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	Login(context context.Context, loginRequest request.UserLoginRequest) (model.User, error)
	UpdateUser(context context.Context, user model.User) (model.User, error)
	GetProfile(context context.Context, userId string) (model.User, error)
	UpdateRole(context context.Context, userRole request.UserUpdateRoleRequest) error
	DeleteUser(context context.Context, userId string) error
	SelectAllUser(context context.Context) ([]model.User, error)
	SelectAllUserByRole(context context.Context, roleId int) ([]model.User, error)
	UpdateLocationUser(context context.Context, rescueUnitId string, lat float64, lng float64) (model.User, error)
}
