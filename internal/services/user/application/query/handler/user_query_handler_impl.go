package userqueryhandler

import (
	"strconv"

	userqueryresponsedto "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/query/dto/response"
	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userQueryHttpHandler struct {
	service userservice.UserQueryService
	logger  *zap.Logger
}

func NewUserQueryHandler(service userservice.UserQueryService, logger *zap.Logger) UserQueryHandler {
	return &userQueryHttpHandler{service: service, logger: logger}
}

// GetUserDetails
// @Summary      Get user details by ID
// @Description  Returns user details data based on ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  response.ResponseData
// @Failure      400  {object}  response.ErrorResponseData
// @Router       /user/{id} [get]
func (ah *userQueryHttpHandler) GetUserDetails(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error(), nil)
		return
	}
	user, err := ah.service.GetUserProfile(ctx.Request.Context(), userID)
	if err != nil {
		ah.logger.Error("Get user details failed",
			zap.String("trace_id", ctx.GetString("trace_id")),
			zap.Uint64("user_id", userID),
			zap.Error(err),
		)
		response.ErrorResponse(ctx, response.ErrServerError, err.Error(), nil)
		return
	}
	if user == nil {
		response.ErrorResponse(ctx, response.ErrCodeUserNotFound, "user not found", nil)
		return
	}
	result := userqueryresponsedto.ToUserInfoResponse(*user)
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
