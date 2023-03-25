package rescue_repository

import (
	"backend/model"
	"context"
)

type RescueUnitRepo interface {
	LoginRescueUnit(context context.Context, phone string, password string) (model.RescueUnit, error)
	SaveRescueUnit(context context.Context, rescueUnit model.RescueUnit) (model.RescueUnit, error)
	GetRescueUnit(context context.Context, rescueUnitId string) (model.RescueUnit, error)
	GetRescueUnits(context context.Context) ([]model.RescueUnit, error)
	GetRescueUnitsByLocation(context context.Context, lat float64, lng float64, radius float64) ([]model.RescueUnit, error)
	GetRescueUnitsByLocationAndType(context context.Context, lat float64, lng float64, radius float64, rescueUnitType int) ([]model.RescueUnit, error)
	GetRescueUnitsByLocationAndTypes(context context.Context, lat float64, lng float64, radius float64, rescueUnitTypes []int) ([]model.RescueUnit, error)
	GetRescueUnitsByLocationAndStatus(context context.Context, lat float64, lng float64, radius float64, rescueUnitStatus int) ([]model.RescueUnit, error)
	GetRescueUnitsByLocationAndTypeAndStatus(context context.Context, lat float64, lng float64, radius float64, rescueUnitType int, rescueUnitStatus int) ([]model.RescueUnit, error)
	UpdateLocationRescueUnit(context context.Context, rescueUnitId string, lat float64, lng float64) (model.RescueUnit, error)
}
