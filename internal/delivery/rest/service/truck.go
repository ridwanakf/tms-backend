package service

import (
	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/app"
	"net/http"
)

type TruckService struct {
	uc internal.TruckUC
}

func NewTruckService(app *app.ShipmentApp) *TruckService {
	return &TruckService{
		uc: app.Usecases.TruckUC,
	}
}

func (s *TruckService) CreateNewTruckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *TruckService) GetTruckListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *TruckService) GetTruckByIDHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *TruckService) UpdateTruckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *TruckService) DeleteTruckHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
