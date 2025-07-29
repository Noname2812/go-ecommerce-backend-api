package transportationmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

// Bus model
type Bus struct {
	BusId           uint64          // primary key
	BusLicensePlate string          // license plate (ex: "51A-12345")
	BusCompany      string          // company name (ex: "Vietnam Airlines", "Vietjet Air")
	BusPrice        decimal.Decimal // price (ex: 1000000 VND)
	BusCapacity     uint8           // capacity (ex: 100)
	BusCreatedAt    time.Time       // created at (ex: 2021-01-01 00:00:00)
	BusUpdatedAt    time.Time       // updated at (ex: 2021-01-01 00:00:00)
	BusDeletedAt    *time.Time      // deleted at (ex: 2021-01-01 00:00:00)

	// Navigation property - lazy loading
	Trips []Trip // one to many
}

func NewBus(licensePlate, company string, price decimal.Decimal, capacity uint8) *Bus {
	return &Bus{
		BusLicensePlate: licensePlate,
		BusCompany:      company,
		BusPrice:        price,
		BusCapacity:     capacity,
		BusCreatedAt:    time.Now(),
		BusUpdatedAt:    time.Now(),
		BusDeletedAt:    nil,
	}
}
