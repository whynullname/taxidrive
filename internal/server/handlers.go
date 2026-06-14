package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/whynullname/taxidrive/internal/car"
)

type HandlersProvider struct {
	carHandlers car.Handlers
}

func NewHandlersProvider(carHandlers car.Handlers) *HandlersProvider {
	return &HandlersProvider{
		carHandlers: carHandlers,
	}
}

func (h *HandlersProvider) ConnectHandlers(router chi.Router) {
	//-------------CARS--------------------
	router.Route("/cars", func(r chi.Router) {
		r.Post("/", h.carHandlers.CreateCar)
		r.Get("/", h.carHandlers.GetAllCars)
		r.Get("/{id}", h.carHandlers.GetCar)
		r.Patch("/{id}", h.carHandlers.UpdateCar)
		r.Delete("/{id}", h.carHandlers.DeleteCar)
	})
}
