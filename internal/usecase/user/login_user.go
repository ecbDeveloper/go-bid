package user

import (
	"context"

	"github.com/ecbDeveloper/go-bid/internal/validator"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserReq) Valid(ctx context.Context) validator.ErrorsValidator {
	var eval validator.ErrorsValidator

	eval.CheckField(validator.NotBlank(req.Email), "email", "this field cannot be empty")
	eval.CheckField(validator.NotBlank(req.Password), "password", "this field cannot be empty")

	eval.CheckField(validator.IsEmail(req.Email), "email", "must be a valid email")

	return eval
}
