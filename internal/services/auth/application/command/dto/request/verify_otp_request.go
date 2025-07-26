package authcommandrequest

import (
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VerifyOTPRequest struct {
	Email string `json:"email" validate:"required,email"` // email
	OTP   string `json:"otp" validate:"required"`         // otp
}

func (v *VerifyOTPRequest) Validate(ctx *gin.Context) map[string]string {
	validate := ctx.Value("validation").(*validator.Validate)
	errors := utils.ValidateStructWithValidatorTags(validate, v)
	if len(errors) == 0 {
		return nil
	}
	return errors
}
