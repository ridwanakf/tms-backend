package usecase

import (
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type ShipmentUsecase struct {
	repo internal.ShipmentRepo
}

func NewShipmentUsecase(repo internal.ShipmentRepo) *ShipmentUsecase {
	return &ShipmentUsecase{
		repo: repo,
	}
}

func (uc *ShipmentUsecase) CreateNewShipment(req entity.CreateNewShipmentRequest) (entity.CreateNewShipmentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *ShipmentUsecase) GetShipmentList(req entity.GetShipmentListRequest) (entity.GetShipmentListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *ShipmentUsecase) GetShipmentByID(req entity.GetShipmentByIDRequest) (entity.GetShipmentByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *ShipmentUsecase) UpdateShipment(req entity.UpdateShipmentRequest) (entity.UpdateShipmentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *ShipmentUsecase) DeleteShipment(req entity.DeleteShipmentRequest) (entity.DeleteShipmentResponse, error) {
	//TODO implement me
	panic("implement me")
}
