package usecase

import (
	"github.com/ridwanakf/tms-backend/internal"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type DriverUsecase struct {
	repo internal.DriverRepo
}

func NewDriverUsecase(repo internal.DriverRepo) *DriverUsecase {
	return &DriverUsecase{
		repo: repo,
	}
}

func (uc *DriverUsecase) CreateNewDriver(req entity.CreateNewDriverRequest) (entity.CreateNewDriverResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *DriverUsecase) GetDriverList(req entity.GetDriverListRequest) (entity.GetDriverListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *DriverUsecase) GetDriverByID(req entity.GetDriverByIDRequest) (entity.GetDriverByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *DriverUsecase) UpdateDriver(req entity.UpdateDriverRequest) (entity.UpdateDriverResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *DriverUsecase) DeleteDriver(req entity.DeleteDriverRequest) (entity.DeleteDriverResponse, error) {
	//TODO implement me
	panic("implement me")
}
