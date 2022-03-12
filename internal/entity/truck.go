package entity

import "time"

type Truck struct {
	ID             int64     `json:"id" db:"id"`
	LicenseNumber  string    `json:"license_number" db:"license_number"`
	TruckType      string    `json:"truck_type" db:"truck_type"`
	PlateType      string    `json:"plate_type" db:"plate_type"`
	ProductionYear int64     `json:"production_year" db:"production_year"`
	Status         bool      `json:"status" db:"status"`
	ShipmentStatus bool      `json:"shipment_status" db:"shipment_status"`
	STNK           string    `json:"stnk" db:"stnk"`
	KIR            string    `json:"kir" db:"kir"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at" db:"deleted_at"`
}

// Requests
type (
	CreateNewTruckRequest struct {
	}

	GetTruckListRequest struct {
	}

	GetTruckByIDRequest struct {
	}

	UpdateTruckRequest struct {
	}

	DeleteTruckRequest struct {
	}
)

// Response
type (
	CreateNewTruckResponse struct {
	}

	GetTruckListResponse struct {
	}

	GetTruckByIDResponse struct {
	}

	UpdateTruckResponse struct {
	}

	DeleteTruckResponse struct {
	}
)
