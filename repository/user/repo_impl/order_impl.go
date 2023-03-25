package userImpl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	user_repo "backend/repository/user"
	"context"

	"gorm.io/gorm/clause"
)

type OrderRepoImpl struct {
	sql *db.Sql
}

func NewOrderRepo(sql *db.Sql) user_repo.OrderUserRepo {
	return &OrderRepoImpl{sql: sql}
}
func (n *OrderRepoImpl) SaveOrder(context context.Context, order model.Order) (model.Order, error) {
	user := model.User{}
	if res := n.sql.Db.Where(&model.User{Id: order.UserId}).First(&user); res.RowsAffected <= 0 {
		return order, biedeptrai.ErrorUserNotFound
	}
	order.Lat = user.Lat
	order.Lng = user.Lng

	err := n.sql.Db.Preload(clause.Associations).Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (n *OrderRepoImpl) PutOrder(context context.Context, order model.Order) (model.Order, error) {
	err := n.sql.Db.Updates(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (n *OrderRepoImpl) GetOrder(context context.Context, orderId int) (model.Order, error) {

	var order model.Order
	if res := n.sql.Db.Where(
		&model.Order{Id: orderId},
	).Preload(clause.Associations).First(&order); res.RowsAffected <= 0 {
		return order, nil
	}
	return order, nil
}

func (n *OrderRepoImpl) GetOrders(context context.Context) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(
		&model.Order{},
	).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}

func (n *OrderRepoImpl) GetOrdersByUserId(context context.Context, userId string) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(
		&model.Order{UserId: userId},
	).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}

/*
func (n *OrderRepoImpl) GetOrdersByRescueUnitId(context context.Context, rescueUnitId string) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(
		&model.Order{RescueUnitId: rescueUnitId},
	).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
} */
func (n *OrderRepoImpl) GetOrdersByStatus(context context.Context, status int) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(
		&model.Order{Status: status},
	).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}
func (n *OrderRepoImpl) GetOrdersByUserIdAndStatus(context context.Context, userId string, status int) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(
		&model.Order{UserId: userId, Status: status},
	).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}
