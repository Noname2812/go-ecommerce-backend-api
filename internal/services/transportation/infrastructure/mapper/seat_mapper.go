package transportationmapper

import (
	transportationqueryresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/application/query/dto/response"
	transportationmodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/transportation/domain/model"
)

// mapper seat lock model to seat item in detail trip response
func SeatLockModelToSeatStatusItem(seat *transportationmodel.TripSeatLock) transportationqueryresponse.Seat {
	return transportationqueryresponse.Seat{
		SeatID:       seat.Seat.SeatId,
		SeatNumber:   seat.Seat.SeatNumber,
		SeatRowNo:    seat.Seat.SeatRowNo,
		SeatColumnNo: seat.Seat.SeatColumnNo,
		SeatFloorNo:  seat.Seat.SeatFloorNo,
		SeatType:     uint8(seat.Seat.SeatType),
		SeatStatus:   uint8(seat.TripSeatLockStatus),
	}
}
