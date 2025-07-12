package userqueryhandler

import (
	"database/sql"
	"strconv"

	userqueryresponse "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/query/dto"
	userservice "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/application/service"
	"github.com/Noname2812/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type userQueryHttpHandler struct {
	service userservice.UserQueryService
}

func NewUserQueryHandler(service userservice.UserQueryService) UserQueryHandler {
	return &userQueryHttpHandler{service: service}
}

// GetUserDetails
// @Summary      Get user details by ID
// @Description  Returns user details data based on ID
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id  path  int  true  "User ID"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/{id} [get]
func (ah *userQueryHttpHandler) GetUserDetails(ctx *gin.Context) {
	idStr := ctx.Param("id")
	userID, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	user, error := ah.service.GetUserProfile(ctx.Request.Context(), userID)
	if error != nil {
		if error == sql.ErrNoRows {
			response.ErrorResponse(ctx, response.ErrCodeUserNotFound, "User not found")
			return
		}
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, error.Error())
		return
	}

	result := userqueryresponse.ToUserInfoResponse(*user)
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
