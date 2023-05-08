package rescue_impl

import (
	"backend/biedeptrai"
	"backend/db"
	"backend/model"
	repository "backend/repository/rescue"
	"backend/security"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

type RescueUnitRepoImpl struct {
	sql *db.Sql
}

func NewRescueUnitRepo(sql *db.Sql) repository.RescueUnitRepo {
	return &RescueUnitRepoImpl{sql: sql}
}

func (n *RescueUnitRepoImpl) LoginRescueUnit(context context.Context, phone string, password string) (model.RescueUnit, error) {
	var rescueUnit model.RescueUnit

	if res := n.sql.Db.Where(
		&model.RescueUnit{Phone: phone},
	).First(&rescueUnit); res.RowsAffected <= 0 {
		return rescueUnit, biedeptrai.ErrorRescueUnitNotFound
	}
	if !security.ComparePasswords(rescueUnit.Password, []byte(password)) {
		return rescueUnit, biedeptrai.ErrorPasswordNotMatch
	}
	return rescueUnit, nil
}

func (n *RescueUnitRepoImpl) SaveRescueUnit(context context.Context, rescueUnit model.RescueUnit) (model.RescueUnit, error) {
	rescueUnit.CreatedAt = time.Now()
	rescueUnit.UpdatedAt = time.Now()

	err := n.sql.Db.Create(rescueUnit).Error
	if err != nil {
		return rescueUnit, biedeptrai.ErrorSaveRescueUnitFail
	}
	return rescueUnit, nil
}
func (n *RescueUnitRepoImpl) GetRescueUnit(context context.Context, rescueUnitId string) (model.RescueUnit, error) {
	var rescueUnit model.RescueUnit
	if res := n.sql.Db.Where(
		&model.RescueUnit{Id: rescueUnitId},
	).Preload(clause.Associations).First(&rescueUnit); res.RowsAffected <= 0 {
		return rescueUnit, biedeptrai.ErrorRescueUnitNotFound
	}
	return rescueUnit, nil
}

func (n *RescueUnitRepoImpl) GetRescueUnits(context context.Context) ([]model.RescueUnit, error) {
	var rescueUnits []model.RescueUnit

	if res := n.sql.Db.Where(
		&model.RescueUnit{},
	).Find(&rescueUnits); res.RowsAffected <= 0 {
		return rescueUnits, biedeptrai.ErrorRescueUnitNotFound
	}

	return rescueUnits, nil
}

func (n *RescueUnitRepoImpl) GetRescueUnitsByLocation(context context.Context, lat float64, lng float64, radius float64) ([]model.RescueUnit, error) {
	var rescueUnits []model.RescueUnit
	fmt.Print(lat)
	fmt.Print(lng)
	if res := n.sql.Db.Raw(`SELECT *, ( 3959 * acos( cos( radians(?) ) * cos( radians( lat ) ) 
    * cos( radians( lng ) - radians(?) ) + sin( radians(?) ) * sin(radians(lat)) ) ) AS distance 
FROM rescue_units
WHERE status = 1

HAVING distance < 5
ORDER BY distance `, lat, lng, lat).Scan(&rescueUnits); res.RowsAffected <= 0 {
		return rescueUnits, biedeptrai.ErrorRescueUnitNotFound
	}

	return rescueUnits, nil
}

func (n *RescueUnitRepoImpl) GetRescueUnitsByLocationAndType(context context.Context, lat float64, lng float64, radius float64, rescueUnitType int) ([]model.RescueUnit, error) {
	var rescueUnits []model.RescueUnit

	if res := n.sql.Db.Where(
		"lat BETWEEN ? AND ? AND lng BETWEEN ? AND ? AND type = ?",
		lat-radius, lat+radius, lng-radius, lng+radius, rescueUnitType,
	).Find(&rescueUnits); res.RowsAffected <= 0 {
		return rescueUnits, biedeptrai.ErrorRescueUnitNotFound
	}

	return rescueUnits, nil
}

func (n *RescueUnitRepoImpl) GetRescueUnitsByLocationAndTypes(context context.Context, lat float64, lng float64, radius float64, rescueUnitTypes []int) ([]model.RescueUnit, error) {
	var rescueUnits []model.RescueUnit

	if res := n.sql.Db.Where(
		"lat BETWEEN ? AND ? AND lng BETWEEN ? AND ? AND type IN (?)",
		lat-radius, lat+radius, lng-radius, lng+radius, rescueUnitTypes,
	).Find(&rescueUnits); res.RowsAffected <= 0 {
		return rescueUnits, biedeptrai.ErrorRescueUnitNotFound
	}

	return rescueUnits, nil
}
func (n *RescueUnitRepoImpl) GetRescueUnitsByLocationAndStatus(context context.Context, lat float64, lng float64, radius float64, rescueUnitStatus int) ([]model.RescueUnit, error) {
	var rescueUnits []model.RescueUnit

	if res := n.sql.Db.Where(
		"lat BETWEEN ? AND ? AND lng BETWEEN ? AND ? AND status = ?",
		lat-radius, lat+radius, lng-radius, lng+radius, rescueUnitStatus,
	).Find(&rescueUnits); res.RowsAffected <= 0 {
		return rescueUnits, biedeptrai.ErrorRescueUnitNotFound
	}

	return rescueUnits, nil
}

func (n *RescueUnitRepoImpl) GetRescueUnitsByLocationAndTypeAndStatus(context context.Context, lat float64, lng float64, radius float64, rescueUnitType int, rescueUnitStatus int) ([]model.RescueUnit, error) {
	var rescueUnits []model.RescueUnit

	if res := n.sql.Db.Where(
		"lat BETWEEN ? AND ? AND lng BETWEEN ? AND ? AND type = ? AND status = ?",
		lat-radius, lat+radius, lng-radius, lng+radius, rescueUnitType, rescueUnitStatus,
	).Find(&rescueUnits); res.RowsAffected <= 0 {
		return rescueUnits, biedeptrai.ErrorRescueUnitNotFound
	}

	return rescueUnits, nil
}

func (n *RescueUnitRepoImpl) UpdateLocationRescueUnit(context context.Context, rescueUnitId string, lat float64, lng float64) (model.RescueUnit, error) {
	var rescueUnit model.RescueUnit

	if res := n.sql.Db.Where(
		&model.RescueUnit{Id: rescueUnitId},
	).First(&rescueUnit); res.RowsAffected <= 0 {
		return rescueUnit, biedeptrai.ErrorRescueUnitNotFound
	}

	rescueUnit.Lat = lat
	rescueUnit.Lng = lng
	rescueUnit.UpdatedAt = time.Now()

	err := n.sql.Db.Save(rescueUnit).Error
	if err != nil {
		return rescueUnit, biedeptrai.ErrorUpdateRescueUnitFail
	}
	return rescueUnit, nil
}
