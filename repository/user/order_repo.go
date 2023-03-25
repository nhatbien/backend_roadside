package user_repo

import (
	"backend/model"
	"context"
)

type OrderUserRepo interface {
	SaveOrder(context context.Context, order model.Order) (model.Order, error)
	PutOrder(context context.Context, order model.Order) (model.Order, error)

	GetOrder(context context.Context, orderId int) (model.Order, error)
	GetOrders(context context.Context) ([]model.Order, error)
	GetOrdersByUserId(context context.Context, userId string) ([]model.Order, error)
	/* 	GetOrdersByRescueUnitId(context context.Context, rescueUnitId string) ([]model.Order, error)
	 */GetOrdersByStatus(context context.Context, status int) ([]model.Order, error)
	GetOrdersByUserIdAndStatus(context context.Context, userId string, status int) ([]model.Order, error)
	/* 	GetOrdersByRescueUnitIdAndStatus(context context.Context, rescueUnitId string, status int) ([]model.Order, error)
	 */
}
