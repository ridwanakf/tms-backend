package internal

import "github.com/ridwanakf/tms-backend/internal/entity"

type TruckRepo interface {
	CreateNewTruck(req entity.CreateNewTruckRequest) (entity.CreateNewTruckResponse, error)
	GetTruckList(req entity.GetTruckListRequest) (entity.GetTruckListResponse, error)
	GetTruckByID(req entity.GetTruckByIDRequest) (entity.GetTruckByIDResponse, error)
	UpdateTruck(req entity.UpdateTruckRequest) (entity.UpdateTruckResponse, error)
	DeleteTruck(req entity.DeleteTruckRequest) (entity.DeleteTruckResponse, error)
}

type DriverRepo interface {
	CreateNewDriver(req entity.CreateNewDriverRequest) (entity.CreateNewDriverResponse, error)
	GetDriverList(req entity.GetDriverListRequest) (entity.GetDriverListResponse, error)
	GetDriverByID(req entity.GetDriverByIDRequest) (entity.GetDriverByIDResponse, error)
	UpdateDriver(req entity.UpdateDriverRequest) (entity.UpdateDriverResponse, error)
	DeleteDriver(req entity.DeleteDriverRequest) (entity.DeleteDriverResponse, error)
}

type ShipmentRepo interface {
	CreateNewShipment(req entity.CreateNewShipmentRequest) (entity.CreateNewShipmentResponse, error)
	GetShipmentList(req entity.GetShipmentListRequest) (entity.GetShipmentListResponse, error)
	GetShipmentByID(req entity.GetShipmentByIDRequest) (entity.GetShipmentByIDResponse, error)
	UpdateShipment(req entity.UpdateShipmentRequest) (entity.UpdateShipmentResponse, error)
	DeleteShipment(req entity.DeleteShipmentRequest) (entity.DeleteShipmentResponse, error)
}
