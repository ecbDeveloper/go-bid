package api

import (
	"net/http"

	"github.com/ecbDeveloper/go-bid/internal/jsonutils"
	"github.com/gorilla/csrf"
)

func (api *Api) AuthMiddeware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.EncodeJson(w, http.StatusUnauthorized, map[string]string{
				"message": "must be logged in",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *Api) HandleGetCSRFtoken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	jsonutils.EncodeJson(w, http.StatusOK, map[string]any{
		"csrf_token": token,
	})
}
