package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type DriverDB struct {
	db *sqlx.DB
}

func NewDriverDB(db *sqlx.DB) *DriverDB {
	return &DriverDB{
		db: db,
	}
}

func (r *DriverDB) CreateNewDriver(req entity.CreateNewDriverRequest) (entity.CreateNewDriverResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DriverDB) GetDriverList(req entity.GetDriverListRequest) (entity.GetDriverListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DriverDB) GetDriverByID(req entity.GetDriverByIDRequest) (entity.GetDriverByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DriverDB) UpdateDriver(req entity.UpdateDriverRequest) (entity.UpdateDriverResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DriverDB) DeleteDriver(req entity.DeleteDriverRequest) (entity.DeleteDriverResponse, error) {
	//TODO implement me
	panic("implement me")
}
