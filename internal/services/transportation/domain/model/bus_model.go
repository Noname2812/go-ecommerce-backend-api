package transportationmodel

import (
	"time"

	"github.com/shopspring/decimal"
)

type Bus struct {
	BusId           uint64
	BusLicensePlate string
	BusCompany      string
	BusPrice        decimal.Decimal
	BusCapacity     uint8
	BusCreatedAt    time.Time
	BusUpdatedAt    time.Time
	BusDeletedAt    *time.Time
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
