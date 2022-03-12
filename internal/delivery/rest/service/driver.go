package service

import (
	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/app"
	"net/http"
)

type DriverService struct {
	uc internal.DriverUC
}

func NewDriverService(app *app.ShipmentApp) *DriverService {
	return &DriverService{
		uc: app.Usecases.DriverUC,
	}
}

func (s *DriverService) CreateNewDriverHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *DriverService) GetDriverListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *DriverService) GetDriverByIDHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *DriverService) UpdateDriverHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *DriverService) DeleteDriverHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
