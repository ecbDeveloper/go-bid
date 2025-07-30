package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	// csrfMiddleware := csrf.Protect(
	// 	[]byte(os.Getenv("GOBID_CSRF_KEY")),
	// 	csrf.Secure(false), // DEV ONLY
	// )

	// api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		// r.Get("/csrftoken", api.HandleGetCSRFtoken)
		api.bindUsersRoutes(r)
		api.bindProductsRoutes(r)
	})
}

func (api *Api) bindUsersRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", api.handleSignupUser)
		r.Post("/login", api.handleLoginUser)
		r.With(api.AuthMiddeware).Post("/logout", api.handleLogoutUser)
	})
}

func (api *Api) bindProductsRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Post("/", api.handleCreateProduct)
	})
}
