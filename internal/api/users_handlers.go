package api

import (
	"errors"
	"net/http"

	"github.com/ecbDeveloper/go-bid/internal/jsonutils"
	"github.com/ecbDeveloper/go-bid/internal/services"
	"github.com/ecbDeveloper/go-bid/internal/usecase/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(r.Context(),
		data.UserName,
		data.Email,
		data.Password,
		data.Bio,
	)
	if err != nil {
		if errors.Is(err, services.ErrDuplciateEmailOrUsername) {
			_ = jsonutils.EncodeJson(w, http.StatusBadRequest, map[string]any{
				"error": "email or password already exists",
			})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, http.StatusOK, map[string]any{
		"user_id": id,
	})
}

func (api *Api) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.LoginUserReq](r)
	if err != nil {
		jsonutils.EncodeJson(w, http.StatusUnprocessableEntity, problems)
		return
	}

	userId, err := api.UserService.AuthenticateUser(r.Context(), data.Email, data.Password)
	if err != nil {
		if errors.Is(err, services.ErrInvalidCredentials) {
			jsonutils.EncodeJson(w, http.StatusBadRequest, map[string]string{
				"error": "invalid email or password",
			})
			return
		}
		jsonutils.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "unexpected internal server error",
		})
		return
	}

	err = api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Put(r.Context(), "AuthenticatedUserId", userId)

	jsonutils.EncodeJson(w, http.StatusOK, map[string]string{
		"message": "logged in successfully",
	})

}

func (api *Api) handleLogoutUser(w http.ResponseWriter, r *http.Request) {
	err := api.Sessions.RenewToken(r.Context())
	if err != nil {
		jsonutils.EncodeJson(w, http.StatusInternalServerError, map[string]string{
			"error": "unexpected internal server error",
		})
		return
	}

	api.Sessions.Remove(r.Context(), "AuthenticatedUserId")
	jsonutils.EncodeJson(w, http.StatusOK, map[string]string{
		"message": "logged out successfully",
	})
}
