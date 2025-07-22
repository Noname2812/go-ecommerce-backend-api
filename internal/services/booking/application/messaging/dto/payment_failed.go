package bookingmessagingdto

import (
	"encoding/json"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
)

type PaymentFailed struct {
	BookingID string                 `json:"booking_id"`
	UserID    *string                `json:"user_id"`
	Status    commonenum.EventStatus `json:"status"`
}

func UnmarshalPaymentFailed(data []byte) (*PaymentFailed, error) {
	var paymentFailed PaymentFailed
	if err := json.Unmarshal(data, &paymentFailed); err != nil {
		return nil, err
	}
	return &paymentFailed, nil
}
