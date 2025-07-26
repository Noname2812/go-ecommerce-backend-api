package authcommandrequest

import (
	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserRegistratorRequest struct {
	Email   string `json:"email" validate:"required,email"`                            // email
	Purpose string `json:"purpose" validate:"required,oneof=TEST_USER CUSTOMER ADMIN"` // purpose (TEST_USER, CUSTOMER, ADMIN, etc.)
}

func (u *UserRegistratorRequest) Validate(ctx *gin.Context) map[string]string {
	validate := ctx.Value("validation").(*validator.Validate)
	errors := utils.ValidateStructWithValidatorTags(validate, u)
	if len(errors) == 0 {
		return nil
	}
	return errors
}
