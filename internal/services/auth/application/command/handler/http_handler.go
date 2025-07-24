package authcommandhandler

import (
	"encoding/json"

	authcommandrequest "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/command/dto/request"
	authservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/auth/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type authCommandHandler struct {
	acs    authservice.AuthCommandService
	logger *zap.Logger
}

// User Base Registration
// @Summary      User Base Registration
// @Description  When user has registered send otp to email
// @Tags         auth management
// @Accept       json
// @Produce      json
// @Param        payload body authcommandrequest.SaveAccountRequest true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Router       /auth/save-account [post]
func (a *authCommandHandler) SaveAccount(ctx *gin.Context) {

	var body authcommandrequest.SaveAccountRequest

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
	code, err := a.acs.SaveAccount(ctx.Request.Context(), &body)
	if err != nil {
		a.logger.Error("Save account failed",
			zap.String("trace_id", ctx.GetString("trace_id")),
			zap.String("email", body.Email),
			zap.Int("err_code", code),
			zap.Error(err),
		)
		response.ErrorResponse(ctx, code, err.Error(), nil)
		return
	}
	response.SuccessResponse(ctx, code, nil)
}

// User Verification OTP documentation
// @Summary      Verify OTP
// @Description  When user is verified otp from email
// @Tags         auth management
// @Accept       json
// @Produce      json
// @Param        payload body authcommandrequest.VerifyOTPRequest true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Router       /auth/verify-account [post]
func (a *authCommandHandler) VerifyOTP(ctx *gin.Context) {
	var body authcommandrequest.VerifyOTPRequest

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

	code, data, err := a.acs.VerifyOTP(ctx.Request.Context(), &body)
	if err != nil {
		a.logger.Error("Verify OTP failed",
			zap.String("trace_id", ctx.GetString("trace_id")),
			zap.String("email", body.Email),
			zap.Int("err_code", code),
			zap.Error(err),
		)
		response.ErrorResponse(ctx, code, err.Error(), nil)
		return
	}
	response.SuccessResponse(ctx, code, data)
}

// User Registration documentation
// @Summary      User Registration
// @Description  When user is registered send otp to email
// @Tags         auth management
// @Accept       json
// @Produce      json
// @Param        payload body authcommandrequest.UserRegistratorRequest true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Router       /auth/register [post]
func (a *authCommandHandler) Register(ctx *gin.Context) {
	traceID := ctx.GetString("trace_id")
	var body authcommandrequest.UserRegistratorRequest

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

	// Call service layer
	code, err := a.acs.Register(ctx.Request.Context(), &body)
	if err != nil {
		a.logger.Error("Register failed",
			zap.String("trace_id", traceID),
			zap.String("email", body.Email),
			zap.Int("err_code", code),
			zap.Error(err),
		)
		response.ErrorResponse(ctx, code, err.Error(), nil)
		return
	}
	response.SuccessResponse(ctx, code, nil)
}

func NewAuthCommandHttpHandler(acs authservice.AuthCommandService, logger *zap.Logger) AuthCommandHttpHandler {
	return &authCommandHandler{acs: acs, logger: logger}
}
