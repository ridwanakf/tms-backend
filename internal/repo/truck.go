package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/ridwanakf/tms-backend/internal/entity"
)

type TruckDB struct {
	db *sqlx.DB
}

func NewTruckDB(db *sqlx.DB) *TruckDB {
	return &TruckDB{
		db: db,
	}
}

func (r *TruckDB) CreateNewTruck(req entity.CreateNewTruckRequest) (entity.CreateNewTruckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TruckDB) GetTruckList(req entity.GetTruckListRequest) (entity.GetTruckListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TruckDB) GetTruckByID(req entity.GetTruckByIDRequest) (entity.GetTruckByIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TruckDB) UpdateTruck(req entity.UpdateTruckRequest) (entity.UpdateTruckResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TruckDB) DeleteTruck(req entity.DeleteTruckRequest) (entity.DeleteTruckResponse, error) {
	//TODO implement me
	panic("implement me")
}
