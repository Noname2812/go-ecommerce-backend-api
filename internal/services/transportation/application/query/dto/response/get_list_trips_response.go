package transportationqueryresponse

import (
	"github.com/shopspring/decimal"
)

type GetListTripsResponse struct {
	Trips []Trip `json:"trips"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
}

type Trip struct {
	ID            uint64          `json:"id"`
	FromLocation  string          `json:"from_location"`
	ToLocation    string          `json:"to_location"`
	DepartureTime string          `json:"departure_time"`
	ArrivalTime   string          `json:"arrival_time"`
	Price         decimal.Decimal `json:"price"`
}
