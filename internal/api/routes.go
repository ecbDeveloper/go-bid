package api

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	csrfMiddleware := csrf.Protect(
		[]byte(os.Getenv("GOBID_CSRF_KEY")),
		csrf.Secure(false), // DEV ONLY
	)

	api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		r.Get("/csrftoken", api.HandleGetCSRFtoken)
		api.bindUsersRoutes(r)
		api.bindProductsRoutes(r)
		api.bindBidsRoutes(r)
	})
}

func (api *Api) bindUsersRoutes(r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/signup", api.handleSignupUser)
		r.Post("/login", api.handleLoginUser)

		r.Group(func(r chi.Router) {
			r.Use(api.AuthMiddeware)

			r.Post("/logout", api.handleLogoutUser)
		})
	})
}

func (api *Api) bindProductsRoutes(r chi.Router) {
	r.Route("/products", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(api.AuthMiddeware)

			r.Post("/", api.handleCreateProductAuction)

			r.Get("/ws/subscriber/{productId}", api.handleSubscribeUserToAuction)
		})

		r.Get("/", api.handleGetAllProducts)
	})
}

func (api *Api) bindBidsRoutes(r chi.Router) {
	r.Route("/bids", func(r chi.Router) {
		r.Get("/highest/{productId}", api.handleGetHighestBidByProductId)
	})
}
