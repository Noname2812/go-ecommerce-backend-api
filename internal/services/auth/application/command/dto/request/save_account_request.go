package authcommandrequest

import (
	"time"

	"github.com/Noname2812/go-ecommerce-backend-api/internal/common/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type SaveAccountRequest struct {
	Token       string `json:"token" validate:"required"`                         // token
	Email       string `json:"email" validate:"required,email"`                   // email
	Password    string `json:"password" validate:"required,min=8"`                // password
	ConfirmPass string `json:"confirm_pass" validate:"required,eqfield=Password"` // confirm password
	Name        string `json:"name" validate:"required"`                          // name
	Phone       string `json:"phone" validate:"omitempty,e164"`                   // phone
	Gender      int8   `json:"gender" validate:"omitempty,oneof=0 1 2"`           // gender
	Birthday    string `json:"birthday" validate:"omitempty"`                     // birthday
}

func (s *SaveAccountRequest) Validate(ctx *gin.Context) map[string]string {
	validate := ctx.Value("validation").(*validator.Validate)
	errors := utils.ValidateStructWithValidatorTags(validate, s)
	if len(errors) == 0 {
		return nil
	}
	s.validateBirthday(errors)
	return errors
}

// validate birthday
func (s *SaveAccountRequest) validateBirthday(errors map[string]string) {
	if s.Birthday == "" {
		return
	}
	_, err := time.Parse("2006-01-02", s.Birthday)
	if err != nil {
		errors["birthday"] = "Birthday must be in format YYYY-MM-DD"
	}
}
