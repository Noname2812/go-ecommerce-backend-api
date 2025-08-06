package transportationqueryresponse

import "github.com/shopspring/decimal"

type TripDetailResponse struct {
	TripID        uint64          `json:"trip_id"`
	DepartureTime string          `json:"departure_time"`
	ArrivalTime   string          `json:"arrival_time"`
	BasePrice     decimal.Decimal `json:"base_price"`
	Seats         []Seat          `json:"seats"`
	Bus           Bus             `json:"bus"`
	Route         Route           `json:"route"`
}

type Seat struct {
	SeatID       uint64 `json:"seat_id"`
	SeatNumber   string `json:"seat_number"`
	SeatRowNo    uint8  `json:"seat_row_no"`
	SeatColumnNo uint8  `json:"seat_column_no"`
	SeatFloorNo  uint8  `json:"seat_floor_no"`
	SeatType     uint8  `json:"seat_type"`
	SeatStatus   uint8  `json:"seat_status"`
}

type Bus struct {
	BusID           uint64 `json:"bus_id"`
	BusLicensePlate string `json:"bus_license_plate"`
	BusCompany      string `json:"bus_company"`
	BusCapacity     uint8  `json:"bus_capacity"`
}

type Route struct {
	RouteID      uint64 `json:"route_id"`
	FromLocation string `json:"from_location"`
	ToLocation   string `json:"to_location"`
}
