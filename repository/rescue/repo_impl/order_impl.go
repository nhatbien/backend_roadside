package rescue_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	repository "backend/repository/rescue"
	"context"

	"gorm.io/gorm/clause"
)

type OrderRepoImpl struct {
	sql *db.Sql
}

func NewOrderRepo(sql *db.Sql) repository.OrderRepo {
	return &OrderRepoImpl{sql: sql}
}
func (n *OrderRepoImpl) SaveOrder(context context.Context, order model.Order) (model.Order, error) {
	err := n.sql.Db.Create(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func (n *OrderRepoImpl) PutOrder(context context.Context, order model.Order) (model.Order, error) {
	var orderModel model.Order

	err := n.sql.Db.Preload(clause.Associations).Find(&orderModel, order.Id).Error
	orderModel.Status = order.Status

	if err != nil {
		return order, err
	}

	if err := n.sql.Db.Save(&orderModel).Error; err != nil {
		return orderModel, err
	}

	return orderModel, nil
}

func (n *OrderRepoImpl) GetOrder(context context.Context, orderId int) (model.Order, error) {
	var order model.Order
	if res := n.sql.Db.Where(
		&model.Order{Id: orderId},
	).Preload(clause.Associations).Find(&order).First(&order); res.RowsAffected <= 0 {
		return order, nil
	}
	return order, nil
}

func (n *OrderRepoImpl) GetOrders(context context.Context) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(
		&model.Order{},
	).Preload(clause.Associations).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}

func (n *OrderRepoImpl) GetOrdersByUserId(context context.Context, userId string) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(
		&model.Order{RescueUnitId: &userId},
	).Preload(clause.Associations).Find(&orders); res.RowsAffected <= 0 {
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
	).Preload(clause.Associations).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}
func (n *OrderRepoImpl) GetOrdersByUserIdAndStatus(context context.Context, userId string, status int) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(
		&model.Order{UserId: userId, Status: status},
	).Preload(clause.Associations).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}

func (n *OrderRepoImpl) GetOrdersNearbyUnit(context context.Context, rescueUnitId string) ([]model.Order, error) {
	var orders []model.Order
	var rescueUnit model.RescueUnit
	if res := n.sql.Db.Where(&model.RescueUnit{Id: rescueUnitId}).First(&rescueUnit); res.RowsAffected <= 0 {
		return orders, nil
	}
	lat := rescueUnit.Lat
	lng := rescueUnit.Lng

	res := n.sql.Db.Raw(`SELECT *, ( 3959 * acos( cos( radians(?) ) * cos( radians( lat ) ) 
    * cos( radians( lng ) - radians(?) ) + sin( radians(?) ) * sin(radians(lat)) ) ) AS distance 
	FROM orders
	WHERE status = 1
	HAVING distance < 20
	ORDER BY distance `, lat, lng, lat).Preload(clause.Associations).Find(&orders).Preload(clause.Associations)

	if res != nil {
		return orders, res.Error
	}

	return orders, nil
}

func (n *OrderRepoImpl) SelectOrder(context context.Context, rescueUnitId string, orderId int) (model.Order, error) {
	var order model.Order

	if res := n.sql.Db.Preload(clause.Associations).Find(
		&model.Order{Id: orderId},
	).First(&order); res.RowsAffected <= 0 {
		return order, nil
	}

	if order.Status != 1 {
		return order, biedeptrai.ErrOrderRecevied
	}

	err := n.sql.Db.Model(&order).Updates(&model.Order{Status: 2, RescueUnitId: &rescueUnitId}).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (n *OrderRepoImpl) GetOrdersPending(context context.Context, rescueUnitId string) ([]model.Order, error) {
	var orders []model.Order
	if res := n.sql.Db.Where(&model.Order{RescueUnitId: &rescueUnitId}).Where("status <> ?", 4).Preload(clause.Associations).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}
