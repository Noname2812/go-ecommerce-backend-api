package transportationqueryresponse

import "github.com/shopspring/decimal"

type TripDetailResponse struct {
	ID                uint64          `json:"id"`
	FromLocation      string          `json:"from_location"`
	ToLocation        string          `json:"to_location"`
	DepartureDate     string          `json:"departure_date"`
	ArrivalDate       string          `json:"arrival_date"`
	EstimatedDuration uint64          `json:"estimated_duration"`
	Price             decimal.Decimal `json:"price"`
	Seats             []Seat          `json:"seats"`
	Bus               Bus             `json:"bus"`
}

type Seat struct {
	SeatNumber string `json:"seat_number"`
	Available  bool   `json:"available"`
	Locked     bool   `json:"locked,omitempty"`
}

type Bus struct {
	BusLicensePlate string `json:"bus_license_plate"`
	BusCompany      string `json:"bus_company"`
	BusCapacity     uint8  `json:"bus_capacity"`
}
