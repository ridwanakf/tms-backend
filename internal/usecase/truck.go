package usecase

import (
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type TruckUsecase struct {
	repo internal.TruckRepo
}

func NewTruckUsecase(repo internal.TruckRepo) *TruckUsecase {
	return &TruckUsecase{
		repo: repo,
	}
}

func (uc *TruckUsecase) CreateNewTruck(req entity.CreateNewTruckRequest) (entity.CreateNewTruckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *TruckUsecase) GetTruckList(req entity.GetTruckListRequest) (entity.GetTruckListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *TruckUsecase) GetTruckByID(req entity.GetTruckByIDRequest) (entity.GetTruckByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *TruckUsecase) UpdateTruck(req entity.UpdateTruckRequest) (entity.UpdateTruckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *TruckUsecase) DeleteTruck(req entity.DeleteTruckRequest) (entity.DeleteTruckResponse, error) {
	//TODO implement me
	panic("implement me")
}
