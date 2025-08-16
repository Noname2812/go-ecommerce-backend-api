package bookingcommandhandler

import (
	"encoding/json"

	bookingcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/command/dto/request"
	bookingservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/booking/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bookingCommandHandler struct {
	bookingCommandService bookingservice.BookingCommandService
	logger                *zap.Logger
}

// Create Booking
// @Summary      Create Booking
// @Description  Create Booking
// @Tags         booking management
// @Accept       json
// @Produce      json
// @Param        payload body bookingcommandrequest.CreateBookingRequest true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Router       /booking/create [post]
func (b *bookingCommandHandler) CreateBooking(ctx *gin.Context) {
	var body bookingcommandrequest.CreateBookingRequest

	// Parse JSON payload
	if err := json.NewDecoder(ctx.Request.Body).Decode(&body); err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidJson, err.Error(), map[string]string{"error": err.Error()})
		return
	}

	// Validate payload
	if err := body.Validate(ctx); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "", err)
		return
	}
	code, err, data := b.bookingCommandService.CreateBooking(ctx.Request.Context(), &body)
	if err != nil {
		b.logger.Error("Create booking failed",
			zap.String("trace_id", ctx.GetString("trace_id")),
			zap.Int("err_code", code),
			zap.Error(err),
		)
		response.ErrorResponse(ctx, code, "", map[string]string{"error": err.Error()})
		return
	}
	response.SuccessResponse(ctx, code, data)
}

func NewBookingCommandHttpHandler(bookingCommandService bookingservice.BookingCommandService, logger *zap.Logger) BookingCommandHttpHandler {
	return &bookingCommandHandler{
		bookingCommandService: bookingCommandService,
		logger:                logger,
	}
}
