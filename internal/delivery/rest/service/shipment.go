package service

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/app"
)

type ShipmentService struct {
	uc internal.ShipmentUC
}

func NewShipmentService(app *app.ShipmentApp) *ShipmentService {
	return &ShipmentService{
		uc: app.Usecases.ShipmentUC,
	}
}

func (s *ShipmentService) CreateNewShipmentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *ShipmentService) GetShipmentListHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *ShipmentService) GetShipmentByIDHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *ShipmentService) UpdateShipmentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (s *ShipmentService) DeleteShipmentHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}
