package authcommandhandler

import (
	authcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto"
	authservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type authCommandHandler struct {
	acs    authservice.AuthCommandService
	logger *zap.Logger
}

// User Registration documentation
// @Summary      User Registration
// @Description  When user is registered send otp to email
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body usercommandrequest.UserRegistratorRequest true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /auth/register [post]
func (a *authCommandHandler) Register(ctx *gin.Context) {
	traceID := ctx.GetString("trace_id")
	var body authcommandrequest.UserRegistratorRequest

	// Parse JSON payload
	if err := ctx.ShouldBindJSON(&body); err != nil {
		a.logger.Warn("Invalid registration payload",
			zap.String("trace_id", traceID),
			zap.Error(err),
		)
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, "Email is invalid")
		return
	}

	// Call service layer
	code, err := a.acs.Register(ctx.Request.Context(), &body)
	if err != nil {
		a.logger.Error("User registration failed",
			zap.String("trace_id", traceID),
			zap.String("email", body.Email),
			zap.Int("err_code", code),
			zap.Error(err),
		)
		response.ErrorResponse(ctx, code, err.Error())
		return
	}
	response.SuccessResponse(ctx, code, nil)
}

func NewAuthCommandHttpHandler(acs authservice.AuthCommandService, logger *zap.Logger) AuthCommandHttpHandler {
	return &authCommandHandler{acs: acs, logger: logger}
}
