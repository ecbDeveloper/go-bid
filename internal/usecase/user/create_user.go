package user

import (
	"context"

	"github.com/ecbDeveloper/go-bid/internal/shared"
)

// useCase cuida das regras de cada corpo de requisição
type CreateUserReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUserReq) Valid(ctx context.Context) shared.ErrorsValidator {
	var eval shared.ErrorsValidator

	eval.CheckField(shared.NotBlank(req.UserName), "user_name", "this field cannot be empty")
	eval.CheckField(shared.NotBlank(req.Email), "email", "this field cannot be empty")
	eval.CheckField(shared.NotBlank(req.Bio), "bio", "this field cannot be empty")

	eval.CheckField(shared.IsEmail(req.Email), "email", "email must be a valid email")

	eval.CheckField(
		shared.MinChars(req.Bio, 10) &&
			shared.MaxChars(req.Bio, 255),
		"bio",
		"this filed must have a length between 10 and 255 chars",
	)

	eval.CheckField(
		shared.MinChars(req.Password, 8) &&
			shared.MaxChars(req.Password, 50),
		"password",
		"this field must have a length between 8 and 50 chars",
	)

	return eval
}
