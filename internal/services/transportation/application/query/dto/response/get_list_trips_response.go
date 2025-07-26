package transportationqueryresponse

import "github.com/shopspring/decimal"

type GetListTripsResponse struct {
	Trips []Trip `json:"trips"`
	Total int    `json:"total"`
	Page  int    `json:"page"`
}

type Trip struct {
	ID            uint64          `json:"id"`
	FromLocation  string          `json:"from_location"`
	ToLocation    string          `json:"to_location"`
	DepartureDate string          `json:"departure_date"`
	ArrivalDate   string          `json:"arrival_date"`
	Price         decimal.Decimal `json:"price"`
}
