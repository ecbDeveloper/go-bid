package api

import (
	"github.com/ecbDeveloper/go-bid/internal/services"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router      *chi.Mux //criamos o Router, ao usarmos r := chi.NewRouter() == var r *chi.Mux
	UserService services.UserService
}
