package api

import (
	"errors"
	"net/http"

	"github.com/ecbDeveloper/go-bid/internal/services"
	"github.com/ecbDeveloper/go-bid/internal/usecase/user"
	"github.com/ecbDeveloper/go-bid/internal/utils"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := utils.DecodeValidJson[user.CreateUserReq](r)
	if err != nil {
		_ = utils.EncodeJson(w, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio,
	)
	if err != nil {
		if errors.Is(err, services.ErrDuplciateEmailOrPassword) {
			_ = utils.EncodeJson(w, http.StatusBadRequest, map[string]any{
				"error": "email or password already exists",
			})
			return
		}
	}

	_ = utils.EncodeJson(w, http.StatusOK, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO - NOT IMPLEMENTED")
}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO - NOT IMPLEMENTED")
}
