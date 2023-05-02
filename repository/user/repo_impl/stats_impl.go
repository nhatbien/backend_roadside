package userImpl

import (
	"backend/db"
	"backend/model"
	user_repo "backend/repository/user"
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

type StatsRepoImpl struct {
	sql *db.Sql
}

func NewStatsRepo(sql *db.Sql) user_repo.StatsRepo {
	return &StatsRepoImpl{sql: sql}
}

func (n *StatsRepoImpl) StatsVehicle(context context.Context) (interface{}, error) {
	now := time.Now()

	// Tính ngày đầu tiên của tháng hiện tại
	startOfMonth := time.Date(now.Year()-1, now.Month(), 1, 0, 0, 0, 0, time.Local)
	fmt.Println(startOfMonth)
	// Tạo slice chứa 12 tháng gần nhất
	var months []string
	for i := 0; i < 12; i++ {
		month := now.AddDate(0, -i, 0)
		months = append(months, month.Format("01/2006"))
	}
	// Đảo ngược thứ tự các tháng để đúng thứ tự thời gian
	for i, j := 0, len(months)-1; i < j; i, j = i+1, j-1 {
		months[i], months[j] = months[j], months[i]
	}

	// Thực hiện truy vấn dữ liệu
	var results []struct {
		Month string
		Count int
	}

	if err := n.sql.Db.Table("users").
		Select("DATE_FORMAT(created_at, '%m/%Y') AS month, COUNT(*) AS count").
		Where("created_at BETWEEN ? AND ?", startOfMonth, now).
		Group("month").
		Order("month DESC").
		Limit(12).
		Find(&results).Error; err != nil {
		return nil, err
	}
	fmt.Println(results)

	// Tạo slice chứa số lượng user đăng ký của 12 tháng gần nhất
	var counts []int
	for i := 0; i < 12; i++ {
		month := now.AddDate(0, -i, 0)
		count := 0
		for _, result := range results {
			if result.Month == month.Format("01/2006") {
				count = result.Count
				break
			}
		}
		counts = append(counts, count)
	}
	// Đảo ngược thứ tự các số liệu để đúng thứ tự thời gian
	for i, j := 0, len(counts)-1; i < j; i, j = i+1, j-1 {
		counts[i], counts[j] = counts[j], counts[i]
	}

	// Trả về kết quả dưới dạng map[string]interface{}
	result := map[string]interface{}{
		"month": months,
		"value": counts,
	}
	return result, nil
}

func (n *StatsRepoImpl) StatsRescueUnit(context context.Context) (interface{}, error) {
	now := time.Now()

	// Tính ngày đầu tiên của tháng hiện tại
	startOfMonth := time.Date(now.Year()-1, now.Month(), 1, 0, 0, 0, 0, time.Local)
	fmt.Println(startOfMonth)
	// Tạo slice chứa 12 tháng gần nhất
	var months []string
	for i := 0; i < 12; i++ {
		month := now.AddDate(0, -i, 0)
		months = append(months, month.Format("01/2006"))
	}
	// Đảo ngược thứ tự các tháng để đúng thứ tự thời gian
	for i, j := 0, len(months)-1; i < j; i, j = i+1, j-1 {
		months[i], months[j] = months[j], months[i]
	}

	// Thực hiện truy vấn dữ liệu
	var results []struct {
		Month string
		Count int
	}

	if err := n.sql.Db.Table("rescue_units").
		Select("DATE_FORMAT(created_at, '%m/%Y') AS month, COUNT(*) AS count").
		Where("created_at BETWEEN ? AND ?", startOfMonth, now).
		Group("month").
		Order("month DESC").
		Limit(12).
		Find(&results).Error; err != nil {
		return nil, err
	}
	fmt.Println(results)

	// Tạo slice chứa số lượng user đăng ký của 12 tháng gần nhất
	var counts []int
	for i := 0; i < 12; i++ {
		month := now.AddDate(0, -i, 0)
		count := 0
		for _, result := range results {
			if result.Month == month.Format("01/2006") {
				count = result.Count
				break
			}
		}
		counts = append(counts, count)
	}
	// Đảo ngược thứ tự các số liệu để đúng thứ tự thời gian
	for i, j := 0, len(counts)-1; i < j; i, j = i+1, j-1 {
		counts[i], counts[j] = counts[j], counts[i]
	}

	// Trả về kết quả dưới dạng map[string]interface{}
	result := map[string]interface{}{
		"month": months,
		"value": counts,
	}
	return result, nil
}

func (n *StatsRepoImpl) StatsOrder(context context.Context) (interface{}, error) {
	now := time.Now()

	// Tính ngày đầu tiên của tháng hiện tại
	startOfMonth := time.Date(now.Year()-1, now.Month(), 1, 0, 0, 0, 0, time.Local)
	fmt.Println(startOfMonth)
	// Tạo slice chứa 12 tháng gần nhất
	var months []string
	for i := 0; i < 12; i++ {
		month := now.AddDate(0, -i, 0)
		months = append(months, month.Format("01/2006"))
	}
	// Đảo ngược thứ tự các tháng để đúng thứ tự thời gian
	for i, j := 0, len(months)-1; i < j; i, j = i+1, j-1 {
		months[i], months[j] = months[j], months[i]
	}

	// Thực hiện truy vấn dữ liệu
	var results []struct {
		Month string
		Count int
	}

	if err := n.sql.Db.Table("orders").
		Select("DATE_FORMAT(created_at, '%m/%Y') AS month, COUNT(*) AS count").
		Where("created_at BETWEEN ? AND ?", startOfMonth, now).
		Group("month").
		Order("month DESC").
		Limit(12).
		Find(&results).Error; err != nil {
		return nil, err
	}
	fmt.Println(results)

	// Tạo slice chứa số lượng user đăng ký của 12 tháng gần nhất
	var counts []int
	for i := 0; i < 12; i++ {
		month := now.AddDate(0, -i, 0)
		count := 0
		for _, result := range results {
			if result.Month == month.Format("01/2006") {
				count = result.Count
				break
			}
		}
		counts = append(counts, count)
	}
	// Đảo ngược thứ tự các số liệu để đúng thứ tự thời gian
	for i, j := 0, len(counts)-1; i < j; i, j = i+1, j-1 {
		counts[i], counts[j] = counts[j], counts[i]
	}

	// Trả về kết quả dưới dạng map[string]interface{}
	result := map[string]interface{}{
		"month": months,
		"value": counts,
	}
	return result, nil
}

func (n *StatsRepoImpl) StatsOrderByDate(context context.Context, startDate time.Time, endDate time.Time) ([]model.Order, error) {

	var orders []model.Order
	if res := n.sql.Db.Where(
		"created_at BETWEEN ? AND ?",
		startDate,
		endDate,
	).Preload(clause.Associations).Find(&orders); res.RowsAffected <= 0 {
		return orders, nil
	}
	return orders, nil
}
func (n *StatsRepoImpl) StatsOrderRating(context context.Context) (interface{}, error) {
	var ratingCount RatingCount

	n.sql.Db.Raw(`SELECT 
    SUM(IF(stats >= 4 AND stats < 5, 1, 0)) AS rating_4_5,
    SUM(IF(stats >= 3 AND stats < 4, 1, 0)) AS rating_3_4,
    SUM(IF(stats >= 2 AND stats < 3, 1, 0)) AS rating_2_3,
    SUM(IF(stats >= 1 AND stats < 2, 1, 0)) AS rating_1_2,
    SUM(IF(stats >= 0 AND stats < 1, 1, 0)) AS rating_0_1 
FROM 
    orders;`).Scan(&ratingCount)
	result := map[string]interface{}{
		"rating_4_5": ratingCount.Rating45,
		"rating_3_4": ratingCount.Rating34,
		"rating_2_3": ratingCount.Rating23,
		"rating_1_2": ratingCount.Rating12,
		"rating_0_1": ratingCount.Rating01,
	}

	return result, nil

}

type RatingCount struct {
	Rating45 int `gorm:"column:rating_4_5"`
	Rating34 int `gorm:"column:rating_3_4"`
	Rating23 int `gorm:"column:rating_2_3"`
	Rating12 int `gorm:"column:rating_1_2"`
	Rating01 int `gorm:"column:rating_0_1"`
}
