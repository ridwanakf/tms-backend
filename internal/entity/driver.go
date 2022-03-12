package entity

import "time"

type Driver struct {
	ID             int64     `json:"id" db:"db"`
	Name           string    `json:"name" db:"name"`
	Phone          string    `json:"phone" db:"phone"`
	CardID         int64     `json:"id_card" db:"id_card"`
	DriverLicense  int64     `json:"driver_license" db:"driver_license"`
	Status         bool      `json:"status" db:"status"`
	ShipmentStatus bool      `json:"shipment_status" db:"shipment_status"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt      time.Time `json:"deleted_at" db:"deleted_at"`
}
