package user

import (
	"context"

	"github.com/ecbDeveloper/go-bid/internal/shared"
)

type LoginUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req LoginUserReq) Valid(ctx context.Context) shared.ErrorsValidator {
	var eval shared.ErrorsValidator

	eval.CheckField(shared.NotBlank(req.Email), "email", "this field cannot be empty")
	eval.CheckField(shared.NotBlank(req.Password), "password", "this field cannot be empty")

	eval.CheckField(shared.IsEmail(req.Email), "email", "must be a valid email")

	return eval
}
