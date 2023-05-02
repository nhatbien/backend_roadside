package user_controller

import (
	"backend/model"
	user_repo "backend/repository/user"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type StatsController struct {
	StatsRepo user_repo.StatsRepo
}

func NewStatsController(statsRepo user_repo.StatsRepo) *StatsController {
	return &StatsController{StatsRepo: statsRepo}
}

func (s *StatsController) StatsVehicle(c echo.Context) error {
	stats, err := s.StatsRepo.StatsVehicle(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Success",
		Data:    stats,
	})
}

func (s *StatsController) StatsRescueUnit(c echo.Context) error {
	stats, err := s.StatsRepo.StatsRescueUnit(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Success",
		Data:    stats,
	})
}
func (s *StatsController) StatsOrder(c echo.Context) error {
	stats, err := s.StatsRepo.StatsOrder(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Success",
		Data:    stats,
	})
}

func (s *StatsController) StatsOrderByDate(c echo.Context) error {
	startDate := c.QueryParam("startDate")
	endDate := c.QueryParam("endDate")
	start, _ := time.Parse("2006-01-02", startDate)
	end, _ := time.Parse("2006-01-02", endDate)
	stats, err := s.StatsRepo.StatsOrderByDate(c.Request().Context(), start, end)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Success",
		Data:    stats,
	})
}

func (s *StatsController) StatsOrderRating(c echo.Context) error {

	stats, err := s.StatsRepo.StatsOrderRating(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			Status:  false,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, model.Response{
		Status:  true,
		Message: "Success",
		Data:    stats,
	})
}
