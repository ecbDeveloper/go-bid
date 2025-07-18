package api

import (
	"github.com/go-chi/chi/v5"
)

func (api *Api) BindRoutes() {
	api.Router.Route("/api", func(r chi.Router) {
		api.bindUsersRoutes(r)
	})
}

func (api *Api) bindUsersRoutes(r chi.Router) {
	r.Route("/users/", func(r chi.Router) {
		r.Post("/signup", api.handleSignupUser)
	})
}
