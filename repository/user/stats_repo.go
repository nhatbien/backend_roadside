package user_repo

import (
	"backend/model"
	"context"
	"time"
)

type StatsRepo interface {
	StatsVehicle(context context.Context) (interface{}, error)
	StatsRescueUnit(context context.Context) (interface{}, error)
	StatsOrderByDate(context context.Context, startDate time.Time, endDate time.Time) ([]model.Order, error)
	StatsOrder(context context.Context) (interface{}, error)
}
